// Use these functions to interact with the user related database tables.
package database

import (
	"fmt"
	"strings"

	pb "capstone.operations_ecosystem/backend/proto"
)

const (
	CAMERA_IOT_DB_TABLE_NAME = "camera_iot"
	CAMERA_IOT_DB_ID         = "camera_iot_id"
)

type CameraIotDbAttributes struct {
	CameraIotId   int64
	Name          string
	CameraUrl     string
	GateId        string
	GateAccessKey string
	FireAlarmId   string
	FireAccessKey string
	CpuId         string
	CpuAccessKey  string
}

// Converts the filters in the camera iot array into a formatted where clause
// that can be parsed into MySQL. If a limit is needed, the LIMIT filter is
// added to the end of the string.
// For example returns: "WHERE id=22 AND num <2 LIMIT 5"
// Returns the formatted SQL filter string.
func getFormattedCamerIotFilters(query *pb.CameraIotQuery, needLimit bool, needOrder bool) string {
	output := ""

	// Get all filters
	whereFilters := make([]string, 0)

	for _, filter := range query.Filters {
		hasQuotes := true
		if filter.Comparisons.Comparison == pb.Filter_CONTAINS {
			filter.Comparisons.Value = FormatLikeQueryValue(filter.Comparisons.Value)
		} else if filter.Comparisons.Comparison == pb.Filter_IN {
			filter.Comparisons.Value = FormatInQueryValue(filter.Comparisons.Value)
			hasQuotes = false
		}
		switch filter.Field {
		case pb.CameraIotFilter_CAMERA_IOT_ID:
			if hasQuotes {
				whereFilters = append(
					whereFilters, fmt.Sprintf("%s %s '%s'", cameraIotFilterToDBCol(filter.Field),
						GetFilterComparisonSign(filter.Comparisons.Comparison), filter.Comparisons.Value),
				)
			} else {
				whereFilters = append(
					whereFilters, fmt.Sprintf("%s %s %s", cameraIotFilterToDBCol(filter.Field),
						GetFilterComparisonSign(filter.Comparisons.Comparison), filter.Comparisons.Value),
				)
			}
		}
	}

	if len(whereFilters) > 0 {
		output += WHERE_KEYWORD + " "
	}

	output += strings.Join(whereFilters, " AND ")

	// Add order
	if needOrder {
		if query.OrderBy != nil {
			output += fmt.Sprintf(" %s %s %s", ORDER_BY_KEYWORD, cameraIotFilterToDBCol(query.OrderBy.Field), orderByProtoToDB(query.OrderBy.OrderBy))
		} else {
			// By default we order incidentReports by the aifs id
			output += fmt.Sprintf(" %s %s %s", ORDER_BY_KEYWORD, cameraIotFilterToDBCol(pb.CameraIotFilter_CAMERA_IOT_ID), DESC_KEYWORD)
		}
	}

	// Add limits
	if needLimit {
		if query.Limit == 0 {
			query.Limit = DEFAULT_LIMIT
		}
		output += fmt.Sprintf(" %s %d, %d", LIMIT_KEYWORD, query.Skip, query.Limit)
	}

	return output
}

func cameraIotFilterToDBCol(filterField pb.CameraIotFilter_Field) string {
	output := ""
	switch filterField {
	case pb.CameraIotFilter_CAMERA_IOT_ID:
		output = CAMERA_IOT_DB_ID
	}

	return output
}

// Helper function to add a new filter to the list of existing
// filters in a camera iot query struct.
// Modifies the camera iot query parameter directly.
func AddCameraIotFilter(query *pb.CameraIotQuery, field pb.CameraIotFilter_Field,
	comparison pb.Filter_Comparisons,
	value string) {
	if query.Filters == nil {
		query.Filters = make([]*pb.CameraIotFilter, 0)
	}
	filter := &pb.Filter{Comparison: comparison, Value: value}
	query.Filters = append(query.Filters, &pb.CameraIotFilter{Field: field, Comparisons: filter})
}
