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
	// AIFS LED light triggers for broadcasting
	AIFS_LED_DEFAULT_URI          = "win&PL=2"
	AIFS_LED_BROADCAST_URGENT_URI = "win&PL=1"
	TEST_BROADCAST_LED_AIFS_ID    = 1
)

// If the broadcast recipient is an AIFS,
// change the recipients to be actual users
// If the
// Modified the broadcast in place
func (s *Server) getDefaultBroadcastFields(broadcast *pb.Broadcast) {
	defaultAck := broadcast.Urgency != pb.Broadcast_HIGH
	var ackTime time.Time

	if defaultAck {
		ackTime = time.Now()
	}

	for _, rec := range broadcast.Recipients {
		newRecipients := make([]*pb.BroadcastRecipient, 0)

		users := s.getAIFSDuty(rec.AifsId)
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
func (s *Server) getAIFSDuty(aifsId int64) []*pb.User {
	users := make([]*pb.User, 0)
	switch aifsId {
	case 1:
		users = append(users, &pb.User{
			UserId: int64(s.Config.Aifs1Id),
		})
	case 2:
		users = append(users, &pb.User{
			UserId: int64(s.Config.Aifs2Id),
		})
	default:
		users = append(users, &pb.User{
			UserId: int64(s.Config.Aifs3Id),
		})
	}

	return users
}

func (s *Server) sendNewBroadcastsOut(broadcastId int64) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered", r)
		}
	}()

	query := &pb.BroadcastQuery{Limit: 1}
	db_pck.AddBroadcastFilter(query, pb.BroadcastFilter_BROADCAST_ID, pb.Filter_EQUAL, strconv.Itoa(int(broadcastId)))
	broadcasts, err := db_pck.GetBroadcasts(s.db, query)

	if err != nil {
		fmt.Println("sendNewBroadcastToTele ERROR:", err)
	}

	if len(broadcasts) == 0 {
		fmt.Println("sendNewBroadcastToTele: No broadcast found for id", broadcastId)
	}
	teleClient := tclient.TelegramClient{}
	teleClient.InsertBroadcast(s.teleServerAddr, s.teleServerPort, broadcasts[0])
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
	resp, err := http.Get(fmt.Sprintf("%s/%s", *s.testLEDAddr, AIFS_LED_DEFAULT_URI))
	if err != nil {
		fmt.Println("notifyAIFSofBroadcastAck ERROR:", err)
	}

	fmt.Println("notifyAIFSofBroadcastAck RESPONSE:", resp)
}
