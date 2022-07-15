// Use these functions to interact with the user related database tables.
package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	pb "capstone.operations_ecosystem/backend/proto"
)

// Get all the details of the camera and IoT devices from the DB
// Returns an array of attributes or errors if any.
func GetCameraIotDetails(db *sql.DB, query *pb.CameraIotQuery) ([]*CameraIotDbAttributes, error) {
	fmt.Println("Getting CameraIotDbAttributes...")
	cameraIotDbAttributes := make([]*CameraIotDbAttributes, 0)

	fields := ALL_COLS
	filters := getFormattedCamerIotFilters(query, true, true)

	rows, err := Query(db, CAMERA_IOT_DB_TABLE_NAME, fields, filters)

	if err != nil {
		return cameraIotDbAttributes, err
	}

	// convert query rows into users
	for rows.Next() {
		var attribute CameraIotDbAttributes

		// cast each row to a user
		err = rows.Scan(
			&attribute.CameraIotId,
			&attribute.Name,
			&attribute.CameraUrl,
			&attribute.GateId,
			&attribute.GateAccessKey,
			&attribute.FireAlarmId,
			&attribute.FireAccessKey,
			&attribute.CpuId,
			&attribute.CpuAccessKey,
		)

		if err != nil {
			fmt.Println("GetCameraIotDetails ERROR::", err)
			break
		}

		cameraIotDbAttributes = append(cameraIotDbAttributes, &attribute)
	}

	return cameraIotDbAttributes, err
}
