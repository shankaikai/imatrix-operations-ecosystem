// This whole file does not work yet. TODO: FIX
// Interact with the database through these functions
package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	pb "capstone.operations_ecosystem/backend/proto"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Create an authorised session with the database.
// The credentials of the database are stored in a .env file.
func GetDB() *sql.DB {
	// load .env file
	envFilePath := filepath.Join("..", ".env")
	err := godotenv.Load(envFilePath)

	if err != nil {
		fmt.Println(err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbAddress := os.Getenv("DB_ADDR")
	dbName := os.Getenv("DB_NAME")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbAddress, dbName)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("getDB", err)
	}
	return db
}

// Insert a new row into the database table.
// This function assumes that we are using the specific table new_table.
// Returns the pk of the last insert value.
func Insert(db *sql.DB, tableName string, fields string, values string, dbLock *sync.Mutex) (int64, error) {
	fmt.Println("Inserting into db", tableName)
	query := fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s)", tableName, fields, values)

	// Ensure that the last insert id is for the corresponding insert
	dbLock.Lock()
	defer dbLock.Unlock()
	insert, err := db.Exec(query)

	if err != nil {
		fmt.Println("INSERT ERROR:", err)
		return -1, err
	}

	pk, err := insert.LastInsertId()

	if err != nil {
		fmt.Println("INSERT ERROR:", err)
		return -1, err
	}

	return pk, err
}

// Get all the broadcast rows in a table that meets specifications.
// fields must be a the exact names of the table columns, separated
// be commas. E.g. "name, cost, creator"
// filters must include the filter keyword that is being used
// such as WHERE, LIMIT, ORDER BY
// filters can be an empty string but there should be a LIMIT.
func Query(db *sql.DB, tableName string, fields string, filters string) (*sql.Rows, error) {
	fmt.Println("Making query to table", tableName)

	query := fmt.Sprintf("SELECT %s FROM %s %s;", fields, tableName, filters)
	results, err := db.Query(query)
	if err != nil {
		fmt.Println("QUERY ERROR:", err)
	}

	return results, err
}

// Update a specific row in the table
func Update(db *sql.DB, tableName string, fields string, filters string) (int64, error) {
	query := fmt.Sprintf("UPDATE %s SET %s %s", tableName, fields, filters)

	result, err := db.Exec(query)
	if err != nil {
		fmt.Println(err)

		return 0, err
	} else {
		return result.RowsAffected()
	}
}

// Delete a particular row in the database
// Returns the number of rows affected
func Delete(db *sql.DB, tableName string, filters string) (int64, error) {
	query := fmt.Sprintf("DELETE FROM  %s %s;", tableName, filters)
	result, err := db.Exec(query)

	if err != nil {
		fmt.Println(err)
		return 0, err
	} else {
		return result.RowsAffected()
	}
}

// Get the biggest primary key in the database table.
// Not really using this function right now.
func GetMaxId(db *sql.DB, tableName string, idColName string) (int, error) {
	max_id := 0
	query := fmt.Sprintf("SELECT max(%s) FROM  %s;", idColName, tableName)
	err := db.QueryRow(query).Scan(&max_id)
	fmt.Println("max id", max_id)

	if err != nil {
		fmt.Println("GET MAX ID ERROR:", err)
	}

	return max_id, err
}

// TODO
func GetFilterComparisonSign(compaison pb.Filter_Comparisons) string {
	switch compaison {
	case pb.Filter_EQUAL:
		return "="
	case pb.Filter_GREATER_EQ:
		return ">="
	default:
		return "<"
	}
}
