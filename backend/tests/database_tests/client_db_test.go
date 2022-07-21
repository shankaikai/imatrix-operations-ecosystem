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
	TEST_CLIENT_NAME   = "test name"
	TEST_CLIENT_ABBR   = "abbr"
	TEST_CLIENT_EMAIL  = "test_email"
	TEST_CLIENT_ADDR   = "test addr"
	TEST_CLIENT_PHONE  = "1231321"
	TEST_CLIENT_POSTAL = 14343
)

// Successful Test
func TestInsertClient(t *testing.T) {
	client := createFakeClient(1)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO client").WillReturnResult(sqlmock.NewResult(1, 1))

	pk, err := db_pck.InsertClient(db, client, &sync.Mutex{})

	if err != nil {
		t.Errorf("error was not expected while inserting client: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 1, int(pk), "The primary key returned should be 1")
}

func TestGetClientsNoFilter(t *testing.T) {
	query := &pb.ClientQuery{}
	fakeClients := []*pb.Client{createFakeClient(1), createFakeClient(2)}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"client_id", "name", "abbreviation", "email", "address", "postal", "phone_number"}).
		AddRow(1, TEST_CLIENT_NAME, TEST_CLIENT_ABBR,
			TEST_CLIENT_EMAIL, TEST_CLIENT_ADDR, TEST_CLIENT_POSTAL, TEST_CLIENT_PHONE,
		).
		AddRow(2, TEST_CLIENT_NAME, TEST_CLIENT_ABBR,
			TEST_CLIENT_EMAIL, TEST_CLIENT_ADDR, TEST_CLIENT_POSTAL, TEST_CLIENT_PHONE,
		)

	mock.ExpectQuery("SELECT \\* FROM client").WillReturnRows(rows)

	clients, err := db_pck.GetClients(db, query)

	if err != nil {
		t.Errorf("error was not expected while selecting clients: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 2, len(clients), "The length of clients returned should be 2")
	assert.Equal(t, true, proto.Equal(clients[0], fakeClients[0]), "The first client returned is not equal to the expected.")
	assert.Equal(t, true, proto.Equal(clients[1], fakeClients[1]), "The second client returned is not equal to the expected.")
}

func TestGetClientsIdFilter(t *testing.T) {
	query := &pb.ClientQuery{Filters: make([]*pb.ClientFilter, 0), Limit: 7}
	query.Filters = append(query.Filters, &pb.ClientFilter{
		Field: pb.ClientFilter_CLIENT_ID,
		Comparisons: &pb.Filter{
			Value: "1", Comparison: pb.Filter_EQUAL,
		},
	})

	fakeClients := []*pb.Client{createFakeClient(1), createFakeClient(2)}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"client_id", "name", "abbreviation", "email", "address", "postal", "phone_number"}).
		AddRow(1, TEST_CLIENT_NAME, TEST_CLIENT_ABBR,
			TEST_CLIENT_EMAIL, TEST_CLIENT_ADDR, TEST_CLIENT_POSTAL, TEST_CLIENT_PHONE,
		).
		AddRow(2, TEST_CLIENT_NAME, TEST_CLIENT_ABBR,
			TEST_CLIENT_EMAIL, TEST_CLIENT_ADDR, TEST_CLIENT_POSTAL, TEST_CLIENT_PHONE,
		)

	mock.ExpectQuery("SELECT \\* FROM client WHERE client_id = '1'.*LIMIT 7").WillReturnRows(rows)

	clients, err := db_pck.GetClients(db, query)

	if err != nil {
		t.Errorf("error was not expected while selecting clients: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 2, len(clients), "The length of clients returned should be 2")
	assert.Equal(t, true, proto.Equal(clients[0], fakeClients[0]), "The first client returned is not equal to the expected.")
	assert.Equal(t, true, proto.Equal(clients[1], fakeClients[1]), "The second client returned is not equal to the expected.")
}

func TestUpdateClientsEmail(t *testing.T) {
	client := &pb.Client{
		ClientId: 1,
		Email:    TEST_CLIENT_EMAIL,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("UPDATE client SET .*email='test_email'.*WHERE client_id = '1'").WillReturnResult(sqlmock.NewResult(1, 1))

	numRows, err := db_pck.UpdateClients(db, client)

	if err != nil {
		t.Errorf("error was not expected while updating clients: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 1, int(numRows), "The number of rows updated 1")
}

func TestDeleteClient(t *testing.T) {
	client := &pb.Client{
		ClientId: 1,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("DELETE FROM client WHERE client_id = '1'").WillReturnResult(sqlmock.NewResult(1, 1))

	numRows, err := db_pck.DeleteClient(db, client)

	if err != nil {
		t.Errorf("error was not expected while deleting clients: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 1, int(numRows), "The number of rows deleted is supposed to be 1")
}

// Add new client filters from empty query
func TestAddClientFilterEmpty(t *testing.T) {
	expectedQuery := &pb.ClientQuery{Filters: make([]*pb.ClientFilter, 0)}
	expectedQuery.Filters = append(expectedQuery.Filters, &pb.ClientFilter{
		Field: pb.ClientFilter_CLIENT_ID,
		Comparisons: &pb.Filter{
			Value: "1", Comparison: pb.Filter_EQUAL,
		},
	})

	inputQuery := &pb.ClientQuery{}
	db_pck.AddClientFilter(inputQuery, pb.ClientFilter_CLIENT_ID, pb.Filter_EQUAL, "1")
	// Check if the filter is added correctly
	assert.Equal(t, 1, len(inputQuery.Filters), "The number of filters expected is not correct")
	assert.Equal(t, true, proto.Equal(inputQuery, expectedQuery), "The query is not equal to the expected.")

}

func createFakeClient(id int) *pb.Client {
	return &pb.Client{
		ClientId:     int64(id),
		Name:         TEST_CLIENT_NAME,
		Abbreviation: TEST_CLIENT_ABBR,
		Email:        TEST_CLIENT_EMAIL,
		Address:      TEST_CLIENT_ADDR,
		PostalCode:   TEST_CLIENT_POSTAL,
		PhoneNumber:  TEST_CLIENT_PHONE,
	}
}

func getSingleClientDbRow(id int) *sqlmock.Rows {
	return sqlmock.NewRows([]string{"client_id", "name", "abbreviation", "email", "address", "postal", "phone_number"}).
		AddRow(id, TEST_CLIENT_NAME, TEST_CLIENT_ABBR,
			TEST_CLIENT_EMAIL, TEST_CLIENT_ADDR, TEST_CLIENT_POSTAL, TEST_CLIENT_PHONE,
		)
}
