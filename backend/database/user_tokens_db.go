// Use these functions to interact with the user related database tables.
package database

import (
	"database/sql"
	"fmt"
	"sync"

	pb "capstone.operations_ecosystem/backend/proto"
	_ "github.com/go-sql-driver/mysql"
)

// Insert a new user token into the database table.
// Returns the primary key of the user from the database or any errors.
func InsertUserToken(db *sql.DB, userToken *pb.UserToken, dbLock *sync.Mutex) (int64, error) {
	fmt.Println("Inserting User Token")

	fields := getUserTokensTableFields()
	values := orderUserTokensFields(userToken)
	pk, err := Insert(db, USER_TOKENS_DB_TABLE_NAME, fields, values, dbLock)

	return pk, err
}

// Get all the user roken rows in a table that meets specifications.
// Returns an array of user tokens or errors if any.
func GetUserTokens(db *sql.DB, query *pb.UserTokenQuery, removeSecrets bool) ([]*pb.UserToken, error) {
	fmt.Println("Getting User Tokens...")
	userTokens := make([]*pb.UserToken, 0)

	fields := ALL_COLS

	// Format filters
	filters := getFormattedUserTokenFilters(query, true, true)

	userTokenRows, err := Query(db, USER_TOKENS_DB_TABLE_NAME, fields, filters)

	if err != nil {
		return userTokens, err
	}

	if userTokenRows != nil {
		// convert query rows into user tokens
		for userTokenRows.Next() {
			var user pb.User
			var userToken pb.UserToken

			// cast each row to a user
			err = userTokenRows.Scan(
				&userToken.UserTokenId,
				&user.UserId,
				&userToken.Token,
				&userToken.CreationDatetime,
				&userToken.ExpiryDatetime,
			)

			if err != nil {
				fmt.Println("GetUsersTokens ERROR::", err)
				break
			}

			userToken.User = &user

			userTokens = append(userTokens, &userToken)
		}
	} else {
		fmt.Println("WARNING GetUsersTokens: user results is null")
	}

	return userTokens, err
}

// Update a specific user token in the table
// The user token id must be filled correctly if there are no other query options
// Only user fields that are not nil will be updated.
// Returns the number of users that were updated and any errors.
// In this case, number of users updated is either 0 or 1.
func UpdateUserTokens(db *sql.DB, userToken *pb.UserToken, query *pb.UserTokenQuery) (int64, error) {
	newFields := getFilledUserTokenFields(userToken)
	var filters string

	// Use filters defined by the caller
	if query.Filters != nil && len(query.Filters) > 0 {
		filters = getFormattedUserTokenFilters(query, false, false)
	} else {
		// Get filter to find the corresponding user in the database
		filters = getUserTokenIdFormattedFilter(int(userToken.UserTokenId))
	}

	rowsAffected, err := Update(db, USER_TOKENS_DB_TABLE_NAME, newFields, filters)

	if err != nil {
		fmt.Println("UpdateUserTokens ERROR::", err)
		return rowsAffected, err
	}

	return rowsAffected, err
}

// Delete a particular user token in the database
// Note that a user is a foreign key of many other tables.
// Deleting a user will have a cascading effect on all other tables.
// Do not delete a user if records from other tables have to be kept.
// Returns the number of rows that were deleted and any errors.
func DeleteUserToken(db *sql.DB, userToken *pb.UserToken) (int64, error) {
	// Get filter to find the corresponding user token in the database
	filters := getUserIdFormattedFilter(int(userToken.UserTokenId))

	rowsAffected, err := Delete(db, USER_TOKENS_DB_TABLE_NAME, filters)
	return rowsAffected, err
}
