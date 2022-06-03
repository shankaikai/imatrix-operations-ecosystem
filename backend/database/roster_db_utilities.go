// Utility functions for database operations related to rostering.
package database

import (
	"database/sql"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"

	"capstone.operations_ecosystem/backend/common"
	pb "capstone.operations_ecosystem/backend/proto"
	_ "github.com/go-sql-driver/mysql"
)

const (
	ROSTER_DB_TABLE_NAME             = "schedule"
	ROSTER_ASSIGNMENT_DB_TABLE_NAME  = "schedule_detail"
	ROSTER_AIFS_CLIENT_DB_TABLE_NAME = "aifs_client_schedule"

	// Roster table fields
	ROSTER_DB_ID         = "schedule_id"
	ROSTER_DB_AIFS_ID    = "aifs_id"
	ROSTER_DB_START_TIME = "start_time"
	ROSTER_DB_END_TIME   = "end_time"

	// Roster guard assignment table fields
	ROSTER_ASGN_DB_ID             = "schedule_detail_id"
	ROSTER_ASGN_DB_RELATED_ROSTER = "schedule"
	ROSTER_ASGN_DB_GUARD_ASSIGNED = "guard_assigned"
	ROSTER_ASGN_DB_START_TIME     = "custom_start_time"
	ROSTER_ASGN_DB_END_TIME       = "custom_end_time"
	ROSTER_ASGN_DB_CONFIRMATION   = "confirmation"
	ROSTER_ASGN_DB_ATTENDED       = "attended"
	ROSTER_ASGN_DB_ATTENDED_TIME  = "attendance_time"

	// AIFS Client Schedule table fields
	AIFS_CLIENT_DB_ID             = "aifs_client_schedule_id"
	AIFS_CLIENT_DB_RELATED_ROSTER = "schedule"
	AIFS_CLIENT_DB_RELATED_CLIENT = "related_client"
)

// UTILITIES

// Returns the fields of the main roster table
// in a specific order.
// Note that IDs are auto incremented and should
// not be modified manually. Ommits ID in resulting string.
func getRosterTableFields() string {
	rosterTableFields := []string{
		ROSTER_DB_AIFS_ID,
		ROSTER_DB_START_TIME,
		ROSTER_DB_END_TIME,
	}

	return strings.Join(rosterTableFields, ",")
}

// This function is highly dependent on the
// order given in getRosterTableFields.
// Returns the values of the roster fields in the
// order that is specified in getRosterTableFields
func orderRosterFields(roster *pb.Roster) string {
	output := ""

	output += "'" + strconv.Itoa(int(roster.AifsId)) + "'" + ", "
	output += "'" + roster.StartTime.AsTime().Format(DATETIME_FORMAT) + "'" + ", "
	output += "'" + roster.EndTime.AsTime().Format(DATETIME_FORMAT) + "'"

	return output
}

// Returns the fields of the roster assignment table
// in a specific order.
// Note that IDs are auto incremented and should
// not be modified manually. Ommits ID in resulting string.
// All nullable columns that we do not expect to be used upon
// first insert shall be omitted.
func getRosterAsgnTableFields() string {
	rosterRecTableFields := []string{
		ROSTER_ASGN_DB_RELATED_ROSTER,
		ROSTER_ASGN_DB_GUARD_ASSIGNED,
		ROSTER_ASGN_DB_START_TIME,
		ROSTER_ASGN_DB_END_TIME,
		ROSTER_ASGN_DB_CONFIRMATION,
		ROSTER_ASGN_DB_ATTENDED,
	}

	return strings.Join(rosterRecTableFields, ",")
}

// This function is highly dependent on the
// order given in getRosterRecTableFields.
// Returns the values of the roster fields in the
// order that is specified in getRosterRecTableFields
func orderRosterAsgnFields(rosterAssignment *pb.RosterAssignement, relatedRosterId int64) string {
	output := ""

	output += strconv.Itoa(int(relatedRosterId)) + ","
	output += strconv.Itoa(int(rosterAssignment.GuardAssigned.Employee.UserId)) + ","
	output += "'" + rosterAssignment.CustomStartTime.AsTime().Format(DATETIME_FORMAT) + "'" + ", "
	output += "'" + rosterAssignment.CustomEndTime.AsTime().Format(DATETIME_FORMAT) + "'" + ", "

	// confirmation and attended are false by default.
	output += "0, 0"

	return output
}

// Returns the fields of the roster rosterAssignments table
// in a specific order.
// Note that IDs are auto incremented and should
// not be modified manually. Ommits ID in resulting string.
// All nullable columns that we do not expect to be used upon
// first insert shall be omitted.
func getAIFSClientTableFields() string {
	aifsClientTableFields := []string{
		AIFS_CLIENT_DB_RELATED_ROSTER,
		AIFS_CLIENT_DB_RELATED_CLIENT,
	}

	return strings.Join(aifsClientTableFields, ",")
}

// This function is highly dependent on the
// order given in getRosterRecTableFields.
// Returns the values of the roster fields in the
// order that is specified in getRosterRecTableFields
func orderAIFSClientFields(aifsClient *pb.AIFSClientRoster, relatedRosterId int64) string {
	output := ""

	output += strconv.Itoa(int(relatedRosterId)) + ","
	output += strconv.Itoa(int(aifsClient.Client.ClientId))

	return output
}

// This function is used in anticipation of an update query.
// Updates use the format "SET field1=val1, field2=val2".
// The return value of this function is for example: "field1=val1, field2=val2"
// ID not included in this because pk should not be manually changed.
// Note this function highly depends on the protocol buffer message definition
// Returns the string fields and values of the filled roster fields
func getFilledRosterFields(roster *pb.Roster) string {
	rosterTableFields := []string{}

	if roster.AifsId > 0 {
		rosterTableFields = append(rosterTableFields, formatFieldEqVal(ROSTER_DB_AIFS_ID, strconv.Itoa(int(roster.AifsId)), true))
	}
	if roster.StartTime != nil {
		rosterTableFields = append(rosterTableFields, formatFieldEqVal(ROSTER_DB_START_TIME, roster.StartTime.AsTime().Format(DATETIME_FORMAT), true))
	}
	if roster.EndTime != nil {
		rosterTableFields = append(rosterTableFields, formatFieldEqVal(ROSTER_DB_END_TIME, roster.EndTime.AsTime().Format(DATETIME_FORMAT), true))
	}

	return strings.Join(rosterTableFields, ",")
}

// This function is used in anticipation of an update query.
// Updates use the format "SET field1=val1, field2=val2".
// The return value of this function is for example: "field1=val1, field2=val2"
// ID not included in this because pk should not be manually changed.
// Note this function highly depends on the protocol buffer message definition
// Returns the string fields and values of the filled roster recipinets fields
func getFilledRosterASGNFields(rosterAssignment *pb.RosterAssignement) string {
	rosterTableFields := []string{}

	if rosterAssignment.GuardAssigned != nil {
		if rosterAssignment.GuardAssigned.Employee != nil {
			rosterTableFields = append(rosterTableFields, formatFieldEqVal(ROSTER_ASGN_DB_GUARD_ASSIGNED, strconv.Itoa(int(rosterAssignment.GuardAssigned.Employee.UserId)), true))
		}
	}
	if rosterAssignment.CustomStartTime != nil {
		rosterTableFields = append(rosterTableFields, formatFieldEqVal(ROSTER_ASGN_DB_START_TIME, rosterAssignment.CustomStartTime.AsTime().Format(DATETIME_FORMAT), true))
	}
	if rosterAssignment.CustomEndTime != nil {
		rosterTableFields = append(rosterTableFields, formatFieldEqVal(ROSTER_ASGN_DB_END_TIME, rosterAssignment.CustomEndTime.AsTime().Format(DATETIME_FORMAT), true))
	}

	rosterTableFields = append(rosterTableFields, formatFieldEqVal(ROSTER_ASGN_DB_CONFIRMATION, strconv.FormatBool(rosterAssignment.Confirmed), false))
	rosterTableFields = append(rosterTableFields, formatFieldEqVal(ROSTER_ASGN_DB_ATTENDED, strconv.FormatBool(rosterAssignment.Attended), false))

	if rosterAssignment.AttendanceTime != nil {
		rosterTableFields = append(rosterTableFields, formatFieldEqVal(ROSTER_ASGN_DB_ATTENDED_TIME, rosterAssignment.AttendanceTime.AsTime().Format(DATETIME_FORMAT), true))
	}

	return strings.Join(rosterTableFields, ",")
}

// This function is used in anticipation of an update query.
// Updates use the format "SET field1=val1, field2=val2".
// The return value of this function is for example: "field1=val1, field2=val2"
// ID not included in this because pk should not be manually changed.
// Note this function highly depends on the protocol buffer message definition
// Returns the string fields and values of the filled roster recipinets fields
func getFilledAIFSClientFields(aifsClient *pb.AIFSClientRoster) string {
	rosterTableFields := []string{}

	if aifsClient.Client != nil {
		rosterTableFields = append(rosterTableFields, formatFieldEqVal(AIFS_CLIENT_DB_RELATED_CLIENT, strconv.Itoa(int(aifsClient.Client.ClientId)), true))
	}

	return strings.Join(rosterTableFields, ",")
}

// Helper function to add a new filter to the list of existing
// filters in a roster query struct.
// Modifies the roster query parameter directly.
func addRosterFilter(query *pb.RosterQuery, field pb.RosterFilter_Field,
	comparison pb.Filter_Comparisons,
	value string) {
	if query.Filters == nil {
		query.Filters = make([]*pb.RosterFilter, 0)
	}
	filter := &pb.Filter{Comparison: comparison, Value: value}
	query.Filters = append(query.Filters, &pb.RosterFilter{Field: field, Comparisons: filter})
}

// Converts the filters in the roster array into a formatted where clause
// that can be parsed into MySQL. If a limit is needed, the LIMIT filter is
// added to the end of the string.
// For example returns: "WHERE id=22 AND num <2 LIMIT 5"
// Returns the formatted SQL filter string.
func getFormattedRosterFilters(query *pb.RosterQuery, table string, needLimit bool, needOrder bool) string {
	output := ""

	// Get all filters
	whereFilters := make([]string, 0)
	groupBy := make([]string, 0)
	haveFilters := make([]string, 0)

	for _, filter := range query.Filters {
		hasQuotes := true
		if filter.Comparisons.Comparison == pb.Filter_CONTAINS {
			filter.Comparisons.Value = FormatLikeQueryValue(filter.Comparisons.Value)
		} else if filter.Comparisons.Comparison == pb.Filter_IN {
			filter.Comparisons.Value = FormatInQueryValue(filter.Comparisons.Value)
			hasQuotes = false
		}
		switch filter.Field {
		case pb.RosterFilter_ROSTER_ID, pb.RosterFilter_AIFS_ID, pb.RosterFilter_GUARD_ASSIGNED_ID,
			pb.RosterFilter_CLIENT_ID, pb.RosterFilter_ROSTER_ASSIGNMENT_ID, pb.RosterFilter_START_TIME,
			pb.RosterFilter_END_TIME, pb.RosterFilter_ROSTER_AIFS_CLIENT_ID:
			if hasQuotes {
				whereFilters = append(
					whereFilters, fmt.Sprintf("%s %s '%s'", rosterFilterToDBCol(filter.Field, table),
						GetFilterComparisonSign(filter.Comparisons.Comparison), filter.Comparisons.Value),
				)
			} else {
				whereFilters = append(
					whereFilters, fmt.Sprintf("%s %s %s", rosterFilterToDBCol(filter.Field, table),
						GetFilterComparisonSign(filter.Comparisons.Comparison), filter.Comparisons.Value),
				)
			}
		case pb.RosterFilter_GUARD_ASSIGNMENT_CONFIRMATION, pb.RosterFilter_GUARD_ASSIGNMENT_ATTENDED:
			whereFilters = append(
				whereFilters, fmt.Sprintf("%s %s %s", rosterFilterToDBCol(filter.Field, table),
					GetFilterComparisonSign(filter.Comparisons.Comparison), filter.Comparisons.Value),
			)
		}
	}

	if len(whereFilters) > 0 {
		output += WHERE_KEYWORD + " "
	}

	output += strings.Join(whereFilters, " AND ")

	if len(groupBy) > 0 {
		output += " GROUP BY " + strings.Join(groupBy, ",")
	}

	if len(haveFilters) > 0 {
		output += " HAVING " + strings.Join(haveFilters, " AND ")
	}

	// Add order
	if needOrder {
		if query.OrderBy != nil {
			output += fmt.Sprintf(" %s %s %s", ORDER_BY_KEYWORD, rosterFilterToDBCol(query.OrderBy.Field, table), orderByProtoToDB(query.OrderBy.OrderBy))
		} else if table == ROSTER_DB_TABLE_NAME {
			// By default we order rosters by the aifs id
			output += fmt.Sprintf(" %s %s %s", ORDER_BY_KEYWORD, rosterFilterToDBCol(pb.RosterFilter_AIFS_ID, table), ASC_KEYWORD)
		}
	}

	// Add limits
	if needLimit {
		if query.Limit == 0 {
			query.Limit = DEFAULT_LIMIT
		}
		output += fmt.Sprintf(" %s %d, %d", LIMIT_KEYWORD, query.Skip, query.Limit)
	}

	return output
}

// This function converts the returned DB rows into Roster objects and
// their corresponding roster recipients.
// These rows come from the join query of both the roster and roster
// recipients table.
// Modifies the roster array in place.
func convertDbRowsToFullRoster(db *sql.DB, rosters *[]*pb.Roster, rows *sql.Rows, query *pb.RosterQuery) error {
	rosterMap := make(map[int64]*pb.Roster)

	for rows.Next() {
		roster := &pb.Roster{
			GuardAssigned: make([]*pb.RosterAssignement, 0),
			Clients:       make([]*pb.AIFSClientRoster, 0),
		}
		rosterAssignment := &pb.RosterAssignement{}
		employeeEval := &pb.EmployeeEvaluation{}
		rosterAssignment.GuardAssigned = employeeEval

		aifsClient := &pb.AIFSClientRoster{}

		assignedUserId := -1
		clientId := -1

		// Redundant Strings
		relatedRoster := ""

		// confirmation can be null
		var confirmation sql.NullBool

		// Datetimes
		startTimeString := ""
		endTimeString := ""

		customStartTimeString := ""
		customEndTimeString := ""
		var attendanceTimeString sql.NullString

		// cast each row to a roster
		err := rows.Scan(
			// Main Roster
			&roster.RosteringId,
			&roster.AifsId,
			&startTimeString,
			&endTimeString,

			// Roster Details
			&rosterAssignment.RosterAssignmentId,
			&relatedRoster,
			&assignedUserId,
			&customStartTimeString,
			&customEndTimeString,
			&rosterAssignment.Confirmed,
			&rosterAssignment.Attended,
			&attendanceTimeString,

			// Client Details
			&aifsClient.AifsClientRosterId,
			&clientId,
		)

		if err != nil {
			fmt.Println("convertDbRowsToFullRoster ERROR:", err)
			continue
		}

		// Check if there was already a roster found with this id
		if existingRoster, ok := rosterMap[roster.RosteringId]; ok {
			roster = existingRoster
		} else {
			// Return only the necessary number of rosters.
			// If the number of rosters have reached the limit,
			// do not add new rosters.
			if int64(len(rosterMap)) >= query.Limit+query.Skip {
				continue
			}
			roster.StartTime, err = DBDatetimeToPB(startTimeString)
			if err != nil {
				fmt.Println("GetRosters:", err.Error())
				continue
			}
			roster.EndTime, err = DBDatetimeToPB(endTimeString)

			if err != nil {
				fmt.Println("GetRosters:", err.Error())
				continue
			}

			rosterMap[roster.RosteringId] = roster
		}

		// Check if there is already a guard assignment for this roster
		guardAssignmentExists := false
		for _, existingAssigned := range roster.GuardAssigned {
			if existingAssigned.RosterAssignmentId == rosterAssignment.RosterAssignmentId {
				guardAssignmentExists = true
				break
			}
		}

		if !guardAssignmentExists {
			convertFromDbRosterAssignment(db, rosterAssignment, assignedUserId,
				customStartTimeString, customEndTimeString,
				attendanceTimeString, confirmation)
			// Add assignment to roster
			roster.GuardAssigned = append(roster.GuardAssigned, rosterAssignment)
		}

		// Check if there is already a client aifs assignment for this roster
		clientExists := false
		for _, existingClient := range roster.Clients {
			if existingClient.AifsClientRosterId == aifsClient.AifsClientRosterId {
				clientExists = true
				break
			}
		}

		if !clientExists {
			convertFromDbRosterAifsClient(db, aifsClient, clientId)
			// Add assignment to roster
			roster.Clients = append(roster.Clients, aifsClient)
		}
	}

	// Add all rosters to the returning array
	for _, roster := range rosterMap {
		*rosters = append(*rosters, roster)
	}

	sort.Slice(*rosters, func(i, j int) bool {
		return (*rosters)[i].AifsId > (*rosters)[j].AifsId
	})

	if query.Skip > 0 {
		*rosters = (*rosters)[query.Skip:]
	}

	return nil
}

// modifies the roster Assignment in place
func convertFromDbRosterAssignment(db *sql.DB, rosterAssignment *pb.RosterAssignement, assignedUserId int,
	customStartTimeString string, customEndTimeString string,
	attendanceTimeString sql.NullString, confirmation sql.NullBool) error {

	var err error
	rosterAssignment.GuardAssigned.Employee, err = idUserByUserId(db, assignedUserId)

	if err != nil {
		fmt.Println("convertFromDbRosterAssignment:", err.Error())
		return err
	}

	// Add Datetimes
	rosterAssignment.CustomStartTime, err = DBDatetimeToPB(customStartTimeString)
	if err != nil {
		fmt.Println("convertFromDbRosterAssignment:", err.Error())
		return err
	}
	rosterAssignment.CustomEndTime, err = DBDatetimeToPB(customEndTimeString)

	if err != nil {
		fmt.Println("convertFromDbRosterAssignment:", err.Error())
		return err
	}

	if attendanceTimeString.Valid {
		rosterAssignment.AttendanceTime, err = DBDatetimeToPB(attendanceTimeString.String)
		if err != nil {
			fmt.Println("convertFromDbRosterAssignment:", err.Error())
			return err
		}
	}

	if confirmation.Valid {
		rosterAssignment.Confirmed = confirmation.Bool
		if err != nil {
			fmt.Println("convertFromDbRosterAssignment:", err.Error())
			return err
		}
	}

	return nil
}

func convertFromDbRosterAifsClient(db *sql.DB, aifsClient *pb.AIFSClientRoster, clientId int) error {
	var err error
	aifsClient.Client, err = IdClientByClientId(db, clientId)

	if err != nil {
		fmt.Println("convertFromDbRosterAifsClient:", err.Error())
		return err
	}
	return nil
}

// This function is different from UpdateRosterAssignements()
// in the idea that this function finds out who the existing
// recipients are and make the necessary changes so that the
// recipients of the main roster will corresponds to the new
// list that is needed. Ie, it inserts and deletes recipients at will.
func updateAssignmentsOfRoster(db *sql.DB, roster *pb.Roster, query *pb.RosterQuery, dbLock *sync.Mutex) error {
	// Get all assignments
	currentAssignments, err := GetRosterAssingments(db, query, roster.RosteringId)
	if err != nil {
		fmt.Println("updateAssignmentsOfRoster ERROR::", err)
		return err
	}

	// create array of current roster recipient ids
	currentAsgnIds := make([]int, 0)

	for _, asgn := range currentAssignments {
		currentAsgnIds = append(currentAsgnIds, int(asgn.GuardAssigned.Employee.UserId))
	}

	sort.Ints(currentAsgnIds)

	// Check if the updated recipients exist within the current ones
	// If they exist, ignore and remove them from the current list.
	// If they do not exist, we need to add them to the db.
	// If the current recipient is not within the list of new recipients
	// delete this rogue recipient.

	// Index of missing recipients from the input roster list
	missingAsgnIndex := make([]int, 0)

	for i, asgn := range roster.GuardAssigned {
		found, index := common.BinarySearch(currentAsgnIds, 0, len(currentAsgnIds)-1, int(asgn.GuardAssigned.Employee.UserId))
		if found {
			fmt.Println("Found updated recipient in current recipient")
			currentAsgnIds = append(currentAsgnIds[:index], currentAsgnIds[index+1:]...)
		} else {
			missingAsgnIndex = append(missingAsgnIndex, i)
		}
	}

	fmt.Println("Missing Assignment Index:", missingAsgnIndex)
	// Add the missing recipients
	for _, asgnIndex := range missingAsgnIndex {
		_, err := InsertRosterASGN(db, roster.GuardAssigned[asgnIndex], roster.RosteringId, dbLock)
		if err != nil {
			fmt.Println("updateAssignmentsOfRoster ERROR::", err)
			return err
		}
	}

	fmt.Println("Deleting Assignment IDs:", currentAsgnIds)
	// See if any need to be deleted
	for _, id := range currentAsgnIds {
		_, err := DeleteRosterAssignment(db, &pb.RosterAssignement{RosterAssignmentId: int64(id)})
		if err != nil {
			fmt.Println("UpdateRoster ERROR::", err)
			return err
		}
	}

	return nil
}

// This function is different from UpdateRosterAssignements()
// in the idea that this function finds out who the existing
// recipients are and make the necessary changes so that the
// recipients of the main roster will corresponds to the new
// list that is needed. Ie, it inserts and deletes recipients at will.
func updateClientsOfRoster(db *sql.DB, roster *pb.Roster, query *pb.RosterQuery, dbLock *sync.Mutex) error {
	// Get all recipients
	currentClients, err := GetRosterAIFSClient(db, query, roster.RosteringId)
	if err != nil {
		fmt.Println("updateClientsOfRoster ERROR::", err)
		return err
	}

	// create array of current roster recipient ids
	currentClientids := make([]int, 0)

	for _, client := range currentClients {
		currentClientids = append(currentClientids, int(client.AifsClientRosterId))
	}

	sort.Ints(currentClientids)

	// Check if the updated recipients exist within the current ones
	// If they exist, ignore and remove them from the current list.
	// If they do not exist, we need to add them to the db.
	// If the current recipient is not within the list of new recipients
	// delete this rogue recipient.

	// Index of missing recipients from the input roster list
	missingClientIndex := make([]int, 0)

	for i, client := range roster.Clients {
		found, index := common.BinarySearch(currentClientids, 0, len(currentClientids)-1, int(client.AifsClientRosterId))
		if found {
			fmt.Println("Found updated recipient in current recipient")
			currentClientids = append(currentClientids[:index], currentClientids[index+1:]...)
		} else {
			missingClientIndex = append(missingClientIndex, i)
		}
	}

	fmt.Println("Missing AIFS Client Index:", missingClientIndex)
	// Add the missing recipients
	for _, cliIndex := range missingClientIndex {
		_, err := InsertAIFSClientRoster(db, roster.Clients[cliIndex], roster.RosteringId, dbLock)
		if err != nil {
			fmt.Println("updateClientsOfRoster ERROR::", err)
			return err
		}
	}

	fmt.Println("Deleting AIFS Client IDs:", currentClientids)
	// See if any need to be deleted
	for _, id := range currentClientids {
		_, err := DeleteRosterAIFSClient(db, &pb.AIFSClientRoster{AifsClientRosterId: int64(id)})
		if err != nil {
			fmt.Println("updateClientsOfRoster ERROR::", err)
			return err
		}
	}

	return nil
}

// This function creates the filter required if
// the only condition is a matching roster id.
func getRosterIdFormattedFilter(rosterId int, table string, isMainRosterId bool) string {
	query := &pb.RosterQuery{}
	if isMainRosterId {
		addRosterFilter(query, pb.RosterFilter_ROSTER_ID, pb.Filter_EQUAL, strconv.Itoa(rosterId))
	} else {
		if table == ROSTER_ASSIGNMENT_DB_TABLE_NAME {
			addRosterFilter(query, pb.RosterFilter_ROSTER_ASSIGNMENT_ID, pb.Filter_EQUAL, strconv.Itoa(rosterId))
		} else {
			addRosterFilter(query, pb.RosterFilter_ROSTER_AIFS_CLIENT_ID, pb.Filter_EQUAL, strconv.Itoa(rosterId))
		}
	}
	return getFormattedRosterFilters(query, table, false, false)
}

func rosterFilterToDBCol(filterField pb.RosterFilter_Field, table string) string {
	output := ""
	switch filterField {
	case pb.RosterFilter_ROSTER_ID:
		if table == ROSTER_DB_TABLE_NAME {
			output = ROSTER_DB_ID
		} else if table == ROSTER_AIFS_CLIENT_DB_TABLE_NAME {
			output = AIFS_CLIENT_DB_RELATED_ROSTER
		} else {
			output = ROSTER_ASGN_DB_RELATED_ROSTER
		}
	case pb.RosterFilter_AIFS_ID:
		output = ROSTER_DB_AIFS_ID
	case pb.RosterFilter_GUARD_ASSIGNED_ID:
		output = ROSTER_ASGN_DB_GUARD_ASSIGNED
	case pb.RosterFilter_CLIENT_ID:
		output = AIFS_CLIENT_DB_RELATED_CLIENT
	case pb.RosterFilter_GUARD_ASSIGNMENT_CONFIRMATION:
		output = ROSTER_ASGN_DB_CONFIRMATION
	case pb.RosterFilter_GUARD_ASSIGNMENT_ATTENDED:
		output = ROSTER_ASGN_DB_ATTENDED
	case pb.RosterFilter_ROSTER_ASSIGNMENT_ID:
		output = ROSTER_ASGN_DB_ID
	case pb.RosterFilter_ROSTER_AIFS_CLIENT_ID:
		output = AIFS_CLIENT_DB_ID
	case pb.RosterFilter_START_TIME:
		if table == ROSTER_DB_TABLE_NAME {
			output = ROSTER_DB_START_TIME
		} else if table == ROSTER_AIFS_CLIENT_DB_TABLE_NAME {
			output = ROSTER_ASGN_DB_START_TIME
		}
	case pb.RosterFilter_END_TIME:
		if table == ROSTER_DB_TABLE_NAME {
			output = ROSTER_DB_END_TIME
		} else if table == ROSTER_AIFS_CLIENT_DB_TABLE_NAME {
			output = ROSTER_ASGN_DB_END_TIME
		}
	}

	return output
}

// Returns the pk of the roster that already exists or -1
// if no such roster exists.
func checkRosterExists(db *sql.DB, roster *pb.Roster) (int64, error) {
	query := &pb.RosterQuery{}
	addRosterFilter(query, pb.RosterFilter_AIFS_ID, pb.Filter_EQUAL, strconv.Itoa(int(roster.AifsId)))
	addRosterFilter(query, pb.RosterFilter_START_TIME, pb.Filter_EQUAL, roster.StartTime.AsTime().Format(DATETIME_FORMAT))
	addRosterFilter(query, pb.RosterFilter_END_TIME, pb.Filter_EQUAL, roster.EndTime.AsTime().Format(DATETIME_FORMAT))
	rosters, err := GetRosters(db, query)

	if err != nil {
		return -1, nil
	}

	if len(rosters) > 0 {
		return rosters[0].RosteringId, nil
	}

	return -1, nil
}

// Removes any undesirable client queries from the original query
// in place.
// Returns the removed client queries in a new RosterQuery.
// unused for now
func RemoveRosteringClientQueries(rosterQuery *pb.RosterQuery) *pb.RosterQuery {
	var clientQueries *pb.RosterQuery
	foundIndexes := make([]int, 0)

	for i, query := range rosterQuery.Filters {
		if query.Field == pb.RosterFilter_CLIENT_ID {
			if clientQueries == nil {
				clientQueries = &pb.RosterQuery{
					Limit:   rosterQuery.Limit,
					Skip:    rosterQuery.Skip,
					OrderBy: rosterQuery.OrderBy,
					Filters: make([]*pb.RosterFilter, 0),
				}
			}

			clientQueries.Filters = append(clientQueries.Filters, query)

			foundIndexes = append(foundIndexes, i)
		}
	}

	if len(foundIndexes) == 0 {
		return clientQueries
	}

	// Remove the client queries from the original query
	for _, index := range foundIndexes {
		rosterQuery.Filters = append(rosterQuery.Filters[:index], rosterQuery.Filters[index+1:]...)
	}

	return clientQueries
}
