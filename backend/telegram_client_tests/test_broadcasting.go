package telegram_client_tests

import (
	fclient "capstone.operations_ecosystem/backend/fake_client"
	operations_ecosys "capstone.operations_ecosystem/backend/proto"
	tclient "capstone.operations_ecosystem/backend/telegram_client"
)

func TestTelegramBroadcasts(serverAddr *string, serverPort *int) {
	testInsertBroadcast(serverAddr, serverPort)
}

func testInsertBroadcast(serverAddr *string, serverPort *int) {
	broadcast := fclient.CreateFakeBroadcast(2, true)
	broadcast.Content = "ahhhhhhhhhhhhhhhhhhhhhhhh"
	broadcast.Urgency = operations_ecosys.Broadcast_HIGH
	for _, rep := range broadcast.Recipients {
		for _, brep := range rep.Recipient {
			// TODO: REMOVE this is emily's chat id
			brep.Recipient.TeleChatId = 29333507
		}
	}
	tclient.InsertBroadcast(serverAddr, serverPort, broadcast)
}