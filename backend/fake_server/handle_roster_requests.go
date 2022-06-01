package fake_server

import (
	"fmt"
	"strconv"

	pb "capstone.operations_ecosystem/backend/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"context"
)

func (s *Server) AddRoster(cxt context.Context, roster *pb.Roster) (*pb.Response, error) {
	fmt.Println("AddRoster")
	res := pb.Response{Type: pb.Response_ACK, PrimaryKey: 1}
	return &res, nil
}

func (s *Server) UpdateRoster(cxt context.Context, roster *pb.Roster) (*pb.Response, error) {
	fmt.Println("UpdateRoster")
	res := pb.Response{Type: pb.Response_ACK, PrimaryKey: 1}
	return &res, nil
}

func (s *Server) DeleteRoster(cxt context.Context, roster *pb.Roster) (*pb.Response, error) {
	fmt.Println("DeleteRoster")
	res := pb.Response{Type: pb.Response_ACK, PrimaryKey: 1}
	return &res, nil
}

func (s *Server) FindRosters(query *pb.RosterQuery, stream pb.RosterServices_FindRostersServer) error {
	fmt.Println("FindRosters")

	res := pb.Response{Type: pb.Response_ACK}
	rosterRes := pb.RosterResponse{Response: &res}

	assignments := make([]*pb.RosterAssignement, 0)
	aifsClientRoster := make([]*pb.AIFSClientRoster, 0)

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
		assignments = append(assignments, &pb.RosterAssignement{
			RosterAssignmentId: int64(i),
			GuardAssigned:      user,
			CustomStartTime:    nil,
			CustomEndTime:      nil,
		})

		aifsClientRoster = append(aifsClientRoster, &pb.AIFSClientRoster{
			AifsClientRosterId: int64(i),
			Client: &pb.Client{
				ClientId:             int64(i),
				Name:                 "test name" + strconv.Itoa(i),
				Email:                "email" + strconv.Itoa(i),
				Address:              "address" + strconv.Itoa(i),
				PhoneNumber:          "1232",
				NumberOfGuardsNeeded: 3,
			},
		})
	}

	for i := 0; i < 3; i++ {
		roster := &pb.Roster{
			RosteringId:   int64(i),
			AifsId:        int64(i),
			StartTime:     timestamppb.Now(),
			EndTime:       &timestamppb.Timestamp{Seconds: timestamppb.Now().Seconds + 1000000},
			Clients:       aifsClientRoster,
			GuardAssigned: assignments,
		}

		rosterRes.Roster = roster

		if err := stream.Send(&rosterRes); err != nil {
			return err
		}
	}

	return nil
}
