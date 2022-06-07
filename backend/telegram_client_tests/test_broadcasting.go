package telegram_client_tests

import (
	fclient "capstone.operations_ecosystem/backend/fake_client"
	tclient "capstone.operations_ecosystem/backend/telegram_client"
)

func TestTelegramBroadcasts(serverAddr *string, serverPort *int) {
	testInsertBroadcast(serverAddr, serverPort)
}

func testInsertBroadcast(serverAddr *string, serverPort *int) {
	broadcast := fclient.CreateFakeBroadcast(2, true)
	tclient.InsertBroadcast(serverAddr, serverPort, broadcast)
}
