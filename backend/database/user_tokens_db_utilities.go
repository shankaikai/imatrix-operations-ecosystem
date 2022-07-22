// Utility functions for database operations related to user tokens
package database

import (
	"fmt"
	"strconv"
	"strings"

	pb "capstone.operations_ecosystem/backend/proto"
	_ "github.com/go-sql-driver/mysql"
)

const (
	USER_TOKENS_DB_TABLE_NAME = "user_tokens"

	// User table fields
	USER_TOKENS_DB_ID       = "user_tokens_id"
	USER_TOKENS_DB_USER     = "user"
	USER_TOKENS_DB_TOKEN    = "token"
	USER_TOKENS_DB_CREATION = "creation"
	USER_TOKENS_DB_EXPIRY   = "expiry"
)

// Returns the fields of the user tokens table
// in a specific order.
// Note that IDs are auto incremented and should
// not be modified manually. Ommits ID in resulting string.
func getUserTokensTableFields() string {
	userTokensTableFields := []string{
		USER_TOKENS_DB_USER,
		USER_TOKENS_DB_TOKEN,
		USER_TOKENS_DB_CREATION,
		USER_TOKENS_DB_EXPIRY,
	}

	return strings.Join(userTokensTableFields, ",")
}

// This function is highly dependent on the
// order given in getUserTableFields.
// Returns the values of the user fields in the
// order that is specified in getUserTableFields
func orderUserTokensFields(userToken *pb.UserToken) string {
	output := ""

	output += "'" + strconv.Itoa(int(userToken.User.UserId)) + "'" + ", "
	output += "'" + userToken.Token + "'" + ", "
	output += "'" + userToken.CreationDatetime + "'" + ", "
	output += "'" + userToken.ExpiryDatetime + "'"

	return output
}

// Converts the filters in the user array into a formatted where clause
// that can be parsed into MySQL. If a limit is needed, the LIMIT filter is
// added to the end of the string.
// For example returns: "WHERE id=22 AND num <2 LIMIT 5"
// Returns the formatted SQL filter string.
func getFormattedUserTokenFilters(query *pb.UserTokenQuery, needLimit bool, needOrder bool) string {
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
		case pb.UserTokenFilter_USER_ID, pb.UserTokenFilter_EXPIRY:
			if hasQuotes {
				filters = append(filters, formatFilterCondition(filter.Comparisons, userTokenFilterToDBCol(filter.Field), true))
			} else {
				filters = append(filters, formatFilterCondition(filter.Comparisons, userTokenFilterToDBCol(filter.Field), false))
			}
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
			output += fmt.Sprintf(" %s %s %s", ORDER_BY_KEYWORD, userTokenFilterToDBCol(query.OrderBy.Field), orderByProtoToDB(query.OrderBy.OrderBy))
		} else {
			// By default we order users by user id
			output += fmt.Sprintf(" %s %s %s", ORDER_BY_KEYWORD, userTokenFilterToDBCol(pb.UserTokenFilter_EXPIRY), DESC_KEYWORD)
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
func getUserTokenIdFormattedFilter(userToken int) string {
	query := &pb.UserTokenQuery{}
	AddUserTokenFilter(query, pb.UserTokenFilter_USER_ID, pb.Filter_EQUAL, strconv.Itoa(userToken))
	return getFormattedUserTokenFilters(query, false, false)
}

// Helper function to add a new filter to the list of existing
// filters in a user token query struct.
// Modifies the user token query parameter directly.
func AddUserTokenFilter(query *pb.UserTokenQuery, field pb.UserTokenFilter_Field,
	comparison pb.Filter_Comparisons,
	value string) {
	if query.Filters == nil {
		query.Filters = make([]*pb.UserTokenFilter, 0)
	}
	filter := &pb.Filter{Comparison: comparison, Value: value}
	query.Filters = append(query.Filters, &pb.UserTokenFilter{Field: field, Comparisons: filter})
}

// This function is used in anticipation of an update query.
// Updates use the format "SET field1=val1, field2=val2".
// The return value of this function is for example: "field1=val1, field2=val2"
// ID not included in this because pk should not be manually changed.
// Note this function highly depends on the protocol buffer message definition
// Returns the formatted fields and values of the filled user fields
func getFilledUserTokenFields(userToken *pb.UserToken) string {
	userTokenTableFields := []string{}

	if userToken.User != nil {
		userTokenTableFields = append(userTokenTableFields, formatFieldEqVal(USER_TOKENS_DB_USER, strconv.Itoa(int(userToken.User.UserId)), true))
	}
	userTokenTableFields = append(userTokenTableFields, formatFieldEqVal(USER_TOKENS_DB_TOKEN, userToken.Token, true))

	if len(userToken.CreationDatetime) > 0 {
		userTokenTableFields = append(userTokenTableFields, formatFieldEqVal(USER_TOKENS_DB_CREATION, userToken.CreationDatetime, true))
	}
	if len(userToken.ExpiryDatetime) > 0 {
		userTokenTableFields = append(userTokenTableFields, formatFieldEqVal(USER_TOKENS_DB_EXPIRY, userToken.ExpiryDatetime, true))
	}
	return strings.Join(userTokenTableFields, ",")
}

func userTokenFilterToDBCol(filterField pb.UserTokenFilter_Field) string {
	output := ""
	switch filterField {
	case pb.UserTokenFilter_USER_ID:
		output = USER_TOKENS_DB_USER
	case pb.UserTokenFilter_EXPIRY:
		output = USER_TOKENS_DB_EXPIRY
	}

	return output
}
