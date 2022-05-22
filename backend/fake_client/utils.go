package client

import (
	"strconv"
	"time"

	pb "capstone.operations_ecosystem/backend/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func createFakeUser(id int) *pb.User {
	return &pb.User{
		UserId:          int64(id),
		UserType:        pb.User_ISPECIALIST,
		Name:            "test name" + strconv.Itoa(id),
		Email:           "email" + strconv.Itoa(id),
		PhoneNumber:     "1232",
		TelegramHandle:  "sfds",
		UserSecurityImg: "dsfds",
		IsPartTimer:     false,
	}
}

func createFakeBroadcast(id int) *pb.Broadcast {
	recipients := make([]*pb.BroadcastRecipient, 0)
	for i := 1; i < 4; i++ {
		recipients = append(recipients, createFakeBroadcastRec(i))
	}

	return &pb.Broadcast{
		BroadcastId:  int64(id),
		Type:         pb.Broadcast_ANNOUNCEMENT,
		Title:        "title name" + strconv.Itoa(id),
		Content:      "content" + strconv.Itoa(id),
		CreationDate: timestamppb.Now(),
		Deadline:     &timestamppb.Timestamp{Seconds: int64(time.Now().Add(30).Unix())},
		Creator:      recipients[0].Recipient,
		Recipients:   recipients,
	}
}

func createFakeBroadcastRec(id int) *pb.BroadcastRecipient {
	user := createFakeUser(id)
	return &pb.BroadcastRecipient{
		BroadcastRecipientsId: int64(id),
		Recipient:             user,
		Acknowledged:          false,
		Rejected:              false,
	}
}
