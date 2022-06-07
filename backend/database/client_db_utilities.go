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
	CLIENT_DB_TABLE_NAME = "client"

	// Client table fields
	CLIENT_DB_ID        = "client_id"
	CLIENT_DB_NAME      = "name"
	CLIENT_DB_ABBR      = "abbreviation"
	CLIENT_DB_EMAIL     = "email"
	CLIENT_DB_ADDR      = "address"
	CLIENT_DB_POSTAL    = "postal_code"
	CLIENT_DB_PHONE_NUM = "phone_number"
)

// Returns the fields of the client table
// in a specific order.
// Note that IDs are auto incremented and should
// not be modified manually. Ommits ID in resulting string.
func getClientTableFields() string {
	clientTableFields := []string{
		CLIENT_DB_NAME,
		CLIENT_DB_ABBR,
		CLIENT_DB_EMAIL,
		CLIENT_DB_ADDR,
		CLIENT_DB_POSTAL,
		CLIENT_DB_PHONE_NUM,
	}

	return strings.Join(clientTableFields, ",")
}

// This function is highly dependent on the
// order given in getClientTableFields.
// Returns the values of the client fields in the
// order that is specified in getClientTableFields
func orderClientFields(client *pb.Client) string {
	output := ""

	output += "'" + client.Name + "'" + ", "
	output += "'" + client.Abbreviation + "'" + ", "
	output += "'" + client.Email + "'" + ", "
	output += "'" + client.Address + "'" + ", "
	output += "'" + strconv.Itoa(int(client.PostalCode)) + "'" + ", "
	output += "'" + client.PhoneNumber + "'"

	return output
}

// Converts the filters in the client array into a formatted where clause
// that can be parsed into MySQL. If a limit is needed, the LIMIT filter is
// added to the end of the string.
// For example returns: "WHERE id=22 AND num <2 LIMIT 5"
// Returns the formatted SQL filter string.
func getFormattedClientFilters(query *pb.ClientQuery, needLimit bool, needOrder bool) string {
	output := ""

	// Store formatted filter conditions in an array to be joined later
	filters := make([]string, 0)

	for _, filter := range query.Filters {
		// Values for contains have to be reformatted
		if filter.Comparisons.Comparison == pb.Filter_CONTAINS {
			filter.Comparisons.Value = FormatLikeQueryValue(filter.Comparisons.Value)
		}

		switch filter.Field {
		case pb.ClientFilter_CLIENT_ID:
			filters = append(filters, formatFilterCondition(filter.Comparisons, clientFilterToDBCol(filter.Field), true))
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
			output += fmt.Sprintf(" %s %s %s", ORDER_BY_KEYWORD, clientFilterToDBCol(query.OrderBy.Field), orderByProtoToDB(query.OrderBy.OrderBy))
		} else {
			// By default we order users by client id
			output += fmt.Sprintf(" %s %s %s", ORDER_BY_KEYWORD, clientFilterToDBCol(pb.ClientFilter_CLIENT_ID), DESC_KEYWORD)
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
// the only condition is a matching client id.
func getClientIdFormattedFilter(clientId int) string {
	query := &pb.ClientQuery{}
	AddClientFilter(query, pb.ClientFilter_CLIENT_ID, pb.Filter_EQUAL, strconv.Itoa(clientId))
	return getFormattedClientFilters(query, false, false)
}

// Helper function to add a new filter to the list of existing
// filters in a user query struct.
// Modifies the user query parameter directly.
func AddClientFilter(query *pb.ClientQuery, field pb.ClientFilter_Field,
	comparison pb.Filter_Comparisons,
	value string) {
	if query.Filters == nil {
		query.Filters = make([]*pb.ClientFilter, 0)
	}
	filter := &pb.Filter{Comparison: comparison, Value: value}
	query.Filters = append(query.Filters, &pb.ClientFilter{Field: field, Comparisons: filter})
}

// This function is used in anticipation of an update query.
// Updates use the format "SET field1=val1, field2=val2".
// The return value of this function is for example: "field1=val1, field2=val2"
// ID not included in this because pk should not be manually changed.
// Note this function highly depends on the protocol buffer message definition
// Returns the formatted fields and values of the filled client fields
func getFilledClientFields(client *pb.Client) string {
	clientTableFields := []string{}

	if len(client.Name) > 0 {
		clientTableFields = append(clientTableFields, formatFieldEqVal(CLIENT_DB_NAME, client.Name, true))
	}
	if len(client.Abbreviation) > 0 {
		clientTableFields = append(clientTableFields, formatFieldEqVal(CLIENT_DB_ABBR, client.Abbreviation, true))
	}
	if len(client.Email) > 0 {
		clientTableFields = append(clientTableFields, formatFieldEqVal(CLIENT_DB_EMAIL, client.Email, true))
	}
	if len(client.Address) > 0 {
		clientTableFields = append(clientTableFields, formatFieldEqVal(CLIENT_DB_ADDR, client.Address, true))
	}
	if client.PostalCode > 0 {
		clientTableFields = append(clientTableFields, formatFieldEqVal(CLIENT_DB_POSTAL, strconv.Itoa(int(client.PostalCode)), true))
	}

	if len(client.PhoneNumber) > 0 {
		clientTableFields = append(clientTableFields, formatFieldEqVal(CLIENT_DB_PHONE_NUM, client.PhoneNumber, true))
	}

	return strings.Join(clientTableFields, ",")
}

func clientFilterToDBCol(filterField pb.ClientFilter_Field) string {
	output := ""
	switch filterField {
	case pb.ClientFilter_CLIENT_ID:
		output = CLIENT_DB_ID
	}

	return output
}

// Get the client corresponding to a particular client id in the db
func IdClientByClientId(db *sql.DB, clientId int) (*pb.Client, error) {
	clientQuery := &pb.ClientQuery{Limit: 1}
	AddClientFilter(clientQuery, pb.ClientFilter_CLIENT_ID, pb.Filter_EQUAL, strconv.Itoa(clientId))

	clients, err := GetClients(db, clientQuery)

	client := &pb.Client{}

	if err == nil {
		client = clients[0]
	}

	return client, err
}
