package server

import (
	pb "capstone.operations_ecosystem/backend/proto"
)

const (
	AIFS1_USER_ID = 8
	AIFS2_USER_ID = 9
	AIFS3_USER_ID = 10
)

// If the broadcast recipient is an AIFS,
// change the recipients to be actual users
// Modified the broadcast in place
func getDefaultRecipients(broadcast *pb.Broadcast) {
	for _, rec := range broadcast.Recipients {
		newRecipients := make([]*pb.BroadcastRecipient, 0)

		users := getAIFSDuty(rec.AifsId)
		for _, user := range users {
			newRecipients = append(newRecipients, &pb.BroadcastRecipient{
				Recipient: user,
				AifsId:    rec.AifsId,
			})
		}

		rec.Recipient = newRecipients
	}
}

// Fixed duty
func getAIFSDuty(aifsId int64) []*pb.User {
	users := make([]*pb.User, 0)
	switch aifsId {
	case 1:
		users = append(users, &pb.User{
			UserId: int64(AIFS1_USER_ID),
		})
	case 2:
		users = append(users, &pb.User{
			UserId: int64(AIFS2_USER_ID),
		})
	default:
		users = append(users, &pb.User{
			UserId: int64(AIFS3_USER_ID),
		})
	}

	return users
}
