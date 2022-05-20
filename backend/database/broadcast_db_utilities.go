// Utility functions for database operations related to broadcasting.
package database

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	pb "capstone.operations_ecosystem/backend/proto"
	_ "github.com/go-sql-driver/mysql"
)

const (
	BROADCAST_DB_TABLE_NAME        = "broadcast"
	BROADCAST_RECIPIENT_TABLE_NAME = "broadcast_recepients"

	// Broadcast table fields
	BC_DB_ID            = "broadcast_id"
	BC_DB_TYPE          = "type"
	BC_DB_TITLE         = "title"
	BC_DB_CONTENT       = "content"
	BC_DB_CREATION_DATE = "creation_date"
	BC_DB_DEADLINE      = "deadline"
	BC_DB_CREATOR       = "creator"

	// Broadcast recipients table fields
	// Broadcast table fields
	BC_REC_DB_ID         = "broadcast_recipients_id"
	BC_REC_DB_RELATED_BC = "related_broadcast"
	BC_REC_DB_RECIPIENT  = "recipient"
	BC_REC_DB_ACK        = "acknowledged"
	BC_REC_DB_REJECTION  = "rejected"
)

// UTILITIES

// Returns the fields of the main broadcast table
// in a specific order.
// Note that IDs are auto incremented and should
// not be modified manually. Ommits ID in resulting string.
func getBroadcastTableFields() string {
	broadcastTableFields := []string{
		BC_DB_TYPE,
		BC_DB_TITLE,
		BC_DB_CONTENT,
		BC_DB_CREATION_DATE,
		BC_DB_DEADLINE,
		BC_DB_CREATOR,
	}

	return strings.Join(broadcastTableFields, ",")
}

// This function is highly dependent on the
// order given in getBroadcastTableFields.
func orderBroadcastFields(broadcast *pb.Broadcast) string {
	output := ""

	output += "'" + getBroadcastDBTypeStringFromProto(broadcast.Type) + "'" + ", "
	output += "'" + broadcast.Title + "'" + ", "
	output += "'" + broadcast.Content + "'" + ", "
	output += "'" + broadcast.CreationDate.AsTime().Format("2006-01-02 15:04:05") + "'" + ", "
	output += "'" + broadcast.Deadline.AsTime().Format("2006-01-02 15:04:05") + "'" + ", "
	output += "'" + strconv.Itoa(int(broadcast.Creator.UserId)) + "'"

	return output
}

// Returns the fields of the broadcast receipeints table
// in a specific order.
// Note that IDs are auto incremented and should
// not be modified manually. Ommits ID in resulting string.
func getBroadcastRecTableFields() string {
	broadcastRecTableFields := []string{
		BC_REC_DB_RELATED_BC,
		BC_REC_DB_RECIPIENT,
		BC_REC_DB_ACK,
		BC_REC_DB_REJECTION,
	}

	return strings.Join(broadcastRecTableFields, ",")
}

// This function is highly dependent on the
// order given in getBroadcastTableFields.
func orderBroadcastRecFields(receipeint *pb.BroadcastRecipient, relatedBCId int64) string {
	output := ""

	output += strconv.Itoa(int(relatedBCId)) + ","
	output += strconv.Itoa(int(receipeint.Recipient.UserId)) + ","

	// Ack and rejection are fale by default.
	output += "0, 0"

	return output
}

// ID not included in this because pk should not be manually changed.
// Note this function highly depends on the protocol buffer message definition
// Returns the string fields and values of the filled broadcast fields
func getFilledBroadcastFields(broadcast *pb.Broadcast) string {
	broadcastTableFields := []string{formatFieldEqVal(BC_DB_TYPE, getBroadcastDBTypeStringFromProto(broadcast.Type))}

	if len(broadcast.Title) > 0 {
		broadcastTableFields = append(broadcastTableFields, formatFieldEqVal(BC_DB_TITLE, broadcast.Title))
	}
	if len(broadcast.Content) > 0 {
		broadcastTableFields = append(broadcastTableFields, formatFieldEqVal(BC_DB_CONTENT, broadcast.Content))
	}
	if broadcast.CreationDate != nil {
		broadcastTableFields = append(broadcastTableFields, formatFieldEqVal(BC_DB_CREATION_DATE, broadcast.CreationDate.AsTime().Format(DATETIME_FORMAT)))
	}
	if broadcast.Deadline != nil {
		broadcastTableFields = append(broadcastTableFields, formatFieldEqVal(BC_DB_DEADLINE, broadcast.Deadline.AsTime().Format(DATETIME_FORMAT)))
	}
	if broadcast.Creator != nil {
		broadcastTableFields = append(broadcastTableFields, formatFieldEqVal(BC_DB_CREATOR, strconv.Itoa(int(broadcast.Creator.UserId))))
	}

	return strings.Join(broadcastTableFields, ",")
}

// Returns the Broadcast Type as expected in the DB
func getBroadcastDBTypeStringFromProto(bcType pb.Broadcast_BroadcastType) string {
	switch bcType {
	case pb.Broadcast_ANNOUNCEMENT:
		return "Announcement"
	default:
		return "Assignment"
	}
}

// Returns the Broadcast Type as expected in the DB
func getBroadcastProtoTypeStringFromDB(bcType string) pb.Broadcast_BroadcastType {
	switch bcType {
	case "Announcement":
		return pb.Broadcast_ANNOUNCEMENT
	default:
		return pb.Broadcast_ASSIGNMENT
	}
}

func addBroadcastFilter(query *pb.BroadcastQuery, field pb.BroadcastFilter_Field,
	comparison pb.Filter_Comparisons,
	value string) {
	if query.Filters == nil {
		query.Filters = make([]*pb.BroadcastFilter, 0)
	}
	filter := &pb.Filter{Comparison: comparison, Value: value}
	query.Filters = append(query.Filters, &pb.BroadcastFilter{Field: field, Comparisons: filter})
}

//TODO add more
func getFormattedBroadcastFilters(query *pb.BroadcastQuery, table string, needLimit bool) string {
	output := ""

	if len(query.Filters) > 0 {
		output += "WHERE "
	}

	// Get all filters
	whereFilters := make([]string, 0)
	groupBy := make([]string, 0)
	haveFilters := make([]string, 0)

	for _, filter := range query.Filters {
		if filter.Comparisons.Comparison == pb.Filter_CONTAINS {
			filter.Comparisons.Value = FormatLikeQueryValue(filter.Comparisons.Value)
		}
		switch filter.Field {
		case pb.BroadcastFilter_BROADCAST_ID:
			if table == BROADCAST_DB_TABLE_NAME {
				whereFilters = append(whereFilters, fmt.Sprintf("%s %s '%s'", BC_DB_ID, GetFilterComparisonSign(filter.Comparisons.Comparison), filter.Comparisons.Value))
			} else {
				whereFilters = append(whereFilters, fmt.Sprintf("%s %s '%s'", BC_REC_DB_RELATED_BC, GetFilterComparisonSign(filter.Comparisons.Comparison), filter.Comparisons.Value))
			}
		case pb.BroadcastFilter_TYPE:
			whereFilters = append(whereFilters, fmt.Sprintf("%s %s '%s'", BC_DB_TYPE, GetFilterComparisonSign(filter.Comparisons.Comparison), filter.Comparisons.Value))
		case pb.BroadcastFilter_TITLE:
			whereFilters = append(whereFilters, fmt.Sprintf("%s %s '%s'", BC_DB_TITLE, GetFilterComparisonSign(filter.Comparisons.Comparison), filter.Comparisons.Value))
		case pb.BroadcastFilter_CONTENT:
			whereFilters = append(whereFilters, fmt.Sprintf("%s %s '%s'", BC_DB_CONTENT, GetFilterComparisonSign(filter.Comparisons.Comparison), filter.Comparisons.Value))
		case pb.BroadcastFilter_CREATION_DATE:
			whereFilters = append(whereFilters, fmt.Sprintf("%s %s '%s'", BC_DB_CREATION_DATE, GetFilterComparisonSign(filter.Comparisons.Comparison), filter.Comparisons.Value))
		case pb.BroadcastFilter_DEADLINE:
			whereFilters = append(whereFilters, fmt.Sprintf("%s %s '%s'", BC_DB_DEADLINE, GetFilterComparisonSign(filter.Comparisons.Comparison), filter.Comparisons.Value))
		case pb.BroadcastFilter_CREATOR_ID:
			whereFilters = append(whereFilters, fmt.Sprintf("%s %s '%s'", BC_DB_CREATOR, GetFilterComparisonSign(filter.Comparisons.Comparison), filter.Comparisons.Value))
		case pb.BroadcastFilter_RECEIPEIENT_ID:
			whereFilters = append(whereFilters, fmt.Sprintf("%s %s '%s'", BC_REC_DB_RECIPIENT, GetFilterComparisonSign(filter.Comparisons.Comparison), filter.Comparisons.Value))

		case pb.BroadcastFilter_NUM_RECEIPIENTS:
			groupBy = append(groupBy, BC_DB_ID)
			haveFilters = append(haveFilters, "COUNT(a.broadcast_id) > 2")
		}
	}

	output += strings.Join(whereFilters, ",")

	if len(groupBy) > 0 {
		output += " GROUP BY " + strings.Join(groupBy, ",")
	}

	if len(haveFilters) > 0 {
		output += " HAVING " + strings.Join(haveFilters, ",")
	}

	// Add limits
	if needLimit {
		if query.Limit == 0 {
			query.Limit = DEFAULT_LIMIT
		}
		output += fmt.Sprintf(" LIMIT %d", query.Limit)
	}

	return output
}

// get the user's corresponding to the id in the db
func idUserByUserId(db *sql.DB, userId int) (*pb.User, error) {
	comparison := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: strconv.Itoa(userId)}
	userFilter := &pb.UserFilter{Field: pb.UserFilter_USER_ID, Comparisons: comparison}
	userQuery := &pb.UserQuery{Limit: 1, Filters: []*pb.UserFilter{userFilter}}

	users, err := GetUsers(db, userQuery)

	user := &pb.User{}

	if err == nil {
		user = users[0]
	}

	return user, err
}

func formatFieldEqVal(field string, val string) string {
	return fmt.Sprintf("%s='%s'", field, val)
}
