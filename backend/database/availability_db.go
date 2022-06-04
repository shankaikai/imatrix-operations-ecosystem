// Use these functions to interact with the availability database table.
package database

import (
	"database/sql"
	"fmt"

	pb "capstone.operations_ecosystem/backend/proto"
	_ "github.com/go-sql-driver/mysql"
)

// Get all the client rows in a table that meets specifications.
// Returns an array of clients or errors if any.
func GetAvailability(db *sql.DB, query *pb.AvailabilityQuery) ([]*Availability, error) {
	fmt.Println("Getting Available Users...")
	availabilities := make([]*Availability, 0)

	fields := ALL_COLS

	// Format filters
	filters := getFormattedAvailabilityFilters(query, true, true)

	availRows, err := Query(db, AVAILABILITY_DB_TABLE_NAME, fields, filters)

	if err != nil {
		return availabilities, err
	}

	if availRows != nil {
		// convert query rows into clients
		for availRows.Next() {
			availability := &Availability{Guard: &pb.User{}}

			// cast each row to a client
			err = availRows.Scan(
				&availability.Availability_id,
				&availability.Week,
				&availability.Year,
				&availability.Guard.UserId,
				&availability.Sun,
				&availability.Mon,
				&availability.Tues,
				&availability.Wed,
				&availability.Thurs,
				&availability.Fri,
				&availability.Sat,
				&availability.NextSun,
			)

			if err != nil {
				fmt.Println("GetAvailability ERROR::", err)
				break
			}

			availabilities = append(availabilities, availability)
		}
	} else {
		fmt.Println("WARNING GetAvailability: results is null")
	}

	return availabilities, err
}
