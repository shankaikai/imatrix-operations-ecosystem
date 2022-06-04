package server

import (
	pb "capstone.operations_ecosystem/backend/proto"
)

// If the broadcast recipient is an AIFS,
// change the recipients to be actual users
// Modified the broadcast in place
func getDefaultRecipients(broadcast *pb.Broadcast) {
	for _, rec := range broadcast.Recipients {
		newRecipients := make([]*pb.BroadcastRecipient, 0)

		users := getFakeAIFSDuty(rec.AifsId)
		for _, user := range users {
			newRecipients = append(newRecipients, &pb.BroadcastRecipient{
				Recipient: user,
				AifsId:    rec.AifsId,
			})
		}

		rec.Recipient = newRecipients
	}
}

// TODO get actual roster for AIFS Groups
func getFakeAIFSDuty(aifsId int64) []*pb.User {
	users := make([]*pb.User, 0)

	for i := 1; i < 3; i++ {
		users = append(users, &pb.User{
			UserId: int64(i),
		})
	}

	return users
}
