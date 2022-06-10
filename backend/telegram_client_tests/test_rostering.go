package telegram_client_tests

import (
	fclient "capstone.operations_ecosystem/backend/fake_client"
	tclient "capstone.operations_ecosystem/backend/telegram_client"

	pb "capstone.operations_ecosystem/backend/proto"
)

func TestTelegramRosters(serverAddr *string, serverPort *int) {
	testInsertRoster(serverAddr, serverPort)
}

func testInsertRoster(serverAddr *string, serverPort *int) {
	rosters := make([]*pb.Roster, 0)
	for i := 0; i < 3; i++ {
		roster := fclient.CreateFakeRoster(2)
		for _, guardAssigned := range roster.GuardAssigned {
			user := guardAssigned.GuardAssigned.Employee
			// TODO: REMOVE this is emily's chat id
			user.TeleChatId = 29333507
		}
		rosters = append(rosters, roster)
	}

	tclient.InsertRoster(serverAddr, serverPort, rosters)
}
