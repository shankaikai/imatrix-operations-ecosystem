// Utility functions for database operations related to broadcasting.
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
	BC_DB_URGENCY       = "urgency"

	// Broadcast recipients table fields
	// Broadcast table fields
	BC_REC_DB_ID           = "broadcast_recipients_id"
	BC_REC_DB_RELATED_BC   = "related_broadcast"
	BC_REC_DB_RECIPIENT    = "recipient"
	BC_REC_DB_ACK          = "acknowledged"
	BC_REC_DB_REJECTION    = "rejected"
	BC_REC_DB_LAST_REPLIED = "last_replied"
	BC_REC_DB_AIDS_ID      = "aifs_id"
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
		BC_DB_URGENCY,
	}

	return strings.Join(broadcastTableFields, ",")
}

// This function is highly dependent on the
// order given in getBroadcastTableFields.
// Returns the values of the broadcast fields in the
// order that is specified in getBroadcastTableFields
func orderBroadcastFields(broadcast *pb.Broadcast) string {
	output := ""

	output += "'" + getBroadcastDBTypeStringFromProto(broadcast.Type) + "'" + ", "
	output += "'" + broadcast.Title + "'" + ", "
	output += "'" + broadcast.Content + "'" + ", "
	output += "'" + broadcast.CreationDate.AsTime().Format(DATETIME_FORMAT) + "'" + ", "
	output += "'" + broadcast.Deadline.AsTime().Format(DATETIME_FORMAT) + "'" + ", "
	output += "'" + strconv.Itoa(int(broadcast.Creator.UserId)) + "'" + ", "
	output += "'" + getBroadcastDBUrgencyStringFromProto(broadcast.Urgency) + "'"

	return output
}

// Returns the fields of the broadcast recipeints table
// in a specific order.
// Note that IDs are auto incremented and should
// not be modified manually. Ommits ID in resulting string.
// All nullable columns that we do not expect to be used upon
// first insert shall be omitted.
func getBroadcastRecTableFields() string {
	broadcastRecTableFields := []string{
		BC_REC_DB_RELATED_BC,
		BC_REC_DB_RECIPIENT,
		BC_REC_DB_ACK,
		BC_REC_DB_REJECTION,
		BC_REC_DB_AIDS_ID,
	}

	return strings.Join(broadcastRecTableFields, ",")
}

// This function is highly dependent on the
// order given in getBroadcastRecTableFields.
// Returns the values of the broadcast fields in the
// order that is specified in getBroadcastRecTableFields
func orderBroadcastRecFields(recipeint *pb.BroadcastRecipient, relatedBCId int64) string {
	output := ""

	output += strconv.Itoa(int(relatedBCId)) + ","
	output += strconv.Itoa(int(recipeint.Recipient.UserId)) + ","

	// Ack and rejection are fale by default.
	output += "0, 0" + ","
	output += strconv.Itoa(int(recipeint.AifsId))

	return output
}

// This function is used in anticipation of an update query.
// Updates use the format "SET field1=val1, field2=val2".
// The return value of this function is for example: "field1=val1, field2=val2"
// ID not included in this because pk should not be manually changed.
// Note this function highly depends on the protocol buffer message definition
// Returns the string fields and values of the filled broadcast fields
func getFilledBroadcastFields(broadcast *pb.Broadcast) string {
	broadcastTableFields := []string{formatFieldEqVal(BC_DB_TYPE, getBroadcastDBTypeStringFromProto(broadcast.Type), true)}

	if len(broadcast.Title) > 0 {
		broadcastTableFields = append(broadcastTableFields, formatFieldEqVal(BC_DB_TITLE, broadcast.Title, true))
	}
	if len(broadcast.Content) > 0 {
		broadcastTableFields = append(broadcastTableFields, formatFieldEqVal(BC_DB_CONTENT, broadcast.Content, true))
	}
	if broadcast.CreationDate != nil {
		broadcastTableFields = append(broadcastTableFields, formatFieldEqVal(BC_DB_CREATION_DATE, broadcast.CreationDate.AsTime().Format(DATETIME_FORMAT), true))
	}
	if broadcast.Deadline != nil {
		broadcastTableFields = append(broadcastTableFields, formatFieldEqVal(BC_DB_DEADLINE, broadcast.Deadline.AsTime().Format(DATETIME_FORMAT), true))
	}
	if broadcast.Creator != nil {
		broadcastTableFields = append(broadcastTableFields, formatFieldEqVal(BC_DB_CREATOR, strconv.Itoa(int(broadcast.Creator.UserId)), true))
	}
	broadcastTableFields = append(broadcastTableFields, formatFieldEqVal(BC_DB_URGENCY, getBroadcastDBUrgencyStringFromProto(broadcast.Urgency), true))

	return strings.Join(broadcastTableFields, ",")
}

// This function is used in anticipation of an update query.
// Updates use the format "SET field1=val1, field2=val2".
// The return value of this function is for example: "field1=val1, field2=val2"
// ID not included in this because pk should not be manually changed.
// Note this function highly depends on the protocol buffer message definition
// Returns the string fields and values of the filled broadcast recipinets fields
func getFilledBroadcastRecFields(bRec *pb.BroadcastRecipient) string {
	broadcastTableFields := []string{}

	if bRec.Recipient != nil {
		broadcastTableFields = append(broadcastTableFields, formatFieldEqVal(BC_REC_DB_RECIPIENT, strconv.Itoa(int(bRec.Recipient.UserId)), true))
	}

	broadcastTableFields = append(broadcastTableFields, formatFieldEqVal(BC_REC_DB_ACK, strconv.FormatBool(bRec.Acknowledged), false))
	broadcastTableFields = append(broadcastTableFields, formatFieldEqVal(BC_REC_DB_REJECTION, strconv.FormatBool(bRec.Rejected), false))

	if bRec.LastReplied != nil {
		broadcastTableFields = append(broadcastTableFields, formatFieldEqVal(BC_REC_DB_LAST_REPLIED, bRec.LastReplied.AsTime().Format(DATETIME_FORMAT), true))
	}

	broadcastTableFields = append(broadcastTableFields, formatFieldEqVal(BC_REC_DB_AIDS_ID, strconv.Itoa(int(bRec.AifsId)), true))

	return strings.Join(broadcastTableFields, ",")
}

// Returns the Broadcast Type as expected in the DB from the protobuf enum
func getBroadcastDBTypeStringFromProto(bcType pb.Broadcast_BroadcastType) string {
	switch bcType {
	case pb.Broadcast_ANNOUNCEMENT:
		return "Announcement"
	default:
		return "Assignment"
	}
}

// Returns the Broadcast Type as expected in the protobuf enum from the DB enum
func getBroadcastProtoTypeStringFromDB(bcType string) pb.Broadcast_BroadcastType {
	switch bcType {
	case "Announcement":
		return pb.Broadcast_ANNOUNCEMENT
	default:
		return pb.Broadcast_ASSIGNMENT
	}
}

// Returns the Broadcast Urgency Type as expected in the DB from the protobuf enum
func getBroadcastDBUrgencyStringFromProto(bcUrgencyType pb.Broadcast_UrgencyType) string {
	switch bcUrgencyType {
	case pb.Broadcast_LOW:
		return "Low"
	case pb.Broadcast_MEDIUM:
		return "Medium"
	default:
		return "High"
	}
}

// Returns the Broadcast Urgency Type as expected in the protobuf enum from the DB enum
func getBroadcastUrgencyProtoTypeStringFromDB(bcUrgencyType string) pb.Broadcast_UrgencyType {
	switch bcUrgencyType {
	case "Low":
		return pb.Broadcast_LOW
	case "Medium":
		return pb.Broadcast_MEDIUM
	default:
		return pb.Broadcast_HIGH
	}
}

// Helper function to add a new filter to the list of existing
// filters in a broadcast query struct.
// Modifies the broadcast query parameter directly.
func addBroadcastFilter(query *pb.BroadcastQuery, field pb.BroadcastFilter_Field,
	comparison pb.Filter_Comparisons,
	value string) {
	if query.Filters == nil {
		query.Filters = make([]*pb.BroadcastFilter, 0)
	}
	filter := &pb.Filter{Comparison: comparison, Value: value}
	query.Filters = append(query.Filters, &pb.BroadcastFilter{Field: field, Comparisons: filter})
}

// Converts the filters in the broadcast array into a formatted where clause
// that can be parsed into MySQL. If a limit is needed, the LIMIT filter is
// added to the end of the string.
// For example returns: "WHERE id=22 AND num <2 LIMIT 5"
// Returns the formatted SQL filter string.
func getFormattedBroadcastFilters(query *pb.BroadcastQuery, table string, needLimit bool, needOrder bool) string {
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
		case pb.BroadcastFilter_BROADCAST_ID, pb.BroadcastFilter_TYPE, pb.BroadcastFilter_TITLE,
			pb.BroadcastFilter_CONTENT, pb.BroadcastFilter_CREATION_DATE, pb.BroadcastFilter_DEADLINE,
			pb.BroadcastFilter_CREATOR_ID, pb.BroadcastFilter_RECEIPEIENT_ID, pb.BroadcastFilter_URGENCY,
			pb.BroadcastFilter_AIFS_ID:
			if hasQuotes {
				whereFilters = append(

					whereFilters, fmt.Sprintf("%s %s '%s'", bcFilterToDBCol(filter.Field, table),
						GetFilterComparisonSign(filter.Comparisons.Comparison), filter.Comparisons.Value),
				)
			} else {
				whereFilters = append(
					whereFilters, fmt.Sprintf("%s %s %s", bcFilterToDBCol(filter.Field, table),
						GetFilterComparisonSign(filter.Comparisons.Comparison), filter.Comparisons.Value),
				)
			}
		case pb.BroadcastFilter_NUM_RECEIPIENTS:
			groupBy = append(groupBy, BC_DB_ID)
			haveFilters = append(haveFilters, fmt.Sprintf("COUNT(%s) > %s", BC_DB_ID, filter.Comparisons.Value))
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
			output += fmt.Sprintf(" %s %s %s", ORDER_BY_KEYWORD, bcFilterToDBCol(query.OrderBy.Field, table), orderByProtoToDB(query.OrderBy.OrderBy))
		} else if table == BROADCAST_DB_TABLE_NAME {
			// By default we order broadcasts by the creation date
			output += fmt.Sprintf(" %s %s %s", ORDER_BY_KEYWORD, bcFilterToDBCol(pb.BroadcastFilter_CREATION_DATE, table), DESC_KEYWORD)
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

// Get the user corresponding to a particular user id in the db
func idUserByUserId(db *sql.DB, userId int) (*pb.User, error) {
	userQuery := &pb.UserQuery{Limit: 1}
	addUserFilter(userQuery, pb.UserFilter_USER_ID, pb.Filter_EQUAL, strconv.Itoa(userId))

	users, err := GetUsers(db, userQuery)

	user := &pb.User{}

	if err == nil {
		user = users[0]
	}

	return user, err
}

// This function converts the returned DB rows into Broadcast objects and
// their corresponding broadcast recipients.
// These rows come from the join query of both the broadcast and broadcast
// recipients table.
// Modifies the broadcast array in place.
func convertDbRowsToBcNBcR(db *sql.DB, broadcasts *[]*pb.Broadcast, rows *sql.Rows, query *pb.BroadcastQuery) error {
	broadcastMap := make(map[int64]*pb.Broadcast)

	for rows.Next() {
		broadcast := &pb.Broadcast{}
		broadcastRecipient := &pb.BroadcastRecipient{}

		creatorUserId := -1
		broadcastType := ""
		creationDateStr := ""
		deadlineStr := ""
		recipientUserId := -1
		urgencyType := ""
		relatedBroadcast := ""
		var lastRepliedString sql.NullString

		// cast each row to a broadcast
		err := rows.Scan(
			&broadcast.BroadcastId,
			&broadcastType,
			&broadcast.Title,
			&broadcast.Content,
			&creationDateStr,
			&deadlineStr,
			&creatorUserId,
			&urgencyType,
			&broadcastRecipient.BroadcastRecipientsId,
			&relatedBroadcast,
			&recipientUserId,
			&broadcastRecipient.Acknowledged,
			&broadcastRecipient.Rejected,
			&lastRepliedString,
			&broadcastRecipient.AifsId,
		)

		if err != nil {
			fmt.Println("convertDbRowsToBcNBcR ERROR:", err)
			continue
		}

		// Check if there was already a broadcast found with this id
		if existingBroadcast, ok := broadcastMap[broadcast.BroadcastId]; ok {
			broadcast = existingBroadcast
		} else {
			// Return only the necessary number of broadcasts.
			// If the number of broadcasts have reached the limit,
			// do not add new broadcasts.
			if int64(len(broadcastMap)) >= query.Limit+query.Skip {
				continue
			}
			broadcast.Type = getBroadcastProtoTypeStringFromDB(broadcastType)
			creator, err := idUserByUserId(db, creatorUserId)
			broadcast.Urgency = getBroadcastUrgencyProtoTypeStringFromDB(urgencyType)

			if err != nil {
				fmt.Println("GetBroadcasts ERROR:", err)
				continue
			}

			broadcast.Creator = creator

			broadcast.CreationDate, err = DBDatetimeToPB(creationDateStr)
			if err != nil {
				fmt.Println("GetBroadcasts:", err.Error())
				continue
			}

			broadcast.Deadline, err = DBDatetimeToPB(deadlineStr)
			if err != nil {
				fmt.Println("GetBroadcasts:", err.Error())
				continue
			}

			broadcastMap[broadcast.BroadcastId] = broadcast
		}

		broadcastRecipient.Recipient, err = idUserByUserId(db, recipientUserId)
		if err != nil {
			fmt.Println("GetBroadcasts:", err.Error())
			continue
		}

		if lastRepliedString.Valid {
			broadcastRecipient.LastReplied, err = DBDatetimeToPB(lastRepliedString.String)
			if err != nil {
				fmt.Println("GetBroadcasts:", err.Error())
				continue
			}
		}

		// Add recipient to broadcast
		if broadcast.Recipients == nil {
			broadcast.Recipients = make([]*pb.BroadcastRecipient, 0)
		}

		broadcast.Recipients = append(broadcast.Recipients, broadcastRecipient)
	}

	// Add all broadcasts to the returning array

	for _, broadcast := range broadcastMap {
		*broadcasts = append(*broadcasts, broadcast)
	}

	sort.Slice(*broadcasts, func(i, j int) bool {
		return (*broadcasts)[i].CreationDate.Seconds > (*broadcasts)[j].CreationDate.Seconds
	})

	if query.Skip > 0 {
		*broadcasts = (*broadcasts)[query.Skip:]
	}

	return nil
}

// This function is different from UpdateBroadcastRecipients()
// in the idea that this function finds out who the existing
// recipients are and make the necessary changes so that the
// recipients of the main broadcast will corresponds to the new
// list that is needed. Ie, it inserts and deletes recipients at will.
func updateRecipientsOfBroadcast(db *sql.DB, broadcast *pb.Broadcast, query *pb.BroadcastQuery, dbLock *sync.Mutex) error {
	// Get all recipients
	currentRecipients, err := GetBroadcastRecipients(db, query, broadcast.BroadcastId)
	if err != nil {
		fmt.Println("UpdateBroadcast ERROR::", err)
		return err
	}

	// create array of current broadcast recipient ids
	currentRecIds := make([]int, 0)

	for _, br := range currentRecipients {
		currentRecIds = append(currentRecIds, int(br.Recipient.UserId))
	}

	sort.Ints(currentRecIds)

	// Check if the updated recipients exist within the current ones
	// If they exist, ignore and remove them from the current list.
	// If they do not exist, we need to add them to the db.
	// If the current recipient is not within the list of new recipients
	// delete this rogue recipient.

	// Index of missing recipients from the input broadcast list
	missingRecIndex := make([]int, 0)

	for i, br := range broadcast.Recipients {
		found, index := common.BinarySearch(currentRecIds, 0, len(currentRecIds)-1, int(br.Recipient.UserId))
		if found {
			fmt.Println("Found updated recipient in current recipient")
			currentRecIds = append(currentRecIds[:index], currentRecIds[index+1:]...)
		} else {
			missingRecIndex = append(missingRecIndex, i)
		}
	}

	fmt.Println("Missing Recipients Index:", missingRecIndex)
	// Add the missing recipients
	for _, recIndex := range missingRecIndex {
		_, err := InsertBroadcastRecipient(db, broadcast.Recipients[recIndex], broadcast.BroadcastId, dbLock)
		if err != nil {
			fmt.Println("UpdateBroadcast ERROR::", err)
			return err
		}
	}

	fmt.Println("Deleting Recipients IDs:", currentRecIds)
	// See if any need to be deleted
	for _, id := range currentRecIds {
		_, err := DeleteBroadcastRecipients(db, &pb.BroadcastRecipient{BroadcastRecipientsId: int64(id)})
		if err != nil {
			fmt.Println("UpdateBroadcast ERROR::", err)
			return err
		}
	}

	return nil
}

// This function creates the filter required if
// the only condition is a matching broadcast id.
func getBroadcastIdFormattedFilter(broadcastId int, table string) string {
	query := &pb.BroadcastQuery{}
	addBroadcastFilter(query, pb.BroadcastFilter_BROADCAST_ID, pb.Filter_EQUAL, strconv.Itoa(broadcastId))
	return getFormattedBroadcastFilters(query, table, false, false)
}

func bcFilterToDBCol(filterField pb.BroadcastFilter_Field, table string) string {
	output := ""
	switch filterField {
	case pb.BroadcastFilter_BROADCAST_ID:
		if table == BROADCAST_DB_TABLE_NAME {
			output = BC_DB_ID
		} else {
			output = BC_REC_DB_RELATED_BC
		}
	case pb.BroadcastFilter_TYPE:
		output = BC_DB_TYPE
	case pb.BroadcastFilter_TITLE:
		output = BC_DB_TITLE
	case pb.BroadcastFilter_CONTENT:
		output = BC_DB_CONTENT
	case pb.BroadcastFilter_CREATION_DATE:
		output = BC_DB_CREATION_DATE
	case pb.BroadcastFilter_DEADLINE:
		output = BC_DB_DEADLINE
	case pb.BroadcastFilter_CREATOR_ID:
		output = BC_DB_CREATOR
	case pb.BroadcastFilter_RECEIPEIENT_ID:
		output = BC_REC_DB_RECIPIENT
	case pb.BroadcastFilter_URGENCY:
		output = BC_DB_URGENCY
	case pb.BroadcastFilter_AIFS_ID:
		output = BC_REC_DB_AIDS_ID
	}

	return output
}