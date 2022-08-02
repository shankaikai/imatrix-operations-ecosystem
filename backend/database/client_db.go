// Use these functions to interact with the client related database tables.
package database

import (
	"database/sql"
	"fmt"
	"sync"

	pb "capstone.operations_ecosystem/backend/proto"
	_ "github.com/go-sql-driver/mysql"
)

// Insert a new client into the database table.
// Returns the primary key of the client from the database or any errors.
func InsertClient(db *sql.DB, client *pb.Client, dbLock *sync.Mutex) (int64, error) {
	fmt.Println("Inserting Client", client.Name)

	fields := getClientTableFields()
	values := orderClientFields(client)
	pk, err := Insert(db, CLIENT_DB_TABLE_NAME, fields, values, dbLock)

	return pk, err
}

// Get all the client rows in a table that meets specifications.
// Returns an array of clients or errors if any.
func GetClients(db *sql.DB, query *pb.ClientQuery) ([]*pb.Client, error) {
	fmt.Println("Getting Clients...")
	clients := make([]*pb.Client, 0)

	fields := ALL_COLS

	// Format filters
	filters := getFormattedClientFilters(query, true, true)

	clientRows, err := Query(db, CLIENT_DB_TABLE_NAME, fields, filters)

	if err != nil {
		return clients, err
	}

	if clientRows != nil {
		// convert query rows into clients
		for clientRows.Next() {
			var client pb.Client

			// cast each row to a client
			err = clientRows.Scan(
				&client.ClientId,
				&client.Name,
				&client.Abbreviation,
				&client.Email,
				&client.Address,
				&client.PostalCode,
				&client.PhoneNumber,
			)

			if err != nil {
				fmt.Println("GetClients ERROR::", err)
				break
			}

			clients = append(clients, &client)
		}
	} else {
		fmt.Println("WARNING GetClients: client results is null")
	}

	return clients, err
}

// Update a specific client in the table
// The client id must be filled correctly.
// Only client fields that are not nil will be updated.
// Returns the number of clients that were updated and any errors.
// In this case, number of clients updated is either 0 or 1.
func UpdateClients(db *sql.DB, client *pb.Client) (int64, error) {
	newFields := getFilledClientFields(client)
	// Get filter to find the corresponding client in the database
	filters := getClientIdFormattedFilter(int(client.ClientId))

	rowsAffected, err := Update(db, CLIENT_DB_TABLE_NAME, newFields, filters)

	if err != nil {
		fmt.Println("UpdateClients ERROR::", err)
		return rowsAffected, err
	}

	return rowsAffected, err
}

// Delete a particular client in the database
// Note that a client is a foreign key of many other tables.
// Deleting a client will have a cascading effect on all other tables.
// Do not delete a client if records from other tables have to be kept.
// Returns the number of rows that were deleted and any errors.
func DeleteClient(db *sql.DB, client *pb.Client) (int64, error) {
	// Get filter to find the corresponding client in the database
	filters := getClientIdFormattedFilter(int(client.ClientId))

	rowsAffected, err := Delete(db, CLIENT_DB_TABLE_NAME, filters)
	return rowsAffected, err
}
