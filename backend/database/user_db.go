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

// Insert a new user into the database table.
func UserInsert(db *sql.DB, user *pb.User, dbLock *sync.Mutex) (int64, error) {
	fmt.Println("Inserting User", user.Name)

	// Create and insert main broadcast first and it's pk
	fields := getUserTableFields()
	values := orderUserFields(user)
	pk, err := Insert(db, USER_DB_TABLE_NAME, fields, values, dbLock)

	return pk, err
}

// Get all the broadcast rows in a table that meets specifications.
func GetUsers(db *sql.DB, query *pb.UserQuery) ([]*pb.User, error) {
	fmt.Println("Getting Users...")
	users := make([]*pb.User, 0)

	// Get the main broadcasts
	fields := "*"

	// Format filters
	filters := getFormattedUserFilters(query, true)

	userRows, err := Query(db, USER_DB_TABLE_NAME, fields, filters)

	if err != nil {
		return users, err
	}

	if userRows != nil {
		// convert query rows into broadcasts
		for userRows.Next() {
			var user pb.User
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

// Update a specific row in the table
func UpdateUser(db *sql.DB, user *pb.User) (int64, error) {
	// Update the main broadcast first
	newFields := getFilledUserFields(user)
	query := &pb.UserQuery{}
	addUserFilter(query, pb.UserFilter_USER_ID, pb.Filter_EQUAL, strconv.Itoa(int(user.UserId)))
	filters := getFormattedUserFilters(query, false)

	rowsAffected, err := Update(db, USER_DB_TABLE_NAME, newFields, filters)

	if err != nil {
		fmt.Println("UpdateUser ERROR::", err)
		return rowsAffected, err
	}

	return rowsAffected, err
}

// Delete a particular user in the database
func DeleteUser(db *sql.DB, user *pb.User) (int64, error) {
	query := &pb.UserQuery{}
	addUserFilter(query, pb.UserFilter_USER_ID, pb.Filter_EQUAL, strconv.Itoa(int(user.UserId)))
	filters := getFormattedUserFilters(query, false)

	rowsAffected, err := Delete(db, USER_DB_TABLE_NAME, filters)
	return rowsAffected, err
}
