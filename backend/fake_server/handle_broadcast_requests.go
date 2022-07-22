package fake_server

import (
	"fmt"
	"strconv"

	pb "capstone.operations_ecosystem/backend/proto"

	"context"
)

func (s *Server) AddBroadcast(cxt context.Context, broadcast *pb.Broadcast) (*pb.Response, error) {
	fmt.Println("AddBroadcast")
	res := pb.Response{Type: pb.Response_ACK, PrimaryKey: 1}
	return &res, nil
}

func (s *Server) UpdateBroadcast(cxt context.Context, broadcast *pb.Broadcast) (*pb.Response, error) {
	fmt.Println("UpdateBroadcast")
	res := pb.Response{Type: pb.Response_ACK, PrimaryKey: 1}
	return &res, nil
}

func (s *Server) DeleteBroadcast(cxt context.Context, broadcast *pb.Broadcast) (*pb.Response, error) {
	fmt.Println("DeleteBroadcast")
	res := pb.Response{Type: pb.Response_ACK, PrimaryKey: 1}
	return &res, nil
}

func (s *Server) FindBroadcasts(query *pb.BroadcastQuery, stream pb.BroadcastServices_FindBroadcastsServer) error {
	fmt.Println("FindBroadcasts")

	res := pb.Response{Type: pb.Response_ACK}
	broadcastRes := pb.BroadcastResponse{Response: &res}

	aifsRecipients := make([]*pb.AIFSBroadcastRecipient, 0)

	for i := 0; i < 3; i++ {
		aifsRecipients = append(aifsRecipients,
			&pb.AIFSBroadcastRecipient{AifsId: int64(i)},
		)

	}

	for j, aifsRec := range aifsRecipients {
		recipients := make([]*pb.BroadcastRecipient, 0)
		for i := 1; i < 4; i++ {
			user := &pb.User{
				UserId:          int64(i * (j + 1)),
				UserType:        pb.User_ISPECIALIST,
				Name:            "test name" + strconv.Itoa(i*(j+1)),
				Email:           "email",
				PhoneNumber:     "1232",
				TelegramHandle:  "sfds",
				UserSecurityImg: "dsfds",
				IsPartTimer:     false,
			}

			recipients = append(recipients, &pb.BroadcastRecipient{
				BroadcastRecipientsId: int64(i * (j + 1)),
				Recipient:             user,
				Acknowledged:          i%2 == 0,
				Rejected:              false,
			})
		}
		aifsRec.Recipient = recipients
	}

	for i := 0; i < 3; i++ {
		urgency := []pb.Broadcast_UrgencyType{pb.Broadcast_HIGH, pb.Broadcast_MEDIUM, pb.Broadcast_LOW}
		broadcast := &pb.Broadcast{
			BroadcastId:  3,
			Type:         pb.Broadcast_ANNOUNCEMENT,
			Content:      "email" + strconv.Itoa(i),
			CreationDate: nil,
			Deadline:     nil,
			Creator: &pb.User{
				UserId:          1,
				UserType:        pb.User_ISPECIALIST,
				Name:            "test name",
				Email:           "email",
				PhoneNumber:     "1232",
				TelegramHandle:  "sfds",
				UserSecurityImg: "dsfds",
				IsPartTimer:     false,
			},
			Recipients: aifsRecipients,
			Urgency:    urgency[i%3],
		}

		broadcastRes.Broadcast = broadcast

		if err := stream.Send(&broadcastRes); err != nil {
			return err
		}
	}

	return nil
}
