// Use these functions to interact with the incidentReport related database tables.
package database

import (
	"database/sql"
	"fmt"
	"strconv"
	"sync"

	pb "capstone.operations_ecosystem/backend/proto"
	_ "github.com/go-sql-driver/mysql"
)

// Insert a new incidentReport into the database table.
// Corresponding incidentReport details are added to their
// respective table as well.
// Returns the primary key of the main incidentReport and errors if any
func InsertIncidentReport(db *sql.DB, incidentReport *pb.IncidentReport, dbLock *sync.Mutex) (int64, error) {
	fmt.Println("Inserting IncidentReport", incidentReport.IncidentReportId)

	// Insert the origial content first
	originalContentPk, err := InsertIncidentReportContent(db, incidentReport.IncidentReportContent, dbLock)
	if err != nil {
		// Do not add the rest if the content fails
		return -1, err
	}

	// Create and insert main incidentReport and get it's pk
	incidentReportTbFields := getIncidentReportTableFields(originalContentPk)
	incidentReportValues := orderIncidentReportFields(incidentReport)

	incidentReportPk, err := Insert(db, ROSTER_DB_TABLE_NAME, incidentReportTbFields, incidentReportValues, dbLock)

	if err != nil {
		// Delete the incidentReport that was just inserted, it should cascade
		// and delete all the other incidentReport details
		incidentReport.IncidentReportContent.ReportContentId = originalContentPk
		DeleteIncidentReportContent(db, incidentReport.IncidentReportContent)
	}

	return incidentReportPk, err
}

// Inserts a new content to the database
// Returns the primary key of the recipient row and any errors.
func InsertIncidentReportContent(db *sql.DB, assignment *pb.IncidentReportContent, dbLock *sync.Mutex) (int64, error) {
	// get fields and values for this particular recipient
	fields := getIncidentReportContentTableFields()
	values := orderIncidentReportContentFields(assignment)

	// Add recipient to DB
	pk, err := Insert(db, ROSTER_ASSIGNMENT_DB_TABLE_NAME, fields, values, dbLock)

	return pk, err
}

// Get all the incidentReport rows in a table that meets specifications.
// Returns an array of incidentReports and any errors.
func GetIncidentReports(db *sql.DB, query *pb.IncidentReportQuery) ([]*pb.IncidentReport, error) {
	fmt.Println("Getting IncidentReports...")
	incidentReports := make([]*pb.IncidentReport, 0)

	// Join the incidentReport and assignment tables and aifs client tables
	// in order to easily filter conditions relating to all tables

	// Set default query limits if needed
	if query.Limit == 0 {
		query.Limit = DEFAULT_LIMIT
	}

	// We ignore any filters to do with the client aifs table first
	requestedLimit := query.Limit

	fields := ALL_COLS

	// tables are joined on the main incidentReport id
	fistOnCondition := formatFieldEqVal(ROSTER_DB_ID, ROSTER_ASSIGNMENT_DB_TABLE_NAME+"."+ROSTER_ASGN_DB_RELATED_ROSTER, false)
	secondOnCondition := formatFieldEqVal(ROSTER_DB_ID, ROSTER_AIFS_CLIENT_DB_TABLE_NAME+"."+AIFS_CLIENT_DB_RELATED_ROSTER, false)

	// Format filters
	// temporarily give the query limit the max
	query.Limit = MAX_LIMIT
	filters := getFormattedIncidentReportFilters(query, ROSTER_DB_TABLE_NAME, true, true)

	rows, err := QueryThreeTablesLeftJoin(db, ROSTER_DB_TABLE_NAME,
		ROSTER_ASSIGNMENT_DB_TABLE_NAME, ROSTER_AIFS_CLIENT_DB_TABLE_NAME,
		fistOnCondition, secondOnCondition, fields, filters)

	if err != nil {
		return incidentReports, err
	}

	// convert query rows into incidentReports
	// give back the query the original limit
	query.Limit = requestedLimit
	err = convertDbRowsToFullIncidentReport(db, &incidentReports, rows, query)

	// Set status of incidentReports
	setIncidentReportStatus(incidentReports)

	return incidentReports, err
}

// Get all the incidentReport recipient rows in a table that meets specifications.
// Returns an array of incidentReport recipients and any errors.
func GetIncidentReportAssingments(db *sql.DB, query *pb.IncidentReportQuery, mainIncidentReportID int64) ([]*pb.IncidentReportContent, error) {
	fmt.Println("Getting IncidentReports Assignments...")
	incidentReportRecipients := make([]*pb.IncidentReportContent, 0)

	fields := ALL_COLS

	// Format filters
	// Get for a specific main incidentReport if needed
	if mainIncidentReportID != -1 {
		AddIncidentReportFilter(query, pb.IncidentReportFilter_ROSTER_ID, pb.Filter_EQUAL, strconv.Itoa(int(mainIncidentReportID)))
	}

	filters := getFormattedIncidentReportFilters(query, ROSTER_ASSIGNMENT_DB_TABLE_NAME, true, true)

	rows, err := Query(db, ROSTER_ASSIGNMENT_DB_TABLE_NAME, fields, filters)

	if err != nil {
		return incidentReportRecipients, err
	}

	// convert query rows into incidentReports assignments
	for rows.Next() {
		assignment := &pb.IncidentReportContent{}
		employeeEval := &pb.EmployeeEvaluation{}
		assignment.GuardAssigned = employeeEval

		// fields that cannot be auto converted
		guardId := -1
		// related incidentReport is not necessary, but for simplicity
		// and for possible future use, we get it back in the query.
		relatedIncidentReport := ""

		// confirmation is nullable
		var confirmation sql.NullBool

		// Datetimes
		startTimeString := ""
		endTimeString := ""
		var attendanceTimeString sql.NullString

		// cast each row to a incidentReport
		err = rows.Scan(
			&assignment.IncidentReportAssignmentId,
			&relatedIncidentReport,
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
			fmt.Println("GetIncidentReportAssingments ERROR::", err)
			break
		}

		// Add Datetimes
		assignment.CustomStartTime, err = DBDatetimeToPB(startTimeString)
		if err != nil {
			fmt.Println("GetIncidentReportAssingments:", err.Error())
			continue
		}
		assignment.CustomEndTime, err = DBDatetimeToPB(endTimeString)
		if err != nil {
			fmt.Println("GetIncidentReportAssingments:", err.Error())
			continue
		}

		if attendanceTimeString.Valid {
			assignment.AttendanceTime, err = DBDatetimeToPB(attendanceTimeString.String)
			if err != nil {
				fmt.Println("GetIncidentReportAssingments:", err.Error())
				continue
			}
		}

		if confirmation.Valid {
			assignment.Confirmed = confirmation.Bool
			if err != nil {
				fmt.Println("GetIncidentReportAssingments:", err.Error())
				continue
			}
		}
		// TODO think about whether I can store the users in cache rather than
		// get the same few users over and over
		assignment.GuardAssigned.Employee, err = idUserByUserId(db, guardId)
		if err != nil {
			fmt.Println("GetIncidentReportAssingments:", err)
			continue
		}

		incidentReportRecipients = append(incidentReportRecipients, assignment)
	}

	return incidentReportRecipients, err
}

// Update a specific incidentReport in the table
// Only fields that have been filled in the incidentReport object will be updated.
// Returns the number of main incidentReport rows updated as well as the a list of
// all the primary keys of newly inserted incidentReport assignments into the db.
// Note that this update does not update the incidentReport's guards's inner status
// such as the acknowledgement or attended status but only if the guard
// is part of the incidentReport. Same for the clients
func UpdateIncidentReport(db *sql.DB, incidentReport *pb.IncidentReport, dbLock *sync.Mutex) (int64, error) {

	// Update the main incidentReport first
	newFields := getFilledIncidentReportFields(incidentReport)

	filters := getIncidentReportIdFormattedFilter(
		int(incidentReport.IncidentReportId),
		ROSTER_DB_TABLE_NAME, true,
	)

	var err error
	rowsAffected := int64(0)

	if len(newFields) > 0 {
		rowsAffected, err = Update(db, ROSTER_DB_TABLE_NAME, newFields, filters)
		if err != nil {
			fmt.Println("UpdateIncidentReport ERROR::", err)
			return rowsAffected, err
		}
	}

	// Update clients if necessary
	if incidentReport.Clients != nil {
		err = updateClientsOfIncidentReport(db, incidentReport, dbLock)
	}

	return rowsAffected, err
}

// Update a specific recipient row in the table
// This function assumes that the incidentReport recipient id is correct.
// Returns the number of rows affected and any errors.
// In this case, number of rows affected is either 0 or 1.
func UpdateIncidentReportAssignments(db *sql.DB, incidentReportContent *pb.IncidentReportContent, query *pb.IncidentReportQuery) (int64, error) {
	newFields := getFilledIncidentReportASGNFields(incidentReportContent)
	filters := getFormattedIncidentReportFilters(query, ROSTER_ASSIGNMENT_DB_TABLE_NAME, false, false)
	// filters := getIncidentReportIdFormattedFilter(
	// 	int(incidentReportAssignment.IncidentReportAssignmentId),
	// 	ROSTER_ASSIGNMENT_DB_TABLE_NAME, false,
	// )

	rowsAffected, err := Update(db, ROSTER_ASSIGNMENT_DB_TABLE_NAME, newFields, filters)

	if err != nil {
		fmt.Println("UpdateIncidentReportRecipients ERROR::", err)
		return rowsAffected, err
	}

	return rowsAffected, err
}

// Delete a particular incidentReport in the database together with
// any corresponding details.
// Details are deleted based on the cascading rule.
func DeleteIncidentReport(db *sql.DB, incidentReport *pb.IncidentReport) (int64, error) {
	filters := getIncidentReportIdFormattedFilter(
		int(incidentReport.IncidentReportId),
		ROSTER_DB_TABLE_NAME, true,
	)

	rowsAffected, err := Delete(db, ROSTER_DB_TABLE_NAME, filters)
	return rowsAffected, err
}

// Delete a particular incidentReport detail
func DeleteIncidentReportContent(db *sql.DB, incidentReportContent *pb.IncidentReportContent) (int64, error) {
	filters := getIncidentReportIdFormattedFilter(
		int(incidentReportContent.ReportContentId),
		ROSTER_ASSIGNMENT_DB_TABLE_NAME, false,
	)
	rowsAffected, err := Delete(db, ROSTER_ASSIGNMENT_DB_TABLE_NAME, filters)
	return rowsAffected, err
}
