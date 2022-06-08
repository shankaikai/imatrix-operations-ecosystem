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
	for _, rep := range broadcast.Recipients {
		for _, brep := range rep.Recipient {
			// TODO: REMOVE this is hannah's chat id
			brep.Recipient.TeleChatId = 223102557
		}
	}
	tclient.InsertBroadcast(serverAddr, serverPort, broadcast)
}
