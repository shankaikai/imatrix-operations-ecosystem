// Utility functions for database operations related to users.
package database

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

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
// order given in getUserTableFields.
// Returns the values of the user fields in the
// order that is specified in getUserTableFields
func orderUserFields(user *pb.User) string {
	output := ""

	output += "'" + getUserDBTypeStringFromProto(user.UserType) + "'" + ", "
	output += "'" + user.Name + "'" + ", "
	output += "'" + user.Email + "'" + ", "
	output += "'" + user.PhoneNumber + "'" + ", "
	output += "'" + user.TelegramHandle + "'" + ", "
	output += "'" + user.UserSecurityImg + "'" + ", "
	output += strconv.FormatBool(user.IsPartTimer)

	return output
}

// Returns the User Type as expected in the DB from the protobuf enum.
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

// Returns the User Type as expected in the proto message from the DB enum.
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

// Converts the filters in the user array into a formatted where clause
// that can be parsed into MySQL. If a limit is needed, the LIMIT filter is
// added to the end of the string.
// For example returns: "WHERE id=22 AND num <2 LIMIT 5"
// Returns the formatted SQL filter string.
func getFormattedUserFilters(query *pb.UserQuery, needLimit bool, needOrder bool) string {
	output := ""

	// Store formatted filter conditions in an array to be joined later
	filters := make([]string, 0)

	for _, filter := range query.Filters {
		hasQuotes := true

		// Values for contains have to be reformatted
		if filter.Comparisons.Comparison == pb.Filter_CONTAINS {
			filter.Comparisons.Value = FormatLikeQueryValue(filter.Comparisons.Value)
		} else if filter.Comparisons.Comparison == pb.Filter_IN || filter.Comparisons.Comparison == pb.Filter_NOT_IN {
			filter.Comparisons.Value = FormatInQueryValue(filter.Comparisons.Value)
			hasQuotes = false
		}

		switch filter.Field {
		case pb.UserFilter_USER_ID, pb.UserFilter_TYPE, pb.UserFilter_NAME,
			pb.UserFilter_EMAIL, pb.UserFilter_PHONE_NUMBER, pb.UserFilter_TELEGRAM_HANDLE:
			if hasQuotes {
				filters = append(filters, formatFilterCondition(filter.Comparisons, userFilterToDBCol(filter.Field), true))
			} else {
				filters = append(filters, formatFilterCondition(filter.Comparisons, userFilterToDBCol(filter.Field), false))
			}
		case pb.UserFilter_IS_PART_TIMER:
			filters = append(filters, formatFilterCondition(filter.Comparisons, userFilterToDBCol(filter.Field), false))
		}
	}

	// Only add WHERE keyword if there are conditions to add.
	if len(filters) > 0 {
		output += fmt.Sprintf("%s ", WHERE_KEYWORD)
	}

	output += strings.Join(filters, " AND ")

	// Add order
	if needOrder {
		if query.OrderBy != nil {
			output += fmt.Sprintf(" %s %s %s", ORDER_BY_KEYWORD, userFilterToDBCol(query.OrderBy.Field), orderByProtoToDB(query.OrderBy.OrderBy))
		} else {
			// By default we order users by user id
			output += fmt.Sprintf(" %s %s %s", ORDER_BY_KEYWORD, userFilterToDBCol(pb.UserFilter_USER_ID), DESC_KEYWORD)
		}
	}

	// Add limits if needed
	if needLimit {
		if query.Limit == 0 {
			query.Limit = DEFAULT_LIMIT
		}
		output += fmt.Sprintf(" %s %d", LIMIT_KEYWORD, query.Limit)
	}

	return output
}

// This function creates the filter required if
// the only condition is a matching user id.
func getUserIdFormattedFilter(userId int) string {
	query := &pb.UserQuery{}
	AddUserFilter(query, pb.UserFilter_USER_ID, pb.Filter_EQUAL, strconv.Itoa(userId))
	return getFormattedUserFilters(query, false, false)
}

// Helper function to add a new filter to the list of existing
// filters in a user query struct.
// Modifies the user query parameter directly.
func AddUserFilter(query *pb.UserQuery, field pb.UserFilter_Field,
	comparison pb.Filter_Comparisons,
	value string) {
	if query.Filters == nil {
		query.Filters = make([]*pb.UserFilter, 0)
	}
	filter := &pb.Filter{Comparison: comparison, Value: value}
	query.Filters = append(query.Filters, &pb.UserFilter{Field: field, Comparisons: filter})
}

// This function is used in anticipation of an update query.
// Updates use the format "SET field1=val1, field2=val2".
// The return value of this function is for example: "field1=val1, field2=val2"
// ID not included in this because pk should not be manually changed.
// Note this function highly depends on the protocol buffer message definition
// Returns the formatted fields and values of the filled user fields
func getFilledUserFields(user *pb.User) string {
	userTableFields := []string{formatFieldEqVal(USER_DB_TYPE, getUserDBTypeStringFromProto(user.UserType), true)}

	if len(user.Name) > 0 {
		userTableFields = append(userTableFields, formatFieldEqVal(USER_DB_NAME, user.Name, true))
	}
	if len(user.Email) > 0 {
		userTableFields = append(userTableFields, formatFieldEqVal(USER_DB_EMAIL, user.Email, true))
	}
	if len(user.PhoneNumber) > 0 {
		userTableFields = append(userTableFields, formatFieldEqVal(USER_DB_PHONE_NUM, user.PhoneNumber, true))
	}
	if len(user.TelegramHandle) > 0 {
		userTableFields = append(userTableFields, formatFieldEqVal(USER_DB_TELE_HANDLE, user.TelegramHandle, true))
	}
	if len(user.UserSecurityImg) > 0 {
		userTableFields = append(userTableFields, formatFieldEqVal(USER_DB_IMG, user.UserSecurityImg, true))
	}

	userTableFields = append(userTableFields, formatFieldEqVal(USER_DB_PART_TIMER, strconv.FormatBool(user.IsPartTimer), false))

	return strings.Join(userTableFields, ",")
}

func userFilterToDBCol(filterField pb.UserFilter_Field) string {
	output := ""
	switch filterField {
	case pb.UserFilter_USER_ID:
		output = USER_DB_ID
	case pb.UserFilter_TYPE:
		output = USER_DB_TYPE
	case pb.UserFilter_NAME:
		output = USER_DB_NAME
	case pb.UserFilter_EMAIL:
		output = USER_DB_EMAIL
	case pb.UserFilter_PHONE_NUMBER:
		output = USER_DB_PHONE_NUM
	case pb.UserFilter_TELEGRAM_HANDLE:
		output = USER_DB_TELE_HANDLE
	case pb.UserFilter_IS_PART_TIMER:
		output = USER_DB_PART_TIMER
	}

	return output
}

// Get the user corresponding to a particular user id in the db
func idUserByUserId(db *sql.DB, userId int) (*pb.User, error) {
	userQuery := &pb.UserQuery{Limit: 1}
	AddUserFilter(userQuery, pb.UserFilter_USER_ID, pb.Filter_EQUAL, strconv.Itoa(userId))

	users, err := GetUsers(db, userQuery)

	user := &pb.User{}

	if err == nil {
		user = users[0]
	}

	return user, err
}
