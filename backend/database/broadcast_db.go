// Use these functions to interact with the broadcast related database tables.
package database

import (
	"database/sql"
	"fmt"
	"strconv"
	"sync"

	pb "capstone.operations_ecosystem/backend/proto"
	_ "github.com/go-sql-driver/mysql"
)

// Insert a new broadcast into the database table.
// Corresponding broadcast recipients are added to their
// respective table as well.
// Returns the primary key of the main broadcast and errors if any.
func InsertBroadcast(db *sql.DB, broadcast *pb.Broadcast, dbLock *sync.Mutex) (int64, error) {
	fmt.Println("Inserting Broadcast", broadcast.BroadcastId)

	// Create and insert main broadcast first and get it's pk
	bcTbFields := getBroadcastTableFields()
	bcValues := orderBroadcastFields(broadcast)

	bcPk, err := Insert(db, BROADCAST_DB_TABLE_NAME, bcTbFields, bcValues, dbLock)

	if err != nil {
		// Do not add recipients if the main broadcast fails
		return bcPk, err
	}

	// Create broadcast recipients rows for corresponding broadcast
	for _, aifsRecipient := range broadcast.Recipients {
		for _, recipient := range aifsRecipient.Recipient {
			_, err = InsertBroadcastRecipient(db, recipient, bcPk, dbLock)

			if err != nil {
				// Delete the broadcast that was just inserted
				broadcast.BroadcastId = bcPk
				DeleteBroadcast(db, broadcast)
				break
			}
		}
	}

	return bcPk, err
}

// Inserts a new recipient to the database and connects it to the appropriate
// main broadcast.
// Assumes the recipient has the correct id that corresponds to its DB row.
// Returns the primary key of the recipient row and any errors.
func InsertBroadcastRecipient(db *sql.DB, recipient *pb.BroadcastRecipient, mainBroadcastID int64, dbLock *sync.Mutex) (int64, error) {
	// get fields and values for this particular recipient
	fields := getBroadcastRecTableFields()
	values := orderBroadcastRecFields(recipient, mainBroadcastID)

	// Add recipient to DB
	pk, err := Insert(db, BROADCAST_RECIPIENT_TABLE_NAME, fields, values, dbLock)

	return pk, err
}

// Get all the broadcast rows in a table that meets specifications.
// Returns an array of broadcasts and any errors.
func GetBroadcasts(db *sql.DB, query *pb.BroadcastQuery) ([]*pb.Broadcast, error) {
	fmt.Println("Getting Broadcasts...")
	broadcasts := make([]*pb.Broadcast, 0)

	// Join the broadcast and recipient tables in order to
	// easily filter conditions relating to both tables together.
	// Conditions that are specified in the broadcast query
	// are used in the inner query to facilitate agregated conditions.
	innerFields := BC_DB_ID
	outerFields := ALL_COLS

	// Format filters
	innerFilters := getFormattedBroadcastFilters(query, BROADCAST_DB_TABLE_NAME, false, false)

	// tables are joined on the main broadcast id
	onCondition := formatFieldEqVal(BC_DB_ID, BC_REC_DB_RELATED_BC, false)

	innerQuery := createLeftJoinQuery(BROADCAST_DB_TABLE_NAME, BROADCAST_RECIPIENT_TABLE_NAME, onCondition, innerFields, innerFilters)

	// It is not easy to find out how many rows a single broadcast
	// will span because of the join, hence limit is the max.
	// The filter for the outer query is any rows that made it through the
	// inner query.
	outerQuery := &pb.BroadcastQuery{Limit: MAX_LIMIT, OrderBy: query.OrderBy}
	addBroadcastFilter(outerQuery, pb.BroadcastFilter_BROADCAST_ID, pb.Filter_IN, innerQuery)
	outerFilter := getFormattedBroadcastFilters(outerQuery, BROADCAST_DB_TABLE_NAME, true, true)

	rows, err := QueryLeftJoin(db, BROADCAST_DB_TABLE_NAME, BROADCAST_RECIPIENT_TABLE_NAME, onCondition, outerFields, outerFilter)

	if err != nil {
		return broadcasts, err
	}

	// convert query rows into broadcasts
	err = convertDbRowsToBcNBcR(db, &broadcasts, rows, query)

	return broadcasts, err
}

// Get all the broadcast recipient rows in a table that meets specifications.
// Returns an array of broadcast recipients and any errors.
func GetBroadcastRecipients(db *sql.DB, query *pb.BroadcastQuery, mainBroadcastID int64) ([]*pb.BroadcastRecipient, error) {
	fmt.Println("Getting Broadcasts Recipients...")
	broadcastRecipients := make([]*pb.BroadcastRecipient, 0)

	fields := ALL_COLS

	// Format filters
	// Get for a specific main broadcast
	addBroadcastFilter(query, pb.BroadcastFilter_BROADCAST_ID, pb.Filter_EQUAL, strconv.Itoa(int(mainBroadcastID)))
	filters := getFormattedBroadcastFilters(query, BROADCAST_RECIPIENT_TABLE_NAME, true, true)

	BCRecRows, err := Query(db, BROADCAST_RECIPIENT_TABLE_NAME, fields, filters)

	if err != nil {
		return broadcastRecipients, err
	}

	// convert query rows into broadcasts recipients
	for BCRecRows.Next() {
		recipient := &pb.BroadcastRecipient{}
		// fields that cannot be auto converted
		recipientId := -1
		// related broadcast is not necessary, but for simplicity
		// and for possible future use, we get it back in the query.
		relatedBroadcast := ""
		var lastRepliedString sql.NullString

		// cast each row to a broadcast
		err = BCRecRows.Scan(
			&recipient.BroadcastRecipientsId,
			&relatedBroadcast,
			&recipientId,
			&recipient.Acknowledged,
			&recipient.Rejected,
			&lastRepliedString,
			&recipient.AifsId,
		)

		if err != nil {
			fmt.Println("GetBroadcastRecipients ERROR::", err)
			break
		}

		if lastRepliedString.Valid {
			recipient.LastReplied, err = DBDatetimeToPB(lastRepliedString.String)
			if err != nil {
				fmt.Println("GetBroadcasts:", err.Error())
				continue
			}
		}

		// TODO think about whether I can store the users in cache rather than
		// get the same few users over and over
		recipient.Recipient, err = idUserByUserId(db, recipientId)
		if err != nil {
			fmt.Println("GetBroadcasts:", err.Error())
			continue
		}

		broadcastRecipients = append(broadcastRecipients, recipient)
	}

	return broadcastRecipients, err
}

// Update a specific broadcast in the table
// Only fields that have been filled in the broadcast object will be updated.
// Note that this update does not update the broadcast's recipient's inner status
// such as the acknowledgement or rejection status but only if the recipient
// is part of the broadcast.
func UpdateBroadcast(db *sql.DB, broadcast *pb.Broadcast, dbLock *sync.Mutex) (int64, error) {
	// Update the main broadcast first
	newFields := getFilledBroadcastFields(broadcast)

	query := &pb.BroadcastQuery{}
	addBroadcastFilter(query, pb.BroadcastFilter_BROADCAST_ID, pb.Filter_EQUAL, strconv.Itoa(int(broadcast.BroadcastId)))
	filters := getFormattedBroadcastFilters(query, BROADCAST_DB_TABLE_NAME, false, false)

	rowsAffected, err := Update(db, BROADCAST_DB_TABLE_NAME, newFields, filters)

	if err != nil {
		fmt.Println("UpdateBroadcast ERROR::", err)
		return rowsAffected, err
	}

	// Update recipients if necessary
	if broadcast.Recipients != nil {
		err = updateRecipientsOfBroadcast(db, broadcast, query, dbLock)
	}
	return rowsAffected, err
}

// Update a specific recipient row in the table
// This function assumes that the broadcast recipient id is correct.
// Returns the number of rows affected and any errors.
// In this case, number of rows affected is either 0 or 1.
func UpdateBroadcastRecipients(db *sql.DB, broadcastRecipient *pb.BroadcastRecipient) (int64, error) {
	newFields := getFilledBroadcastRecFields(broadcastRecipient)
	filters := getBroadcastIdFormattedFilter(
		int(broadcastRecipient.BroadcastRecipientsId),
		BROADCAST_RECIPIENT_TABLE_NAME,
	)

	rowsAffected, err := Update(db, BROADCAST_RECIPIENT_TABLE_NAME, newFields, filters)

	if err != nil {
		fmt.Println("UpdateBroadcastRecipients ERROR::", err)
		return rowsAffected, err
	}

	return rowsAffected, err
}

// Delete a particular broadcast in the database together with
// any corresponding recipients.
// Recipients are deleted based on the cascading rule.
func DeleteBroadcast(db *sql.DB, broadcast *pb.Broadcast) (int64, error) {
	filters := getBroadcastIdFormattedFilter(
		int(broadcast.BroadcastId),
		BROADCAST_DB_TABLE_NAME,
	)

	rowsAffected, err := Delete(db, BROADCAST_DB_TABLE_NAME, filters)
	return rowsAffected, err
}

// Delete a particular broadcast recipient
func DeleteBroadcastRecipients(db *sql.DB, broadcastRecipient *pb.BroadcastRecipient) (int64, error) {
	filters := fmt.Sprintf("WHERE %s=%d", BC_REC_DB_ID, broadcastRecipient.BroadcastRecipientsId)

	rowsAffected, err := Delete(db, BROADCAST_RECIPIENT_TABLE_NAME, filters)
	return rowsAffected, err
}

// Delete all recipients belonging to a particular broadcast
// Currently not in use.
func DeleteAllBCRecipientsOfMainBC(db *sql.DB, mainBroadcastID int) (int64, error) {
	filters := getBroadcastIdFormattedFilter(mainBroadcastID, BROADCAST_RECIPIENT_TABLE_NAME)
	rowsAffected, err := Delete(db, BROADCAST_RECIPIENT_TABLE_NAME, filters)
	return rowsAffected, err
}
