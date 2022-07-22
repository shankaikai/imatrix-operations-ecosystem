// Utility functions for database operations related to users.
package database

import (
	"database/sql"
	"fmt"
	"strings"

	pb "capstone.operations_ecosystem/backend/proto"
	_ "github.com/go-sql-driver/mysql"
)

const (
	AVAILABILITY_DB_TABLE_NAME = "availability"

	// Availability table fields
	AVAILABILITY_DB_ID       = "availability_id"
	AVAILABILITY_DB_WEEK     = "week"
	AVAILABILITY_DB_YEAR     = "year"
	AVAILABILITY_DB_GUARD    = "guard"
	AVAILABILITY_DB_SUN      = "sunday"
	AVAILABILITY_DB_MON      = "monday"
	AVAILABILITY_DB_TUES     = "tuesday"
	AVAILABILITY_DB_WED      = "wednesday"
	AVAILABILITY_DB_THURS    = "thursday"
	AVAILABILITY_DB_FRI      = "friday"
	AVAILABILITY_DB_SAT      = "saturday"
	AVAILABILITY_DB_NEXT_SUN = "next_sunday"

	// The limit for availability should be much higher
	// than the actual default to accomodate for the number of guards
	AVAILABILITY_DEFAULT_LIMIT = 1000
)

type Availability struct {
	Availability_id int
	Week            int
	Year            int
	Guard           *pb.User
	Sun             sql.NullString
	Mon             sql.NullString
	Tues            sql.NullString
	Wed             sql.NullString
	Thurs           sql.NullString
	Fri             sql.NullString
	Sat             sql.NullString
	NextSun         sql.NullString
}

// Converts the filters in the availability query array into a formatted where clause
// that can be parsed into MySQL. If a limit is needed, the LIMIT filter is
// added to the end of the string.
// For example returns: "WHERE id=22 AND num <2 LIMIT 5"
// Returns the formatted SQL filter string.
func getFormattedAvailabilityFilters(query *pb.AvailabilityQuery, needLimit bool, needOrder bool) string {
	output := ""

	// Store formatted filter conditions in an array to be joined later
	filters := make([]string, 0)

	for _, filter := range query.Filters {
		// Values for contains have to be reformatted
		if filter.Comparisons.Comparison == pb.Filter_CONTAINS {
			filter.Comparisons.Value = FormatLikeQueryValue(filter.Comparisons.Value)
		}

		switch filter.Field {
		case pb.AvailabilityFilter_AVAILABILITY_ID, pb.AvailabilityFilter_WEEK,
			pb.AvailabilityFilter_YEAR, pb.AvailabilityFilter_GUARD_ID:
			filters = append(filters, formatFilterCondition(filter.Comparisons, availabilityFilterToDBCol(filter.Field), true))

			// DAYS just check for not null
		case pb.AvailabilityFilter_SUN, pb.AvailabilityFilter_MON,
			pb.AvailabilityFilter_TUES, pb.AvailabilityFilter_WED,
			pb.AvailabilityFilter_THURS, pb.AvailabilityFilter_FRI,
			pb.AvailabilityFilter_SAT, pb.AvailabilityFilter_NEXT_SUN:
			filters = append(filters, fmt.Sprintf("%s IS NOT NULL", availabilityFilterToDBCol(filter.Field)))
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
			output += fmt.Sprintf(" %s %s %s", ORDER_BY_KEYWORD, availabilityFilterToDBCol(query.OrderBy.Field), orderByProtoToDB(query.OrderBy.OrderBy))
		} else {
			// By default we order users by client id
			output += fmt.Sprintf(" %s %s %s", ORDER_BY_KEYWORD, availabilityFilterToDBCol(pb.AvailabilityFilter_GUARD_ID), DESC_KEYWORD)
		}
	}

	// Add limits if needed
	if needLimit {
		if query.Limit == 0 {
			query.Limit = AVAILABILITY_DEFAULT_LIMIT
		}
		output += fmt.Sprintf(" %s %d", LIMIT_KEYWORD, query.Limit)
	}

	return output
}

// Helper function to add a new filter to the list of existing
// filters in an AvailabilityQuery struct.
// Modifies the AvailabilityQuery query parameter directly.
func AddAvailabilityFilter(query *pb.AvailabilityQuery, field pb.AvailabilityFilter_Field,
	comparison pb.Filter_Comparisons,
	value string) {
	if query.Filters == nil {
		query.Filters = make([]*pb.AvailabilityFilter, 0)
	}
	filter := &pb.Filter{Comparison: comparison, Value: value}
	query.Filters = append(query.Filters, &pb.AvailabilityFilter{Field: field, Comparisons: filter})
}

func availabilityFilterToDBCol(filterField pb.AvailabilityFilter_Field) string {
	output := ""
	switch filterField {
	case pb.AvailabilityFilter_AVAILABILITY_ID:
		output = AVAILABILITY_DB_ID
	case pb.AvailabilityFilter_WEEK:
		output = AVAILABILITY_DB_WEEK
	case pb.AvailabilityFilter_YEAR:
		output = AVAILABILITY_DB_YEAR
	case pb.AvailabilityFilter_GUARD_ID:
		output = AVAILABILITY_DB_GUARD

		// DAYS
	case pb.AvailabilityFilter_SUN:
		output = AVAILABILITY_DB_SUN
	case pb.AvailabilityFilter_MON:
		output = AVAILABILITY_DB_MON
	case pb.AvailabilityFilter_TUES:
		output = AVAILABILITY_DB_TUES
	case pb.AvailabilityFilter_WED:
		output = AVAILABILITY_DB_WED
	case pb.AvailabilityFilter_THURS:
		output = AVAILABILITY_DB_THURS
	case pb.AvailabilityFilter_FRI:
		output = AVAILABILITY_DB_FRI
	case pb.AvailabilityFilter_SAT:
		output = AVAILABILITY_DB_SAT
	case pb.AvailabilityFilter_NEXT_SUN:
		output = AVAILABILITY_DB_NEXT_SUN

	}

	return output
}
