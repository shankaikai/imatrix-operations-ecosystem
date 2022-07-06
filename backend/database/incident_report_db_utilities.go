// Utility functions for database operations related to incidentReporting.
package database

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"capstone.operations_ecosystem/backend/common"
	pb "capstone.operations_ecosystem/backend/proto"
	_ "github.com/go-sql-driver/mysql"
)

const (
	INCIDENT_REPORT_DB_TABLE_NAME         = "incident_report"
	INCIDENT_REPORT_CONTENT_DB_TABLE_NAME = "incident_report_content"

	// Incident Report table fields
	INCIDENT_REPORT_DB_ID               = "report_id"
	INCIDENT_REPORT_DB_REPORT_TYPE      = "report_type"
	INCIDENT_REPORT_DB_ORIGINAL_CONTENT = "original_content"
	INCIDENT_REPORT_DB_MODIFIED_CONTENT = "modified_content"
	INCIDENT_REPORT_DB_IS_APPROVED      = "is_approved"
	INCIDENT_REPORT_DB_SIGNATURE        = "signature"
	INCIDENT_REPORT_DB_APPROVAL_DATE    = "approval_date"

	// Incident Report content table fields
	INCIDENT_REPORT_CONTENT_DB_ID                      = "report_content_id"
	INCIDENT_REPORT_CONTENT_DB_LAST_MODIFIED_DATE      = "last_modified_date"
	INCIDENT_REPORT_CONTENT_DB_LAST_MODIFIED_USER      = "last_modifed_user"
	INCIDENT_REPORT_CONTENT_DB_ADDRESS                 = "address"
	INCIDENT_REPORT_CONTENT_DB_INCIDENT_TIME           = "incident_time"
	INCIDENT_REPORT_CONTENT_DB_TITLE                   = "title"
	INCIDENT_REPORT_CONTENT_DB_IS_POLICE_NOTIFIED      = "is_police_notified"
	INCIDENT_REPORT_CONTENT_DB_DESCRIPTION             = "description"
	INCIDENT_REPORT_CONTENT_DB_HAS_ACTION_TAKEN        = "has_action_taken"
	INCIDENT_REPORT_CONTENT_DB_ACTION_TAKEN            = "action_taken"
	INCIDENT_REPORT_CONTENT_DB_HAS_INJURY              = "has_injury"
	INCIDENT_REPORT_CONTENT_DB_INJURY_DESCRIPTION      = "injury_description"
	INCIDENT_REPORT_CONTENT_DB_HAS_STOLEN_ITEM         = "has_stolen_item"
	INCIDENT_REPORT_CONTENT_DB_STOLEN_ITEM_DESCRIPTION = "stolen_item_description"
	INCIDENT_REPORT_CONTENT_DB_REPORT_IMG              = "report_image"

	// alias for the report content tables during joins
	ORIGINAL_INCIDENT_REPORT_TABLE_ALIAS = "original"
	MODIFIED_INCIDENT_REPORT_TABLE_ALIAS = "modified"
)

// UTILITIES

// Returns the fields of the main incidentReport table
// in a specific order.
// Note that IDs are auto incremented and should
// not be modified manually. Ommits ID in resulting string.
// Does not include fields that have default values on insertion
func getIncidentReportTableFields() string {
	incidentReportTableFields := []string{
		INCIDENT_REPORT_DB_REPORT_TYPE,
		INCIDENT_REPORT_DB_ORIGINAL_CONTENT,
		INCIDENT_REPORT_DB_MODIFIED_CONTENT,
	}

	return strings.Join(incidentReportTableFields, ",")
}

// This function is highly dependent on the
// order given in getIncidentReportTableFields.
// Returns the values of the incidentReport fields in the
// order that is specified in getIncidentReportTableFields
func orderIncidentReportFields(incidentReport *pb.IncidentReport, originalContentId int64) string {
	output := ""

	output += "'" + getIncidentReportDBTypeStringFromProto(incidentReport.Type) + "'" + ", "
	output += "'" + strconv.Itoa(int(originalContentId)) + "'" + ","
	output += "'" + strconv.Itoa(int(originalContentId)) + "'"

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
		INCIDENT_REPORT_CONTENT_DB_LAST_MODIFIED_DATE,
		INCIDENT_REPORT_CONTENT_DB_LAST_MODIFIED_USER,
		INCIDENT_REPORT_CONTENT_DB_ADDRESS,
		INCIDENT_REPORT_CONTENT_DB_INCIDENT_TIME,
		INCIDENT_REPORT_CONTENT_DB_TITLE,
		INCIDENT_REPORT_CONTENT_DB_IS_POLICE_NOTIFIED,
		INCIDENT_REPORT_CONTENT_DB_DESCRIPTION,
		INCIDENT_REPORT_CONTENT_DB_HAS_ACTION_TAKEN,
		INCIDENT_REPORT_CONTENT_DB_ACTION_TAKEN,
		INCIDENT_REPORT_CONTENT_DB_HAS_INJURY,
		INCIDENT_REPORT_CONTENT_DB_INJURY_DESCRIPTION,
		INCIDENT_REPORT_CONTENT_DB_HAS_STOLEN_ITEM,
		INCIDENT_REPORT_CONTENT_DB_STOLEN_ITEM_DESCRIPTION,
		INCIDENT_REPORT_CONTENT_DB_REPORT_IMG,
	}

	return strings.Join(incidentReportRecTableFields, ",")
}

// This function is highly dependent on the
// order given in getIncidentReportRecTableFields.
// Returns the values of the incidentReport fields in the
// order that is specified in getIncidentReportRecTableFields
func orderIncidentReportContentFields(IncidentReportContent *pb.IncidentReportContent) string {
	output := ""

	output += "'" + IncidentReportContent.LastModifiedDate + "'" + ", "
	output += strconv.Itoa(int(IncidentReportContent.LastModifedUser.UserId)) + ","
	output += "'" + IncidentReportContent.Address + "'" + ", "
	output += "'" + IncidentReportContent.IncidentTime + "'" + ", "
	output += "\"" + strings.ReplaceAll(IncidentReportContent.Title, "\"", "'") + "\"" + ", "
	output += strconv.FormatBool(IncidentReportContent.IsPoliceNotified) + ", "
	output += "\"" + strings.ReplaceAll(IncidentReportContent.Description, "\"", "'") + "\"" + ", "
	output += strconv.FormatBool(IncidentReportContent.HasActionTaken) + ", "
	output += "\"" + strings.ReplaceAll(IncidentReportContent.ActionTaken, "\"", "'") + "\"" + ", "
	output += strconv.FormatBool(IncidentReportContent.HasInjury) + ", "
	output += "\"" + strings.ReplaceAll(IncidentReportContent.InjuryDescription, "\"", "'") + "\"" + ", "
	output += strconv.FormatBool(IncidentReportContent.HasStolenItem) + ", "
	output += "\"" + strings.ReplaceAll(IncidentReportContent.StolenItemDescription, "\"", "'") + "\"" + ", "
	output += "'" + IncidentReportContent.ReportImageLink + "'"

	return output
}

// This function is used in anticipation of an update query.
// Updates use the format "SET field1=val1, field2=val2".
// The return value of this function is for example: "field1=val1, field2=val2"
// ID not included in this because pk should not be manually changed.
// Note this function highly depends on the protocol buffer message definition
// Returns the string fields and values of the filled incidentReport fields
func getFilledIncidentReportFields(incidentReport *pb.IncidentReport, originalContentId int64, modifiedContentId int64) string {
	incidentReportTableFields := []string{}

	incidentReportTableFields = append(incidentReportTableFields, formatFieldEqVal(INCIDENT_REPORT_DB_REPORT_TYPE, strconv.Itoa(int(incidentReport.Type)), true))

	if originalContentId > -1 {
		incidentReportTableFields = append(incidentReportTableFields, formatFieldEqVal(INCIDENT_REPORT_DB_ORIGINAL_CONTENT, strconv.Itoa(int(originalContentId)), true))
	}
	if modifiedContentId > -1 {
		incidentReportTableFields = append(incidentReportTableFields, formatFieldEqVal(INCIDENT_REPORT_DB_MODIFIED_CONTENT, strconv.Itoa(int(modifiedContentId)), true))
	}

	incidentReportTableFields = append(incidentReportTableFields, formatFieldEqVal(INCIDENT_REPORT_DB_IS_APPROVED, strconv.FormatBool(incidentReport.IsApproved), false))

	if incidentReport.Signature != nil {
		incidentReportTableFields = append(incidentReportTableFields, formatFieldEqVal(INCIDENT_REPORT_DB_SIGNATURE, strconv.Itoa(int(incidentReport.Signature.UserId)), true))
	}

	if len(incidentReport.ApprovalDate) > 0 {
		incidentReportTableFields = append(incidentReportTableFields, formatFieldEqVal(INCIDENT_REPORT_DB_APPROVAL_DATE, incidentReport.ApprovalDate, true))
	}

	return strings.Join(incidentReportTableFields, ",")
}

func fillUpdatedIncidentReport(newReportContent *pb.IncidentReportContent, oldReportContent *pb.IncidentReportContent) error {
	if len(newReportContent.LastModifiedDate) == 0 {
		newReportContent.LastModifiedDate = time.Now().Format(common.DATETIME_FORMAT)
	}
	if len(newReportContent.Address) == 0 {
		newReportContent.Address = oldReportContent.Address
	}
	if len(newReportContent.IncidentTime) == 0 {
		newReportContent.IncidentTime = oldReportContent.IncidentTime
	}
	if len(newReportContent.Title) == 0 {
		newReportContent.Title = oldReportContent.Title
	}
	if len(newReportContent.Description) == 0 {
		newReportContent.Description = oldReportContent.Description
	}
	if len(newReportContent.ActionTaken) == 0 {
		newReportContent.ActionTaken = oldReportContent.ActionTaken
	}
	if len(newReportContent.InjuryDescription) == 0 {
		newReportContent.InjuryDescription = oldReportContent.InjuryDescription
	}
	if len(newReportContent.StolenItemDescription) == 0 {
		newReportContent.StolenItemDescription = oldReportContent.StolenItemDescription
	}
	if len(newReportContent.ReportImageLink) == 0 {
		newReportContent.ReportImageLink = oldReportContent.ReportImageLink
	}
	return nil
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
		case pb.IncidentReportFilter_REPORT_ID, pb.IncidentReportFilter_REPORT_TYPE,
			pb.IncidentReportFilter_SIGNATURE, pb.IncidentReportFilter_APPROVAL_DATE:
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
		case pb.IncidentReportFilter_IS_APPROVED:
			whereFilters = append(
				whereFilters, fmt.Sprintf("%s %s %s", incidentReportFilterToDBCol(filter.Field, table),
					GetFilterComparisonSign(filter.Comparisons.Comparison), filter.Comparisons.Value),
			)
		// This can either refer to the original or the new content
		case pb.IncidentReportFilter_MODIFIER, pb.IncidentReportFilter_LAST_MODIFIED_DATE,
			pb.IncidentReportFilter_REPORT_CONTENT_ID:
			if table == INCIDENT_REPORT_DB_TABLE_NAME {
				whereFilters = append(
					whereFilters, fmt.Sprintf("(%s.%s %s '%s' OR %s.%s %s '%s')",
						ORIGINAL_INCIDENT_REPORT_TABLE_ALIAS, incidentReportFilterToDBCol(filter.Field, table),
						GetFilterComparisonSign(filter.Comparisons.Comparison), filter.Comparisons.Value,
						MODIFIED_INCIDENT_REPORT_TABLE_ALIAS, incidentReportFilterToDBCol(filter.Field, table),
						GetFilterComparisonSign(filter.Comparisons.Comparison), filter.Comparisons.Value),
				)
			} else {
				whereFilters = append(
					whereFilters, fmt.Sprintf("%s %s '%s'", incidentReportFilterToDBCol(filter.Field, table),
						GetFilterComparisonSign(filter.Comparisons.Comparison), filter.Comparisons.Value),
				)
			}
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
			output += fmt.Sprintf(" %s %s.%s %s", ORDER_BY_KEYWORD, MODIFIED_INCIDENT_REPORT_TABLE_ALIAS, incidentReportFilterToDBCol(pb.IncidentReportFilter_LAST_MODIFIED_DATE, table), DESC_KEYWORD)
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
	retrievedUsers := make(map[int64]*pb.User)

	for rows.Next() {
		incidentReport := &pb.IncidentReport{}
		originalIncidentReportContent := &pb.IncidentReportContent{}
		modifiedIncidentReportContent := &pb.IncidentReportContent{}

		// main table
		// Protocol buffer enums
		reportType := ""
		// Redundant Strings
		originalContentPk := -1
		modifiedContentPk := -1
		// nullable content
		var signature sql.NullInt64
		var approvalDate sql.NullString

		// original content
		originalContentUser := -1
		// nullable content
		var originalContentActionTaken sql.NullString
		var originalContentInjury sql.NullString
		var originalContentStolenItem sql.NullString
		var originalContentImg sql.NullString

		// modified content
		modifiedContentUser := -1
		// nullable content
		var modifiedContentActionTaken sql.NullString
		var modifiedContentInjury sql.NullString
		var modifiedContentStolenItem sql.NullString
		var modifiedContentImg sql.NullString

		// cast each row to a incidentReport
		err := rows.Scan(
			// Main IncidentReport
			&incidentReport.IncidentReportId,
			&reportType,
			&originalContentPk,
			&modifiedContentPk,
			&incidentReport.IsApproved,
			&signature,
			&approvalDate,

			// Original IncidentReport Details
			&originalIncidentReportContent.ReportContentId,
			&originalIncidentReportContent.LastModifiedDate,
			&originalContentUser,
			&originalIncidentReportContent.Address,
			&originalIncidentReportContent.IncidentTime,
			&originalIncidentReportContent.Title,
			&originalIncidentReportContent.IsPoliceNotified,
			&originalIncidentReportContent.Description,
			&originalIncidentReportContent.HasActionTaken,
			&originalContentActionTaken,
			&originalIncidentReportContent.HasInjury,
			&originalContentInjury,
			&originalIncidentReportContent.HasStolenItem,
			&originalContentStolenItem,
			&originalContentImg,

			// Modified IncidentReport Details
			&modifiedIncidentReportContent.ReportContentId,
			&modifiedIncidentReportContent.LastModifiedDate,
			&modifiedContentUser,
			&modifiedIncidentReportContent.Address,
			&modifiedIncidentReportContent.IncidentTime,
			&modifiedIncidentReportContent.Title,
			&modifiedIncidentReportContent.IsPoliceNotified,
			&modifiedIncidentReportContent.Description,
			&modifiedIncidentReportContent.HasActionTaken,
			&modifiedContentActionTaken,
			&modifiedIncidentReportContent.HasInjury,
			&modifiedContentInjury,
			&modifiedIncidentReportContent.HasStolenItem,
			&modifiedContentStolenItem,
			&modifiedContentImg,
		)

		if err != nil {
			fmt.Println("convertDbRowsToFullIncidentReport ERROR:", err)
			continue
		}

		// main report
		incidentReport.Type = getIncidentReportTypeProtoTypeStringFromDB(reportType)
		// nullable content
		if signature.Valid {
			incidentReport.Signature, err = getUserFromCache(db, &retrievedUsers, signature.Int64)
			if err != nil {
				fmt.Println("convertDbRowsToFullIncidentReport ERROR:", err)
				continue
			}
		}

		if approvalDate.Valid {
			incidentReport.ApprovalDate = approvalDate.String
		}

		// fields that are necessary regardless of content type returned
		originalIncidentReportContent.LastModifedUser, err = getUserFromCache(db, &retrievedUsers, int64(originalContentUser))
		if err != nil {
			fmt.Println("convertDbRowsToFullIncidentReport ERROR:", err)
			continue
		}

		modifiedIncidentReportContent.LastModifedUser, err = getUserFromCache(db, &retrievedUsers, int64(modifiedContentUser))
		if err != nil {
			fmt.Println("convertDbRowsToFullIncidentReport ERROR:", err)
			continue
		}

		// Fill up high level incident reporting fields
		incidentReport.Creator = originalIncidentReportContent.LastModifedUser
		incidentReport.CreationDate = originalIncidentReportContent.LastModifiedDate
		incidentReport.LastModifedUser = modifiedIncidentReportContent.LastModifedUser
		incidentReport.LastModifiedDate = modifiedIncidentReportContent.LastModifiedDate

		// Get the original content the user wants it or
		// if there is no modified content
		// there is no modified content if the original and the modified are the same
		if wantsOriginalIncidentReportContent(query) ||
			(originalIncidentReportContent.ReportContentId == modifiedIncidentReportContent.ReportContentId) {
			// get the original content details
			if originalContentActionTaken.Valid {
				originalIncidentReportContent.ActionTaken = originalContentActionTaken.String
			}
			if originalContentInjury.Valid {
				originalIncidentReportContent.InjuryDescription = originalContentInjury.String
			}
			if originalContentStolenItem.Valid {
				originalIncidentReportContent.StolenItemDescription = originalContentStolenItem.String
			}
			if originalContentImg.Valid {
				originalIncidentReportContent.ReportImageLink = originalContentImg.String
			}

			incidentReport.IsOriginal = true
			incidentReport.IncidentReportContent = originalIncidentReportContent

		} else {
			// get the modified content details
			if modifiedContentActionTaken.Valid {
				modifiedIncidentReportContent.ActionTaken = modifiedContentActionTaken.String
			}
			if modifiedContentInjury.Valid {
				modifiedIncidentReportContent.InjuryDescription = modifiedContentInjury.String
			}
			if modifiedContentStolenItem.Valid {
				modifiedIncidentReportContent.StolenItemDescription = modifiedContentStolenItem.String
			}
			if modifiedContentImg.Valid {
				modifiedIncidentReportContent.ReportImageLink = modifiedContentImg.String
			}
			incidentReport.IncidentReportContent = modifiedIncidentReportContent
		}
		*incidentReports = append(*incidentReports, incidentReport)
	}

	if query.Skip > 0 {
		*incidentReports = (*incidentReports)[query.Skip:]
	}

	return nil
}

// This function creates the filter required if
// the only condition is a matching incidentReport id.
func getIncidentReportIdFormattedFilter(incidentReportId int, table string) string {
	query := &pb.IncidentReportQuery{}
	if table == INCIDENT_REPORT_DB_TABLE_NAME {
		AddIncidentReportFilter(query, pb.IncidentReportFilter_REPORT_ID, pb.Filter_EQUAL, strconv.Itoa(incidentReportId))
	} else {
		AddIncidentReportFilter(query, pb.IncidentReportFilter_REPORT_CONTENT_ID, pb.Filter_EQUAL, strconv.Itoa(incidentReportId))

	}
	return getFormattedIncidentReportFilters(query, table, false, false)
}

func incidentReportFilterToDBCol(filterField pb.IncidentReportFilter_Field, table string) string {
	output := ""
	switch filterField {
	case pb.IncidentReportFilter_REPORT_ID:
		output = INCIDENT_REPORT_DB_ID
	case pb.IncidentReportFilter_REPORT_CONTENT_ID:
		output = INCIDENT_REPORT_CONTENT_DB_ID
	case pb.IncidentReportFilter_REPORT_TYPE:
		output = INCIDENT_REPORT_DB_REPORT_TYPE
	case pb.IncidentReportFilter_MODIFIER:
		output = INCIDENT_REPORT_CONTENT_DB_LAST_MODIFIED_USER
	case pb.IncidentReportFilter_LAST_MODIFIED_DATE:
		output = INCIDENT_REPORT_CONTENT_DB_LAST_MODIFIED_DATE
	case pb.IncidentReportFilter_SIGNATURE:
		output = INCIDENT_REPORT_DB_SIGNATURE
	case pb.IncidentReportFilter_APPROVAL_DATE:
		output = INCIDENT_REPORT_DB_APPROVAL_DATE
	case pb.IncidentReportFilter_IS_APPROVED:
		output = INCIDENT_REPORT_DB_IS_APPROVED
	}

	return output
}

// Returns the Broadcast Urgency Type as expected in the DB from the protobuf enum
func getIncidentReportDBTypeStringFromProto(reportType pb.IncidentReport_ReportType) string {
	switch reportType {
	case pb.IncidentReport_FIRE_ALARM:
		return "Fire Alarm"
	case pb.IncidentReport_INTRUDER:
		return "Intruder"
	default:
		return "Others"
	}
}

// Returns the Broadcast Urgency Type as expected in the protobuf enum from the DB enum
func getIncidentReportTypeProtoTypeStringFromDB(reportType string) pb.IncidentReport_ReportType {
	switch reportType {
	case "Fire Alarm":
		return pb.IncidentReport_FIRE_ALARM
	case "Intruder":
		return pb.IncidentReport_INTRUDER
	default:
		return pb.IncidentReport_OTHERS
	}
}

// Loops through the incident report query to see
// if the original content is wanted
func wantsOriginalIncidentReportContent(query *pb.IncidentReportQuery) bool {
	for _, filter := range query.Filters {
		if filter.Field == pb.IncidentReportFilter_GET_ORIGINAL {
			return true
		}
	}
	return false
}
