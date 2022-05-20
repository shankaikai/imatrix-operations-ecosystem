// Use these functions to interact with the broadcast related database tables.
package database

import (
	"database/sql"
	"fmt"
	"strings"
	"sync"

	pb "capstone.operations_ecosystem/backend/proto"
	_ "github.com/go-sql-driver/mysql"
)

const (
	USER_DB_TABLE_NAME = "user"

	// User table fields
	USER_DB_ID          = "user_id"
	USER_DB_TYPE        = "user_type"
	USER_DB_NAME        = "name"
	USER_DB_EMAIL       = "email"
	USER_DB_PHONE_NUM   = "phone_number"
	USER_DB_TELE_HANDLE = "telegram_handle"
	USER_DB_IMG         = "user_security_img"
	USER_DB_PART_TIMER  = "is_part_timer"
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
	filters := getFormattedUserWhereFilters(query)
	// Add limits
	filters += fmt.Sprintf(" LIMIT %d", query.Limit)

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

// Returns the fields of the user table
// in a specific order.
// Note that IDs are auto incremented and should
// not be modified manually. Ommits ID in resulting string.
func getUserTableFields() string {
	userTableFields := []string{
		USER_DB_TYPE,
		USER_DB_NAME,
		USER_DB_EMAIL,
		USER_DB_PHONE_NUM,
		USER_DB_TELE_HANDLE,
		USER_DB_IMG,
		USER_DB_PART_TIMER,
	}

	return strings.Join(userTableFields, ",")
}

// This function is highly dependent on the
// order given in getBroadcastTableFields.
func orderUserFields(user *pb.User) string {
	output := ""

	output += "'" + getUserDBTypeStringFromProto(user.UserType) + "'" + ", "
	output += "'" + user.Name + "'" + ", "
	output += "'" + user.Email + "'" + ", "
	output += "'" + user.PhoneNumber + "'" + ", "
	output += "'" + user.TelegramHandle + "'" + ", "
	output += "'" + user.UserSecurityImg + "'" + ", "

	if user.IsPartTimer {
		output += "1"
	} else {
		output += "0"
	}

	return output
}

// Returns the User Type as expected in the DB
func getUserDBTypeStringFromProto(userType pb.User_UserType) string {
	switch userType {
	case pb.User_ISPECIALIST:
		return "I-Specialist"
	case pb.User_CONTROLLER:
		return "Controller"
	case pb.User_MANAGER:
		return "Manager"
	default:
		return "Security Guard"
	}
}

// Returns the User Type as expected in the proto message
func getUserProtoTypeStringFromDB(userType string) pb.User_UserType {
	switch userType {
	case "I-Specialist":
		return pb.User_ISPECIALIST
	case "Controller":
		return pb.User_CONTROLLER
	case "Manager":
		return pb.User_MANAGER
	default:
		return pb.User_SECURITY_GUARD
	}
}

//TODO
func getFormattedUserWhereFilters(query *pb.UserQuery) string {
	output := ""

	if len(query.Filters) > 0 {
		output += "WHERE "
	}

	for _, filter := range query.Filters {
		switch filter.Field {
		case pb.UserFilter_USER_ID:
			output += USER_DB_ID + GetFilterComparisonSign(filter.Comparisons.Comparison) + "'" + filter.Comparisons.Value + "'"
		}

		// TODO: should add a comma after the filter
	}

	return output
}
