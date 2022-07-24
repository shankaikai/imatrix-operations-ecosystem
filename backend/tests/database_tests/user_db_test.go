package tests

import (
	"sync"
	"testing"

	db_pck "capstone.operations_ecosystem/backend/database"
	"google.golang.org/protobuf/proto"

	pb "capstone.operations_ecosystem/backend/proto"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

const (
	TEST_USER_NAME        = "test name"
	TEST_USER_EMAIL       = "test_email"
	TEST_USER_PHONE       = "1231321"
	TEST_USER_TELE_HANDLE = "telegram_handle"
	TEST_USER_IMG         = "img"
	TEST_TELE_USER_ID     = 1
)

// Successful Test
func TestInsertUser(t *testing.T) {
	user := createFakeUser(1)
	fullUser := &pb.FullUser{User: user, HashedPassword: "hashedpassword", SecurityString: "fdsfdfds"}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO user").WillReturnResult(sqlmock.NewResult(1, 1))

	pk, err := db_pck.InsertUser(db, fullUser, &sync.Mutex{})

	if err != nil {
		t.Errorf("error was not expected while inserting user: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 1, int(pk), "The primary key returned should be 1")
}

func TestGetUsersNoFilter(t *testing.T) {
	query := &pb.UserQuery{}
	fakeUsers := []*pb.User{createFakeUser(1), createFakeUser(2)}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"user_id", "user_type", "name", "email", "phone_number", "telegram_handle", "user_security_img", "is_part_timer", "tele_user_id"}).
		AddRow(1, "I-Specialist", TEST_USER_NAME, TEST_USER_EMAIL,
			TEST_USER_PHONE, TEST_USER_TELE_HANDLE, TEST_USER_IMG,
			false, TEST_TELE_USER_ID,
		).
		AddRow(2, "I-Specialist", TEST_USER_NAME, TEST_USER_EMAIL,
			TEST_USER_PHONE, TEST_USER_TELE_HANDLE, TEST_USER_IMG,
			false, TEST_TELE_USER_ID,
		)

	mock.ExpectQuery("SELECT \\* FROM user").WillReturnRows(rows)

	users, err := db_pck.GetUsers(db, query, false)

	if err != nil {
		t.Errorf("error was not expected while selecting users: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 2, len(users), "The length of users returned should be 2")
	assert.Equal(t, true, proto.Equal(users[0], fakeUsers[0]), "The first user returned is not equal to the expected.")
	assert.Equal(t, true, proto.Equal(users[1], fakeUsers[1]), "The second user returned is not equal to the expected.")
}

func TestGetUsersIdFilter(t *testing.T) {
	query := &pb.UserQuery{Filters: make([]*pb.UserFilter, 0), Limit: 7}
	query.Filters = append(query.Filters, &pb.UserFilter{
		Field: pb.UserFilter_USER_ID,
		Comparisons: &pb.Filter{
			Value: "1", Comparison: pb.Filter_EQUAL,
		},
	})

	fakeUsers := []*pb.User{createFakeUser(1), createFakeUser(2)}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"user_id", "user_type", "name", "email", "phone_number", "telegram_handle", "user_security_img", "is_part_timer", "tele_user_id"}).
		AddRow(1, "I-Specialist", TEST_USER_NAME, TEST_USER_EMAIL,
			TEST_USER_PHONE, TEST_USER_TELE_HANDLE, TEST_USER_IMG,
			false, TEST_TELE_USER_ID,
		).
		AddRow(2, "I-Specialist", TEST_USER_NAME, TEST_USER_EMAIL,
			TEST_USER_PHONE, TEST_USER_TELE_HANDLE, TEST_USER_IMG,
			false, TEST_TELE_USER_ID,
		)

	mock.ExpectQuery("SELECT \\* FROM user WHERE user_id = '1'.*LIMIT 7").WillReturnRows(rows)

	users, err := db_pck.GetUsers(db, query, false)

	if err != nil {
		t.Errorf("error was not expected while selecting users: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 2, len(users), "The length of users returned should be 2")
	assert.Equal(t, true, proto.Equal(users[0], fakeUsers[0]), "The first user returned is not equal to the expected.")
	assert.Equal(t, true, proto.Equal(users[1], fakeUsers[1]), "The second user returned is not equal to the expected.")
}

func TestGetUsersTeleHandleFilter(t *testing.T) {
	query := &pb.UserQuery{Filters: make([]*pb.UserFilter, 0), Limit: 1}
	query.Filters = append(query.Filters, &pb.UserFilter{
		Field: pb.UserFilter_TELEGRAM_HANDLE,
		Comparisons: &pb.Filter{
			Value: "1", Comparison: pb.Filter_EQUAL,
		},
	})

	fakeUsers := []*pb.User{createFakeUser(1), createFakeUser(2)}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"user_id", "user_type", "name", "email", "phone_number", "telegram_handle", "user_security_img", "is_part_timer", "tele_user_id"}).
		AddRow(1, "I-Specialist", TEST_USER_NAME, TEST_USER_EMAIL,
			TEST_USER_PHONE, TEST_USER_TELE_HANDLE, TEST_USER_IMG,
			false, TEST_TELE_USER_ID,
		)

	mock.ExpectQuery("SELECT \\* FROM user WHERE telegram_handle = '1'.*LIMIT 1").WillReturnRows(rows)

	users, err := db_pck.GetUsers(db, query, false)

	if err != nil {
		t.Errorf("error was not expected while selecting users: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 1, len(users), "The length of users returned should be 2")
	assert.Equal(t, true, proto.Equal(users[0], fakeUsers[0]), "The first user returned is not equal to the expected.")
}

func TestUpdateUsersTeleUserId(t *testing.T) {
	user := &pb.User{
		UserId:     1,
		TeleUserId: TEST_TELE_USER_ID,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("UPDATE user SET .*tele_user_id='1'.*WHERE user_id = '1'").WillReturnResult(sqlmock.NewResult(1, 1))

	numRows, err := db_pck.UpdateUser(db, user, &pb.UserQuery{})

	if err != nil {
		t.Errorf("error was not expected while updating users: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 1, int(numRows), "The number of rows updated 1")
}

func TestDeleteUser(t *testing.T) {
	user := &pb.User{
		UserId: 1,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("DELETE FROM user WHERE user_id = '1'").WillReturnResult(sqlmock.NewResult(1, 1))

	numRows, err := db_pck.DeleteUser(db, user)

	if err != nil {
		t.Errorf("error was not expected while deleting users: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 1, int(numRows), "The number of rows deleted is supposed to be 1")
}

// Add new user filters from empty query
func TestAddUserFilterEmpty(t *testing.T) {
	expectedQuery := &pb.UserQuery{Filters: make([]*pb.UserFilter, 0)}
	expectedQuery.Filters = append(expectedQuery.Filters, &pb.UserFilter{
		Field: pb.UserFilter_USER_ID,
		Comparisons: &pb.Filter{
			Value: "1", Comparison: pb.Filter_EQUAL,
		},
	})

	inputQuery := &pb.UserQuery{}
	db_pck.AddUserFilter(inputQuery, pb.UserFilter_USER_ID, pb.Filter_EQUAL, "1")
	// Check if the filter is added correctly
	assert.Equal(t, 1, len(inputQuery.Filters), "The number of filters expected is not correct")
	assert.Equal(t, true, proto.Equal(inputQuery, expectedQuery), "The query is not equal to the expected.")

}

// Add new user filters from query that already has something
func TestAddUserFilterAlreadyContains(t *testing.T) {
	expectedQuery := &pb.UserQuery{Filters: make([]*pb.UserFilter, 0), Limit: 7}
	expectedQuery.Filters = append(expectedQuery.Filters, &pb.UserFilter{
		Field: pb.UserFilter_USER_ID,
		Comparisons: &pb.Filter{
			Value: "1", Comparison: pb.Filter_EQUAL,
		},
	})
	expectedQuery.Filters = append(expectedQuery.Filters, &pb.UserFilter{
		Field: pb.UserFilter_TELEGRAM_HANDLE,
		Comparisons: &pb.Filter{
			Value: "5", Comparison: pb.Filter_GREATER,
		},
	})

	inputQuery := &pb.UserQuery{Filters: make([]*pb.UserFilter, 0), Limit: 7}
	inputQuery.Filters = append(inputQuery.Filters, &pb.UserFilter{
		Field: pb.UserFilter_USER_ID,
		Comparisons: &pb.Filter{
			Value: "1", Comparison: pb.Filter_EQUAL,
		},
	})

	db_pck.AddUserFilter(inputQuery, pb.UserFilter_TELEGRAM_HANDLE, pb.Filter_GREATER, "5")
	// Check if the filter is added correctly
	assert.Equal(t, 2, len(inputQuery.Filters), "The number of filters expected is not correct")
	assert.Equal(t, true, proto.Equal(inputQuery, expectedQuery), "The query is not equal to the expected.")

}

func createFakeUser(id int) *pb.User {
	return &pb.User{
		UserId:          int64(id),
		UserType:        pb.User_ISPECIALIST,
		Name:            TEST_USER_NAME,
		Email:           TEST_USER_EMAIL,
		PhoneNumber:     TEST_USER_PHONE,
		TelegramHandle:  TEST_USER_TELE_HANDLE,
		UserSecurityImg: TEST_USER_IMG,
		TeleUserId:      TEST_TELE_USER_ID,
		IsPartTimer:     false,
	}
}

func getSingleUserDbRow(id int) *sqlmock.Rows {
	return sqlmock.NewRows([]string{"user_id", "user_type", "name", "email", "phone_number", "telegram_handle", "user_security_img", "is_part_timer", "tele_user_id"}).
		AddRow(id, "I-Specialist", TEST_USER_NAME, TEST_USER_EMAIL,
			TEST_USER_PHONE, TEST_USER_TELE_HANDLE, TEST_USER_IMG,
			false, TEST_TELE_USER_ID,
		)
}
