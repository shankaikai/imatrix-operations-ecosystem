// TODO: Add validation
package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"capstone.operations_ecosystem/backend/common"
	db_pck "capstone.operations_ecosystem/backend/database"
	rs "capstone.operations_ecosystem/backend/rating_system"
	tclient "capstone.operations_ecosystem/backend/telegram_client"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "capstone.operations_ecosystem/backend/proto"
)

// Send back all the available users that are not yet assigned for the particular day
// Returns the available and unavailable ones in two separate lists
func (s *Server) GetAvailableIspecialistsWithScore(query *pb.AvailabilityQuery) ([]*pb.EmployeeEvaluation, []*pb.EmployeeEvaluation, error) {
	fmt.Println("GetAvailableUsersUtil")

	employeeEvals := make([]*pb.EmployeeEvaluation, 0)
	availEmployeeEvals := make([]*pb.EmployeeEvaluation, 0)
	unavailEmployeeEvals := make([]*pb.EmployeeEvaluation, 0)

	// First get all users who are not yet assigned
	unassignedUsers, err := s.getUnassignedIspecialists(query)
	if err != nil {
		fmt.Println("getUnassignedIspecialists err here", err)
		return availEmployeeEvals, unavailEmployeeEvals, err
	}

	// Get the scores for all these users
	err = getParallelEmployeeScores(unassignedUsers, &employeeEvals)
	if err != nil {
		fmt.Println("getParallelEmployeeScores err here", err)
		return availEmployeeEvals, unavailEmployeeEvals, err
	}

	// Find out which of these users are available
	err = s.updateAvailability(query, employeeEvals, &availEmployeeEvals, &unavailEmployeeEvals)
	if err != nil {
		fmt.Println("updateAvailability err here", err)
		return availEmployeeEvals, unavailEmployeeEvals, err
	}

	return availEmployeeEvals, unavailEmployeeEvals, err
}

func (s *Server) getUnassignedIspecialists(query *pb.AvailabilityQuery) ([]*pb.User, error) {
	unassignedUsers := make([]*pb.User, 0)

	// Get a list of all the guard assignments for this day
	// who are still assigned
	rosterAssignmentQuery := &pb.RosterQuery{Limit: 20}
	// still assigned
	db_pck.AddRosterFilter(rosterAssignmentQuery, pb.RosterFilter_IS_ASSIGNED, pb.Filter_EQUAL, "1")
	// Start and end times of roster
	db_pck.AddRosterFilter(rosterAssignmentQuery, pb.RosterFilter_START_TIME, pb.Filter_GREATER_EQ, query.StartTime)
	db_pck.AddRosterFilter(rosterAssignmentQuery, pb.RosterFilter_END_TIME, pb.Filter_LESSER_EQ, query.EndTime)
	rosterAssignements, err := db_pck.GetRosterAssingments(s.db, rosterAssignmentQuery, -1)

	if err != nil {
		return unassignedUsers, err
	}

	// Check if found assignments is empty, if so get the default users for that day
	if len(rosterAssignements) == 0 {
		defaultRosterQuery := &pb.RosterQuery{Limit: 5}
		db_pck.AddRosterFilter(defaultRosterQuery, pb.RosterFilter_START_TIME, pb.Filter_EQUAL, query.StartTime)
		defaultRosters, err := db_pck.GetDefaultRosters(s.db, defaultRosterQuery)
		if err != nil {
			return unassignedUsers, err
		}
		for _, defaultRoster := range defaultRosters {
			rosterAssignements = append(rosterAssignements, defaultRoster.GuardAssigned...)
		}
	}

	// Create a string list of all the assigned users
	assignedUsers := make([]string, 0)
	for _, assignment := range rosterAssignements {
		assignedUsers = append(assignedUsers, strconv.Itoa(int(assignment.GuardAssigned.Employee.UserId)))
	}

	assignedUsersString := strings.Join(assignedUsers, ",")

	// Get all ISpecialists not in this list
	userQuery := &pb.UserQuery{Limit: db_pck.AVAILABILITY_DEFAULT_LIMIT}
	if len(assignedUsers) > 0 {
		db_pck.AddUserFilter(userQuery, pb.UserFilter_USER_ID, pb.Filter_NOT_IN, assignedUsersString)
	}
	db_pck.AddUserFilter(userQuery, pb.UserFilter_TYPE, pb.Filter_EQUAL, "I-Specialist")
	return db_pck.GetUsers(s.db, userQuery)
}

func getParallelEmployeeScores(unassignedUsers []*pb.User, employeeEvals *[]*pb.EmployeeEvaluation) error {
	fmt.Println("Getting user scores for unassigned users...")
	userScoreChannel := make(chan rs.ScoreStruct, len(unassignedUsers))
	for i, user := range unassignedUsers {
		employeeEval := &pb.EmployeeEvaluation{Employee: user}
		*employeeEvals = append(*employeeEvals, employeeEval)
		go rs.GetUserScoreFromChan(user, userScoreChannel, i)
	}

	for i := 0; i < len(unassignedUsers); i++ {
		scoreStruct := <-userScoreChannel
		if scoreStruct.Err != nil {
			return scoreStruct.Err
		}

		(*employeeEvals)[scoreStruct.ChannelId].EmployeeScore = scoreStruct.Score
	}

	return nil
}

func (s *Server) updateAvailability(query *pb.AvailabilityQuery,
	employeeEvals []*pb.EmployeeEvaluation, availEmployeeEvals *[]*pb.EmployeeEvaluation,
	unavailEmployeeEvals *[]*pb.EmployeeEvaluation) error {
	// Add week and year to the availability filter
	startTime, err := time.Parse(common.DATETIME_FORMAT, query.StartTime)

	if err != nil {
		fmt.Println("updateAvailability ERROR:", err)
		return err
	}

	endTime, err := time.Parse(common.DATETIME_FORMAT, query.EndTime)
	if err != nil {
		fmt.Println("updateAvailability ERROR:", err)
		return err
	}

	year, week := startTime.ISOWeek()
	db_pck.AddAvailabilityFilter(query, pb.AvailabilityFilter_YEAR, pb.Filter_EQUAL, strconv.Itoa(year))
	db_pck.AddAvailabilityFilter(query, pb.AvailabilityFilter_WEEK, pb.Filter_EQUAL, strconv.Itoa(week))

	// add day NOT NULL to the availability filter

	db_pck.AddAvailabilityFilter(query, getDayAvailabilityQueryEnum(startTime), pb.Filter_EQUAL, "")
	// check if the end time is on a diff day
	if endTime.Weekday() != startTime.Weekday() {
		// if the start day is sat, the end day will be next sun
		if startTime.Weekday() == time.Saturday {
			db_pck.AddAvailabilityFilter(query, pb.AvailabilityFilter_NEXT_SUN, pb.Filter_EQUAL, "")
		} else {
			db_pck.AddAvailabilityFilter(query, getDayAvailabilityQueryEnum(endTime), pb.Filter_EQUAL, "")
		}
	}

	// Get all availability for users who are available on that day
	availabilities, err := db_pck.GetAvailability(s.db, query)
	if err != nil {
		return err
	}

	availableUserIds := make([]int, 0)

	// For each availability, check if the timing matches with the shift
	// If the timing matches, add it to the list of available users id
	// create int list of available users
	for _, avail := range availabilities {
		if checkAvailabilityTiming(avail, query) {
			availableUserIds = append(availableUserIds, int(avail.Guard.UserId))
		}
	}

	sort.Ints(availableUserIds)

	for _, eval := range employeeEvals {
		// binary search is employee is in the available list
		isAvail, _ := common.BinarySearch(availableUserIds, 0, len(availableUserIds)-1, int(eval.Employee.UserId))
		eval.IsAvailable = isAvail
		fmt.Println(isAvail)
		if isAvail {
			*availEmployeeEvals = append(*availEmployeeEvals, eval)
		} else {
			*unavailEmployeeEvals = append(*unavailEmployeeEvals, eval)
		}
	}

	return nil
}

func getDayAvailabilityQueryEnum(date time.Time) pb.AvailabilityFilter_Field {
	switch date.Weekday() {
	case 0:
		return pb.AvailabilityFilter_SUN
	case 1:
		return pb.AvailabilityFilter_MON
	case 2:
		return pb.AvailabilityFilter_TUES
	case 3:
		return pb.AvailabilityFilter_WED
	case 4:
		return pb.AvailabilityFilter_THURS
	case 5:
		return pb.AvailabilityFilter_FRI
	default:
		return pb.AvailabilityFilter_SAT
	}
}

// Returns if the particular person is available
// based on their json array time
func checkAvailabilityTiming(availability *db_pck.Availability, availQuery *pb.AvailabilityQuery) bool {
	startAvail := false
	endAvail := false

	availStartArrayString, availEndArrayString := getStartEndTimeStrings(availability, availQuery)
	fmt.Println("Availability string", availStartArrayString, availEndArrayString)

	var startAvailArray []string
	var endAvailArray []string

	_ = json.Unmarshal([]byte(availStartArrayString), &startAvailArray)
	_ = json.Unmarshal([]byte(availEndArrayString), &endAvailArray)

	var sameDayEndTime time.Time
	startTimeQuery, err := time.Parse(common.DATETIME_FORMAT, availQuery.StartTime)

	if err != nil {
		fmt.Println("updateAvailability ERROR:", err)
		return false
	}

	endTimeQuery, err := time.Parse(common.DATETIME_FORMAT, availQuery.EndTime)

	if err != nil {
		fmt.Println("updateAvailability ERROR:", err)
		return false
	}

	for _, timeRange := range startAvailArray {
		// expect eg "18:00:00-23:59:59"
		startEnd := strings.Split(timeRange, "-")
		start := startEnd[0]
		end := startEnd[1]

		// parse string to actual time
		startTime, err := time.Parse(common.TIME_FORMAT, start)
		if err != nil {
			fmt.Println("checkAvailabilityTiming ERROR:", err)
			continue
		}

		endTime, err := time.Parse(common.TIME_FORMAT, end)
		if err != nil {
			fmt.Println("checkAvailabilityTiming ERROR:", err)
			continue
		}

		// check if the start time before the particular start time
		if startTime.Hour() > startTimeQuery.Hour() {
			continue
		}

		// check if the end of the range is after the specified end time
		if endTime.Hour() >= endTimeQuery.Hour() {
			startAvail = true
			sameDayEndTime = endTime
			fmt.Println("found starting availablity", startTime, endTime)
			break
		}
	}

	// Check for ending times
	// if on same day, check that the end time of the same index is ok
	if startTimeQuery.Day() == endTimeQuery.Day() {
		if sameDayEndTime.Hour() > endTimeQuery.Hour() {
			endAvail = true
		} else if sameDayEndTime.Hour() == endTimeQuery.Hour() {
			// check min
			endAvail = sameDayEndTime.Minute() >= endTimeQuery.Minute()
		}
	} else {
		// next day
		for _, timeRange := range endAvailArray {
			// expect eg "18:00:00-23:59:59"
			startEnd := strings.Split(timeRange, "-")
			start := startEnd[0]
			end := startEnd[1]

			// parse string to actual time
			startTime, err := time.Parse(common.TIME_FORMAT, start)
			if err != nil {
				fmt.Println("checkAvailabilityTiming ERROR:", err)
				continue
			}

			endTime, err := time.Parse(common.TIME_FORMAT, end)
			if err != nil {
				fmt.Println("checkAvailabilityTiming ERROR:", err)
				continue
			}
			// if next day, check that the start time for the next day is 00:00:00
			if startTime.Hour() != 0 || startTime.Minute() != 0 || startTime.Second() != 0 {
				continue
			}
			// if it is, check that the end time is after the specified end time or eq
			if endTime.Hour() > endTimeQuery.Hour() {
				endAvail = true
				fmt.Println("found ending availablity", startTime, endTime)
				break
			} else if endTime.Hour() == endTimeQuery.Hour() {
				// check min
				endAvail = endTime.Minute() >= endTimeQuery.Minute()
				fmt.Println("found ending availablity", startTime, endTime)
				break
			}
		}
	}

	return startAvail && endAvail
}

// return start and end time strings for a particular availability
// given the start and end datetimes
func getStartEndTimeStrings(availability *db_pck.Availability, availQuery *pb.AvailabilityQuery) (string, string) {
	availStartArrayString := "[]"
	availEndArrayString := "[]"

	startTime, err := time.Parse(common.DATETIME_FORMAT, availQuery.StartTime)

	if err != nil {
		fmt.Println("updateAvailability ERROR:", err)
		return availStartArrayString, availEndArrayString
	}

	endTime, err := time.Parse(common.DATETIME_FORMAT, availQuery.EndTime)

	if err != nil {
		fmt.Println("updateAvailability ERROR:", err)
		return availStartArrayString, availEndArrayString
	}

	switch startTime.Weekday() {
	case 0:
		availStartArrayString = availability.Sun.String
		if endTime.Weekday() != startTime.Weekday() {
			availEndArrayString = availability.Mon.String
		}
	case 1:
		availStartArrayString = availability.Mon.String
		if endTime.Weekday() != startTime.Weekday() {
			availEndArrayString = availability.Tues.String
		}
	case 2:
		availStartArrayString = availability.Tues.String
		if endTime.Weekday() != startTime.Weekday() {
			availEndArrayString = availability.Wed.String
		}
	case 3:
		availStartArrayString = availability.Wed.String
		if endTime.Weekday() != startTime.Weekday() {
			availEndArrayString = availability.Thurs.String
		}
	case 4:
		availStartArrayString = availability.Thurs.String
		if endTime.Weekday() != startTime.Weekday() {
			availEndArrayString = availability.Fri.String
		}
	case 5:
		availStartArrayString = availability.Fri.String
		if endTime.Weekday() != startTime.Weekday() {
			availEndArrayString = availability.Sat.String
		}
	default:
		availStartArrayString = availability.Sat.String
		if endTime.Weekday() != startTime.Weekday() {
			availEndArrayString = availability.NextSun.String
		}
	}

	// same day
	if endTime.Weekday() == startTime.Weekday() {
		availEndArrayString = availStartArrayString
	}

	return availStartArrayString, availEndArrayString
}

// Adds the default values for the roster
func (s *Server) insertDefaultRosterValues(roster *pb.Roster) error {
	if roster.Clients == nil {
		err := fillDefaultClients(roster, s.db)
		if err != nil {
			return err
		}
	}

	// fill up custom time
	for _, assignment := range roster.GuardAssigned {
		if assignment.CustomStartTime == nil {
			startTime, err := time.Parse(common.DATETIME_FORMAT, roster.StartTime)
			if err != nil {
				fmt.Println("insertDefaultRosterValues", err)
				return err
			}
			assignment.CustomStartTime = &timestamppb.Timestamp{Seconds: startTime.Unix()}
		}
		if assignment.CustomEndTime == nil {
			endTime, err := time.Parse(common.DATETIME_FORMAT, roster.EndTime)
			if err != nil {
				fmt.Println("insertDefaultRosterValues", err)
				return err
			}
			assignment.CustomEndTime = &timestamppb.Timestamp{Seconds: endTime.Unix()}
		}
	}

	// check if clients are nil
	if roster.Clients == nil || len(roster.Clients) == 0 {
		clients := make([]*pb.AIFSClientRoster, 0)
		clients = append(clients, &pb.AIFSClientRoster{
			Client:      &pb.Client{ClientId: 1},
			PatrolOrder: 1,
		})
		clients = append(clients, &pb.AIFSClientRoster{
			Client:      &pb.Client{ClientId: 2},
			PatrolOrder: 2,
		})
		clients = append(clients, &pb.AIFSClientRoster{
			Client:      &pb.Client{ClientId: 3},
			PatrolOrder: 3,
		})

		if roster.AifsId == 2 {
			clients[0].PatrolOrder = 2
			clients[1].PatrolOrder = 3
			clients[2].PatrolOrder = 1
		} else if roster.AifsId == 3 {
			clients[0].PatrolOrder = 3
			clients[1].PatrolOrder = 1
			clients[2].PatrolOrder = 2
		}
		roster.Clients = clients
	}

	return nil
}

func fillDefaultClients(roster *pb.Roster, db *sql.DB) error {
	clientQuery := &pb.ClientQuery{}
	aifsClientRosters := make([]*pb.AIFSClientRoster, 0)
	var clients []*pb.Client
	var err error

	switch roster.AifsId {
	case 1:
		db_pck.AddClientFilter(clientQuery, pb.ClientFilter_CLIENT_ID,
			pb.Filter_EQUAL, "1")
		db_pck.AddClientFilter(clientQuery, pb.ClientFilter_CLIENT_ID,
			pb.Filter_EQUAL, "2")
		clients, err = db_pck.GetClients(db, clientQuery)
		if err != nil {
			return err
		}
	case 2:
		db_pck.AddClientFilter(clientQuery, pb.ClientFilter_CLIENT_ID,
			pb.Filter_EQUAL, "3")
		db_pck.AddClientFilter(clientQuery, pb.ClientFilter_CLIENT_ID,
			pb.Filter_EQUAL, "4")
		clients, err = db_pck.GetClients(db, clientQuery)
		if err != nil {
			return err
		}
	default:
		db_pck.AddClientFilter(clientQuery, pb.ClientFilter_CLIENT_ID,
			pb.Filter_EQUAL, "5")
		db_pck.AddClientFilter(clientQuery, pb.ClientFilter_CLIENT_ID,
			pb.Filter_EQUAL, "6")
		clients, err = db_pck.GetClients(db, clientQuery)
		if err != nil {
			return err
		}
	}

	for _, client := range clients {
		aifsClientRosters = append(aifsClientRosters, &pb.AIFSClientRoster{Client: client})
	}

	roster.Clients = aifsClientRosters

	return nil
}

func validateStartTime(query *pb.RosterQuery) error {
	foundStartTime := true
	startTimeString := ""
	for _, query := range query.Filters {
		if query.Field == pb.RosterFilter_START_TIME {
			startTimeString = query.Comparisons.Value
			foundStartTime = true
			break
		}
	}

	if !foundStartTime {
		return status.Errorf(codes.NotFound, "start time must be set in roster query")
	}

	_, err := time.Parse(common.DATETIME_FORMAT, startTimeString)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, err.Error()+". Note: start time must be in the format:"+common.DATETIME_FORMAT)
	}

	return nil
}

func getDefaultEndTime(query *pb.AvailabilityQuery) error {
	if len(query.EndTime) == 0 {
		startTime, err := time.Parse(common.DATETIME_FORMAT, query.StartTime)
		if err != nil {
			fmt.Println("getDefaultEndTime ERROR:", err)
			return err
		}
		query.EndTime = startTime.Add(time.Hour * 12).Format(common.DATETIME_FORMAT)
	}
	return nil
}

func (s *Server) sendNewRostersToTele(rosterIds []int64) {
	query := &pb.RosterQuery{Limit: int64(len(rosterIds))}
	idStrings := make([]string, 0)

	for _, id := range rosterIds {
		idStrings = append(idStrings, strconv.Itoa(int(id)))
	}

	db_pck.AddRosterFilter(query, pb.RosterFilter_ROSTER_ID, pb.Filter_IN, strings.Join(idStrings, ","))
	// Only find rosters assignments that are still assigned
	db_pck.AddRosterFilter(query, pb.RosterFilter_IS_ASSIGNED, pb.Filter_EQUAL, "1")
	// Only find rosters where the confirmed is false
	db_pck.AddRosterFilter(query, pb.RosterFilter_GUARD_ASSIGNMENT_CONFIRMATION, pb.Filter_EQUAL, "0")

	rosters, err := db_pck.GetRosters(s.db, query)

	if err != nil {
		fmt.Println("sendNewRostersToTele ERROR:", err)
	}

	if len(rosters) == 0 {
		fmt.Println("sendNewRostersToTele: No Rosters found for ids", rosterIds)
	}

	tclient.InsertRoster(s.teleServerAddr, s.teleServerPort, rosters)
}
