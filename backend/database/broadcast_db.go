// Use these functions to interact with the broadcast related database tables.
package database

import (
	"database/sql"
	"fmt"
	"strconv"
	"sync"
	"time"

	pb "capstone.operations_ecosystem/backend/proto"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	for _, recipient := range broadcast.Receipients {
		_, err = BroadcastReceipientInsert(db, recipient, bc_pk, dbLock)

		if err != nil {
			break
		}
	}

	return bc_pk, err
}

// Assumes the user has the correct user id that corresponds to its DB row.
func BroadcastReceipientInsert(db *sql.DB, receipient *pb.BroadcastRecipient, mainBroadcastID int64, dbLock *sync.Mutex) (int64, error) {
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

	// Join all the broadcast
	innerFields := BC_DB_ID
	outerFields := "*"
	// Format filters
	innerFilters := getFormattedBroadcastFilters(query, BROADCAST_DB_TABLE_NAME, false)
	onCondition := fmt.Sprintf("%s = %s", BC_DB_ID, BC_REC_DB_RELATED_BC)

	innerQuery := createLeftJoinQuery(BROADCAST_DB_TABLE_NAME, BROADCAST_RECIPIENT_TABLE_NAME, onCondition, innerFields, innerFilters)

	// It is not easy to find out how many rows a single broadcast will span because of the join.
	outerFilter := fmt.Sprintf("WHERE %s IN (%s) LIMIT %d", BC_DB_ID, innerQuery, MAX_LIMIT)

	rows, err := QueryLeftJoin(db, BROADCAST_DB_TABLE_NAME, BROADCAST_RECIPIENT_TABLE_NAME, onCondition, outerFields, outerFilter)

	if err != nil {
		return broadcasts, err
	}

	// convert query rows into broadcasts
	broadcastMap := make(map[int64]*pb.Broadcast)

	for rows.Next() {
		broadcast := &pb.Broadcast{}
		broadcastReceipient := &pb.BroadcastRecipient{}

		creatorUserId := -1
		broadcastType := ""
		creationDateStr := ""
		deadlineStr := ""
		recipientUserId := -1
		relatedBroadcast := ""

		// cast each row to a broadcast
		err = rows.Scan(
			&broadcast.BroadcastId,
			&broadcastType,
			&broadcast.Title,
			&broadcast.Content,
			&creationDateStr,
			&deadlineStr,
			&creatorUserId,
			&broadcastReceipient.BroadcastRecipientsId,
			&relatedBroadcast,
			&recipientUserId,
			&broadcastReceipient.Acknowledged,
			&broadcastReceipient.Rejected,
		)

		if err != nil {
			fmt.Println("GetBroadcasts ERROR:", err)
			continue
		}

		// Check if there was already a broadcast found with this id
		if existingBroadcast, ok := broadcastMap[broadcast.BroadcastId]; ok {
			broadcast = existingBroadcast
		} else {
			// Return only the necessary number of broadcasts.
			// If the number of broadcasts have reached the limit,
			// do not add new broadcasts.
			if len(broadcastMap) >= int(query.Limit) {
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

		broadcastReceipient.Recipient, err = idUserByUserId(db, recipientUserId)
		if err != nil {
			fmt.Println("GetBroadcasts:", err.Error())
			continue
		}

		// Add recipient to broadcast
		if broadcast.Receipients == nil {
			broadcast.Receipients = make([]*pb.BroadcastRecipient, 0)
		}

		broadcast.Receipients = append(broadcast.Receipients, broadcastReceipient)
	}

	// Add all broadcasts to the returning array
	for _, broadcast := range broadcastMap {
		broadcasts = append(broadcasts, broadcast)
	}

	return broadcasts, err
}

// Get all the broadcast rows in a table that meets specifications.
func GetBroadcastRecipients(db *sql.DB, query *pb.BroadcastQuery, mainBroadcastID int64) ([]*pb.BroadcastRecipient, error) {
	fmt.Println("Getting Broadcasts Recipients...")
	broadcastReceipients := make([]*pb.BroadcastRecipient, 0)

	// We are currently only interested in the corresponding user
	// TODO: update this assumption
	fields := BC_REC_DB_RECIPIENT

	// Format filters
	// Get for a specific main broadcast
	addBroadcastFilter(query, pb.BroadcastFilter_BROADCAST_ID, pb.Filter_EQUAL, strconv.Itoa(int(mainBroadcastID)))
	filters := getFormattedBroadcastFilters(query, BROADCAST_RECIPIENT_TABLE_NAME, true)

	BCRecRows, err := Query(db, BROADCAST_RECIPIENT_TABLE_NAME, fields, filters)

	if err != nil {
		return broadcastReceipients, err
	}

	// convert query rows into broadcasts
	for BCRecRows.Next() {
		receipient := &pb.BroadcastRecipient{}
		receipientId := -1
		relatedBroadcast := ""
		// cast each row to a broadcast
		err = BCRecRows.Scan(
			receipient.BroadcastRecipientsId,
			&relatedBroadcast,
			&receipientId,
			receipient.Acknowledged,
			receipient.Rejected,
		)

		if err != nil {
			fmt.Println("GetBroadcastRecipients ERROR::", err)
			break
		}

		// TODO think about whether I can store the users in cache rather than
		// get the same few users over and over
		receipient.Recipient, err = idUserByUserId(db, receipientId)
		if err != nil {
			fmt.Println("GetBroadcasts:", err.Error())
			continue
		}

		broadcastReceipients = append(broadcastReceipients, receipient)
	}

	return broadcastReceipients, err
}

// Update a specific broadcast in the table
// Only fields that have been filled in the broadcast object will be updated.
// Note that this update does not update the broadcast's recipient's inner status
// such as the acknowledgement or rejection status but only if the recipient
// is part of the broadcast.
func UpdateBroadcast(db *sql.DB, broadcast *pb.Broadcast) (int64, error) {
	// Update the main broadcast first
	newFields := getFilledBroadcastFields(broadcast)
	query := &pb.BroadcastQuery{}
	addBroadcastFilter(query, pb.BroadcastFilter_BROADCAST_ID, pb.Filter_EQUAL, strconv.Itoa(int(broadcast.BroadcastId)))
	filters := getFormattedBroadcastFilters(query, BROADCAST_DB_TABLE_NAME, false)

	rowsAffected, err := Update(db, BROADCAST_DB_TABLE_NAME, newFields, filters)

	if err != nil {
		fmt.Println("UpdateBroadcast ERROR::", err)
		return rowsAffected, err
	}

	//TODO
	// Update recipients if necessary
	// Get all recipients
	// See if any are missing
	// See if any need to be deleted

	return rowsAffected, err
}

//TODO
// Update a specific row in the table
func UpdateBroadcastReceipients(db *sql.DB, tableName string, broadcast *pb.Broadcast) (int64, error) {
	// Update the main broadcast first
	newFields := getFilledBroadcastFields(broadcast)
	query := &pb.BroadcastQuery{}
	addBroadcastFilter(query, pb.BroadcastFilter_BROADCAST_ID, pb.Filter_EQUAL, strconv.Itoa(int(broadcast.BroadcastId)))
	filters := getFormattedBroadcastFilters(query, BROADCAST_DB_TABLE_NAME, false)

	rowsAffected, err := Update(db, BROADCAST_DB_TABLE_NAME, newFields, filters)

	if err != nil {
		fmt.Println("UpdateBroadcast ERROR::", err)
		return rowsAffected, err
	}

	return rowsAffected, err
}

// Delete a particular broadcast in the database together with
// any corresponding recipients.
// Recipients are deleted based on the cascading rule.
func DeleteBroadcast(db *sql.DB, broadcast *pb.Broadcast) (int64, error) {
	// delete main broadcast
	query := &pb.BroadcastQuery{}
	addBroadcastFilter(query, pb.BroadcastFilter_BROADCAST_ID, pb.Filter_EQUAL, strconv.Itoa(int(broadcast.BroadcastId)))
	filters := getFormattedBroadcastFilters(query, BROADCAST_DB_TABLE_NAME, false)

	rowsAffected, err := Delete(db, BROADCAST_DB_TABLE_NAME, filters)
	return rowsAffected, err
}

// Delete a particular broadcast recipient
func DeleteBroadcastRecipients(db *sql.DB, broadcastReceipient *pb.BroadcastRecipient) (int64, error) {
	query := &pb.BroadcastQuery{}
	addBroadcastFilter(query, pb.BroadcastFilter_RECEIPEIENT_ID, pb.Filter_EQUAL, strconv.Itoa(int(broadcastReceipient.BroadcastRecipientsId)))
	filters := getFormattedBroadcastFilters(query, BROADCAST_RECIPIENT_TABLE_NAME, false)

	rowsAffected, err := Delete(db, BROADCAST_DB_TABLE_NAME, filters)
	return rowsAffected, err
}

func DeleteAllBCRecipientsOfMainBC(db *sql.DB, mainBroadcastID int) (int64, error) {
	query := &pb.BroadcastQuery{}
	addBroadcastFilter(query, pb.BroadcastFilter_BROADCAST_ID, pb.Filter_EQUAL, strconv.Itoa(mainBroadcastID))
	filters := getFormattedBroadcastFilters(query, BROADCAST_RECIPIENT_TABLE_NAME, false)

	rowsAffected, err := Delete(db, BROADCAST_RECIPIENT_TABLE_NAME, filters)
	return rowsAffected, err
}
