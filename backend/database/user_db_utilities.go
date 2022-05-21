// Utility functions for database operations related to users.
package database

import (
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
// order given in getBroadcastTableFields.
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
func getFormattedUserFilters(query *pb.UserQuery, needLimit bool) string {
	output := ""

	if len(query.Filters) > 0 {
		output += "WHERE "
	}

	// Get all filters
	filters := make([]string, 0)
	for _, filter := range query.Filters {
		if filter.Comparisons.Comparison == pb.Filter_CONTAINS {
			filter.Comparisons.Value = FormatLikeQueryValue(filter.Comparisons.Value)
		}

		switch filter.Field {
		case pb.UserFilter_USER_ID:
			filters = append(filters, fmt.Sprintf("%s %s '%s'", USER_DB_ID, GetFilterComparisonSign(filter.Comparisons.Comparison), filter.Comparisons.Value))
		}

	}

	output += strings.Join(filters, ",")

	// Add limits
	if needLimit {
		if query.Limit == 0 {
			query.Limit = DEFAULT_LIMIT
		}
		output += fmt.Sprintf(" LIMIT %d", query.Limit)
	}

	return output
}

func addUserFilter(query *pb.UserQuery, field pb.UserFilter_Field,
	comparison pb.Filter_Comparisons,
	value string) {
	if query.Filters == nil {
		query.Filters = make([]*pb.UserFilter, 0)
	}
	filter := &pb.Filter{Comparison: comparison, Value: value}
	query.Filters = append(query.Filters, &pb.UserFilter{Field: field, Comparisons: filter})
}

// ID not included in this because pk should not be manually changed.
// Note this function highly depends on the protocol buffer message definition
// Returns the string fields and values of the filled broadcast fields
func getFilledUserFields(user *pb.User) string {
	userTableFields := []string{formatFieldEqVal(USER_DB_TYPE, getUserDBTypeStringFromProto(user.UserType))}

	if len(user.Name) > 0 {
		userTableFields = append(userTableFields, formatFieldEqVal(USER_DB_NAME, user.Name))
	}
	if len(user.Email) > 0 {
		userTableFields = append(userTableFields, formatFieldEqVal(USER_DB_EMAIL, user.Email))
	}
	if len(user.PhoneNumber) > 0 {
		userTableFields = append(userTableFields, formatFieldEqVal(USER_DB_PHONE_NUM, user.PhoneNumber))
	}
	if len(user.TelegramHandle) > 0 {
		userTableFields = append(userTableFields, formatFieldEqVal(USER_DB_TELE_HANDLE, user.TelegramHandle))
	}
	if len(user.UserSecurityImg) > 0 {
		userTableFields = append(userTableFields, formatFieldEqVal(USER_DB_IMG, user.UserSecurityImg))
	}

	userTableFields = append(userTableFields, formatFieldEqVal(USER_DB_PART_TIMER, strconv.FormatBool(user.IsPartTimer)))

	return strings.Join(userTableFields, ",")
}
