// Utility functions for database operations related to incidentReporting.
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
	rs "capstone.operations_ecosystem/backend/rating_system"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	INCIDENT_REPORT_DB_TABLE_NAME             = "schedule"
	INCIDENT_REPORT_ASSIGNMENT_DB_TABLE_NAME  = "schedule_detail"
	INCIDENT_REPORT_AIFS_CLIENT_DB_TABLE_NAME = "aifs_client_schedule"
	INCIDENT_REPORT_DEFAULT_DB_TABLE_NAME     = "default_incidentReporting"
	// IncidentReport table fields
	INCIDENT_REPORT_DB_ID         = "schedule_id"
	INCIDENT_REPORT_DB_AIFS_ID    = "aifs_id"
	INCIDENT_REPORT_DB_START_TIME = "start_time"
	INCIDENT_REPORT_DB_END_TIME   = "end_time"

	// IncidentReport guard assignment table fields
	INCIDENT_REPORT_ASGN_DB_ID                      = "schedule_detail_id"
	INCIDENT_REPORT_ASGN_DB_RELATED_INCIDENT_REPORT = "schedule"
	INCIDENT_REPORT_ASGN_DB_GUARD_ASSIGNED          = "guard_assigned"
	INCIDENT_REPORT_ASGN_DB_START_TIME              = "custom_start_time"
	INCIDENT_REPORT_ASGN_DB_END_TIME                = "custom_end_time"
	INCIDENT_REPORT_ASGN_DB_CONFIRMATION            = "confirmation"
	INCIDENT_REPORT_ASGN_DB_ATTENDED                = "attended"
	INCIDENT_REPORT_ASGN_DB_ATTENDED_TIME           = "attendance_time"
	INCIDENT_REPORT_ASGN_DB_IS_ASSIGNED             = "is_assigned"
	INCIDENT_REPORT_ASGN_DB_REJECTED                = "rejected"
)

// UTILITIES

// Returns the fields of the main incidentReport table
// in a specific order.
// Note that IDs are auto incremented and should
// not be modified manually. Ommits ID in resulting string.
func getIncidentReportTableFields(originalReportPk int64) string {
	incidentReportTableFields := []string{
		INCIDENT_REPORT_DB_AIFS_ID,
		INCIDENT_REPORT_DB_START_TIME,
		INCIDENT_REPORT_DB_END_TIME,
	}

	return strings.Join(incidentReportTableFields, ",")
}

// This function is highly dependent on the
// order given in getIncidentReportTableFields.
// Returns the values of the incidentReport fields in the
// order that is specified in getIncidentReportTableFields
func orderIncidentReportFields(incidentReport *pb.IncidentReport) string {
	output := ""

	output += "'" + strconv.Itoa(int(incidentReport.AifsId)) + "'" + ", "
	output += "'" + incidentReport.StartTime + "'" + ", "
	output += "'" + incidentReport.EndTime + "'"

	return output
}

// Returns the fields of the incidentReport assignment table
// in a specific order.
// Note that IDs are auto incremented and should
// not be modified manually. Ommits ID in resulting string.
// All nullable columns that we do not expect to be used upon
// first insert shall be omitted.
func getIncidentReportContentTableFields() string {
	incidentReportRecTableFields := []string{
		INCIDENT_REPORT_ASGN_DB_RELATED_INCIDENT_REPORT,
		INCIDENT_REPORT_ASGN_DB_GUARD_ASSIGNED,
		INCIDENT_REPORT_ASGN_DB_START_TIME,
		INCIDENT_REPORT_ASGN_DB_END_TIME,
		INCIDENT_REPORT_ASGN_DB_CONFIRMATION,
		INCIDENT_REPORT_ASGN_DB_ATTENDED,
		INCIDENT_REPORT_ASGN_DB_IS_ASSIGNED,
		INCIDENT_REPORT_ASGN_DB_REJECTED,
	}

	return strings.Join(incidentReportRecTableFields, ",")
}

// This function is highly dependent on the
// order given in getIncidentReportRecTableFields.
// Returns the values of the incidentReport fields in the
// order that is specified in getIncidentReportRecTableFields
func orderIncidentReportContentFields(incidentReportAssignment *pb.IncidentReportContent) string {
	output := ""

	output += strconv.Itoa(int(relatedIncidentReportId)) + ","
	output += strconv.Itoa(int(incidentReportAssignment.GuardAssigned.Employee.UserId)) + ","
	output += "'" + incidentReportAssignment.CustomStartTime.AsTime().Format(common.DATETIME_FORMAT) + "'" + ", "
	output += "'" + incidentReportAssignment.CustomEndTime.AsTime().Format(common.DATETIME_FORMAT) + "'" + ", "

	// confirmation and attended are false by default.
	output += "0, 0" + ", "
	// is assigned is true by default
	output += "1" + ", "
	// rejection are false by default
	output += "0"

	return output
}

// This function is used in anticipation of an update query.
// Updates use the format "SET field1=val1, field2=val2".
// The return value of this function is for example: "field1=val1, field2=val2"
// ID not included in this because pk should not be manually changed.
// Note this function highly depends on the protocol buffer message definition
// Returns the string fields and values of the filled incidentReport fields
func getFilledIncidentReportFields(incidentReport *pb.IncidentReport) string {
	incidentReportTableFields := []string{}

	if incidentReport.AifsId > 0 {
		incidentReportTableFields = append(incidentReportTableFields, formatFieldEqVal(INCIDENT_REPORT_DB_AIFS_ID, strconv.Itoa(int(incidentReport.AifsId)), true))
	}
	if len(incidentReport.StartTime) > 0 {
		incidentReportTableFields = append(incidentReportTableFields, formatFieldEqVal(INCIDENT_REPORT_DB_START_TIME, incidentReport.StartTime, true))
	}
	if len(incidentReport.EndTime) > 0 {
		incidentReportTableFields = append(incidentReportTableFields, formatFieldEqVal(INCIDENT_REPORT_DB_END_TIME, incidentReport.EndTime, true))
	}

	return strings.Join(incidentReportTableFields, ",")
}

// This function is used in anticipation of an update query.
// Updates use the format "SET field1=val1, field2=val2".
// The return value of this function is for example: "field1=val1, field2=val2"
// ID not included in this because pk should not be manually changed.
// Note this function highly depends on the protocol buffer message definition
// Returns the string fields and values of the filled incidentReport recipinets fields
func getFilledIncidentReportASGNFields(incidentReportAssignment *pb.IncidentReportContent) string {
	incidentReportTableFields := []string{}

	if incidentReportAssignment.GuardAssigned != nil {
		if incidentReportAssignment.GuardAssigned.Employee != nil {
			incidentReportTableFields = append(incidentReportTableFields, formatFieldEqVal(INCIDENT_REPORT_ASGN_DB_GUARD_ASSIGNED, strconv.Itoa(int(incidentReportAssignment.GuardAssigned.Employee.UserId)), true))
		}
	}
	if incidentReportAssignment.CustomStartTime != nil {
		incidentReportTableFields = append(incidentReportTableFields, formatFieldEqVal(INCIDENT_REPORT_ASGN_DB_START_TIME, incidentReportAssignment.CustomStartTime.AsTime().Format(common.DATETIME_FORMAT), true))
	}
	if incidentReportAssignment.CustomEndTime != nil {
		incidentReportTableFields = append(incidentReportTableFields, formatFieldEqVal(INCIDENT_REPORT_ASGN_DB_END_TIME, incidentReportAssignment.CustomEndTime.AsTime().Format(common.DATETIME_FORMAT), true))
	}

	incidentReportTableFields = append(incidentReportTableFields, formatFieldEqVal(INCIDENT_REPORT_ASGN_DB_CONFIRMATION, strconv.FormatBool(incidentReportAssignment.Confirmed), false))
	incidentReportTableFields = append(incidentReportTableFields, formatFieldEqVal(INCIDENT_REPORT_ASGN_DB_ATTENDED, strconv.FormatBool(incidentReportAssignment.Attended), false))

	if incidentReportAssignment.AttendanceTime != nil {
		incidentReportTableFields = append(incidentReportTableFields, formatFieldEqVal(INCIDENT_REPORT_ASGN_DB_ATTENDED_TIME, incidentReportAssignment.AttendanceTime.AsTime().Format(common.DATETIME_FORMAT), true))
	}

	incidentReportTableFields = append(incidentReportTableFields, formatFieldEqVal(INCIDENT_REPORT_ASGN_DB_IS_ASSIGNED, strconv.FormatBool(incidentReportAssignment.IsAssigned), false))
	incidentReportTableFields = append(incidentReportTableFields, formatFieldEqVal(INCIDENT_REPORT_ASGN_DB_REJECTED, strconv.FormatBool(incidentReportAssignment.Rejected), false))

	return strings.Join(incidentReportTableFields, ",")
}

// This function is used in anticipation of an update query.
// Updates use the format "SET field1=val1, field2=val2".
// The return value of this function is for example: "field1=val1, field2=val2"
// ID not included in this because pk should not be manually changed.
// Note this function highly depends on the protocol buffer message definition
// Returns the string fields and values of the filled incidentReport recipinets fields
func getFilledAIFSClientFields(aifsClient *pb.AIFSClientIncidentReport) string {
	incidentReportTableFields := []string{}

	if aifsClient.Client != nil {
		incidentReportTableFields = append(incidentReportTableFields, formatFieldEqVal(AIFS_CLIENT_DB_RELATED_CLIENT, strconv.Itoa(int(aifsClient.Client.ClientId)), true))
	}

	if aifsClient.PatrolOrder != 0 {
		incidentReportTableFields = append(incidentReportTableFields, formatFieldEqVal(AIFS_CLIENT_DB_PATROL_ORDER, strconv.Itoa(int(aifsClient.PatrolOrder)), true))
	}

	return strings.Join(incidentReportTableFields, ",")
}

// Helper function to add a new filter to the list of existing
// filters in a incidentReport query struct.
// Modifies the incidentReport query parameter directly.
func AddIncidentReportFilter(query *pb.IncidentReportQuery, field pb.IncidentReportFilter_Field,
	comparison pb.Filter_Comparisons,
	value string) {
	if query.Filters == nil {
		query.Filters = make([]*pb.IncidentReportFilter, 0)
	}
	filter := &pb.Filter{Comparison: comparison, Value: value}
	query.Filters = append(query.Filters, &pb.IncidentReportFilter{Field: field, Comparisons: filter})
}

// Converts the filters in the incidentReport array into a formatted where clause
// that can be parsed into MySQL. If a limit is needed, the LIMIT filter is
// added to the end of the string.
// For example returns: "WHERE id=22 AND num <2 LIMIT 5"
// Returns the formatted SQL filter string.
func getFormattedIncidentReportFilters(query *pb.IncidentReportQuery, table string, needLimit bool, needOrder bool) string {
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
		case pb.IncidentReportFilter_INCIDENT_REPORT_ID, pb.IncidentReportFilter_AIFS_ID, pb.IncidentReportFilter_GUARD_ASSIGNED_ID,
			pb.IncidentReportFilter_CLIENT_ID, pb.IncidentReportFilter_INCIDENT_REPORT_ASSIGNMENT_ID, pb.IncidentReportFilter_START_TIME,
			pb.IncidentReportFilter_END_TIME, pb.IncidentReportFilter_INCIDENT_REPORT_AIFS_CLIENT_ID, pb.IncidentReportFilter_DEFAULT_INCIDENT_REPORTING_DAY_OF_WEEK:
			if hasQuotes {
				whereFilters = append(
					whereFilters, fmt.Sprintf("%s %s '%s'", incidentReportFilterToDBCol(filter.Field, table),
						GetFilterComparisonSign(filter.Comparisons.Comparison), filter.Comparisons.Value),
				)
			} else {
				whereFilters = append(
					whereFilters, fmt.Sprintf("%s %s %s", incidentReportFilterToDBCol(filter.Field, table),
						GetFilterComparisonSign(filter.Comparisons.Comparison), filter.Comparisons.Value),
				)
			}
		case pb.IncidentReportFilter_GUARD_ASSIGNMENT_CONFIRMATION, pb.IncidentReportFilter_GUARD_ASSIGNMENT_ATTENDED,
			pb.IncidentReportFilter_IS_ASSIGNED:
			whereFilters = append(
				whereFilters, fmt.Sprintf("%s %s %s", incidentReportFilterToDBCol(filter.Field, table),
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
			output += fmt.Sprintf(" %s %s %s", ORDER_BY_KEYWORD, incidentReportFilterToDBCol(query.OrderBy.Field, table), orderByProtoToDB(query.OrderBy.OrderBy))
		} else if table == INCIDENT_REPORT_DB_TABLE_NAME {
			// By default we order incidentReports by the aifs id
			output += fmt.Sprintf(" %s %s %s", ORDER_BY_KEYWORD, incidentReportFilterToDBCol(pb.IncidentReportFilter_AIFS_ID, table), ASC_KEYWORD)
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

// This function converts the returned DB rows into IncidentReport objects and
// their corresponding incidentReport recipients.
// These rows come from the join query of both the incidentReport and incidentReport
// recipients table.
// Modifies the incidentReport array in place.
func convertDbRowsToFullIncidentReport(db *sql.DB, incidentReports *[]*pb.IncidentReport, rows *sql.Rows, query *pb.IncidentReportQuery) error {
	incidentReportMap := make(map[int64]*pb.IncidentReport)

	for rows.Next() {
		incidentReport := &pb.IncidentReport{
			GuardAssigned: make([]*pb.IncidentReportContent, 0),
			Clients:       make([]*pb.AIFSClientIncidentReport, 0),
		}
		incidentReportAssignment := &pb.IncidentReportContent{}
		employeeEval := &pb.EmployeeEvaluation{}
		incidentReportAssignment.GuardAssigned = employeeEval

		aifsClient := &pb.AIFSClientIncidentReport{}

		assignedUserId := -1
		clientId := -1

		// Redundant Strings
		assignmentRelatedIncidentReport := ""
		clientRelatedIncidentReport := ""

		// confirmation can be null
		var confirmation sql.NullBool

		// Datetimes
		customStartTimeString := ""
		customEndTimeString := ""
		var attendanceTimeString sql.NullString

		// cast each row to a incidentReport
		err := rows.Scan(
			// Main IncidentReport
			&incidentReport.IncidentReportingId,
			&incidentReport.AifsId,
			&incidentReport.StartTime,
			&incidentReport.EndTime,

			// IncidentReport Details
			&incidentReportAssignment.IncidentReportAssignmentId,
			&assignmentRelatedIncidentReport,
			&assignedUserId,
			&customStartTimeString,
			&customEndTimeString,
			&incidentReportAssignment.Confirmed,
			&incidentReportAssignment.Attended,
			&attendanceTimeString,
			&incidentReportAssignment.IsAssigned,
			&incidentReportAssignment.Rejected,

			// Client Details
			&aifsClient.AifsClientIncidentReportId,
			&clientRelatedIncidentReport,
			&clientId,
			&aifsClient.PatrolOrder,
		)

		if err != nil {
			fmt.Println("convertDbRowsToFullIncidentReport ERROR:", err)
			continue
		}

		// Check if there was already a incidentReport found with this id
		if existingIncidentReport, ok := incidentReportMap[incidentReport.IncidentReportingId]; ok {
			incidentReport = existingIncidentReport
		} else {
			// Return only the necessary number of incidentReports.
			// If the number of incidentReports have reached the limit,
			// do not add new incidentReports.
			if int64(len(incidentReportMap)) >= query.Limit+query.Skip {
				continue
			}

			incidentReportMap[incidentReport.IncidentReportingId] = incidentReport
		}

		// Check if there is already a guard assignment for this incidentReport
		guardAssignmentExists := false
		for _, existingAssigned := range incidentReport.GuardAssigned {
			if existingAssigned.IncidentReportAssignmentId == incidentReportAssignment.IncidentReportAssignmentId {
				guardAssignmentExists = true
				break
			}
		}

		if !guardAssignmentExists {
			convertFromDbIncidentReportAssignment(db, incidentReportAssignment, assignedUserId,
				customStartTimeString, customEndTimeString,
				attendanceTimeString, confirmation)
			// Add assignment to incidentReport
			incidentReport.GuardAssigned = append(incidentReport.GuardAssigned, incidentReportAssignment)
		}

		// Check if there is already a client aifs assignment for this incidentReport
		clientExists := false
		for _, existingClient := range incidentReport.Clients {
			if existingClient.AifsClientIncidentReportId == aifsClient.AifsClientIncidentReportId {
				clientExists = true
				break
			}
		}

		if !clientExists {
			convertFromDbIncidentReportAifsClient(db, aifsClient, clientId)
			// Add assignment to incidentReport
			incidentReport.Clients = append(incidentReport.Clients, aifsClient)
		}
	}

	// Add all incidentReports to the returning array
	for _, incidentReport := range incidentReportMap {
		*incidentReports = append(*incidentReports, incidentReport)
	}

	sort.Slice(*incidentReports, func(i, j int) bool {
		return (*incidentReports)[i].AifsId < (*incidentReports)[j].AifsId
	})

	// sort the aifs patrol order
	for _, incidentReport := range *incidentReports {
		sort.Slice(incidentReport.Clients, func(i, j int) bool {
			return incidentReport.Clients[i].PatrolOrder < incidentReport.Clients[j].PatrolOrder
		})
	}

	if query.Skip > 0 {
		*incidentReports = (*incidentReports)[query.Skip:]
	}

	return nil
}

// modifies the incidentReport Assignment in place
func convertFromDbIncidentReportAssignment(db *sql.DB, incidentReportAssignment *pb.IncidentReportContent, assignedUserId int,
	customStartTimeString string, customEndTimeString string,
	attendanceTimeString sql.NullString, confirmation sql.NullBool) error {

	var err error
	incidentReportAssignment.GuardAssigned.Employee, err = idUserByUserId(db, assignedUserId)

	if err != nil {
		fmt.Println("convertFromDbIncidentReportAssignment:", err.Error())
		return err
	}

	// Add Datetimes
	incidentReportAssignment.CustomStartTime, err = DBDatetimeToPB(customStartTimeString)
	if err != nil {
		fmt.Println("convertFromDbIncidentReportAssignment:", err.Error())
		return err
	}
	incidentReportAssignment.CustomEndTime, err = DBDatetimeToPB(customEndTimeString)

	if err != nil {
		fmt.Println("convertFromDbIncidentReportAssignment:", err.Error())
		return err
	}

	if attendanceTimeString.Valid {
		incidentReportAssignment.AttendanceTime, err = DBDatetimeToPB(attendanceTimeString.String)
		if err != nil {
			fmt.Println("convertFromDbIncidentReportAssignment:", err.Error())
			return err
		}
	}

	if confirmation.Valid {
		incidentReportAssignment.Confirmed = confirmation.Bool
		if err != nil {
			fmt.Println("convertFromDbIncidentReportAssignment:", err.Error())
			return err
		}
	}

	// Get the user scores
	incidentReportAssignment.GuardAssigned.EmployeeScore, err = rs.GetUserScore(incidentReportAssignment.GuardAssigned.Employee)
	if err != nil {
		fmt.Println("convertFromDbIncidentReportAssignment:", err.Error())
		return err
	}

	return nil
}

func convertFromDbIncidentReportAifsClient(db *sql.DB, aifsClient *pb.AIFSClientIncidentReport, clientId int) error {
	var err error
	aifsClient.Client, err = IdClientByClientId(db, clientId)

	if err != nil {
		fmt.Println("convertFromDbIncidentReportAifsClient:", err.Error())
		return err
	}
	return nil
}

// This function is different from UpdateIncidentReportContents()
// in the idea that this function finds out who the existing
// recipients are and make the necessary changes so that the
// recipients of the main incidentReport will corresponds to the new
// list that is needed. Ie, it inserts and deletes recipients at will.
func updateAssignmentsOfIncidentReport(db *sql.DB, incidentReport *pb.IncidentReport, dbLock *sync.Mutex) ([]int64, error) {
	newIncidentReportAssignmentsPk := make([]int64, 0)

	// Get all current assignments that are assigned
	query := &pb.IncidentReportQuery{}
	AddIncidentReportFilter(query, pb.IncidentReportFilter_IS_ASSIGNED, pb.Filter_EQUAL, "1")
	currentAssignments, err := GetIncidentReportAssingments(db, query, incidentReport.IncidentReportingId)
	if err != nil {
		fmt.Println("updateAssignmentsOfIncidentReport ERROR::", err)
		return newIncidentReportAssignmentsPk, err
	}

	// create array of current incidentReport recipient ids
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

	// Index of missing recipients from the input incidentReport list
	missingAsgnIndex := make([]int, 0)

	for i, asgn := range incidentReport.GuardAssigned {
		// Ensure not nil
		if asgn.GuardAssigned.Employee == nil {
			return newIncidentReportAssignmentsPk, status.Errorf(codes.InvalidArgument, "Employee field of guards in incidentReport assignment must not be nil")
		}

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
		// Ensure the start and end custom times of the guard assigned
		// is the same as the incidentReport
		incidentReport.GuardAssigned[asgnIndex].CustomStartTime, err = DBDatetimeToPB(incidentReport.StartTime)
		if err != nil {
			fmt.Println("updateAssignmentsOfIncidentReport ERROR::", err)
			return newIncidentReportAssignmentsPk, err
		}
		incidentReport.GuardAssigned[asgnIndex].CustomEndTime, err = DBDatetimeToPB(incidentReport.EndTime)
		if err != nil {
			fmt.Println("updateAssignmentsOfIncidentReport ERROR::", err)
			return newIncidentReportAssignmentsPk, err
		}
		incidentReportAsgnPk, err := InsertIncidentReportASGN(db, incidentReport.GuardAssigned[asgnIndex], incidentReport.IncidentReportingId, dbLock)
		if err != nil {
			fmt.Println("updateAssignmentsOfIncidentReport ERROR::", err)
			return newIncidentReportAssignmentsPk, err
		}
		newIncidentReportAssignmentsPk = append(newIncidentReportAssignmentsPk, incidentReportAsgnPk)
	}

	fmt.Println("Removing is_assigned flag for Assignment IDs:", currentAsgnIds)
	// See if any need to be removed
	for _, id := range currentAsgnIds {
		// change the is_assigned flag to false
		query = &pb.IncidentReportQuery{} //TODO HERE
		AddIncidentReportFilter(query, pb.IncidentReportFilter_GUARD_ASSIGNED_ID, pb.Filter_EQUAL, strconv.Itoa(id))
		AddIncidentReportFilter(query, pb.IncidentReportFilter_INCIDENT_REPORT_ID, pb.Filter_EQUAL, strconv.Itoa(int(incidentReport.IncidentReportingId)))
		_, err := UpdateIncidentReportAssignments(db, &pb.IncidentReportContent{
			IncidentReportAssignmentId: int64(id),
			Confirmed:                  false,
			Attended:                   false,
			IsAssigned:                 false,
		}, query)
		if err != nil {
			fmt.Println("UpdateIncidentReport ERROR::", err)
			return newIncidentReportAssignmentsPk, err
		}
	}

	return newIncidentReportAssignmentsPk, nil
}

// NOTE: UNTESTED because unused
// This function is different from UpdateIncidentReportContents()
// in the idea that this function finds out who the existing
// recipients are and make the necessary changes so that the
// recipients of the main incidentReport will corresponds to the new
// list that is needed. Ie, it inserts and deletes recipients at will.
func updateClientsOfIncidentReport(db *sql.DB, incidentReport *pb.IncidentReport, dbLock *sync.Mutex) error {
	// Get all recipients
	currentClients, err := GetIncidentReportAIFSClient(db, &pb.IncidentReportQuery{}, incidentReport.IncidentReportingId)
	if err != nil {
		fmt.Println("updateClientsOfIncidentReport ERROR::", err)
		return err
	}

	// create array of current incidentReport recipient ids
	currentClientids := make([]int, 0)

	for _, client := range currentClients {
		currentClientids = append(currentClientids, int(client.AifsClientIncidentReportId))
	}

	sort.Ints(currentClientids)

	// Check if the updated recipients exist within the current ones
	// If they exist, ignore and remove them from the current list.
	// If they do not exist, we need to add them to the db.
	// If the current recipient is not within the list of new recipients
	// delete this rogue recipient.

	// Index of missing recipients from the input incidentReport list
	missingClientIndex := make([]int, 0)

	for i, client := range incidentReport.Clients {
		found, index := common.BinarySearch(currentClientids, 0, len(currentClientids)-1, int(client.AifsClientIncidentReportId))
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
		_, err := InsertAIFSClientIncidentReport(db, incidentReport.Clients[cliIndex], incidentReport.IncidentReportingId, dbLock)
		if err != nil {
			fmt.Println("updateClientsOfIncidentReport ERROR::", err)
			return err
		}
	}

	fmt.Println("Deleting AIFS Client IDs:", currentClientids)
	// See if any need to be deleted
	for _, id := range currentClientids {
		_, err := DeleteIncidentReportAIFSClient(db, &pb.AIFSClientIncidentReport{AifsClientIncidentReportId: int64(id)})
		if err != nil {
			fmt.Println("updateClientsOfIncidentReport ERROR::", err)
			return err
		}
	}

	return nil
}

// This function creates the filter required if
// the only condition is a matching incidentReport id.
func getIncidentReportIdFormattedFilter(incidentReportId int, table string, isMainIncidentReportId bool) string {
	query := &pb.IncidentReportQuery{}
	if isMainIncidentReportId {
		AddIncidentReportFilter(query, pb.IncidentReportFilter_INCIDENT_REPORT_ID, pb.Filter_EQUAL, strconv.Itoa(incidentReportId))
	} else {
		if table == INCIDENT_REPORT_ASSIGNMENT_DB_TABLE_NAME {
			AddIncidentReportFilter(query, pb.IncidentReportFilter_INCIDENT_REPORT_ASSIGNMENT_ID, pb.Filter_EQUAL, strconv.Itoa(incidentReportId))
		} else {
			AddIncidentReportFilter(query, pb.IncidentReportFilter_INCIDENT_REPORT_AIFS_CLIENT_ID, pb.Filter_EQUAL, strconv.Itoa(incidentReportId))
		}
	}
	return getFormattedIncidentReportFilters(query, table, false, false)
}

func incidentReportFilterToDBCol(filterField pb.IncidentReportFilter_Field, table string) string {
	output := ""
	switch filterField {
	case pb.IncidentReportFilter_INCIDENT_REPORT_ID:
		if table == INCIDENT_REPORT_DB_TABLE_NAME {
			output = INCIDENT_REPORT_DB_ID
		} else if table == INCIDENT_REPORT_AIFS_CLIENT_DB_TABLE_NAME {
			output = AIFS_CLIENT_DB_RELATED_INCIDENT_REPORT
		} else {
			output = INCIDENT_REPORT_ASGN_DB_RELATED_INCIDENT_REPORT
		}
	case pb.IncidentReportFilter_AIFS_ID:
		output = INCIDENT_REPORT_DB_AIFS_ID
	case pb.IncidentReportFilter_GUARD_ASSIGNED_ID:
		output = INCIDENT_REPORT_ASGN_DB_GUARD_ASSIGNED
	case pb.IncidentReportFilter_CLIENT_ID:
		output = AIFS_CLIENT_DB_RELATED_CLIENT
	case pb.IncidentReportFilter_GUARD_ASSIGNMENT_CONFIRMATION:
		output = INCIDENT_REPORT_ASGN_DB_CONFIRMATION
	case pb.IncidentReportFilter_GUARD_ASSIGNMENT_ATTENDED:
		output = INCIDENT_REPORT_ASGN_DB_ATTENDED
	case pb.IncidentReportFilter_INCIDENT_REPORT_ASSIGNMENT_ID:
		output = INCIDENT_REPORT_ASGN_DB_ID
	case pb.IncidentReportFilter_INCIDENT_REPORT_AIFS_CLIENT_ID:
		output = AIFS_CLIENT_DB_ID
	case pb.IncidentReportFilter_START_TIME:
		if table == INCIDENT_REPORT_DB_TABLE_NAME {
			output = INCIDENT_REPORT_DB_START_TIME
		} else if table == INCIDENT_REPORT_ASSIGNMENT_DB_TABLE_NAME {
			output = INCIDENT_REPORT_ASGN_DB_START_TIME
		}
	case pb.IncidentReportFilter_END_TIME:
		if table == INCIDENT_REPORT_DB_TABLE_NAME {
			output = INCIDENT_REPORT_DB_END_TIME
		} else if table == INCIDENT_REPORT_ASSIGNMENT_DB_TABLE_NAME {
			output = INCIDENT_REPORT_ASGN_DB_END_TIME
		}
	case pb.IncidentReportFilter_IS_ASSIGNED:
		output = INCIDENT_REPORT_ASGN_DB_IS_ASSIGNED
	case pb.IncidentReportFilter_DEFAULT_INCIDENT_REPORTING_DAY_OF_WEEK:
		output = DEFAULT_INCIDENT_REPORTING_DB_DAY_OF_WEEK
	}

	return output
}

// Returns the pk of the incidentReport that already exists or -1
// if no such incidentReport exists.
func checkIncidentReportExists(db *sql.DB, incidentReport *pb.IncidentReport) (int64, error) {
	query := &pb.IncidentReportQuery{}
	AddIncidentReportFilter(query, pb.IncidentReportFilter_AIFS_ID, pb.Filter_EQUAL, strconv.Itoa(int(incidentReport.AifsId)))
	AddIncidentReportFilter(query, pb.IncidentReportFilter_START_TIME, pb.Filter_EQUAL, incidentReport.StartTime)
	AddIncidentReportFilter(query, pb.IncidentReportFilter_END_TIME, pb.Filter_EQUAL, incidentReport.EndTime)
	incidentReports, err := GetIncidentReports(db, query)

	if err != nil {
		return -1, nil
	}

	if len(incidentReports) > 0 {
		return incidentReports[0].IncidentReportingId, nil
	}

	return -1, nil
}

// Removes any undesirable client queries from the original query
// in place.
// Returns the removed client queries in a new IncidentReportQuery.
// unused for now
func RemoveIncidentReportingClientQueries(incidentReportQuery *pb.IncidentReportQuery) *pb.IncidentReportQuery {
	var clientQueries *pb.IncidentReportQuery
	foundIndexes := make([]int, 0)

	for i, query := range incidentReportQuery.Filters {
		if query.Field == pb.IncidentReportFilter_CLIENT_ID {
			if clientQueries == nil {
				clientQueries = &pb.IncidentReportQuery{
					Limit:   incidentReportQuery.Limit,
					Skip:    incidentReportQuery.Skip,
					OrderBy: incidentReportQuery.OrderBy,
					Filters: make([]*pb.IncidentReportFilter, 0),
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
		incidentReportQuery.Filters = append(incidentReportQuery.Filters[:index], incidentReportQuery.Filters[index+1:]...)
	}

	return clientQueries
}

func setIncidentReportStatus(incidentReports []*pb.IncidentReport) {
	for _, incidentReport := range incidentReports {
		has_unconfirmed := false
		has_rejected := false
		status := pb.IncidentReport_CONFIRMED
		for _, incidentReportAssignement := range incidentReport.GuardAssigned {
			if incidentReportAssignement.Rejected {
				has_rejected = true
				break
			}
			if !incidentReportAssignement.Confirmed {
				has_unconfirmed = true
			}
		}
		if has_rejected {
			status = pb.IncidentReport_REJECTED
		} else if has_unconfirmed {
			status = pb.IncidentReport_PENDING
		}
		incidentReport.Status = status
	}
}
