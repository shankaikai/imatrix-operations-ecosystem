package tests

import (
	"sync"
	"testing"
	"time"

	"capstone.operations_ecosystem/backend/common"
	db_pck "capstone.operations_ecosystem/backend/database"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "capstone.operations_ecosystem/backend/proto"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

const (
	TEST_ROSTER_FAKE_START_TIME = "2022-06-21 18:00:00"
	TEST_ROSTER_FAKE_END_TIME   = "2022-06-22 06:00:00"

	TEST_ROSTER_ASGN_DB_CONFIRMATION = false
	TEST_ROSTER_ASGN_DB_ATTENDED     = false
	TEST_ROSTER_ASGN_DB_IS_ASSIGNED  = false
	TEST_ROSTER_ASGN_DB_REJECTED     = false
	TEST_AIFS_CLIENT_DB_PATROL_ORDER = 1
)

func TestInsertRoster(t *testing.T) {
	roster := CreateFakeRoster(1)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rosterRows := sqlmock.NewRows(getFullRosterCols())

	mock.ExpectQuery("SELECT \\* FROM schedule LEFT JOIN schedule_detail ON schedule_id=schedule_detail.schedule LEFT " +
		"JOIN aifs_client_schedule ON schedule_id=aifs_client_schedule.schedule WHERE aifs_id = '1' AND start_time =.* AND " +
		"end_time =").WillReturnRows(rosterRows)
	mock.ExpectExec("INSERT INTO schedule").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO schedule_detail").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO aifs_client_schedule").WillReturnResult(sqlmock.NewResult(1, 1))

	pk, err := db_pck.InsertRoster(db, roster, &sync.Mutex{})

	if err != nil {
		t.Errorf("error was not expected while inserting roster: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 1, int(pk), "The primary key returned should be 1")
}

func TestInsertRosterASGN(t *testing.T) {
	assignment := createFakeRosterAssignment(1)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO schedule_detail").WillReturnResult(sqlmock.NewResult(1, 1))

	pk, err := db_pck.InsertRosterASGN(db, assignment, 1, &sync.Mutex{})

	if err != nil {
		t.Errorf("error was not expected while inserting roster assignment: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 1, int(pk), "The primary key returned should be 1")
}

func TestInsertAIFSClientRoster(t *testing.T) {
	aifsClient := createFakeAifsAssignment(1)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO aifs_client_schedule").WillReturnResult(sqlmock.NewResult(1, 1))

	pk, err := db_pck.InsertAIFSClientRoster(db, aifsClient, 1, &sync.Mutex{})

	if err != nil {
		t.Errorf("error was not expected while inserting aifs roster: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 1, int(pk), "The primary key returned should be 1")
}

func TestGetRostersNoFilter(t *testing.T) {
	query := &pb.RosterQuery{}
	fakeRosters := []*pb.Roster{CreateFakeRoster(1), CreateFakeRoster(2)}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rosterRows := sqlmock.NewRows(getFullRosterCols()).
		AddRow(1, 1, TEST_ROSTER_FAKE_START_TIME, TEST_ROSTER_FAKE_END_TIME, 1, 1, 1, TEST_ROSTER_FAKE_START_TIME, TEST_ROSTER_FAKE_END_TIME,
			TEST_ROSTER_ASGN_DB_CONFIRMATION, TEST_ROSTER_ASGN_DB_ATTENDED, nil, TEST_ROSTER_ASGN_DB_IS_ASSIGNED,
			TEST_ROSTER_ASGN_DB_REJECTED, 1, 1, 1, TEST_AIFS_CLIENT_DB_PATROL_ORDER,
		).
		AddRow(2, 2, TEST_ROSTER_FAKE_START_TIME, TEST_ROSTER_FAKE_END_TIME, 1, 1, 1, TEST_ROSTER_FAKE_START_TIME, TEST_ROSTER_FAKE_END_TIME,
			TEST_ROSTER_ASGN_DB_CONFIRMATION, TEST_ROSTER_ASGN_DB_ATTENDED, nil, TEST_ROSTER_ASGN_DB_IS_ASSIGNED,
			TEST_ROSTER_ASGN_DB_REJECTED, 1, 1, 1, TEST_AIFS_CLIENT_DB_PATROL_ORDER,
		)

	mock.ExpectQuery("SELECT \\* FROM schedule LEFT JOIN schedule_detail ON schedule_id=schedule_detail.schedule LEFT " +
		"JOIN aifs_client_schedule ON schedule_id=aifs_client_schedule.schedule").WillReturnRows(rosterRows)
	mock.ExpectQuery("SELECT \\* FROM user WHERE user_id =").WillReturnRows(getSingleUserDbRow(1))
	mock.ExpectQuery("SELECT \\* FROM client WHERE client_id =").WillReturnRows(getSingleClientDbRow(1))
	mock.ExpectQuery("SELECT \\* FROM user WHERE user_id =").WillReturnRows(getSingleUserDbRow(1))
	mock.ExpectQuery("SELECT \\* FROM client WHERE client_id =").WillReturnRows(getSingleClientDbRow(1))

	rosters, err := db_pck.GetRosters(db, query)

	if err != nil {
		t.Errorf("error was not expected while selecting rosters: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 2, len(rosters), "The length of rosters returned should be 2")
	t.Log("returned ", rosters[1])
	t.Log("fake", fakeRosters[1])
	assert.Equal(t, true, proto.Equal(rosters[0], fakeRosters[0]), "The first roster returned is not equal to the expected.")
	assert.Equal(t, true, proto.Equal(rosters[1], fakeRosters[1]), "The second roster returned is not equal to the expected.")
}

func TestGetRostersIdFilter(t *testing.T) {
	query := &pb.RosterQuery{Filters: make([]*pb.RosterFilter, 0), Limit: 7}
	query.Filters = append(query.Filters, &pb.RosterFilter{
		Field: pb.RosterFilter_ROSTER_ID,
		Comparisons: &pb.Filter{
			Value: "1", Comparison: pb.Filter_EQUAL,
		},
	})

	fakeRosters := []*pb.Roster{CreateFakeRoster(1), CreateFakeRoster(2)}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rosterRows := sqlmock.NewRows(getFullRosterCols()).
		AddRow(1, 1, TEST_ROSTER_FAKE_START_TIME, TEST_ROSTER_FAKE_END_TIME, 1, 1, 1, TEST_ROSTER_FAKE_START_TIME, TEST_ROSTER_FAKE_END_TIME,
			TEST_ROSTER_ASGN_DB_CONFIRMATION, TEST_ROSTER_ASGN_DB_ATTENDED, nil, TEST_ROSTER_ASGN_DB_IS_ASSIGNED,
			TEST_ROSTER_ASGN_DB_REJECTED, 1, 1, 1, TEST_AIFS_CLIENT_DB_PATROL_ORDER,
		).
		AddRow(2, 2, TEST_ROSTER_FAKE_START_TIME, TEST_ROSTER_FAKE_END_TIME, 1, 1, 1, TEST_ROSTER_FAKE_START_TIME, TEST_ROSTER_FAKE_END_TIME,
			TEST_ROSTER_ASGN_DB_CONFIRMATION, TEST_ROSTER_ASGN_DB_ATTENDED, nil, TEST_ROSTER_ASGN_DB_IS_ASSIGNED,
			TEST_ROSTER_ASGN_DB_REJECTED, 1, 1, 1, TEST_AIFS_CLIENT_DB_PATROL_ORDER,
		)

	mock.ExpectQuery("SELECT \\* FROM schedule LEFT JOIN schedule_detail ON schedule_id=schedule_detail.schedule LEFT " +
		"JOIN aifs_client_schedule ON schedule_id=aifs_client_schedule.schedule WHERE schedule_id = '1'").
		WillReturnRows(rosterRows)
	mock.ExpectQuery("SELECT \\* FROM user WHERE user_id =").WillReturnRows(getSingleUserDbRow(1))
	mock.ExpectQuery("SELECT \\* FROM client WHERE client_id =").WillReturnRows(getSingleClientDbRow(1))
	mock.ExpectQuery("SELECT \\* FROM user WHERE user_id =").WillReturnRows(getSingleUserDbRow(1))
	mock.ExpectQuery("SELECT \\* FROM client WHERE client_id =").WillReturnRows(getSingleClientDbRow(1))

	rosters, err := db_pck.GetRosters(db, query)

	if err != nil {
		t.Errorf("error was not expected while selecting rosters: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 2, len(rosters), "The length of rosters returned should be 2")
	t.Log("returned ", rosters[1])
	t.Log("fake", fakeRosters[1])
	assert.Equal(t, true, proto.Equal(rosters[0], fakeRosters[0]), "The first roster returned is not equal to the expected.")
	assert.Equal(t, true, proto.Equal(rosters[1], fakeRosters[1]), "The second roster returned is not equal to the expected.")
}

func TestGetRostersStartDayFilter(t *testing.T) {
	query := &pb.RosterQuery{Filters: make([]*pb.RosterFilter, 0), Limit: 7}
	query.Filters = append(query.Filters, &pb.RosterFilter{
		Field: pb.RosterFilter_START_TIME,
		Comparisons: &pb.Filter{
			Value: TEST_ROSTER_FAKE_START_TIME, Comparison: pb.Filter_EQUAL,
		},
	})

	fakeRosters := []*pb.Roster{CreateFakeRoster(1), CreateFakeRoster(2)}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rosterRows := sqlmock.NewRows(getFullRosterCols()).
		AddRow(1, 1, TEST_ROSTER_FAKE_START_TIME, TEST_ROSTER_FAKE_END_TIME, 1, 1, 1, TEST_ROSTER_FAKE_START_TIME, TEST_ROSTER_FAKE_END_TIME,
			TEST_ROSTER_ASGN_DB_CONFIRMATION, TEST_ROSTER_ASGN_DB_ATTENDED, nil, TEST_ROSTER_ASGN_DB_IS_ASSIGNED,
			TEST_ROSTER_ASGN_DB_REJECTED, 1, 1, 1, TEST_AIFS_CLIENT_DB_PATROL_ORDER,
		).
		AddRow(2, 2, TEST_ROSTER_FAKE_START_TIME, TEST_ROSTER_FAKE_END_TIME, 1, 1, 1, TEST_ROSTER_FAKE_START_TIME, TEST_ROSTER_FAKE_END_TIME,
			TEST_ROSTER_ASGN_DB_CONFIRMATION, TEST_ROSTER_ASGN_DB_ATTENDED, nil, TEST_ROSTER_ASGN_DB_IS_ASSIGNED,
			TEST_ROSTER_ASGN_DB_REJECTED, 1, 1, 1, TEST_AIFS_CLIENT_DB_PATROL_ORDER,
		)

	mock.ExpectQuery("SELECT \\* FROM schedule LEFT JOIN schedule_detail ON schedule_id=schedule_detail.schedule LEFT " +
		"JOIN aifs_client_schedule ON schedule_id=aifs_client_schedule.schedule WHERE start_time = '" +
		TEST_ROSTER_FAKE_START_TIME + "'").
		WillReturnRows(rosterRows)
	mock.ExpectQuery("SELECT \\* FROM user WHERE user_id =").WillReturnRows(getSingleUserDbRow(1))
	mock.ExpectQuery("SELECT \\* FROM client WHERE client_id =").WillReturnRows(getSingleClientDbRow(1))
	mock.ExpectQuery("SELECT \\* FROM user WHERE user_id =").WillReturnRows(getSingleUserDbRow(1))
	mock.ExpectQuery("SELECT \\* FROM client WHERE client_id =").WillReturnRows(getSingleClientDbRow(1))

	rosters, err := db_pck.GetRosters(db, query)

	if err != nil {
		t.Errorf("error was not expected while selecting rosters: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 2, len(rosters), "The length of rosters returned should be 2")
	t.Log("returned ", rosters[1])
	t.Log("fake", fakeRosters[1])
	assert.Equal(t, true, proto.Equal(rosters[0], fakeRosters[0]), "The first roster returned is not equal to the expected.")
	assert.Equal(t, true, proto.Equal(rosters[1], fakeRosters[1]), "The second roster returned is not equal to the expected.")
}

func TestGetDefaultRostersStartDayFilter(t *testing.T) {
	query := &pb.RosterQuery{Filters: make([]*pb.RosterFilter, 0), Limit: 7}
	query.Filters = append(query.Filters, &pb.RosterFilter{
		Field: pb.RosterFilter_START_TIME,
		Comparisons: &pb.Filter{
			Value: TEST_ROSTER_FAKE_START_TIME, Comparison: pb.Filter_EQUAL,
		},
	})

	fakeRosters := []*pb.Roster{CreateFakeRoster(1), CreateFakeRoster(2), CreateFakeRoster(3)}
	fakeRosters[0].Status = pb.Roster_IS_DEFAULT
	fakeRosters[1].Status = pb.Roster_IS_DEFAULT
	fakeRosters[2].Status = pb.Roster_IS_DEFAULT

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	defaultRosteringRows := sqlmock.NewRows([]string{"defaultRosteringId", "day_of_week", "schedule1", "schedule2", "schedule3"}).
		AddRow(1, 1, 1, 2, 3)

	rosterRows := sqlmock.NewRows(getFullRosterCols()).
		AddRow(1, 1, TEST_ROSTER_FAKE_START_TIME, TEST_ROSTER_FAKE_END_TIME, 1, 1, 1, TEST_ROSTER_FAKE_START_TIME, TEST_ROSTER_FAKE_END_TIME,
			TEST_ROSTER_ASGN_DB_CONFIRMATION, TEST_ROSTER_ASGN_DB_ATTENDED, nil, TEST_ROSTER_ASGN_DB_IS_ASSIGNED,
			TEST_ROSTER_ASGN_DB_REJECTED, 1, 1, 1, TEST_AIFS_CLIENT_DB_PATROL_ORDER,
		).
		AddRow(2, 2, TEST_ROSTER_FAKE_START_TIME, TEST_ROSTER_FAKE_END_TIME, 1, 1, 1, TEST_ROSTER_FAKE_START_TIME, TEST_ROSTER_FAKE_END_TIME,
			TEST_ROSTER_ASGN_DB_CONFIRMATION, TEST_ROSTER_ASGN_DB_ATTENDED, nil, TEST_ROSTER_ASGN_DB_IS_ASSIGNED,
			TEST_ROSTER_ASGN_DB_REJECTED, 1, 1, 1, TEST_AIFS_CLIENT_DB_PATROL_ORDER,
		)

	mock.ExpectQuery("SELECT \\* FROM default_rostering WHERE day_of_week =").WillReturnRows(defaultRosteringRows)
	mock.ExpectQuery("SELECT \\* FROM schedule LEFT JOIN schedule_detail ON schedule_id=schedule_detail.schedule LEFT " +
		"JOIN aifs_client_schedule ON schedule_id=aifs_client_schedule.schedule WHERE schedule_id IN").
		WillReturnRows(rosterRows)
	mock.ExpectQuery("SELECT \\* FROM user WHERE user_id =").WillReturnRows(getSingleUserDbRow(1))
	mock.ExpectQuery("SELECT \\* FROM client WHERE client_id =").WillReturnRows(getSingleClientDbRow(1))
	mock.ExpectQuery("SELECT \\* FROM user WHERE user_id =").WillReturnRows(getSingleUserDbRow(1))
	mock.ExpectQuery("SELECT \\* FROM client WHERE client_id =").WillReturnRows(getSingleClientDbRow(1))

	rosters, err := db_pck.GetDefaultRosters(db, query)

	if err != nil {
		t.Errorf("error was not expected while selecting rosters: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 2, len(rosters), "The length of rosters returned should be 2")
	t.Log("returned ", rosters[1])
	t.Log("fake", fakeRosters[1])
	assert.Equal(t, true, proto.Equal(rosters[0], fakeRosters[0]), "The first roster returned is not equal to the expected.")
	assert.Equal(t, true, proto.Equal(rosters[1], fakeRosters[1]), "The second roster returned is not equal to the expected.")
}

func TestGetRostersAssingmentsIdFilter(t *testing.T) {
	query := &pb.RosterQuery{Filters: make([]*pb.RosterFilter, 0), Limit: 7}
	query.Filters = append(query.Filters, &pb.RosterFilter{
		Field: pb.RosterFilter_ROSTER_ASSIGNMENT_ID,
		Comparisons: &pb.Filter{
			Value: "1", Comparison: pb.Filter_EQUAL,
		},
	})

	fakeRosters := []*pb.RosterAssignement{createFakeRosterAssignment(1), createFakeRosterAssignment(2)}
	fakeRosters[0].GuardAssigned.EmployeeScore = 0
	fakeRosters[1].GuardAssigned.EmployeeScore = 0

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rosterRows := sqlmock.NewRows(getRosterAssignmentCols()).
		AddRow(1, 1, 1, TEST_ROSTER_FAKE_START_TIME, TEST_ROSTER_FAKE_END_TIME,
			TEST_ROSTER_ASGN_DB_CONFIRMATION, TEST_ROSTER_ASGN_DB_ATTENDED, nil, TEST_ROSTER_ASGN_DB_IS_ASSIGNED,
			TEST_ROSTER_ASGN_DB_REJECTED,
		).
		AddRow(2, 1, 1, TEST_ROSTER_FAKE_START_TIME, TEST_ROSTER_FAKE_END_TIME,
			TEST_ROSTER_ASGN_DB_CONFIRMATION, TEST_ROSTER_ASGN_DB_ATTENDED, nil, TEST_ROSTER_ASGN_DB_IS_ASSIGNED,
			TEST_ROSTER_ASGN_DB_REJECTED,
		)

	mock.ExpectQuery("SELECT \\* FROM schedule_detail WHERE schedule_detail_id = '1'").
		WillReturnRows(rosterRows)
	mock.ExpectQuery("SELECT \\* FROM user WHERE user_id =").WillReturnRows(getSingleUserDbRow(1))
	mock.ExpectQuery("SELECT \\* FROM user WHERE user_id =").WillReturnRows(getSingleUserDbRow(2))

	rosters, err := db_pck.GetRosterAssingments(db, query, -1)

	if err != nil {
		t.Errorf("error was not expected while selecting rosters: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 2, len(rosters), "The length of rosters returned should be 2")
	t.Log("returned ", rosters[1])
	t.Log("fake", fakeRosters[1])
	assert.Equal(t, true, proto.Equal(rosters[0], fakeRosters[0]), "The first roster returned is not equal to the expected.")
	assert.Equal(t, true, proto.Equal(rosters[1], fakeRosters[1]), "The second roster returned is not equal to the expected.")
}

func TestGetRostersAssingmentsGuardFilter(t *testing.T) {
	query := &pb.RosterQuery{Filters: make([]*pb.RosterFilter, 0), Limit: 7}
	query.Filters = append(query.Filters, &pb.RosterFilter{
		Field: pb.RosterFilter_GUARD_ASSIGNED_ID,
		Comparisons: &pb.Filter{
			Value: "1", Comparison: pb.Filter_EQUAL,
		},
	})

	fakeRosters := []*pb.RosterAssignement{createFakeRosterAssignment(1), createFakeRosterAssignment(2)}
	fakeRosters[0].GuardAssigned.EmployeeScore = 0
	fakeRosters[1].GuardAssigned.EmployeeScore = 0

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rosterRows := sqlmock.NewRows(getRosterAssignmentCols()).
		AddRow(1, 1, 1, TEST_ROSTER_FAKE_START_TIME, TEST_ROSTER_FAKE_END_TIME,
			TEST_ROSTER_ASGN_DB_CONFIRMATION, TEST_ROSTER_ASGN_DB_ATTENDED, nil, TEST_ROSTER_ASGN_DB_IS_ASSIGNED,
			TEST_ROSTER_ASGN_DB_REJECTED,
		).
		AddRow(2, 1, 1, TEST_ROSTER_FAKE_START_TIME, TEST_ROSTER_FAKE_END_TIME,
			TEST_ROSTER_ASGN_DB_CONFIRMATION, TEST_ROSTER_ASGN_DB_ATTENDED, nil, TEST_ROSTER_ASGN_DB_IS_ASSIGNED,
			TEST_ROSTER_ASGN_DB_REJECTED,
		)

	mock.ExpectQuery("SELECT \\* FROM schedule_detail WHERE guard_assigned = '1'").
		WillReturnRows(rosterRows)
	mock.ExpectQuery("SELECT \\* FROM user WHERE user_id =").WillReturnRows(getSingleUserDbRow(1))
	mock.ExpectQuery("SELECT \\* FROM user WHERE user_id =").WillReturnRows(getSingleUserDbRow(2))

	rosters, err := db_pck.GetRosterAssingments(db, query, -1)

	if err != nil {
		t.Errorf("error was not expected while selecting rosters: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 2, len(rosters), "The length of rosters returned should be 2")
	t.Log("returned ", rosters[1])
	t.Log("fake", fakeRosters[1])
	assert.Equal(t, true, proto.Equal(rosters[0], fakeRosters[0]), "The first roster returned is not equal to the expected.")
	assert.Equal(t, true, proto.Equal(rosters[1], fakeRosters[1]), "The second roster returned is not equal to the expected.")
}

func TestGetRostersAssingmentsMainRosterId(t *testing.T) {
	query := &pb.RosterQuery{Filters: make([]*pb.RosterFilter, 0), Limit: 7}
	query.Filters = append(query.Filters, &pb.RosterFilter{
		Field: pb.RosterFilter_ROSTER_ID,
		Comparisons: &pb.Filter{
			Value: "1", Comparison: pb.Filter_EQUAL,
		},
	})

	fakeRosters := []*pb.RosterAssignement{createFakeRosterAssignment(1), createFakeRosterAssignment(2)}
	fakeRosters[0].GuardAssigned.EmployeeScore = 0
	fakeRosters[1].GuardAssigned.EmployeeScore = 0

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rosterRows := sqlmock.NewRows(getRosterAssignmentCols()).
		AddRow(1, 1, 1, TEST_ROSTER_FAKE_START_TIME, TEST_ROSTER_FAKE_END_TIME,
			TEST_ROSTER_ASGN_DB_CONFIRMATION, TEST_ROSTER_ASGN_DB_ATTENDED, nil, TEST_ROSTER_ASGN_DB_IS_ASSIGNED,
			TEST_ROSTER_ASGN_DB_REJECTED,
		).
		AddRow(2, 1, 1, TEST_ROSTER_FAKE_START_TIME, TEST_ROSTER_FAKE_END_TIME,
			TEST_ROSTER_ASGN_DB_CONFIRMATION, TEST_ROSTER_ASGN_DB_ATTENDED, nil, TEST_ROSTER_ASGN_DB_IS_ASSIGNED,
			TEST_ROSTER_ASGN_DB_REJECTED,
		)

	mock.ExpectQuery("SELECT \\* FROM schedule_detail WHERE schedule = '1'").
		WillReturnRows(rosterRows)
	mock.ExpectQuery("SELECT \\* FROM user WHERE user_id =").WillReturnRows(getSingleUserDbRow(1))
	mock.ExpectQuery("SELECT \\* FROM user WHERE user_id =").WillReturnRows(getSingleUserDbRow(2))

	rosters, err := db_pck.GetRosterAssingments(db, query, -1)

	if err != nil {
		t.Errorf("error was not expected while selecting rosters: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 2, len(rosters), "The length of rosters returned should be 2")
	t.Log("returned ", rosters[1])
	t.Log("fake", fakeRosters[1])
	assert.Equal(t, true, proto.Equal(rosters[0], fakeRosters[0]), "The first roster returned is not equal to the expected.")
	assert.Equal(t, true, proto.Equal(rosters[1], fakeRosters[1]), "The second roster returned is not equal to the expected.")
}

func TestGetRosterAIFSClientNoFilter(t *testing.T) {
	query := &pb.RosterQuery{Filters: make([]*pb.RosterFilter, 0), Limit: 7}

	fakeRosters := []*pb.AIFSClientRoster{createFakeAifsAssignment(1), createFakeAifsAssignment(2)}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rosterRows := sqlmock.NewRows(getRosterClientCols()).
		AddRow(1, 1, 1, TEST_AIFS_CLIENT_DB_PATROL_ORDER).
		AddRow(2, 1, 1, TEST_AIFS_CLIENT_DB_PATROL_ORDER)

	mock.ExpectQuery("SELECT \\* FROM aifs_client_schedule WHERE schedule = '1'").
		WillReturnRows(rosterRows)
	mock.ExpectQuery("SELECT \\* FROM client WHERE client_id =").WillReturnRows(getSingleClientDbRow(1))
	mock.ExpectQuery("SELECT \\* FROM client WHERE client_id =").WillReturnRows(getSingleClientDbRow(2))

	rosters, err := db_pck.GetRosterAIFSClient(db, query, 1)

	if err != nil {
		t.Errorf("error was not expected while selecting rosters: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 2, len(rosters), "The length of rosters returned should be 2")
	t.Log("returned ", rosters[1])
	t.Log("fake", fakeRosters[1])
	assert.Equal(t, true, proto.Equal(rosters[0], fakeRosters[0]), "The first roster returned is not equal to the expected.")
	assert.Equal(t, true, proto.Equal(rosters[1], fakeRosters[1]), "The second roster returned is not equal to the expected.")
}

func TestGetDefaultRosterDetails(t *testing.T) {
	query := &pb.RosterQuery{}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	defaultRosteringRows := sqlmock.NewRows([]string{"defaultRosteringId", "day_of_week", "schedule1", "schedule2", "schedule3"}).
		AddRow(1, 1, 1, 2, 3)

	mock.ExpectQuery("SELECT \\* FROM default_rostering WHERE day_of_week = '2'").WillReturnRows(defaultRosteringRows)

	schedules, err := db_pck.GetDefaultRosterDetails(db, query, 2)

	if err != nil {
		t.Errorf("error was not expected while selecting rosters: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 3, len(schedules), "The length of rosters returned should be 2")
	assert.Equal(t, []int{1, 2, 3}, schedules, "The first roster returned is not equal to the expected.")
}

func TestUpdateRosterChangeUsers(t *testing.T) {
	roster := &pb.Roster{RosteringId: 1, GuardAssigned: make([]*pb.RosterAssignement, 0)}
	roster.GuardAssigned = append(roster.GuardAssigned, &pb.RosterAssignement{
		GuardAssigned: &pb.EmployeeEvaluation{Employee: createFakeUser(1)},
	})
	roster.GuardAssigned = append(roster.GuardAssigned, &pb.RosterAssignement{
		GuardAssigned: &pb.EmployeeEvaluation{Employee: createFakeUser(2)},
	})

	t.Log("created", roster)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rosterRows := sqlmock.NewRows(getRosterAssignmentCols()).
		AddRow(1, 1, 3, TEST_ROSTER_FAKE_START_TIME, TEST_ROSTER_FAKE_END_TIME,
			TEST_ROSTER_ASGN_DB_CONFIRMATION, TEST_ROSTER_ASGN_DB_ATTENDED, nil, TEST_ROSTER_ASGN_DB_IS_ASSIGNED,
			TEST_ROSTER_ASGN_DB_REJECTED,
		)

	mock.ExpectQuery("SELECT \\* FROM schedule_detail WHERE is_assigned = 1 AND schedule = '1'").
		WillReturnRows(rosterRows)

	mock.ExpectQuery("SELECT \\* FROM user WHERE user_id =").WillReturnRows(getSingleUserDbRow(3))
	mock.ExpectExec("INSERT INTO schedule_detail.*VALUES \\(1,1").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO schedule_detail.*VALUES \\(1,2").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("UPDATE schedule_detail SET.*is_assigned=false.*WHERE guard_assigned = '3' AND schedule = '1'").WillReturnResult(sqlmock.NewResult(1, 1))

	_, newlyInsertedRosters, err := db_pck.UpdateRoster(db, roster, &sync.Mutex{})

	if err != nil {
		t.Errorf("error was not expected while updating broadcast: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 2, len(newlyInsertedRosters), "The number of rows updated 1")
}

func TestUpdateRosterAssignmentsConfirmed(t *testing.T) {
	rosterAssignement := &pb.RosterAssignement{
		RosterAssignmentId: 1,
		Confirmed:          true,
	}
	query := &pb.RosterQuery{Filters: make([]*pb.RosterFilter, 0), Limit: 7}
	query.Filters = append(query.Filters, &pb.RosterFilter{
		Field: pb.RosterFilter_ROSTER_ASSIGNMENT_ID,
		Comparisons: &pb.Filter{
			Value: "1", Comparison: pb.Filter_EQUAL,
		},
	})

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("UPDATE schedule_detail SET confirmation=true.*WHERE schedule_detail_id = '1'").WillReturnResult(sqlmock.NewResult(1, 1))

	numRows, err := db_pck.UpdateRosterAssignments(db, rosterAssignement, query)

	if err != nil {
		t.Errorf("error was not expected while updating broadcast: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 1, int(numRows), "The number of rows updated 1")
}

func TestUpdateRosterAssignmentsRejected(t *testing.T) {
	rosterAssignement := &pb.RosterAssignement{
		RosterAssignmentId: 1,
		Rejected:           true,
	}
	query := &pb.RosterQuery{Filters: make([]*pb.RosterFilter, 0), Limit: 7}
	query.Filters = append(query.Filters, &pb.RosterFilter{
		Field: pb.RosterFilter_ROSTER_ASSIGNMENT_ID,
		Comparisons: &pb.Filter{
			Value: "1", Comparison: pb.Filter_EQUAL,
		},
	})

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("UPDATE schedule_detail SET.*rejected=true.*WHERE schedule_detail_id = '1'").WillReturnResult(sqlmock.NewResult(1, 1))

	numRows, err := db_pck.UpdateRosterAssignments(db, rosterAssignement, query)

	if err != nil {
		t.Errorf("error was not expected while updating broadcast: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 1, int(numRows), "The number of rows updated 1")
}

func TestDeleteRoster(t *testing.T) {
	roster := &pb.Roster{
		RosteringId: 1,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("DELETE FROM schedule WHERE schedule_id = '1'").WillReturnResult(sqlmock.NewResult(1, 1))

	numRows, err := db_pck.DeleteRoster(db, roster)

	if err != nil {
		t.Errorf("error was not expected while deleting roster: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 1, int(numRows), "The number of rows deleted is supposed to be 1")
}

func TestDeleteRosterAssignment(t *testing.T) {
	assignment := &pb.RosterAssignement{
		RosterAssignmentId: 1,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("DELETE FROM schedule_detail WHERE schedule_detail_id = '1'").WillReturnResult(sqlmock.NewResult(1, 1))

	query := &pb.BroadcastQuery{}
	db_pck.AddBroadcastFilter(query, pb.BroadcastFilter_BROADCAST_RECIPIENT_TABLE_ID, pb.Filter_EQUAL, "1")

	numRows, err := db_pck.DeleteRosterAssignment(db, assignment)

	if err != nil {
		t.Errorf("error was not expected while deleting assignment: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 1, int(numRows), "The number of rows deleted is supposed to be 1")
}

func TestDeleteRosterClient(t *testing.T) {
	clientRoster := &pb.AIFSClientRoster{
		AifsClientRosterId: 1,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("DELETE FROM aifs_client_schedule WHERE aifs_client_schedule_id = '1'").WillReturnResult(sqlmock.NewResult(1, 1))

	query := &pb.BroadcastQuery{}
	db_pck.AddBroadcastFilter(query, pb.BroadcastFilter_BROADCAST_RECIPIENT_TABLE_ID, pb.Filter_EQUAL, "1")

	numRows, err := db_pck.DeleteRosterAIFSClient(db, clientRoster)

	if err != nil {
		t.Errorf("error was not expected while deleting AIFSClientRoster: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 1, int(numRows), "The number of rows deleted is supposed to be 1")
}

// Add new broadcast filters from empty query
func TestAddRosterFilterEmpty(t *testing.T) {
	expectedQuery := &pb.RosterQuery{}
	expectedQuery.Filters = append(expectedQuery.Filters, &pb.RosterFilter{
		Field: pb.RosterFilter_ROSTER_ID,
		Comparisons: &pb.Filter{
			Value: "1", Comparison: pb.Filter_EQUAL,
		},
	})

	inputQuery := &pb.RosterQuery{}
	db_pck.AddRosterFilter(inputQuery, pb.RosterFilter_ROSTER_ID, pb.Filter_EQUAL, "1")
	// Check if the filter is added correctly
	assert.Equal(t, 1, len(inputQuery.Filters), "The number of filters expected is not correct")
	assert.Equal(t, true, proto.Equal(inputQuery, expectedQuery), "The query is not equal to the expected.")

}

// Add new broadcast filters from query that already has something
func TestAddRosterFilterAlreadyContains(t *testing.T) {
	expectedQuery := &pb.RosterQuery{Filters: make([]*pb.RosterFilter, 0), Limit: 7}
	expectedQuery.Filters = append(expectedQuery.Filters, &pb.RosterFilter{
		Field: pb.RosterFilter_ROSTER_ID,
		Comparisons: &pb.Filter{
			Value: "1", Comparison: pb.Filter_EQUAL,
		},
	})
	expectedQuery.Filters = append(expectedQuery.Filters, &pb.RosterFilter{
		Field: pb.RosterFilter_AIFS_ID,
		Comparisons: &pb.Filter{
			Value: "1", Comparison: pb.Filter_EQUAL,
		},
	})

	inputQuery := &pb.RosterQuery{Filters: make([]*pb.RosterFilter, 0), Limit: 7}
	inputQuery.Filters = append(inputQuery.Filters, &pb.RosterFilter{
		Field: pb.RosterFilter_ROSTER_ID,
		Comparisons: &pb.Filter{
			Value: "1", Comparison: pb.Filter_EQUAL,
		},
	})

	db_pck.AddRosterFilter(inputQuery, pb.RosterFilter_AIFS_ID, pb.Filter_EQUAL, "1")
	// Check if the filter is added correctly
	assert.Equal(t, 2, len(inputQuery.Filters), "The number of filters expected is not correct")
	assert.Equal(t, true, proto.Equal(inputQuery, expectedQuery), "The query is not equal to the expected.")

}

func CreateFakeRoster(id int) *pb.Roster {
	rosterAssignments := make([]*pb.RosterAssignement, 0)
	aifsAssignments := make([]*pb.AIFSClientRoster, 0)

	for i := 1; i < 2; i++ {
		rosterAssignments = append(rosterAssignments, createFakeRosterAssignment(i))
		aifsAssignments = append(aifsAssignments, createFakeAifsAssignment(i))
	}

	startFakeTime := time.Date(2022, 6, 21, 18, 0, 0, 0, time.UTC)
	endFakeTime := time.Date(2022, 6, 22, 6, 0, 0, 0, time.UTC)

	return &pb.Roster{
		RosteringId:   int64(id),
		AifsId:        int64(id),
		StartTime:     startFakeTime.Format(common.DATETIME_FORMAT),
		EndTime:       endFakeTime.Format(common.DATETIME_FORMAT),
		Clients:       aifsAssignments,
		GuardAssigned: rosterAssignments,
		Status:        pb.Roster_PENDING,
	}
}

func createFakeRosterAssignment(id int) *pb.RosterAssignement {
	startFakeTime := time.Date(2022, 6, 21, 18, 0, 0, 0, time.UTC)
	endFakeTime := time.Date(2022, 6, 22, 6, 0, 0, 0, time.UTC)

	return &pb.RosterAssignement{
		RosterAssignmentId: int64(id),
		GuardAssigned:      createFakeEmployeeEval(id),
		CustomStartTime:    &timestamppb.Timestamp{Seconds: startFakeTime.Unix()},
		CustomEndTime:      &timestamppb.Timestamp{Seconds: endFakeTime.Unix()},
		Confirmed:          false,
		Attended:           false,
		AttendanceTime:     nil,
	}
}

func createFakeAifsAssignment(id int) *pb.AIFSClientRoster {
	return &pb.AIFSClientRoster{
		AifsClientRosterId: int64(id),
		Client:             createFakeClient(id),
		PatrolOrder:        TEST_AIFS_CLIENT_DB_PATROL_ORDER,
	}
}

func createFakeEmployeeEval(id int) *pb.EmployeeEvaluation {
	return &pb.EmployeeEvaluation{
		Employee:      createFakeUser(id),
		EmployeeScore: float32(100 - float32(id)),
		IsAvailable:   false,
	}
}

func getFullRosterCols() []string {
	cols := []string{
		"schedule_id",
		"aifs_id",
		"start_time",
		"end_time",
	}
	cols = append(cols, getRosterAssignmentCols()...)
	cols = append(cols, getRosterClientCols()...)

	return cols
}

func getRosterAssignmentCols() []string {
	return []string{
		"schedule_detail_id",
		"schedule",
		"guard_assigned",
		"custom_start_time",
		"custom_end_time",
		"confirmation",
		"attended",
		"attendance_time",
		"is_assigned",
		"rejected",
	}
}

func getRosterClientCols() []string {
	return []string{
		"aifs_client_schedule_id",
		"schedule",
		"related_client",
		"patrol_order",
	}
}
