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
	incidentReportTbFields := getIncidentReportTableFields()
	incidentReportValues := orderIncidentReportFields(incidentReport, originalContentPk)

	incidentReportPk, err := Insert(db, INCIDENT_REPORT_DB_TABLE_NAME, incidentReportTbFields, incidentReportValues, dbLock)

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
func InsertIncidentReportContent(db *sql.DB, content *pb.IncidentReportContent, dbLock *sync.Mutex) (int64, error) {
	// get fields and values for this particular recipient
	fields := getIncidentReportContentTableFields()
	values := orderIncidentReportContentFields(content)

	// Add recipient to DB
	pk, err := Insert(db, INCIDENT_REPORT_CONTENT_DB_TABLE_NAME, fields, values, dbLock)

	return pk, err
}

// Get all the incidentReport rows in a table that meets specifications.
// Returns an array of incidentReports and any errors.
func GetIncidentReports(db *sql.DB, query *pb.IncidentReportQuery) ([]*pb.IncidentReport, error) {
	fmt.Println("Getting IncidentReports...")
	incidentReports := make([]*pb.IncidentReport, 0)

	// Join the incident Report, original and modified content tables
	// in order to easily filter conditions relating to all tables

	// Set default query limits if needed
	if query.Limit == 0 {
		query.Limit = DEFAULT_LIMIT
	}

	fields := ALL_COLS

	// tables are joined on the original and modified content table ids
	fistOnCondition := formatFieldEqVal(INCIDENT_REPORT_DB_ORIGINAL_CONTENT, ORIGINAL_INCIDENT_REPORT_TABLE_ALIAS+"."+INCIDENT_REPORT_CONTENT_DB_ID, false)
	secondOnCondition := formatFieldEqVal(INCIDENT_REPORT_DB_MODIFIED_CONTENT, MODIFIED_INCIDENT_REPORT_TABLE_ALIAS+"."+INCIDENT_REPORT_CONTENT_DB_ID, false)

	// Format filters
	filters := getFormattedIncidentReportFilters(query, INCIDENT_REPORT_DB_TABLE_NAME, true, true)

	rows, err := QueryThreeTablesLeftJoin(db, INCIDENT_REPORT_DB_TABLE_NAME,
		INCIDENT_REPORT_CONTENT_DB_TABLE_NAME+" AS "+ORIGINAL_INCIDENT_REPORT_TABLE_ALIAS,
		INCIDENT_REPORT_CONTENT_DB_TABLE_NAME+" AS "+MODIFIED_INCIDENT_REPORT_TABLE_ALIAS,
		fistOnCondition, secondOnCondition, fields, filters)

	if err != nil {
		return incidentReports, err
	}

	// convert query rows into incidentReports
	// give back the query the original limit
	err = convertDbRowsToFullIncidentReport(db, &incidentReports, rows, query)

	return incidentReports, err
}

// Get all the incidentReport recipient rows in a table that meets specifications.
// Returns an array of incidentReport recipients and any errors.
func GetIncidentReportContents(db *sql.DB, query *pb.IncidentReportQuery, mainIncidentReportID int64) ([]*pb.IncidentReportContent, error) {
	fmt.Println("Getting IncidentReports Contents...")
	incidentReportContents := make([]*pb.IncidentReportContent, 0)

	fields := ALL_COLS

	// Format filters
	// Get for a specific main incidentReport if needed
	if mainIncidentReportID != -1 {
		AddIncidentReportFilter(query, pb.IncidentReportFilter_REPORT_ID, pb.Filter_EQUAL, strconv.Itoa(int(mainIncidentReportID)))
	}

	filters := getFormattedIncidentReportFilters(query, INCIDENT_REPORT_CONTENT_DB_TABLE_NAME, true, true)

	rows, err := Query(db, INCIDENT_REPORT_CONTENT_DB_TABLE_NAME, fields, filters)

	if err != nil {
		return incidentReportContents, err
	}

	retrievedUsers := make(map[int64]*pb.User)

	// convert query rows into incidentReports assignments
	for rows.Next() {
		content := &pb.IncidentReportContent{}

		// original content
		contentUser := -1
		// nullable content
		var actionTaken sql.NullString
		var injury sql.NullString
		var stolenItem sql.NullString
		var img sql.NullString

		// cast each row to a incidentReport
		err := rows.Scan(
			// Original IncidentReport Details
			&content.ReportContentId,
			&content.LastModifiedDate,
			&contentUser,
			&content.Address,
			&content.IncidentTime,
			&content.Title,
			&content.IsPoliceNotified,
			&content.Description,
			&content.HasActionTaken,
			&actionTaken,
			&content.HasInjury,
			&injury,
			&content.HasStolenItem,
			&stolenItem,
			&img,
		)

		if err != nil {
			fmt.Println("GetIncidentReportContents ERROR:", err)
			continue
		}

		// fields that are necessary regardless of content type returned
		content.LastModifedUser, err = getUserFromCache(db, &retrievedUsers, int64(contentUser))
		if err != nil {
			fmt.Println("GetIncidentReportContents ERROR:", err)
			continue
		}

		if actionTaken.Valid {
			content.ActionTaken = actionTaken.String
		}
		if injury.Valid {
			content.InjuryDescription = injury.String
		}
		if stolenItem.Valid {
			content.StolenItemDescription = stolenItem.String
		}
		if img.Valid {
			content.ReportImageLink = img.String
		}

		incidentReportContents = append(incidentReportContents, content)
	}

	return incidentReportContents, err
}

// Update a specific incidentReport in the table
// Only fields that have been filled in the incidentReport object will be updated.
// Returns the number of main incidentReport rows updated as well as the a list of
// all the primary keys of newly inserted incidentReport assignments into the db.
// For this implementation, only the original content and the newest updated content is kept
// To update the content of the report, delete any existing non-original contents
// and update the modified content id in the original report with the new content.
func UpdateIncidentReport(db *sql.DB, incidentReport *pb.IncidentReport, dbLock *sync.Mutex) (int64, error) {
	rowsAffected := int64(0)
	newContentPk := int64(-1)
	var err error

	query := &pb.IncidentReportQuery{Limit: 1}
	AddIncidentReportFilter(query, pb.IncidentReportFilter_REPORT_ID,
		pb.Filter_EQUAL, strconv.Itoa(int(incidentReport.IncidentReportId)))

	// Insert the new incident report if any
	if incidentReport.IncidentReportContent != nil {
		newContentPk, err = InsertIncidentReportContent(db, incidentReport.IncidentReportContent, dbLock)
		if err != nil {
			return 0, err
		}

		// Get old incident report
		reports, err := GetIncidentReports(db, query)
		if err != nil {
			// Delete the new content that was just inserted
			DeleteIncidentReportContent(db, &pb.IncidentReportContent{ReportContentId: newContentPk})
			return 0, err
		}

		// check if the content has any existing modified content

		if len(reports) > 0 && !reports[0].IsOriginal {
			// Delete the old incident report content
			_, err := DeleteIncidentReportContent(db, reports[0].IncidentReportContent)
			if err != nil {
				// Delete the new content that was just inserted
				DeleteIncidentReportContent(db, &pb.IncidentReportContent{ReportContentId: newContentPk})
				return 0, err
			}
		}

	}

	// Update the main incidentReport
	newFields := getFilledIncidentReportFields(incidentReport, -1, int64(newContentPk))
	filters := getFormattedIncidentReportFilters(query, INCIDENT_REPORT_DB_TABLE_NAME, false, false)

	if len(newFields) > 0 {
		rowsAffected, err = Update(db, INCIDENT_REPORT_DB_TABLE_NAME, newFields, filters)
		if err != nil {
			fmt.Println("UpdateIncidentReport ERROR::", err)
			return rowsAffected, err
		}
	}

	return rowsAffected, err
}

// Delete a particular incidentReport in the database together with
// any corresponding details.
// Details are deleted based on the cascading rule.
func DeleteIncidentReport(db *sql.DB, incidentReport *pb.IncidentReport) (int64, error) {
	filters := getIncidentReportIdFormattedFilter(
		int(incidentReport.IncidentReportId),
		INCIDENT_REPORT_DB_TABLE_NAME,
	)

	rowsAffected, err := Delete(db, INCIDENT_REPORT_DB_TABLE_NAME, filters)
	return rowsAffected, err
}

// Delete a particular incidentReport detail
func DeleteIncidentReportContent(db *sql.DB, incidentReportContent *pb.IncidentReportContent) (int64, error) {
	filters := getIncidentReportIdFormattedFilter(
		int(incidentReportContent.ReportContentId),
		INCIDENT_REPORT_CONTENT_DB_TABLE_NAME,
	)
	rowsAffected, err := Delete(db, INCIDENT_REPORT_CONTENT_DB_TABLE_NAME, filters)
	return rowsAffected, err
}
