package tests

import (
	"fmt"
	"sync"
	"testing"
	"time"

	db_pck "capstone.operations_ecosystem/backend/database"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "capstone.operations_ecosystem/backend/proto"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

const (
	TEST_BC_CONTENT   = "test content"
	TEST_BC_ACK       = true
	TEST_BC_REJECTION = false
	TEST_BC_FAKE_TIME = "2022-06-21 18:00:00"
)

func TestInsertBroadcast(t *testing.T) {
	broadcast := createFakeBroadcast(1, true, 2)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO broadcast").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO broadcast_recepients").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO broadcast_recepients").WillReturnResult(sqlmock.NewResult(1, 1))

	pk, err := db_pck.InsertBroadcast(db, broadcast, &sync.Mutex{})

	if err != nil {
		t.Errorf("error was not expected while inserting broadcast: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 1, int(pk), "The primary key returned should be 1")
}

func TestInsertBroadcastRecipients(t *testing.T) {
	broadcastRec := createFakeBroadcastRec(1, true)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO broadcast_recepients").WillReturnResult(sqlmock.NewResult(1, 1))

	pk, err := db_pck.InsertBroadcastRecipient(db, broadcastRec, 1, &sync.Mutex{})

	if err != nil {
		t.Errorf("error was not expected while inserting broadcast recipient: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 1, int(pk), "The primary key returned should be 1")
}

func TestGetBroadcastNoFilter(t *testing.T) {
	query := &pb.BroadcastQuery{}
	fakeBroadcasts := []*pb.Broadcast{createFakeBroadcast(1, true, 1), createFakeBroadcast(2, true, 1)}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	broadcast_rows := sqlmock.NewRows([]string{"broadcast_id", "type", "content", "creation_date", "deadline", "creator",
		"urgency", "broadcast_recipients_id", "related_broadcast", "recipient", "acknowledged",
		"rejected", "last_replied", "aifs_id"}).
		AddRow(1, "Announcement", TEST_BC_CONTENT, TEST_BC_FAKE_TIME, TEST_BC_FAKE_TIME, 1,
			"Low", 1, 1, 1, TEST_BC_ACK, TEST_BC_REJECTION, TEST_BC_FAKE_TIME, 1,
		).
		AddRow(2, "Announcement", TEST_BC_CONTENT, TEST_BC_FAKE_TIME, TEST_BC_FAKE_TIME, 1,
			"Low", 1, 2, 1, TEST_BC_ACK, TEST_BC_REJECTION, TEST_BC_FAKE_TIME, 1,
		)

	mock.ExpectQuery("SELECT \\* FROM broadcast LEFT JOIN broadcast_recepients ON broadcast_id=related_broadcast WHERE broadcast_id IN").
		WillReturnRows(broadcast_rows)
	mock.ExpectQuery("SELECT \\* FROM user WHERE user_id =").WillReturnRows(getSingleUserDbRow(1))
	mock.ExpectQuery("SELECT \\* FROM user WHERE user_id =").WillReturnRows(getSingleUserDbRow(1))
	mock.ExpectQuery("SELECT \\* FROM user WHERE user_id =").WillReturnRows(getSingleUserDbRow(1))
	mock.ExpectQuery("SELECT \\* FROM user WHERE user_id =").WillReturnRows(getSingleUserDbRow(1))

	broadcasts, err := db_pck.GetBroadcasts(db, query)

	if err != nil {
		t.Errorf("error was not expected while selecting broadcast: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 2, len(broadcasts), "The length of broadcast returned should be 2")
	t.Log("returned ", broadcasts[1])
	t.Log("fake", fakeBroadcasts[1])
	assert.Equal(t, true, proto.Equal(broadcasts[0], fakeBroadcasts[0]), "The first broadcast returned is not equal to the expected.")
	assert.Equal(t, true, proto.Equal(broadcasts[1], fakeBroadcasts[1]), "The second broadcast returned is not equal to the expected.")
}

func TestGetBroadcastIdFilter(t *testing.T) {
	query := &pb.BroadcastQuery{Filters: make([]*pb.BroadcastFilter, 0), Limit: 7}
	query.Filters = append(query.Filters, &pb.BroadcastFilter{
		Field: pb.BroadcastFilter_BROADCAST_ID,
		Comparisons: &pb.Filter{
			Value: "1", Comparison: pb.Filter_EQUAL,
		},
	})

	fakeBroadcasts := []*pb.Broadcast{createFakeBroadcast(1, true, 1), createFakeBroadcast(2, true, 1)}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	broadcast_rows := sqlmock.NewRows([]string{"broadcast_id", "type", "content", "creation_date", "deadline", "creator",
		"urgency", "broadcast_recipients_id", "related_broadcast", "recipient", "acknowledged",
		"rejected", "last_replied", "aifs_id"}).
		AddRow(1, "Announcement", TEST_BC_CONTENT, TEST_BC_FAKE_TIME, TEST_BC_FAKE_TIME, 1,
			"Low", 1, 1, 1, TEST_BC_ACK, TEST_BC_REJECTION, TEST_BC_FAKE_TIME, 1,
		).
		AddRow(2, "Announcement", TEST_BC_CONTENT, TEST_BC_FAKE_TIME, TEST_BC_FAKE_TIME, 2,
			"Low", 1, 2, 1, TEST_BC_ACK, TEST_BC_REJECTION, TEST_BC_FAKE_TIME, 1,
		)

	mock.ExpectQuery("SELECT \\* FROM broadcast LEFT JOIN broadcast_recepients ON broadcast_id=related_broadcast WHERE " +
		"broadcast_id IN.*broadcast_id = '1'").
		WillReturnRows(broadcast_rows)
	mock.ExpectQuery("SELECT \\* FROM user WHERE user_id =").WillReturnRows(getSingleUserDbRow(1))
	mock.ExpectQuery("SELECT \\* FROM user WHERE user_id =").WillReturnRows(getSingleUserDbRow(1))
	mock.ExpectQuery("SELECT \\* FROM user WHERE user_id =").WillReturnRows(getSingleUserDbRow(1))
	mock.ExpectQuery("SELECT \\* FROM user WHERE user_id =").WillReturnRows(getSingleUserDbRow(1))

	broadcasts, err := db_pck.GetBroadcasts(db, query)

	if err != nil {
		t.Errorf("error was not expected while selecting broadcast: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 2, len(broadcasts), "The length of broadcast returned should be 2")
	t.Log("returned ", broadcasts[1])
	t.Log("fake", fakeBroadcasts[1])
	assert.Equal(t, true, proto.Equal(broadcasts[0], fakeBroadcasts[0]), "The first broadcast returned is not equal to the expected.")
	assert.Equal(t, true, proto.Equal(broadcasts[1], fakeBroadcasts[1]), "The second broadcast returned is not equal to the expected.")
}

func TestGetBroadcastRecipientsNoFilter(t *testing.T) {
	query := &pb.BroadcastQuery{}

	fakeBroadcasts := []*pb.BroadcastRecipient{createFakeBroadcastRec(1, true), createFakeBroadcastRec(2, true)}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	broadcast_rows := sqlmock.NewRows([]string{"broadcast_recipients_id", "related_broadcast", "recipient", "acknowledged",
		"rejected", "last_replied", "aifs_id"}).
		AddRow(1, 1, 1, TEST_BC_ACK, TEST_BC_REJECTION, TEST_BC_FAKE_TIME, 1).
		AddRow(2, 2, 2, TEST_BC_ACK, TEST_BC_REJECTION, TEST_BC_FAKE_TIME, 2)

	mock.ExpectQuery("SELECT \\* FROM broadcast_recepients").
		WillReturnRows(broadcast_rows)
	mock.ExpectQuery("SELECT \\* FROM user WHERE user_id =").WillReturnRows(getSingleUserDbRow(1))
	mock.ExpectQuery("SELECT \\* FROM user WHERE user_id =").WillReturnRows(getSingleUserDbRow(2))

	broadcasts, err := db_pck.GetBroadcastRecipients(db, query, -1)

	if err != nil {
		t.Errorf("error was not expected while selecting broadcast: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 2, len(broadcasts), "The length of broadcast returned should be 2")
	t.Log("returned ", broadcasts[1])
	t.Log("fake", fakeBroadcasts[1])
	assert.Equal(t, true, proto.Equal(broadcasts[0], fakeBroadcasts[0]), "The first broadcast returned is not equal to the expected.")
	assert.Equal(t, true, proto.Equal(broadcasts[1], fakeBroadcasts[1]), "The second broadcast returned is not equal to the expected.")
}

func TestGetBroadcastRecipientsTableId(t *testing.T) {
	query := &pb.BroadcastQuery{Filters: make([]*pb.BroadcastFilter, 0), Limit: 7}
	query.Filters = append(query.Filters, &pb.BroadcastFilter{
		Field: pb.BroadcastFilter_BROADCAST_RECIPIENT_TABLE_ID,
		Comparisons: &pb.Filter{
			Value: "1", Comparison: pb.Filter_EQUAL,
		},
	})

	fakeBroadcasts := []*pb.BroadcastRecipient{createFakeBroadcastRec(1, true), createFakeBroadcastRec(2, true)}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	broadcast_rows := sqlmock.NewRows([]string{"broadcast_recipients_id", "related_broadcast", "recipient", "acknowledged",
		"rejected", "last_replied", "aifs_id"}).
		AddRow(1, 1, 1, TEST_BC_ACK, TEST_BC_REJECTION, TEST_BC_FAKE_TIME, 1).
		AddRow(2, 2, 2, TEST_BC_ACK, TEST_BC_REJECTION, TEST_BC_FAKE_TIME, 2)

	mock.ExpectQuery("SELECT \\* FROM broadcast_recepients WHERE broadcast_recipients_id = '1'").
		WillReturnRows(broadcast_rows)
	mock.ExpectQuery("SELECT \\* FROM user WHERE user_id =").WillReturnRows(getSingleUserDbRow(1))
	mock.ExpectQuery("SELECT \\* FROM user WHERE user_id =").WillReturnRows(getSingleUserDbRow(2))

	broadcasts, err := db_pck.GetBroadcastRecipients(db, query, -1)

	if err != nil {
		t.Errorf("error was not expected while selecting broadcast: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 2, len(broadcasts), "The length of broadcast returned should be 2")
	t.Log("returned ", broadcasts[1])
	t.Log("fake", fakeBroadcasts[1])
	assert.Equal(t, true, proto.Equal(broadcasts[0], fakeBroadcasts[0]), "The first broadcast returned is not equal to the expected.")
	assert.Equal(t, true, proto.Equal(broadcasts[1], fakeBroadcasts[1]), "The second broadcast returned is not equal to the expected.")
}

func TestGetBroadcastRecipientsMainBroadcastId(t *testing.T) {
	query := &pb.BroadcastQuery{}

	fakeBroadcasts := []*pb.BroadcastRecipient{createFakeBroadcastRec(1, true), createFakeBroadcastRec(2, true)}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	broadcast_rows := sqlmock.NewRows([]string{"broadcast_recipients_id", "related_broadcast", "recipient", "acknowledged",
		"rejected", "last_replied", "aifs_id"}).
		AddRow(1, 1, 1, TEST_BC_ACK, TEST_BC_REJECTION, TEST_BC_FAKE_TIME, 1).
		AddRow(2, 2, 2, TEST_BC_ACK, TEST_BC_REJECTION, TEST_BC_FAKE_TIME, 2)

	mock.ExpectQuery("SELECT \\* FROM broadcast_recepients WHERE related_broadcast = '1'").
		WillReturnRows(broadcast_rows)
	mock.ExpectQuery("SELECT \\* FROM user WHERE user_id =").WillReturnRows(getSingleUserDbRow(1))
	mock.ExpectQuery("SELECT \\* FROM user WHERE user_id =").WillReturnRows(getSingleUserDbRow(2))

	broadcasts, err := db_pck.GetBroadcastRecipients(db, query, 1)

	if err != nil {
		t.Errorf("error was not expected while selecting broadcast: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 2, len(broadcasts), "The length of broadcast returned should be 2")
	t.Log("returned ", broadcasts[1])
	t.Log("fake", fakeBroadcasts[1])
	assert.Equal(t, true, proto.Equal(broadcasts[0], fakeBroadcasts[0]), "The first broadcast returned is not equal to the expected.")
	assert.Equal(t, true, proto.Equal(broadcasts[1], fakeBroadcasts[1]), "The second broadcast returned is not equal to the expected.")
}

func TestGetBroadcastRecipientUserId(t *testing.T) {
	query := &pb.BroadcastQuery{Filters: make([]*pb.BroadcastFilter, 0), Limit: 7}
	query.Filters = append(query.Filters, &pb.BroadcastFilter{
		Field: pb.BroadcastFilter_RECEIPEIENT_ID,
		Comparisons: &pb.Filter{
			Value: "1", Comparison: pb.Filter_EQUAL,
		},
	})

	fakeBroadcasts := []*pb.BroadcastRecipient{createFakeBroadcastRec(1, true), createFakeBroadcastRec(2, true)}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	broadcast_rows := sqlmock.NewRows([]string{"broadcast_recipients_id", "related_broadcast", "recipient", "acknowledged",
		"rejected", "last_replied", "aifs_id"}).
		AddRow(1, 1, 1, TEST_BC_ACK, TEST_BC_REJECTION, TEST_BC_FAKE_TIME, 1).
		AddRow(2, 2, 2, TEST_BC_ACK, TEST_BC_REJECTION, TEST_BC_FAKE_TIME, 2)

	mock.ExpectQuery("SELECT \\* FROM broadcast_recepients WHERE recipient = '1'").
		WillReturnRows(broadcast_rows)
	mock.ExpectQuery("SELECT \\* FROM user WHERE user_id =").WillReturnRows(getSingleUserDbRow(1))
	mock.ExpectQuery("SELECT \\* FROM user WHERE user_id =").WillReturnRows(getSingleUserDbRow(2))

	broadcasts, err := db_pck.GetBroadcastRecipients(db, query, -1)

	if err != nil {
		t.Errorf("error was not expected while selecting broadcast: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 2, len(broadcasts), "The length of broadcast returned should be 2")
	t.Log("returned ", broadcasts[1])
	t.Log("fake", fakeBroadcasts[1])
	assert.Equal(t, true, proto.Equal(broadcasts[0], fakeBroadcasts[0]), "The first broadcast returned is not equal to the expected.")
	assert.Equal(t, true, proto.Equal(broadcasts[1], fakeBroadcasts[1]), "The second broadcast returned is not equal to the expected.")
}

func TestUpdateBroadcastNotRecipients(t *testing.T) {
	broadcast := &pb.Broadcast{
		BroadcastId: 1,
		Content:     TEST_BC_CONTENT,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("UPDATE broadcast SET .*content=.*WHERE broadcast_id = '1'").WillReturnResult(sqlmock.NewResult(1, 1))

	numRows, err := db_pck.UpdateBroadcast(db, broadcast, &sync.Mutex{})

	if err != nil {
		t.Errorf("error was not expected while updating broadcast: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 1, int(numRows), "The number of rows updated 1")
}

// TODO:
// func UpdateBroadcast(db *sql.DB, broadcast *pb.Broadcast, dbLock *sync.Mutex) (int64, error) {
// 			add users and delete users

func TestUpdateBroadcastChangeRecipients(t *testing.T) {
	broadcast := createFakeBroadcast(1, true, 2)
	t.Log("created", broadcast)
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	broadcast_rows := sqlmock.NewRows([]string{"broadcast_recipients_id", "related_broadcast", "recipient", "acknowledged",
		"rejected", "last_replied", "aifs_id"}).
		AddRow(1, 1, 3, TEST_BC_ACK, TEST_BC_REJECTION, TEST_BC_FAKE_TIME, 1)

	mock.ExpectExec("UPDATE broadcast SET .*WHERE broadcast_id = '1'").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectQuery("SELECT \\* FROM broadcast_recepients WHERE related_broadcast = '1'").
		WillReturnRows(broadcast_rows)
	mock.ExpectQuery("SELECT \\* FROM user WHERE user_id =").WillReturnRows(getSingleUserDbRow(3))
	mock.ExpectExec("INSERT INTO broadcast_recepients.*VALUES \\(1,1").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO broadcast_recepients.*VALUES \\(1,2").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("DELETE FROM broadcast_recepients WHERE broadcast_recipients_id=1").WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("UPDATE broadcast SET .*content=.*WHERE broadcast_id = '1'").WillReturnResult(sqlmock.NewResult(1, 1))

	numRows, err := db_pck.UpdateBroadcast(db, broadcast, &sync.Mutex{})

	if err != nil {
		t.Errorf("error was not expected while updating broadcast: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 1, int(numRows), "The number of rows updated 1")
}

func TestUpdateBroadcastRecipients(t *testing.T) {
	broadcastRec := &pb.BroadcastRecipient{
		BroadcastRecipientsId: 1,
		Acknowledged:          true,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("UPDATE broadcast_recepients SET acknowledged=true.*WHERE broadcast_recipients_id = '1'").WillReturnResult(sqlmock.NewResult(1, 1))

	numRows, err := db_pck.UpdateBroadcastRecipients(db, broadcastRec)

	if err != nil {
		t.Errorf("error was not expected while updating broadcast: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 1, int(numRows), "The number of rows updated 1")
}

func TestDeleteBroadcast(t *testing.T) {
	broadcast := &pb.Broadcast{
		BroadcastId: 1,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("DELETE FROM broadcast WHERE broadcast_id = '1'").WillReturnResult(sqlmock.NewResult(1, 1))

	numRows, err := db_pck.DeleteBroadcast(db, broadcast)

	if err != nil {
		t.Errorf("error was not expected while deleting broadcast: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 1, int(numRows), "The number of rows deleted is supposed to be 1")
}

func TestDeleteBroadcastRecipient(t *testing.T) {
	broadcast := &pb.BroadcastRecipient{
		BroadcastRecipientsId: 1,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("DELETE FROM broadcast_recepients WHERE broadcast_recipients_id=1").WillReturnResult(sqlmock.NewResult(1, 1))

	numRows, err := db_pck.DeleteBroadcastRecipients(db, broadcast)

	if err != nil {
		t.Errorf("error was not expected while deleting broadcast: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 1, int(numRows), "The number of rows deleted is supposed to be 1")
}

// Add new broadcast filters from empty query
func TestAddBroadcastFilterEmpty(t *testing.T) {
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

// Add new broadcast filters from query that already has something
func TestAddBroadcastFilterAlreadyContains(t *testing.T) {
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

func createFakeBroadcast(id int, hasRecipient bool, numRec int) *pb.Broadcast {
	recipients := make([]*pb.BroadcastRecipient, 0)
	aifsRecipients := make([]*pb.AIFSBroadcastRecipient, 0)
	for i := 1; i < numRec+1; i++ {
		recipients = append(recipients, createFakeBroadcastRec(i, hasRecipient))
	}

	for i := 1; i < numRec+1; i++ {
		aifsRecipients = append(aifsRecipients, &pb.AIFSBroadcastRecipient{
			AifsId:    int64(i),
			Recipient: recipients[i-1 : i],
		})
	}
	fmt.Println("HERE", aifsRecipients[0].Recipient)
	fakeTime := time.Date(2022, 6, 21, 18, 0, 0, 0, time.UTC)

	return &pb.Broadcast{
		BroadcastId:  int64(id),
		Type:         pb.Broadcast_ANNOUNCEMENT,
		Content:      TEST_BC_CONTENT,
		CreationDate: &timestamppb.Timestamp{Seconds: fakeTime.Unix()},
		Deadline:     &timestamppb.Timestamp{Seconds: fakeTime.Unix()},
		Creator:      createFakeUser(1),
		Recipients:   aifsRecipients,
		Urgency:      pb.Broadcast_LOW,
	}
}

func createFakeBroadcastRec(id int, hasUser bool) *pb.BroadcastRecipient {
	var user *pb.User
	if hasUser {
		user = createFakeUser(id)
	}

	fakeTime := time.Date(2022, 6, 21, 18, 0, 0, 0, time.UTC)

	return &pb.BroadcastRecipient{
		BroadcastRecipientsId: int64(id),
		Recipient:             user,
		Acknowledged:          TEST_BC_ACK,
		Rejected:              TEST_BC_REJECTION,
		AifsId:                int64(id),
		LastReplied:           &timestamppb.Timestamp{Seconds: fakeTime.Unix()},
	}
}
