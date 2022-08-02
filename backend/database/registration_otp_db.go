// Use these functions to interact with the user related database tables.
package database

import (
	"database/sql"
	"fmt"
	"sync"

	pb "capstone.operations_ecosystem/backend/proto"
	_ "github.com/go-sql-driver/mysql"
)

// Insert a new registration otp into the database table.
// Returns the primary key of the user from the database or any errors.
func InsertRegOTP(db *sql.DB, regOtp *pb.RegistrationOTP, dbLock *sync.Mutex) (int64, error) {
	fmt.Println("Inserting User Token")

	fields := getregOtpTableFields()
	values := orderRegOtpFields(regOtp)
	pk, err := Insert(db, REG_OTP_DB_TABLE_NAME, fields, values, dbLock)

	return pk, err
}

// Get all the reg otp rows in a table that meets specifications.
// Returns an array of reg otp or errors if any.
func GetRegOtp(db *sql.DB, query *pb.RegistrationOTPQuery) ([]*pb.RegistrationOTP, error) {
	fmt.Println("Getting Reg OTPs...")
	regOtps := make([]*pb.RegistrationOTP, 0)

	fields := ALL_COLS

	// Format filters
	filters := getFormattedRegOtpFilters(query, true, true)

	userTokenRows, err := Query(db, REG_OTP_DB_TABLE_NAME, fields, filters)

	if err != nil {
		return regOtps, err
	}

	if userTokenRows != nil {
		// convert query rows into reg otps
		for userTokenRows.Next() {
			var regOtp pb.RegistrationOTP
			var creator pb.User
			var userType string

			// cast each row to a user
			err = userTokenRows.Scan(
				&regOtp.RegistrationOtpId,
				&regOtp.Token,
				&userType,
				&regOtp.CreationDatetime,
				&creator.UserId,
				&regOtp.IsUsed,
			)

			if err != nil {
				fmt.Println("GetRegOtp ERROR::", err)
				break
			}
			regOtp.UserType = getUserProtoTypeStringFromDB(userType)
			regOtp.Creator, err = idUserByUserId(db, int(creator.UserId))
			if err != nil {
				fmt.Println("GetRegOtp ERROR::", err)
				break
			}

			regOtps = append(regOtps, &regOtp)
		}
	} else {
		fmt.Println("WARNING GetRegOtp: db row results is null")
	}
	fmt.Println("ppdfdpsofds")

	return regOtps, err
}

// Update a specific reg otp in the table
// The reg otp id must be filled correctly if there are no other query options
// Only user fields that are not nil will be updated.
// Returns the number of users that were updated and any errors.
// In this case, number of users updated is either 0 or 1.
func UpdateRegOtp(db *sql.DB, regOtp *pb.RegistrationOTP, query *pb.RegistrationOTPQuery) (int64, error) {
	newFields := getFilledRefOtpFields(regOtp)

	// Format filters
	var filters string

	// Use filters defined by the caller
	if query != nil && query.Filters != nil && len(query.Filters) > 0 {
		filters = getFormattedRegOtpFilters(query, false, false)
	} else {
		// Get filter to find the corresponding reg otp in the database
		filters = getRegOtpIdFormattedFilter(int(regOtp.RegistrationOtpId))
	}

	rowsAffected, err := Update(db, REG_OTP_DB_TABLE_NAME, newFields, filters)

	if err != nil {
		fmt.Println("UpdateRegOtp ERROR::", err)
		return rowsAffected, err
	}

	return rowsAffected, err
}

// Delete a particular reg otp in the database
// Note that a user is a foreign key of many other tables.
// Deleting a user will have a cascading effect on all other tables.
// Do not delete a user if records from other tables have to be kept.
// Returns the number of rows that were deleted and any errors.
func DeleteRegOtp(db *sql.DB, regOtp *pb.RegistrationOTP) (int64, error) {
	// Get filter to find the corresponding reg otp in the database
	filters := getRegOtpIdFormattedFilter(int(regOtp.RegistrationOtpId))

	rowsAffected, err := Delete(db, REG_OTP_DB_TABLE_NAME, filters)
	return rowsAffected, err
}
