package server

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	db_pck "capstone.operations_ecosystem/backend/database"
	tclient "capstone.operations_ecosystem/backend/telegram_client"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "capstone.operations_ecosystem/backend/proto"
)

const (
	AIFS1_USER_ID = 8
	AIFS2_USER_ID = 9
	AIFS3_USER_ID = 10

	// AIFS LED light triggers for broadcasting
	AIFS_LED_DEFAULT_URI          = "win&PL=2"
	AIFS_LED_BROADCAST_URGENT_URI = "win&PL=1"
	TEST_BROADCAST_LED_AIFS_ID    = 1
)

// If the broadcast recipient is an AIFS,
// change the recipients to be actual users
// If the
// Modified the broadcast in place
func getDefaultBroadcastFields(broadcast *pb.Broadcast) {
	defaultAck := broadcast.Urgency != pb.Broadcast_HIGH
	var ackTime time.Time

	if defaultAck {
		ackTime = time.Now()
	}

	for _, rec := range broadcast.Recipients {
		newRecipients := make([]*pb.BroadcastRecipient, 0)

		users := getAIFSDuty(rec.AifsId)
		for _, user := range users {
			newRecipients = append(newRecipients, &pb.BroadcastRecipient{
				Recipient:    user,
				AifsId:       rec.AifsId,
				Acknowledged: defaultAck,
				LastReplied:  &timestamppb.Timestamp{Seconds: ackTime.Unix()},
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

func (s *Server) sendNewBroadcastsOut(broadcastId int64) {
	query := &pb.BroadcastQuery{Limit: 1}
	db_pck.AddBroadcastFilter(query, pb.BroadcastFilter_BROADCAST_ID, pb.Filter_EQUAL, strconv.Itoa(int(broadcastId)))
	broadcasts, err := db_pck.GetBroadcasts(s.db, query)

	if err != nil {
		fmt.Println("sendNewBroadcastToTele ERROR:", err)
	}

	if len(broadcasts) == 0 {
		fmt.Println("sendNewBroadcastToTele: No broadcast found for id", broadcastId)
	}

	tclient.InsertBroadcast(s.teleServerAddr, s.teleServerPort, broadcasts[0])
	s.notifyAIFSofNewBroadcast(broadcasts[0])
}

func (s *Server) notifyAIFSofNewBroadcast(broadcast *pb.Broadcast) {
	for _, aifsRec := range broadcast.Recipients {
		for _, rec := range aifsRec.Recipient {
			switch rec.AifsId {
			case TEST_BROADCAST_LED_AIFS_ID:
				resp, err := http.Get(fmt.Sprintf("%s/%s", *s.testLEDAddr, AIFS_LED_BROADCAST_URGENT_URI))
				if err != nil {
					fmt.Println("notifyAIFSofNewBroadcast ERROR:", err)
				}

				fmt.Println("notifyAIFSofNewBroadcast RESPONSE:", resp)
			}
		}
	}
}

func (s *Server) notifyAIFSofBroadcastAck(broadcastRecipient *pb.BroadcastRecipient) {
	switch broadcastRecipient.AifsId {
	case TEST_BROADCAST_LED_AIFS_ID:
		resp, err := http.Get(fmt.Sprintf("%s/%s", *s.testLEDAddr, AIFS_LED_BROADCAST_URGENT_URI))
		if err != nil {
			fmt.Println("notifyAIFSofBroadcastAck ERROR:", err)
		}

		fmt.Println("notifyAIFSofBroadcastAck RESPONSE:", resp)
	}
}
