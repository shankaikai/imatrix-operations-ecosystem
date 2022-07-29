// Utility functions for database operations related to reg Otps
package database

import (
	"fmt"
	"strconv"
	"strings"

	pb "capstone.operations_ecosystem/backend/proto"
	_ "github.com/go-sql-driver/mysql"
)

const (
	REG_OTP_DB_TABLE_NAME = "registration_otp"

	// Registration OTP fields
	REG_OTP_DB_ID        = "registration_otp_id"
	REG_OTP_DB_TOKEN     = "token"
	REG_OTP_DB_USER_TYPE = "user_type"
	REG_OTP_DB_CREATION  = "creation_date"
	REG_OTP_DB_CREATOR   = "creator"
	REG_OTP_DB_IS_USED   = "is_used"
)

// Returns the fields of the reg Otps table
// in a specific order.
// Note that IDs are auto incremented and should
// not be modified manually. Ommits ID in resulting string.
func getregOtpTableFields() string {
	regOtpTableFields := []string{
		REG_OTP_DB_TOKEN,
		REG_OTP_DB_USER_TYPE,
		REG_OTP_DB_CREATION,
		REG_OTP_DB_CREATOR,
		REG_OTP_DB_IS_USED,
	}

	return strings.Join(regOtpTableFields, ",")
}

// This function is highly dependent on the
// order given in getUserTableFields.
// Returns the values of the user fields in the
// order that is specified in getUserTableFields
func orderRegOtpFields(regOtp *pb.RegistrationOTP) string {
	output := ""

	output += "'" + regOtp.Token + "'" + ", "
	output += "'" + getUserDBTypeStringFromProto(regOtp.UserType) + "'" + ", "
	output += "'" + regOtp.CreationDatetime + "'" + ", "
	output += "'" + strconv.Itoa(int(regOtp.Creator.UserId)) + "'" + ", "
	output += "'" + strconv.FormatBool(regOtp.IsUsed) + "'"

	return output
}

// Converts the filters in the user array into a formatted where clause
// that can be parsed into MySQL. If a limit is needed, the LIMIT filter is
// added to the end of the string.
// For example returns: "WHERE id=22 AND num <2 LIMIT 5"
// Returns the formatted SQL filter string.
func getFormattedRegOtpFilters(query *pb.RegistrationOTPQuery, needLimit bool, needOrder bool) string {
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
		case pb.RegistrationOTPFilter_REG_OTP_ID, pb.RegistrationOTPFilter_TOKEN, pb.RegistrationOTPFilter_USER_TYPE,
			pb.RegistrationOTPFilter_CREATION_DATE, pb.RegistrationOTPFilter_CREATOR_ID:
			if hasQuotes {
				filters = append(filters, formatFilterCondition(filter.Comparisons, regOtpFilterToDBCol(filter.Field), true))
			} else {
				filters = append(filters, formatFilterCondition(filter.Comparisons, regOtpFilterToDBCol(filter.Field), false))
			}
		case pb.RegistrationOTPFilter_IS_USED:
			filters = append(filters, formatFilterCondition(filter.Comparisons, regOtpFilterToDBCol(filter.Field), false))
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
			output += fmt.Sprintf(" %s %s %s", ORDER_BY_KEYWORD, regOtpFilterToDBCol(query.OrderBy.Field), orderByProtoToDB(query.OrderBy.OrderBy))
		} else {
			// By default we order users by user id
			output += fmt.Sprintf(" %s %s %s", ORDER_BY_KEYWORD, regOtpFilterToDBCol(pb.RegistrationOTPFilter_REG_OTP_ID), DESC_KEYWORD)
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
func getRegOtpIdFormattedFilter(regOtp int) string {
	query := &pb.RegistrationOTPQuery{}
	AddRegOtpFilter(query, pb.RegistrationOTPFilter_REG_OTP_ID, pb.Filter_EQUAL, strconv.Itoa(regOtp))
	return getFormattedRegOtpFilters(query, false, false)
}

// Helper function to add a new filter to the list of existing
// filters in a reg Otp query struct.
// Modifies the reg Otp query parameter directly.
func AddRegOtpFilter(query *pb.RegistrationOTPQuery, field pb.RegistrationOTPFilter_Field,
	comparison pb.Filter_Comparisons,
	value string) {
	if query.Filters == nil {
		query.Filters = make([]*pb.RegistrationOTPFilter, 0)
	}
	filter := &pb.Filter{Comparison: comparison, Value: value}
	query.Filters = append(query.Filters, &pb.RegistrationOTPFilter{Field: field, Comparisons: filter})
}

// This function is used in anticipation of an update query.
// Updates use the format "SET field1=val1, field2=val2".
// The return value of this function is for example: "field1=val1, field2=val2"
// ID not included in this because pk should not be manually changed.
// Note this function highly depends on the protocol buffer message definition
// Returns the formatted fields and values of the filled user fields
func getFilledRefOtpFields(regOtp *pb.RegistrationOTP) string {
	regOtpTableFields := []string{}

	if len(regOtp.Token) != 0 {
		regOtpTableFields = append(regOtpTableFields, formatFieldEqVal(REG_OTP_DB_TOKEN, regOtp.Token, true))
	}
	regOtpTableFields = append(regOtpTableFields, formatFieldEqVal(REG_OTP_DB_USER_TYPE, getUserDBTypeStringFromProto(regOtp.UserType), true))

	if len(regOtp.CreationDatetime) > 0 {
		regOtpTableFields = append(regOtpTableFields, formatFieldEqVal(REG_OTP_DB_CREATION, regOtp.CreationDatetime, true))
	}
	if regOtp.Creator.UserId > 0 {
		regOtpTableFields = append(regOtpTableFields, formatFieldEqVal(REG_OTP_DB_CREATOR, strconv.Itoa(int(regOtp.Creator.UserId)), true))
	}
	regOtpTableFields = append(regOtpTableFields, formatFieldEqVal(REG_OTP_DB_IS_USED, strconv.FormatBool(regOtp.IsUsed), true))

	return strings.Join(regOtpTableFields, ",")
}

func regOtpFilterToDBCol(filterField pb.RegistrationOTPFilter_Field) string {
	output := ""
	switch filterField {
	case pb.RegistrationOTPFilter_REG_OTP_ID:
		output = REG_OTP_DB_ID
	case pb.RegistrationOTPFilter_TOKEN:
		output = REG_OTP_DB_TOKEN
	case pb.RegistrationOTPFilter_USER_TYPE:
		output = REG_OTP_DB_USER_TYPE
	case pb.RegistrationOTPFilter_CREATION_DATE:
		output = REG_OTP_DB_CREATION
	case pb.RegistrationOTPFilter_CREATOR_ID:
		output = REG_OTP_DB_CREATOR
	case pb.RegistrationOTPFilter_IS_USED:
		output = REG_OTP_DB_IS_USED
	}

	return output
}
