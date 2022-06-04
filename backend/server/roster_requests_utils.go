// TODO: Add validation
package server

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"capstone.operations_ecosystem/backend/common"
	db_pck "capstone.operations_ecosystem/backend/database"
	rs "capstone.operations_ecosystem/backend/rating_system"

	pb "capstone.operations_ecosystem/backend/proto"
)

// Send back all the available users that are not yet assigned for the particular day
func (s *Server) GetAvailableIspecialistsWithScore(query *pb.AvailabilityQuery) ([]*pb.EmployeeEvaluation, error) {
	fmt.Println("GetAvailableUsersUtil")

	employeeEvals := make([]*pb.EmployeeEvaluation, 0)

	// First get all users who are not yet assigned
	unassignedUsers, err := s.getUnassignedIspecialists(query)
	if err != nil {
		return employeeEvals, err
	}

	// Get the scores for all these users
	err = getParallelEmployeeScores(unassignedUsers, &employeeEvals)
	if err != nil {
		return employeeEvals, err
	}

	// Find out which of these users are available
	err = s.updateAvailability(query, employeeEvals)
	if err != nil {
		return employeeEvals, err
	}

	return employeeEvals, nil
}

func (s *Server) getUnassignedIspecialists(query *pb.AvailabilityQuery) ([]*pb.User, error) {
	unassignedUsers := make([]*pb.User, 0)

	// Get a list of all the guard assignments for this day
	// who are still assigned
	rosterAssignmentQuery := &pb.RosterQuery{Limit: 20}
	// still assigned
	db_pck.AddRosterFilter(rosterAssignmentQuery, pb.RosterFilter_IS_ASSIGNED, pb.Filter_EQUAL, "1")
	// Start and end times of roster
	db_pck.AddRosterFilter(rosterAssignmentQuery, pb.RosterFilter_START_TIME, pb.Filter_EQUAL, query.StartTime.AsTime().Format(db_pck.DATETIME_FORMAT))
	db_pck.AddRosterFilter(rosterAssignmentQuery, pb.RosterFilter_END_TIME, pb.Filter_EQUAL, query.EndTime.AsTime().Format(db_pck.DATETIME_FORMAT))
	rosterAssignements, err := db_pck.GetRosterAssingments(s.db, rosterAssignmentQuery, -1)

	if err != nil {
		return unassignedUsers, err
	}

	// Create a string list of all the assigned users
	assignedUsers := make([]string, 0)
	for _, assignment := range rosterAssignements {
		assignedUsers = append(assignedUsers, strconv.Itoa(int(assignment.GuardAssigned.Employee.UserId)))
	}

	assignedUsersString := strings.Join(assignedUsers, ",")

	// Get all ISpecialists not in this list
	userQuery := &pb.UserQuery{Limit: db_pck.AVAILABILITY_DEFAULT_LIMIT}
	db_pck.AddUserFilter(userQuery, pb.UserFilter_USER_ID, pb.Filter_NOT_IN, assignedUsersString)
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

func (s *Server) updateAvailability(query *pb.AvailabilityQuery, employeeEvals []*pb.EmployeeEvaluation) error {
	// Add week and year to the availability filter
	year, week := query.StartTime.AsTime().ISOWeek()
	db_pck.AddAvailabilityFilter(query, pb.AvailabilityFilter_YEAR, pb.Filter_EQUAL, strconv.Itoa(year))
	db_pck.AddAvailabilityFilter(query, pb.AvailabilityFilter_WEEK, pb.Filter_EQUAL, strconv.Itoa(week))

	// add day NOT NULL to the availability filter
	db_pck.AddAvailabilityFilter(query, getDayAvailabilityQueryEnum(query.StartTime.AsTime()), pb.Filter_EQUAL, "")
	// if the start day is sat, the end day will be next sun
	if query.StartTime.AsTime().Day() == int(time.Saturday) {
		db_pck.AddAvailabilityFilter(query, pb.AvailabilityFilter_NEXT_SUN, pb.Filter_EQUAL, "")
	} else {
		db_pck.AddAvailabilityFilter(query, getDayAvailabilityQueryEnum(query.StartTime.AsTime()), pb.Filter_EQUAL, "")
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

	for _, eval := range employeeEvals {
		// binary search is employee is in the available list
		isAvail, _ := common.BinarySearch(availableUserIds, 0, len(availableUserIds)-1, int(eval.Employee.UserId))
		eval.IsAvailable = isAvail
	}

	return nil
}

func getDayAvailabilityQueryEnum(date time.Time) pb.AvailabilityFilter_Field {
	switch date.Day() {
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
	availStartArrayString := ""
	endArrayString := ""

	switch availQuery.StartTime.AsTime().Day() {
	case 0:
		availStartArrayString = availability.Sun.String
		endArrayString = availability.Mon.String
	case 1:
		availStartArrayString = availability.Mon.String
		endArrayString = availability.Tues.String

	case 2:
		availStartArrayString = availability.Tues.String
		endArrayString = availability.Wed.String

	case 3:
		availStartArrayString = availability.Wed.String
		endArrayString = availability.Thurs.String

	case 4:
		availStartArrayString = availability.Thurs.String
		endArrayString = availability.Fri.String

	case 5:
		availStartArrayString = availability.Fri.String
		endArrayString = availability.Sat.String

	default:
		availStartArrayString = availability.Sat.String
		endArrayString = availability.NextSun.String
	}

	fmt.Println("Availability string", availStartArrayString, endArrayString)
	//TODO actually check this
	return true
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
			assignment.CustomStartTime = roster.StartTime
		}
		if assignment.CustomEndTime == nil {
			assignment.CustomEndTime = roster.EndTime
		}
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
