// Use these functions to interact with the broadcast related database tables.
package database

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

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

// Insert a new broadcast into the database table.
// TODO: decide what should be done if the adding rows for the
// receipients fail. Should the main broadcast and all others be deleted as well?
func BroadcastInsert(db *sql.DB, broadcast *pb.Broadcast, dbLock *sync.Mutex) (int64, error) {
	fmt.Println("Inserting Broadcast", broadcast.Title)

	// Create and insert main broadcast first and it's pk
	bcTbFields := getBroadcastTableFields()
	bcValues := orderBroadcastFields(broadcast)

	bc_pk, err := Insert(db, BROADCAST_DB_TABLE_NAME, bcTbFields, bcValues, dbLock)

	if err != nil {
		// Do not add receipients if the main broadcast fails
		return bc_pk, err
	}

	// Create broadcast recipients rows for corresponding broadcast
	for _, user := range broadcast.Receipients {
		_, err = BroadcastReceipientInsert(db, user, bc_pk, dbLock)

		if err != nil {
			break
		}
	}

	return bc_pk, err
}

// Assumes the user has the correct user id that corresponds to its DB row.
func BroadcastReceipientInsert(db *sql.DB, receipient *pb.User, mainBroadcastID int64, dbLock *sync.Mutex) (int64, error) {
	// get fields and values for this particular receipient
	fields := getBroadcastRecTableFields()
	values := orderBroadcastRecFields(receipient, mainBroadcastID)

	// Add receipient to DB
	pk, err := Insert(db, BROADCAST_RECIPIENT_TABLE_NAME, fields, values, dbLock)

	return pk, err
}

// Get all the broadcast rows in a table that meets specifications.
func GetBroadcasts(db *sql.DB, query *pb.BroadcastQuery) ([]*pb.Broadcast, error) {
	fmt.Println("Getting Broadcasts...")
	broadcasts := make([]*pb.Broadcast, 0)

	// Get the main broadcasts
	fields := "*"
	// Format filters
	filters := getFormattedBroadcastWhereFilters(query)
	// Add limits
	filters += fmt.Sprintf(" LIMIT %d", query.Limit)

	mainBCRows, err := Query(db, BROADCAST_DB_TABLE_NAME, fields, filters)

	if err != nil {
		return broadcasts, err
	}

	// convert query rows into broadcasts
	for mainBCRows.Next() {
		var broadcast pb.Broadcast
		creatorUserId := -1
		broadcastType := ""
		creationDateStr := ""
		deadlineStr := ""

		// cast each row to a broadcast
		err = mainBCRows.Scan(
			&broadcast.BroadcastId,
			&broadcastType,
			&broadcast.Title,
			&broadcast.Content,
			&creationDateStr,
			&deadlineStr,
			&creatorUserId)

		if err != nil {
			fmt.Println("GetBroadcasts ERROR:", err)
			continue
		}

		broadcast.Type = getBroadcastProtoTypeStringFromDB(broadcastType)
		creator, err := idUserByUserId(db, creatorUserId)

		if err != nil {
			fmt.Println("GetBroadcasts ERROR:", err)
			continue
		}

		broadcast.Creator = creator
		fmt.Println("creation date", creationDateStr)
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

		// For each broadcast, get the recepients
		receipients, err := GetBroadcastRecipients(db, query, broadcast.BroadcastId)
		if err != nil {
			fmt.Println("GetBroadcasts:", err.Error())
			continue
		}
		broadcast.Receipients = receipients

		// append broadcast to the output list
		broadcasts = append(broadcasts, &broadcast)
	}

	return broadcasts, err
}

// Get all the broadcast rows in a table that meets specifications.
func GetBroadcastRecipients(db *sql.DB, query *pb.BroadcastQuery, mainBroadcastID int64) ([]*pb.User, error) {
	fmt.Println("Getting Broadcasts Recipients...")
	broadcastReceipients := make([]*pb.User, 0)

	// We are currently only interested in the corresponding user
	// TODO: update this assumption
	fields := BC_REC_DB_RECIPIENT

	// Format filters
	filters := getFormattedBroadcastWhereFilters(query)
	// Get for a specific main broadcast
	if len(filters) == 0 {
		filters = "WHERE "
	} else {
		filters += ", "
	}

	filters += fmt.Sprintf("%s='%d'", BC_REC_DB_RELATED_BC, mainBroadcastID)
	// Add limits
	filters += fmt.Sprintf(" LIMIT %d", query.Limit)

	BCRecRows, err := Query(db, BROADCAST_RECIPIENT_TABLE_NAME, fields, filters)

	if err != nil {
		return broadcastReceipients, err
	}

	// convert query rows into broadcasts
	for BCRecRows.Next() {
		var receipient *pb.User
		receipientId := -1

		err = BCRecRows.Scan(&receipientId)

		if err != nil {
			fmt.Println("GetBroadcastRecipients ERROR::", err)
			break
		}

		receipient, err = idUserByUserId(db, receipientId)
		if err != nil {
			fmt.Println("GetBroadcastRecipients ERROR::", err)
			break
		}

		broadcastReceipients = append(broadcastReceipients, receipient)
	}

	return broadcastReceipients, err
}

//TODO
// Update a specific row in the table
// This function assumes that we are using the specific table new_table.
func UpdateBroadcast(db *sql.DB, tableName string, Broadcast *pb.Broadcast) (int64, error) {
	query := fmt.Sprintf("update %s set id = ?, Broadcast = ?, Creator = ? where id = ?", tableName)
	result, err := db.Exec(query,
		Broadcast.BroadcastId,
		Broadcast.Content,
		Broadcast.Creator,
		Broadcast.BroadcastId,
	)
	if err != nil {
		fmt.Println(err)

		return 0, err
	} else {
		return result.RowsAffected()
	}
}

//TODO
// Delete a particular row in the database
// This function assumes that we are using the specific table new_table.
func DeleteBroadcast(db *sql.DB, tableName string, BroadcastId int) (int64, error) {
	query := fmt.Sprintf("DELETE FROM  %s WHERE id = ?;", tableName)
	result, err := db.Exec(query, BroadcastId)

	if err != nil {
		fmt.Println(err)

		return 0, err
	} else {
		return result.RowsAffected()
	}
}

//TODO
// Update a specific row in the table
// This function assumes that we are using the specific table new_table.
func UpdateBroadcastReceipients(db *sql.DB, tableName string, Broadcast *pb.Broadcast) (int64, error) {
	query := fmt.Sprintf("update %s set id = ?, Broadcast = ?, Creator = ? where id = ?", tableName)
	result, err := db.Exec(query,
		Broadcast.BroadcastId,
		Broadcast.Content,
		Broadcast.Creator,
		Broadcast.BroadcastId,
	)
	if err != nil {
		fmt.Println(err)

		return 0, err
	} else {
		return result.RowsAffected()
	}
}

//TODO
// Delete a particular row in the database
// This function assumes that we are using the specific table new_table.
func DeleteBroadcastReceipients(db *sql.DB, tableName string, BroadcastId int) (int64, error) {
	query := fmt.Sprintf("DELETE FROM  %s WHERE id = ?;", tableName)
	result, err := db.Exec(query, BroadcastId)

	if err != nil {
		fmt.Println(err)

		return 0, err
	} else {
		return result.RowsAffected()
	}
}

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
func orderBroadcastRecFields(receipeint *pb.User, relatedBCId int64) string {
	output := ""

	output += strconv.Itoa(int(relatedBCId)) + ","
	output += strconv.Itoa(int(receipeint.UserId)) + ","

	// Ack and rejection are fale by default.
	output += "0, 0"

	return output
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

//TODO
func getFormattedBroadcastWhereFilters(query *pb.BroadcastQuery) string {
	// return "WHERE"
	return ""
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
