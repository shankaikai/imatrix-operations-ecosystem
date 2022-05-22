// Utility functions for database operations related to broadcasting.
package database

import (
	"database/sql"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"capstone.operations_ecosystem/backend/common"
	pb "capstone.operations_ecosystem/backend/proto"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/protobuf/types/known/timestamppb"
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

// Returns the fields of the broadcast recipeints table
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
func orderBroadcastRecFields(recipeint *pb.BroadcastRecipient, relatedBCId int64) string {
	output := ""

	output += strconv.Itoa(int(relatedBCId)) + ","
	output += strconv.Itoa(int(recipeint.Recipient.UserId)) + ","

	// Ack and rejection are fale by default.
	output += "0, 0"

	return output
}

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

	return strings.Join(broadcastTableFields, ",")
}

// ID not included in this because pk should not be manually changed.
// Note this function highly depends on the protocol buffer message definition
// Returns the string fields and values of the filled broadcast fields
func getFilledBroadcastRecFields(bRec *pb.BroadcastRecipient) string {
	broadcastTableFields := []string{}

	if bRec.Recipient != nil {
		broadcastTableFields = append(broadcastTableFields, formatFieldEqVal(BC_REC_DB_RECIPIENT, strconv.Itoa(int(bRec.Recipient.UserId)), true))
	}

	broadcastTableFields = append(broadcastTableFields, formatFieldEqVal(BC_REC_DB_ACK, strconv.FormatBool(bRec.Acknowledged), false))
	broadcastTableFields = append(broadcastTableFields, formatFieldEqVal(BC_REC_DB_REJECTION, strconv.FormatBool(bRec.Rejected), false))

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

func getFormattedBroadcastFilters(query *pb.BroadcastQuery, table string, needLimit bool) string {
	output := ""

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
			haveFilters = append(haveFilters, fmt.Sprintf("COUNT(%s) > %s", BC_DB_ID, filter.Comparisons.Value))
		}
	}

	if len(whereFilters) > 0 {
		output += "WHERE "
	}

	output += strings.Join(whereFilters, " AND ")

	if len(groupBy) > 0 {
		output += " GROUP BY " + strings.Join(groupBy, ",")
	}

	if len(haveFilters) > 0 {
		output += " HAVING " + strings.Join(haveFilters, " AND ")
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

func convertDbRowsToBcNBcR(db *sql.DB, broadcasts *[]*pb.Broadcast, rows *sql.Rows, queryLimit int) error {
	broadcastMap := make(map[int64]*pb.Broadcast)

	for rows.Next() {
		broadcast := &pb.Broadcast{}
		broadcastRecipient := &pb.BroadcastRecipient{}

		creatorUserId := -1
		broadcastType := ""
		creationDateStr := ""
		deadlineStr := ""
		recipientUserId := -1
		relatedBroadcast := ""

		// cast each row to a broadcast
		err := rows.Scan(
			&broadcast.BroadcastId,
			&broadcastType,
			&broadcast.Title,
			&broadcast.Content,
			&creationDateStr,
			&deadlineStr,
			&creatorUserId,
			&broadcastRecipient.BroadcastRecipientsId,
			&relatedBroadcast,
			&recipientUserId,
			&broadcastRecipient.Acknowledged,
			&broadcastRecipient.Rejected,
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
			if len(broadcastMap) >= queryLimit {
				continue
			}
			broadcast.Type = getBroadcastProtoTypeStringFromDB(broadcastType)
			creator, err := idUserByUserId(db, creatorUserId)

			if err != nil {
				fmt.Println("GetBroadcasts ERROR:", err)
				continue
			}

			broadcast.Creator = creator
			dateFormat := "2006-01-02 15:04:05"

			creationDate, err := time.Parse(dateFormat, creationDateStr)
			if err != nil {
				fmt.Println("GetBroadcasts:", err.Error())
				continue
			}

			deadline, err := time.Parse(dateFormat, deadlineStr)
			if err != nil {
				fmt.Println("GetBroadcasts:", err.Error())
				continue
			}

			broadcast.CreationDate = &timestamppb.Timestamp{Seconds: creationDate.Unix()}
			broadcast.Deadline = &timestamppb.Timestamp{Seconds: deadline.Unix()}
			broadcastMap[broadcast.BroadcastId] = broadcast
		}

		broadcastRecipient.Recipient, err = idUserByUserId(db, recipientUserId)
		if err != nil {
			fmt.Println("GetBroadcasts:", err.Error())
			continue
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
	return getFormattedBroadcastFilters(query, table, false)
}
