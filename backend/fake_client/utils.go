package client

import (
	"strconv"
	"time"

	"capstone.operations_ecosystem/backend/common"
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

func createFakeClient(id int) *pb.Client {
	return &pb.Client{
		ClientId:     int64(id),
		Name:         "test name" + strconv.Itoa(id),
		Abbreviation: "ABC" + strconv.Itoa(id),
		Email:        "email" + strconv.Itoa(id),
		Address:      "address" + strconv.Itoa(id),
		PostalCode:   int64(123213 + id),
		PhoneNumber:  "1232",
	}
}

func CreateFakeBroadcast(id int, hasRecipient bool) *pb.Broadcast {
	recipients := make([]*pb.BroadcastRecipient, 0)
	aifsRecipients := make([]*pb.AIFSBroadcastRecipient, 0)
	for i := 1; i < 6; i++ {
		recipients = append(recipients, createFakeBroadcastRec(i, hasRecipient))
	}

	for i := 1; i < 3; i++ {
		aifsRecipients = append(aifsRecipients, &pb.AIFSBroadcastRecipient{
			AifsId:    int64(i),
			Recipient: recipients,
		})
	}

	return &pb.Broadcast{
		BroadcastId:  int64(id),
		Type:         pb.Broadcast_ANNOUNCEMENT,
		Content:      "content" + strconv.Itoa(id),
		CreationDate: timestamppb.Now(),
		Deadline:     &timestamppb.Timestamp{Seconds: int64(time.Now().Add(30).Unix())},
		Creator:      createFakeUser(1),
		Recipients:   aifsRecipients,
		Urgency:      pb.Broadcast_LOW,
	}
}

func createFakeBroadcastRec(id int, hasUser bool) *pb.BroadcastRecipient {
	var user *pb.User
	if hasUser {
		user = createFakeUser(id)
	}

	return &pb.BroadcastRecipient{
		BroadcastRecipientsId: int64(id),
		Recipient:             user,
		Acknowledged:          false,
		Rejected:              false,
		AifsId:                int64(id),
	}
}

func CreateFakeRoster(id int) *pb.Roster {
	rosterAssignments := make([]*pb.RosterAssignement, 0)
	aifsAssignments := make([]*pb.AIFSClientRoster, 0)

	for i := 2; i < 3; i++ {
		rosterAssignments = append(rosterAssignments, createFakeRosterAssignment(i))
		aifsAssignments = append(aifsAssignments, createFakeAifsAssignment(i))
	}

	startTimeTime := time.Date(2022, 1, 22, 18, 0, 0, 0, time.UTC)
	endTimeTime := time.Date(2022, 1, 22, 6, 0, 0, 0, time.UTC)
	// startTime := &timestamppb.Timestamp{Seconds: startTimeTime.Unix()}
	// endTime := &timestamppb.Timestamp{Seconds: endTimeTime.Unix()}

	return &pb.Roster{
		RosteringId:   int64(id),
		AifsId:        int64(id),
		StartTime:     startTimeTime.Format(common.DATETIME_FORMAT),
		EndTime:       endTimeTime.Format(common.DATETIME_FORMAT),
		Clients:       aifsAssignments,
		GuardAssigned: rosterAssignments,
	}
}

func createFakeRosterAssignment(id int) *pb.RosterAssignement {
	return &pb.RosterAssignement{
		RosterAssignmentId: int64(id),
		GuardAssigned:      createFakeEmployeeEval(id),
		CustomStartTime:    timestamppb.Now(),
		CustomEndTime:      &timestamppb.Timestamp{Seconds: timestamppb.Now().AsTime().Add(time.Hour * 12).Unix()},
		Confirmed:          false,
		Attended:           false,
		AttendanceTime:     nil,
	}
}

func createFakeAifsAssignment(id int) *pb.AIFSClientRoster {
	return &pb.AIFSClientRoster{
		AifsClientRosterId: int64(id),
		Client:             createFakeClient(id),
		PatrolOrder:        int64(id),
	}
}

func createFakeEmployeeEval(id int) *pb.EmployeeEvaluation {
	return &pb.EmployeeEvaluation{
		Employee:      createFakeUser(id),
		EmployeeScore: float32(5 - 0.01*float32(id)),
		IsAvailable:   true,
	}
}
