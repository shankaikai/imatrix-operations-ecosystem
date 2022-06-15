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
	TEST_TELE_CHAT_ID     = 1
)

// Successful Test
func TestInsertUser(t *testing.T) {
	user := createFakeUser(1)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO user").WillReturnResult(sqlmock.NewResult(1, 1))

	pk, err := db_pck.InsertUser(db, user, &sync.Mutex{})

	// now we execute our method
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

	rows := sqlmock.NewRows([]string{"user_id", "user_type", "name", "email", "phone_number", "telegram_handle", "user_security_img", "is_part_timer", "tele_chat_id"}).
		AddRow(1, "I-Specialist", TEST_USER_NAME, TEST_USER_EMAIL,
			TEST_USER_PHONE, TEST_USER_TELE_HANDLE, TEST_USER_IMG,
			false, TEST_TELE_CHAT_ID,
		).
		AddRow(2, "I-Specialist", TEST_USER_NAME, TEST_USER_EMAIL,
			TEST_USER_PHONE, TEST_USER_TELE_HANDLE, TEST_USER_IMG,
			false, TEST_TELE_CHAT_ID,
		)

	mock.ExpectQuery("SELECT \\* FROM user").WillReturnRows(rows)

	users, err := db_pck.GetUsers(db, query)

	// now we execute our method
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

// TODO:
// GetUsers with user id filter
// GetUsers with telegram handle
// Update user with chat id
// Delete user just test
// Add user filter empty query
// Add user filter with existing query

// Add new user filters from empty query
func TestAddUserFilter(t *testing.T) {
	query := &pb.UserQuery{}
	db_pck.AddUserFilter(query, pb.UserFilter_USER_ID, pb.Filter_EQUAL, "1")
	// TODO: check if the query has 1 new filter
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
		TeleChatId:      TEST_TELE_CHAT_ID,
		IsPartTimer:     false,
	}
}
