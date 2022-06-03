// Use these functions to interact with the roster related database tables.
package database

import (
	"database/sql"
	"fmt"
	"strconv"
	"sync"

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
		err = status.Errorf(codes.AlreadyExists, "roster for AIFS "+strconv.Itoa(int(roster.AifsId))+" at "+roster.StartTime.String()+"already exists")
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
// TODO: return roster assignments that are not false confirmation
func GetRosters(db *sql.DB, query *pb.RosterQuery) ([]*pb.Roster, error) {
	fmt.Println("Getting Rosters...")
	rosters := make([]*pb.Roster, 0)

	// Join the roster and assignment tables in order to
	// easily filter conditions relating to both tables together.

	// Set default query limits if needed
	if query.Limit == 0 {
		query.Limit = DEFAULT_LIMIT
	}

	// We ignore any filters to do with the client aifs table first
	requestedLimit := query.Limit
	clientQueries := removeRosteringClientQueries(query)

	fields := ALL_COLS

	// tables are joined on the main roster id
	onCondition := formatFieldEqVal(ROSTER_DB_ID, ROSTER_ASGN_DB_RELATED_ROSTER, false)

	// Format filters
	// temporarily give the query limit the max
	query.Limit = MAX_LIMIT
	filters := getFormattedRosterFilters(query, ROSTER_DB_TABLE_NAME, true, true)

	rows, err := QueryLeftJoin(db, ROSTER_DB_TABLE_NAME, ROSTER_ASSIGNMENT_DB_TABLE_NAME, onCondition, fields, filters)

	if err != nil {
		return rosters, err
	}

	// convert query rows into rosters
	// give back the query the original limit
	query.Limit = requestedLimit
	err = convertDbRowsToFullRoster(db, &rosters, rows, query, clientQueries)

	return rosters, err
}

// Get all the roster recipient rows in a table that meets specifications.
// Returns an array of roster recipients and any errors.
func GetRosterAssingments(db *sql.DB, query *pb.RosterQuery, mainRosterID int64) ([]*pb.RosterAssignement, error) {
	fmt.Println("Getting Rosters Assignments...")
	rosterRecipients := make([]*pb.RosterAssignement, 0)

	fields := ALL_COLS

	// Format filters
	// Get for a specific main roster
	addRosterFilter(query, pb.RosterFilter_ROSTER_ID, pb.Filter_EQUAL, strconv.Itoa(int(mainRosterID)))
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
	addRosterFilter(query, pb.RosterFilter_ROSTER_ID, pb.Filter_EQUAL, strconv.Itoa(int(mainRosterID)))
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
		)

		if err != nil {
			fmt.Println("GetRosterAIFSClient ERROR::", err)
			break
		}

		aifsClient.Client, err = idClientByClientId(db, clientId)
		if err != nil {
			fmt.Println("GetRosterAIFSClient:", err)
			continue
		}

		aifsClients = append(aifsClients, aifsClient)
	}

	return aifsClients, err
}

// Update a specific roster in the table
// Only fields that have been filled in the roster object will be updated.
// Note that this update does not update the roster's guards's inner status
// such as the acknowledgement or attended status but only if the guard
// is part of the roster. Same for the clients
func UpdateRoster(db *sql.DB, roster *pb.Roster, dbLock *sync.Mutex) (int64, error) {
	// Update the main roster first
	newFields := getFilledRosterFields(roster)
	query := &pb.RosterQuery{}
	addRosterFilter(query, pb.RosterFilter_ROSTER_ID, pb.Filter_EQUAL, strconv.Itoa(int(roster.RosteringId)))
	filters := getFormattedRosterFilters(query, ROSTER_DB_TABLE_NAME, false, false)

	var err error
	rowsAffected := int64(0)

	if len(newFields) > 0 {
		rowsAffected, err = Update(db, ROSTER_DB_TABLE_NAME, newFields, filters)
		if err != nil {
			fmt.Println("UpdateRoster ERROR::", err)
			return rowsAffected, err
		}
	}

	// Update assignments if necessary
	if roster.GuardAssigned != nil {
		err = updateAssignmentsOfRoster(db, roster, query, dbLock)
	}

	// Update clients if necessary
	if roster.Clients != nil {
		err = updateClientsOfRoster(db, roster, query, dbLock)
	}

	return rowsAffected, err
}

// Update a specific recipient row in the table
// This function assumes that the roster recipient id is correct.
// Returns the number of rows affected and any errors.
// In this case, number of rows affected is either 0 or 1.
func UpdateRosterRecipients(db *sql.DB, rosterAssignment *pb.RosterAssignement) (int64, error) {
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
