package fake_server

import (
	"fmt"

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

func (s *Server) FindBroadcasts(cxt context.Context, query *pb.BroadcastQuery) (*pb.BulkBroadcasts, error) {
	fmt.Println("FindBroadcasts")

	res := pb.Response{Type: pb.Response_ACK}
	broadcasts := pb.BulkBroadcasts{Response: &res}

	foundBroadcasts := make([]*pb.Broadcast, 0)
	recipients := make([]*pb.BroadcastRecipient, 0)

	for i := 0; i < 3; i++ {
		user := &pb.User{
			UserId:          int64(i),
			UserType:        pb.User_ISPECIALIST,
			Name:            "test name",
			Email:           "email",
			PhoneNumber:     "1232",
			TelegramHandle:  "sfds",
			UserSecurityImg: "dsfds",
			IsPartTimer:     false,
		}
		recipients = append(recipients, &pb.BroadcastRecipient{
			BroadcastRecipientsId: int64(i),
			Recipient:             user,
			Acknowledged:          i%2 == 0,
			Rejected:              false,
		})
	}

	for i := 0; i < 3; i++ {
		foundBroadcasts = append(foundBroadcasts, &pb.Broadcast{
			BroadcastId:  3,
			Type:         pb.Broadcast_ANNOUNCEMENT,
			Title:        "test name",
			Content:      "email",
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
			Recipients: recipients,
		})
	}

	broadcasts.Broadcasts = foundBroadcasts

	return &broadcasts, nil
}
