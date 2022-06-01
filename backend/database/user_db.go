// Use these functions to interact with the user related database tables.
package database

import (
	"database/sql"
	"fmt"
	"sync"

	pb "capstone.operations_ecosystem/backend/proto"
	_ "github.com/go-sql-driver/mysql"
)

// Insert a new user into the database table.
// Returns the primary key of the user from the database or any errors.
func InsertUser(db *sql.DB, user *pb.User, dbLock *sync.Mutex) (int64, error) {
	fmt.Println("Inserting User", user.Name)

	fields := getUserTableFields()
	values := orderUserFields(user)
	pk, err := Insert(db, USER_DB_TABLE_NAME, fields, values, dbLock)

	return pk, err
}

// Get all the user rows in a table that meets specifications.
// Returns an array of users or errors if any.
func GetUsers(db *sql.DB, query *pb.UserQuery) ([]*pb.User, error) {
	fmt.Println("Getting Users...")
	users := make([]*pb.User, 0)

	fields := ALL_COLS

	// Format filters
	filters := getFormattedUserFilters(query, true, true)

	userRows, err := Query(db, USER_DB_TABLE_NAME, fields, filters)

	if err != nil {
		return users, err
	}

	if userRows != nil {
		// convert query rows into users
		for userRows.Next() {
			var user pb.User
			// Get the string user type and convert it to an enum later.
			userType := ""

			// cast each row to a user
			err = userRows.Scan(
				&user.UserId,
				&userType,
				&user.Name,
				&user.Email,
				&user.PhoneNumber,
				&user.TelegramHandle,
				&user.UserSecurityImg,
				&user.IsPartTimer,
			)

			if err != nil {
				fmt.Println("GetUsers ERROR::", err)
				break
			}

			user.UserType = getUserProtoTypeStringFromDB(userType)

			users = append(users, &user)
		}
	} else {
		fmt.Println("WARNING GetUsers: user results is null")
	}

	return users, err
}

// Update a specific user in the table
// The user id must be filled correctly.
// Only user fields that are not nil will be updated.
// Returns the number of users that were updated and any errors.
// In this case, number of users updated is either 0 or 1.
func UpdateUser(db *sql.DB, user *pb.User) (int64, error) {
	newFields := getFilledUserFields(user)
	// Get filter to find the corresponding user in the database
	filters := getUserIdFormattedFilter(int(user.UserId))

	rowsAffected, err := Update(db, USER_DB_TABLE_NAME, newFields, filters)

	if err != nil {
		fmt.Println("UpdateUser ERROR::", err)
		return rowsAffected, err
	}

	return rowsAffected, err
}

// Delete a particular user in the database
// Note that a user is a foreign key of many other tables.
// Deleting a user will have a cascading effect on all other tables.
// Do not delete a user if records from other tables have to be kept.
// Returns the number of rows that were deleted and any errors.
func DeleteUser(db *sql.DB, user *pb.User) (int64, error) {
	// Get filter to find the corresponding user in the database
	filters := getUserIdFormattedFilter(int(user.UserId))

	rowsAffected, err := Delete(db, USER_DB_TABLE_NAME, filters)
	return rowsAffected, err
}
