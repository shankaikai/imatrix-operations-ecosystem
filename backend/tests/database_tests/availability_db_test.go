package tests

import (
	"database/sql"
	"testing"

	db_pck "capstone.operations_ecosystem/backend/database"
	"google.golang.org/protobuf/proto"

	pb "capstone.operations_ecosystem/backend/proto"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

const (
	TEST_AVAILABILITY_DEFAULT_TIME_ARRAY = "['2021-02-22 18:00:00']"
)

func TestGetAvailabilityNoFilter(t *testing.T) {
	query := &pb.AvailabilityQuery{}
	fakeAvailabilities := []*db_pck.Availability{getDefaultAvailabilities(1), getDefaultAvailabilities(2)}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"availability_id", "week", "year", "guard", "sunday", "monday",
		"tuesday", "wednesday", "thursday", "friday", "saturday", "next_sunday"}).
		AddRow(1, 1, 1, 1, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY,
			TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY,
			TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY,
		).
		AddRow(2, 1, 1, 2, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY,
			TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY,
			TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY,
		)

	mock.ExpectQuery("SELECT \\* FROM availability").WillReturnRows(rows)

	availabilities, err := db_pck.GetAvailability(db, query)

	if err != nil {
		t.Errorf("error was not expected while selecting users: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 2, len(availabilities), "The length of availabilities returned should be 2")
	t.Log(availabilities, fakeAvailabilities)
	assert.Equal(t, true, compareAvailabilities(availabilities[0], fakeAvailabilities[0]), "The first avail returned is not equal to the expected.")
	assert.Equal(t, true, compareAvailabilities(availabilities[1], fakeAvailabilities[1]), "The second avail returned is not equal to the expected.")
}

func TestGetAvailabilityID(t *testing.T) {
	query := &pb.AvailabilityQuery{}
	query.Filters = append(query.Filters, &pb.AvailabilityFilter{
		Field: pb.AvailabilityFilter_AVAILABILITY_ID,
		Comparisons: &pb.Filter{
			Value: "1", Comparison: pb.Filter_EQUAL,
		},
	})

	fakeAvailabilities := []*db_pck.Availability{getDefaultAvailabilities(1)}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"availability_id", "week", "year", "guard", "sunday", "monday",
		"tuesday", "wednesday", "thursday", "friday", "saturday", "next_sunday"}).
		AddRow(1, 1, 1, 1, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY,
			TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY,
			TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY,
		)

	mock.ExpectQuery("SELECT \\* FROM availability WHERE availability_id = '1'.*LIMIT").WillReturnRows(rows)

	availabilities, err := db_pck.GetAvailability(db, query)

	if err != nil {
		t.Errorf("error was not expected while selecting users: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 1, len(availabilities), "The length of availabilities returned should be 1")
	assert.Equal(t, true, compareAvailabilities(availabilities[0], fakeAvailabilities[0]), "The first avail returned is not equal to the expected.")
}

func TestGetAvailabilityGuardID(t *testing.T) {
	query := &pb.AvailabilityQuery{}
	query.Filters = append(query.Filters, &pb.AvailabilityFilter{
		Field: pb.AvailabilityFilter_GUARD_ID,
		Comparisons: &pb.Filter{
			Value: "1", Comparison: pb.Filter_EQUAL,
		},
	})

	fakeAvailabilities := []*db_pck.Availability{getDefaultAvailabilities(1)}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"availability_id", "week", "year", "guard", "sunday", "monday",
		"tuesday", "wednesday", "thursday", "friday", "saturday", "next_sunday"}).
		AddRow(1, 1, 1, 1, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY,
			TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY,
			TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY,
		)

	mock.ExpectQuery("SELECT \\* FROM availability WHERE guard = '1'.*LIMIT").WillReturnRows(rows)

	availabilities, err := db_pck.GetAvailability(db, query)

	if err != nil {
		t.Errorf("error was not expected while selecting users: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 1, len(availabilities), "The length of availabilities returned should be 1")
	assert.Equal(t, true, compareAvailabilities(availabilities[0], fakeAvailabilities[0]), "The first avail returned is not equal to the expected.")
}

func TestGetAvailabilityWeekYear(t *testing.T) {
	query := &pb.AvailabilityQuery{}
	query.Filters = append(query.Filters, &pb.AvailabilityFilter{
		Field: pb.AvailabilityFilter_WEEK,
		Comparisons: &pb.Filter{
			Value: "1", Comparison: pb.Filter_EQUAL,
		},
	})

	query.Filters = append(query.Filters, &pb.AvailabilityFilter{
		Field: pb.AvailabilityFilter_YEAR,
		Comparisons: &pb.Filter{
			Value: "1", Comparison: pb.Filter_EQUAL,
		},
	})

	fakeAvailabilities := []*db_pck.Availability{getDefaultAvailabilities(1), getDefaultAvailabilities(2)}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"availability_id", "week", "year", "guard", "sunday", "monday",
		"tuesday", "wednesday", "thursday", "friday", "saturday", "next_sunday"}).
		AddRow(1, 1, 1, 1, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY,
			TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY,
			TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY,
		).
		AddRow(2, 1, 1, 2, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY,
			TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY,
			TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY,
		)

	mock.ExpectQuery("SELECT \\* FROM availability WHERE week = '1' AND year = '1'.*LIMIT").WillReturnRows(rows)

	availabilities, err := db_pck.GetAvailability(db, query)

	if err != nil {
		t.Errorf("error was not expected while selecting users: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 2, len(availabilities), "The length of availabilities returned should be 2")
	assert.Equal(t, true, compareAvailabilities(availabilities[0], fakeAvailabilities[0]), "The first avail returned is not equal to the expected.")
	assert.Equal(t, true, compareAvailabilities(availabilities[1], fakeAvailabilities[1]), "The second avail returned is not equal to the expected.")
}

func TestGetAvailabilityDays(t *testing.T) {
	query := &pb.AvailabilityQuery{}

	query.Filters = append(query.Filters, &pb.AvailabilityFilter{
		Field: pb.AvailabilityFilter_SUN,
		Comparisons: &pb.Filter{
			Value: "", Comparison: pb.Filter_EQUAL,
		},
	})

	query.Filters = append(query.Filters, &pb.AvailabilityFilter{
		Field: pb.AvailabilityFilter_MON,
		Comparisons: &pb.Filter{
			Value: "1", Comparison: pb.Filter_EQUAL,
		},
	})

	query.Filters = append(query.Filters, &pb.AvailabilityFilter{
		Field: pb.AvailabilityFilter_TUES,
		Comparisons: &pb.Filter{
			Value: "", Comparison: pb.Filter_EQUAL,
		},
	})

	query.Filters = append(query.Filters, &pb.AvailabilityFilter{
		Field: pb.AvailabilityFilter_WED,
		Comparisons: &pb.Filter{
			Value: "", Comparison: pb.Filter_EQUAL,
		},
	})

	query.Filters = append(query.Filters, &pb.AvailabilityFilter{
		Field: pb.AvailabilityFilter_THURS,
		Comparisons: &pb.Filter{
			Value: "", Comparison: pb.Filter_EQUAL,
		},
	})

	query.Filters = append(query.Filters, &pb.AvailabilityFilter{
		Field: pb.AvailabilityFilter_FRI,
		Comparisons: &pb.Filter{
			Value: "", Comparison: pb.Filter_EQUAL,
		},
	})

	query.Filters = append(query.Filters, &pb.AvailabilityFilter{
		Field: pb.AvailabilityFilter_SAT,
		Comparisons: &pb.Filter{
			Value: "", Comparison: pb.Filter_EQUAL,
		},
	})

	query.Filters = append(query.Filters, &pb.AvailabilityFilter{
		Field: pb.AvailabilityFilter_NEXT_SUN,
		Comparisons: &pb.Filter{
			Value: "", Comparison: pb.Filter_EQUAL,
		},
	})

	fakeAvailabilities := []*db_pck.Availability{getDefaultAvailabilities(1), getDefaultAvailabilities(2)}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"availability_id", "week", "year", "guard", "sunday", "monday",
		"tuesday", "wednesday", "thursday", "friday", "saturday", "next_sunday"}).
		AddRow(1, 1, 1, 1, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY,
			TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY,
			TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY,
		).
		AddRow(2, 1, 1, 2, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY,
			TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY,
			TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, TEST_AVAILABILITY_DEFAULT_TIME_ARRAY,
		)

	mock.ExpectQuery("SELECT \\* FROM availability WHERE sunday IS NOT NULL AND monday IS NOT NULL " +
		"AND tuesday IS NOT NULL AND wednesday IS NOT NULL AND thursday IS NOT NULL AND " +
		"friday IS NOT NULL AND saturday IS NOT NULL AND next_sunday IS NOT NULL.*LIMIT").WillReturnRows(rows)

	availabilities, err := db_pck.GetAvailability(db, query)

	if err != nil {
		t.Errorf("error was not expected while selecting users: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 2, len(availabilities), "The length of availabilities returned should be 2")
	assert.Equal(t, true, compareAvailabilities(availabilities[0], fakeAvailabilities[0]), "The first avail returned is not equal to the expected.")
	assert.Equal(t, true, compareAvailabilities(availabilities[1], fakeAvailabilities[1]), "The second avail returned is not equal to the expected.")
}

func getDefaultAvailabilities(id int) *db_pck.Availability {
	return &db_pck.Availability{
		Availability_id: id,
		Week:            1,
		Year:            1,
		Guard:           &pb.User{UserId: int64(id)},
		Sun:             sql.NullString{String: TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, Valid: true},
		Mon:             sql.NullString{String: TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, Valid: true},
		Tues:            sql.NullString{String: TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, Valid: true},
		Wed:             sql.NullString{String: TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, Valid: true},
		Thurs:           sql.NullString{String: TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, Valid: true},
		Fri:             sql.NullString{String: TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, Valid: true},
		Sat:             sql.NullString{String: TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, Valid: true},
		NextSun:         sql.NullString{String: TEST_AVAILABILITY_DEFAULT_TIME_ARRAY, Valid: true},
	}
}

func compareAvailabilities(first *db_pck.Availability, second *db_pck.Availability) bool {
	if first.Availability_id != second.Availability_id {
		return false
	}
	if first.Week != second.Week {
		return false
	}
	if first.Year != second.Year {
		return false
	}
	if !proto.Equal(first.Guard, second.Guard) {
		return false
	}
	if first.Sun.String != second.Sun.String || first.Sun.Valid != second.Sun.Valid {
		return false
	}
	if first.Tues.String != second.Tues.String {
		return false
	}
	if first.Wed.String != second.Wed.String {
		return false
	}
	if first.Thurs.String != second.Thurs.String {
		return false
	}
	if first.Fri.String != second.Fri.String {
		return false
	}
	if first.Sat.String != second.Sat.String {
		return false
	}
	if first.NextSun.String != second.NextSun.String {
		return false
	}
	return true
}
