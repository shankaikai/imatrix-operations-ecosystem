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
	broadcast.Content = "This is a message :)"
	// broadcast.Urgency = operations_ecosys.Broadcast_HIGH
	// broadcast.Urgency = operations_ecosys.Broadcast_MEDIUM
	broadcast.Urgency = operations_ecosys.Broadcast_LOW
	for _, rep := range broadcast.Recipients {
		for i, brep := range rep.Recipient {
			// TODO: REMOVE this is emily's chat id
			brep.Recipient.TeleUserId = 223102557 // emily: 29333507
			brep.BroadcastRecipientsId = int64(20 + i)
		}
	}
	teleClient := tclient.TelegramClient{}
	teleClient.InsertBroadcast(serverAddr, serverPort, broadcast)
}
