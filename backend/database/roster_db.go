// Use these functions to interact with the roster related database tables.
package database

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"capstone.operations_ecosystem/backend/common"
	pb "capstone.operations_ecosystem/backend/proto"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Insert a new roster into the database table.
// Corresponding roster details are added to their
// respective table as well.
// Returns the primary key of the main roster and errors if any
func InsertRoster(db *sql.DB, roster *pb.Roster, dbLock *sync.Mutex) (int64, error) {

	fmt.Println("Checking if Roster for this AIFS and time already exists...", "AIFS:", roster.AifsId, roster.StartTime)

	// Do not add the roster if it already exists
	existingPk, err := checkRosterExists(db, roster)
	if err != nil {
		// Do not add the rest if the main roster fails
		return -1, err
	}
	if existingPk != -1 {
		err = status.Errorf(codes.AlreadyExists, "roster for AIFS "+strconv.Itoa(int(roster.AifsId))+" at "+roster.StartTime+"already exists")
		return -1, err
	}

	fmt.Println("Inserting Roster", roster.RosteringId, "for AIFS", roster.AifsId)

	// Create and insert main roster first and get it's pk
	rosterTbFields := getRosterTableFields()
	rosterValues := orderRosterFields(roster)

	rosterPk, err := Insert(db, ROSTER_DB_TABLE_NAME, rosterTbFields, rosterValues, dbLock)

	if err != nil {
		// Do not add the rest if the main roster fails
		return rosterPk, err
	}

	// Create roster guard assignments rows for corresponding roster
	for _, guard := range roster.GuardAssigned {
		_, err = InsertRosterASGN(db, guard, rosterPk, dbLock)

		if err != nil {
			// Delete the roster that was just inserted
			roster.RosteringId = rosterPk
			DeleteRoster(db, roster)
			break
		}
	}

	// Create roster guard assignments rows for corresponding roster
	for _, aifsClient := range roster.Clients {
		_, err = InsertAIFSClientRoster(db, aifsClient, rosterPk, dbLock)

		if err != nil {
			// Delete the roster that was just inserted, it should cascade
			// and delete all the other roster details
			roster.RosteringId = rosterPk
			DeleteRoster(db, roster)
			break
		}
	}
	return rosterPk, err
}

// Inserts a new assignment to the database and connects it to the appropriate
// main roster.
// Assumes the assignment guard has the correct id that corresponds to its DB row.
// Returns the primary key of the recipient row and any errors.
func InsertRosterASGN(db *sql.DB, assignment *pb.RosterAssignement, mainRosterID int64, dbLock *sync.Mutex) (int64, error) {
	// get fields and values for this particular recipient
	fields := getRosterAsgnTableFields()
	values := orderRosterAsgnFields(assignment, mainRosterID)

	// Add recipient to DB
	pk, err := Insert(db, ROSTER_ASSIGNMENT_DB_TABLE_NAME, fields, values, dbLock)

	return pk, err
}

// Inserts a new aifs client pair to the database and connects it to the appropriate
// main roster.
// Assumes the client has the correct id that corresponds to its DB row.
// Returns the primary key of the recipient row and any errors.
func InsertAIFSClientRoster(db *sql.DB, aifsClient *pb.AIFSClientRoster, mainRosterID int64, dbLock *sync.Mutex) (int64, error) {
	// get fields and values for this particular recipient
	fields := getAIFSClientTableFields()
	values := orderAIFSClientFields(aifsClient, mainRosterID)

	// Add aifs client to DB
	pk, err := Insert(db, ROSTER_AIFS_CLIENT_DB_TABLE_NAME, fields, values, dbLock)

	return pk, err
}

// Get all the roster rows in a table that meets specifications.
// Returns an array of rosters and any errors.
func GetRosters(db *sql.DB, query *pb.RosterQuery) ([]*pb.Roster, error) {
	fmt.Println("Getting Rosters...")
	rosters := make([]*pb.Roster, 0)

	// Join the roster and assignment tables and aifs client tables
	// in order to easily filter conditions relating to all tables

	// Set default query limits if needed
	if query.Limit == 0 {
		query.Limit = DEFAULT_LIMIT
	}

	// We ignore any filters to do with the client aifs table first
	requestedLimit := query.Limit

	fields := ALL_COLS

	// tables are joined on the main roster id
	fistOnCondition := formatFieldEqVal(ROSTER_DB_ID, ROSTER_ASSIGNMENT_DB_TABLE_NAME+"."+ROSTER_ASGN_DB_RELATED_ROSTER, false)
	secondOnCondition := formatFieldEqVal(ROSTER_DB_ID, ROSTER_AIFS_CLIENT_DB_TABLE_NAME+"."+AIFS_CLIENT_DB_RELATED_ROSTER, false)

	// Format filters
	// temporarily give the query limit the max
	query.Limit = MAX_LIMIT
	filters := getFormattedRosterFilters(query, ROSTER_DB_TABLE_NAME, true, true)

	rows, err := QueryThreeTablesLeftJoin(db, ROSTER_DB_TABLE_NAME,
		ROSTER_ASSIGNMENT_DB_TABLE_NAME, ROSTER_AIFS_CLIENT_DB_TABLE_NAME,
		fistOnCondition, secondOnCondition, fields, filters)

	if err != nil {
		return rosters, err
	}

	// convert query rows into rosters
	// give back the query the original limit
	query.Limit = requestedLimit
	err = convertDbRowsToFullRoster(db, &rosters, rows, query)

	// Set status of rosters
	setRosterStatus(rosters)

	return rosters, err
}

func GetDefaultRosters(db *sql.DB, query *pb.RosterQuery) ([]*pb.Roster, error) {
	fmt.Println("Getting Default Rosters...")

	rosters := make([]*pb.Roster, 0)

	startTimeString := ""

	for _, query := range query.Filters {
		if query.Field == pb.RosterFilter_START_TIME {
			startTimeString = query.Comparisons.Value
			break
		}
	}

	startTime, err := time.Parse(common.DATETIME_FORMAT, startTimeString)
	if err != nil {
		fmt.Println("GetDefaultRosters ERROR:", err)
		return rosters, err
	}

	dayOfWeek := startTime.Weekday()

	defaultAifsRosterIds, err := GetDefaultRosterDetails(db, &pb.RosterQuery{}, int(dayOfWeek))
	if err != nil {
		return rosters, err
	}
	idStringArray := make([]string, 0)
	for _, id := range defaultAifsRosterIds {
		idStringArray = append(idStringArray, strconv.Itoa(id))
	}

	query = &pb.RosterQuery{}
	AddRosterFilter(query, pb.RosterFilter_ROSTER_ID, pb.Filter_IN, strings.Join(idStringArray, ","))

	rosters, err = GetRosters(db, query)

	// Set start and end time correctly, also set as default
	for _, roster := range rosters {
		roster.StartTime = startTimeString
		// Shifts are 12 hours long
		roster.EndTime = startTime.Add(time.Hour * 12).Format(common.DATETIME_FORMAT)
		roster.Status = pb.Roster_IS_DEFAULT
	}

	return rosters, err
}

// Get all the roster recipient rows in a table that meets specifications.
// Returns an array of roster recipients and any errors.
func GetRosterAssingments(db *sql.DB, query *pb.RosterQuery, mainRosterID int64) ([]*pb.RosterAssignement, error) {
	fmt.Println("Getting Rosters Assignments...")
	rosterRecipients := make([]*pb.RosterAssignement, 0)

	fields := ALL_COLS

	// Format filters
	// Get for a specific main roster if needed
	if mainRosterID != -1 {
		AddRosterFilter(query, pb.RosterFilter_ROSTER_ID, pb.Filter_EQUAL, strconv.Itoa(int(mainRosterID)))
	}

	filters := getFormattedRosterFilters(query, ROSTER_ASSIGNMENT_DB_TABLE_NAME, true, true)

	rows, err := Query(db, ROSTER_ASSIGNMENT_DB_TABLE_NAME, fields, filters)

	if err != nil {
		return rosterRecipients, err
	}

	// convert query rows into rosters assignments
	for rows.Next() {
		assignment := &pb.RosterAssignement{}
		employeeEval := &pb.EmployeeEvaluation{}
		assignment.GuardAssigned = employeeEval

		// fields that cannot be auto converted
		guardId := -1
		// related roster is not necessary, but for simplicity
		// and for possible future use, we get it back in the query.
		relatedRoster := ""

		// confirmation is nullable
		var confirmation sql.NullBool

		// Datetimes
		startTimeString := ""
		endTimeString := ""
		var attendanceTimeString sql.NullString

		// cast each row to a roster
		err = rows.Scan(
			&assignment.RosterAssignmentId,
			&relatedRoster,
			&guardId,
			&startTimeString,
			&endTimeString,
			&confirmation,
			&assignment.Attended,
			&attendanceTimeString,
			&assignment.IsAssigned,
			&assignment.Rejected,
		)

		if err != nil {
			fmt.Println("GetRosterAssingments ERROR::", err)
			break
		}

		// Add Datetimes
		assignment.CustomStartTime, err = DBDatetimeToPB(startTimeString)
		if err != nil {
			fmt.Println("GetRosterAssingments:", err.Error())
			continue
		}
		assignment.CustomEndTime, err = DBDatetimeToPB(endTimeString)
		if err != nil {
			fmt.Println("GetRosterAssingments:", err.Error())
			continue
		}

		if attendanceTimeString.Valid {
			assignment.AttendanceTime, err = DBDatetimeToPB(attendanceTimeString.String)
			if err != nil {
				fmt.Println("GetRosterAssingments:", err.Error())
				continue
			}
		}

		if confirmation.Valid {
			assignment.Confirmed = confirmation.Bool
			if err != nil {
				fmt.Println("GetRosterAssingments:", err.Error())
				continue
			}
		}
		// TODO think about whether I can store the users in cache rather than
		// get the same few users over and over
		assignment.GuardAssigned.Employee, err = idUserByUserId(db, guardId)
		if err != nil {
			fmt.Println("GetRosterAssingments:", err)
			continue
		}

		rosterRecipients = append(rosterRecipients, assignment)
	}

	return rosterRecipients, err
}

// Get all the roster recipient rows in a table that meets specifications.
// Returns an array of roster recipients and any errors.
func GetRosterAIFSClient(db *sql.DB, query *pb.RosterQuery, mainRosterID int64) ([]*pb.AIFSClientRoster, error) {
	fmt.Println("Getting Rosters AIFS Clients...")
	aifsClients := make([]*pb.AIFSClientRoster, 0)

	fields := ALL_COLS

	// Format filters
	// Get for a specific main roster
	AddRosterFilter(query, pb.RosterFilter_ROSTER_ID, pb.Filter_EQUAL, strconv.Itoa(int(mainRosterID)))
	filters := getFormattedRosterFilters(query, ROSTER_AIFS_CLIENT_DB_TABLE_NAME, true, true)

	rows, err := Query(db, ROSTER_AIFS_CLIENT_DB_TABLE_NAME, fields, filters)

	if err != nil {
		return aifsClients, err
	}

	// convert query rows into roster aifs clients
	for rows.Next() {
		aifsClient := &pb.AIFSClientRoster{}
		// fields that cannot be auto converted
		clientId := -1
		// related roster is not necessary, but for simplicity
		// and for possible future use, we get it back in the query.
		relatedRoster := ""

		// cast each row to a roster
		err = rows.Scan(
			&aifsClient.AifsClientRosterId,
			&relatedRoster,
			&clientId,
			&aifsClient.PatrolOrder,
		)

		if err != nil {
			fmt.Println("GetRosterAIFSClient ERROR::", err)
			break
		}

		aifsClient.Client, err = IdClientByClientId(db, clientId)
		if err != nil {
			fmt.Println("GetRosterAIFSClient:", err)
			continue
		}

		aifsClients = append(aifsClients, aifsClient)
	}

	return aifsClients, err
}

// Get the default roster ids for the default 3 aifs
func GetDefaultRosterDetails(db *sql.DB, query *pb.RosterQuery, dayOfWeek int) ([]int, error) {
	fmt.Println("GetDefaultRosterDetails...")
	scheduleIds := make([]int, 3)

	fields := ALL_COLS

	// Format filters
	// Get for a specific day of week
	query.Limit = 1
	AddRosterFilter(query, pb.RosterFilter_DEFAULT_ROSTERING_DAY_OF_WEEK, pb.Filter_EQUAL, strconv.Itoa(dayOfWeek))
	filters := getFormattedRosterFilters(query, ROSTER_AIFS_CLIENT_DB_TABLE_NAME, true, true)

	rows, err := Query(db, ROSTER_DEFAULT_DB_TABLE_NAME, fields, filters)

	if err != nil {
		fmt.Println("GetDefaultRosterDetails ERROR:", err)
		return scheduleIds, err
	}

	// convert query rows into roster aifs clients
	for rows.Next() {
		// unnecessary fields
		defaultRosteringId := -1

		// cast each row to a roster
		err = rows.Scan(
			&defaultRosteringId,
			&dayOfWeek,
			&scheduleIds[0],
			&scheduleIds[1],
			&scheduleIds[2],
		)

		if err != nil {
			fmt.Println("GetDefaultRosterDetails ERROR::", err)
			break
		}
	}

	return scheduleIds, nil
}

// Update a specific roster in the table
// Only fields that have been filled in the roster object will be updated.
// Returns the number of main roster rows updated as well as the a list of
// all the primary keys of newly inserted roster assignments into the db.
// Note that this update does not update the roster's guards's inner status
// such as the acknowledgement or attended status but only if the guard
// is part of the roster. Same for the clients
func UpdateRoster(db *sql.DB, roster *pb.Roster, dbLock *sync.Mutex) (int64, []int64, error) {
	newRosterAssignmentsPk := make([]int64, 0)

	// Update the main roster first
	newFields := getFilledRosterFields(roster)

	filters := getRosterIdFormattedFilter(
		int(roster.RosteringId),
		ROSTER_DB_TABLE_NAME, true,
	)

	var err error
	rowsAffected := int64(0)

	if len(newFields) > 0 {
		rowsAffected, err = Update(db, ROSTER_DB_TABLE_NAME, newFields, filters)
		if err != nil {
			fmt.Println("UpdateRoster ERROR::", err)
			return rowsAffected, newRosterAssignmentsPk, err
		}
	}

	// Update assignments if necessary
	if roster.GuardAssigned != nil {
		newRosterAssignmentsPk, err = updateAssignmentsOfRoster(db, roster, dbLock)
	}

	// Update clients if necessary
	if roster.Clients != nil {
		err = updateClientsOfRoster(db, roster, dbLock)
	}

	return rowsAffected, newRosterAssignmentsPk, err
}

// Update a specific recipient row in the table
// This function assumes that the roster recipient id is correct.
// Returns the number of rows affected and any errors.
// In this case, number of rows affected is either 0 or 1.
func UpdateRosterAssignments(db *sql.DB, rosterAssignment *pb.RosterAssignement) (int64, error) {
	newFields := getFilledRosterASGNFields(rosterAssignment)
	filters := getRosterIdFormattedFilter(
		int(rosterAssignment.RosterAssignmentId),
		ROSTER_ASSIGNMENT_DB_TABLE_NAME, false,
	)

	rowsAffected, err := Update(db, ROSTER_ASSIGNMENT_DB_TABLE_NAME, newFields, filters)

	if err != nil {
		fmt.Println("UpdateRosterRecipients ERROR::", err)
		return rowsAffected, err
	}

	return rowsAffected, err
}

// Update a specific recipient row in the table
// This function assumes that the roster recipient id is correct.
// Returns the number of rows affected and any errors.
// In this case, number of rows affected is either 0 or 1.
func UpdateRosterAIFSClient(db *sql.DB, aifsClient *pb.AIFSClientRoster) (int64, error) {
	newFields := getFilledAIFSClientFields(aifsClient)
	filters := getRosterIdFormattedFilter(
		int(aifsClient.AifsClientRosterId),
		ROSTER_AIFS_CLIENT_DB_TABLE_NAME, false,
	)

	rowsAffected, err := Update(db, ROSTER_AIFS_CLIENT_DB_TABLE_NAME, newFields, filters)

	if err != nil {
		fmt.Println("UpdateRosterRecipients ERROR::", err)
		return rowsAffected, err
	}

	return rowsAffected, err
}

// Delete a particular roster in the database together with
// any corresponding details.
// Details are deleted based on the cascading rule.
func DeleteRoster(db *sql.DB, roster *pb.Roster) (int64, error) {
	filters := getRosterIdFormattedFilter(
		int(roster.RosteringId),
		ROSTER_DB_TABLE_NAME, true,
	)

	rowsAffected, err := Delete(db, ROSTER_DB_TABLE_NAME, filters)
	return rowsAffected, err
}

// Delete a particular roster detail
func DeleteRosterAssignment(db *sql.DB, rosterAssignment *pb.RosterAssignement) (int64, error) {
	filters := getRosterIdFormattedFilter(
		int(rosterAssignment.RosterAssignmentId),
		ROSTER_ASSIGNMENT_DB_TABLE_NAME, false,
	)
	rowsAffected, err := Delete(db, ROSTER_ASSIGNMENT_DB_TABLE_NAME, filters)
	return rowsAffected, err
}

// Delete a particular roster aifs client
func DeleteRosterAIFSClient(db *sql.DB, aifsClient *pb.AIFSClientRoster) (int64, error) {
	filters := getRosterIdFormattedFilter(
		int(aifsClient.AifsClientRosterId),
		ROSTER_AIFS_CLIENT_DB_TABLE_NAME, false,
	)

	fmt.Println("fsdfds filters", filters)
	rowsAffected, err := Delete(db, ROSTER_AIFS_CLIENT_DB_TABLE_NAME, filters)
	return rowsAffected, err
}
