// This whole file does not work yet. TODO: FIX
// Interact with the database through these functions
package database

import (
	"database/sql"
	"fmt"
	"os"
	"sync"
	"time"

	pb "capstone.operations_ecosystem/backend/proto"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	DATETIME_FORMAT  = "2006-01-02 15:04:05"
	DEFAULT_LIMIT    = 10
	MAX_LIMIT        = 1000
	ALL_COLS         = "*"
	WHERE_KEYWORD    = "WHERE"
	LIMIT_KEYWORD    = "LIMIT"
	ORDER_BY_KEYWORD = "ORDER BY"
	ASC_KEYWORD      = "ASC"
	DESC_KEYWORD     = "DESC"
)

// Create an authorised session with the database.
// The credentials of the database are stored in a .env file.
func GetDB() *sql.DB {
	// load .env file
	envFilePath := ".env"
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
	fmt.Println(query)

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

// Get all the rows in a table that meets specifications.
// fields must be a the exact names of the table columns, separated
// be commas. E.g. "name, cost, creator"
// filters must include the filter keyword that is being used
// such as WHERE, LIMIT, ORDER BY, GROUP BY, HAVING
// filters can be an empty string but there should be a LIMIT.
func Query(db *sql.DB, tableName string, fields string, filters string) (*sql.Rows, error) {
	fmt.Println("Making query to table", tableName)

	query := fmt.Sprintf("SELECT %s FROM %s %s;", fields, tableName, filters)
	fmt.Println(query)

	results, err := db.Query(query)
	if err != nil {
		fmt.Println("QUERY ERROR:", err)
	}

	return results, err
}

// Get all the rows from 2 joined table that meets specifications.
// fields must be a the exact names of the table columns, separated
// be commas. E.g. "name, cost, creator"
// filters must include the filter keyword that is being used
// such as WHERE, LIMIT, ORDER BY, GROUP BY, HAVING
// filters can be an empty string but there should be a LIMIT.
func QueryLeftJoin(db *sql.DB, leftTableName string, rightTableName string,
	onCondition string, fields string, filters string) (*sql.Rows, error) {
	fmt.Println("Making left join query to tables", leftTableName, rightTableName)

	query := createLeftJoinQuery(leftTableName, rightTableName, onCondition, fields, filters)

	fmt.Println(query)

	// query := fmt.Sprintf("SELECT %s FROM %s %s;", fields, tableName, filters)
	results, err := db.Query(query)
	if err != nil {
		fmt.Println("QUERY ERROR:", err)
	}

	return results, err
}

func createLeftJoinQuery(leftTableName string, rightTableName string,
	onCondition string, fields string, filters string) string {
	fmt.Println("Creating left join query to tables", leftTableName, rightTableName)
	query := fmt.Sprintf("SELECT %s FROM %s LEFT JOIN %s ON %s %s",
		fields,
		leftTableName,
		rightTableName,
		onCondition,
		filters,
	)
	return query

}

// Update a specific row in the table
// newFields must be a the exact names of the table columns with their new values,
// separated be commas. E.g. "name='mark', cost='22', creator='2'"
// filters must include the filter keyword that is being used
// such as WHERE. For updates filters must not be an empty string.
// Returns the number of rows affected
func Update(db *sql.DB, tableName string, newFields string, filters string) (int64, error) {
	if len(filters) == 0 {
		return 0, fmt.Errorf("filters must not be empty")
	}

	query := fmt.Sprintf("UPDATE %s SET %s %s", tableName, newFields, filters)
	fmt.Println(query)

	result, err := db.Exec(query)
	if err != nil {
		fmt.Println(err)
		return 0, err
	} else {
		return result.RowsAffected()
	}
}

// Delete a particular row in the database
// filters must include the filter keyword that is being used
// such as WHERE. For updates filters must not be an empty string.
// Returns the number of rows affected
func Delete(db *sql.DB, tableName string, filters string) (int64, error) {
	if len(filters) == 0 {
		return 0, fmt.Errorf("filters must not be empty")
	}

	query := fmt.Sprintf("DELETE FROM %s %s;", tableName, filters)
	fmt.Println(query)

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

// Return the string comparison sign for a comparison.
func GetFilterComparisonSign(compaison pb.Filter_Comparisons) string {
	switch compaison {
	case pb.Filter_GREATER:
		return ">"
	case pb.Filter_GREATER_EQ:
		return ">="
	case pb.Filter_LESSER_EQ:
		return "<="
	case pb.Filter_LESSER:
		return "<"
	case pb.Filter_CONTAINS:
		return "LIKE"
	case pb.Filter_IN:
		return "IN"
	default:
		return "="
	}
}

func FormatLikeQueryValue(value string) string {
	return "%" + value + "%"
}

func FormatInQueryValue(value string) string {
	return "(" + value + ")"
}

func formatFieldEqVal(field string, val string, encloseVal bool) string {
	if encloseVal {
		return fmt.Sprintf("%s='%s'", field, val)
	} else {
		return fmt.Sprintf("%s=%s", field, val)
	}
}

func formatFilterCondition(filter *pb.Filter, fieldName string, encloseVal bool) string {
	if encloseVal {
		return fmt.Sprintf("%s %s '%s'", fieldName, GetFilterComparisonSign(filter.Comparison), filter.Value)
	} else {
		return fmt.Sprintf("%s %s %s", fieldName, GetFilterComparisonSign(filter.Comparison), filter.Value)
	}
}

func DBDatetimeToPB(datetimeString string) (*timestamppb.Timestamp, error) {
	creationDate, err := time.Parse(DATETIME_FORMAT, datetimeString)
	if err != nil {
		return &timestamppb.Timestamp{}, err
	}

	return &timestamppb.Timestamp{Seconds: creationDate.Unix()}, nil
}

func orderByProtoToDB(order pb.OrderBy) string {
	if order == pb.OrderBy_ASC {
		return ASC_KEYWORD
	}
	return DESC_KEYWORD
}
