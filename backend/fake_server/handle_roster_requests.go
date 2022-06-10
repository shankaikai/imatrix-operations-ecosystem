package fake_server

import (
	"fmt"
	"strconv"
	"time"

	"capstone.operations_ecosystem/backend/common"
	pb "capstone.operations_ecosystem/backend/proto"

	"context"
)

func (s *Server) AddRoster(cxt context.Context, rosters *pb.BulkRosters) (*pb.Response, error) {
	fmt.Println("AddRoster")
	res := pb.Response{Type: pb.Response_ACK, PrimaryKey: 1}
	return &res, nil
}

func (s *Server) UpdateRoster(cxt context.Context, roster *pb.BulkRosters) (*pb.Response, error) {
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

		employeeEval := &pb.EmployeeEvaluation{
			Employee:      user,
			EmployeeScore: float32(5 - 0.01*float32(i)),
			IsAvailable:   true,
		}

		assignments = append(assignments, &pb.RosterAssignement{
			RosterAssignmentId: int64(i),
			GuardAssigned:      employeeEval,
			CustomStartTime:    nil,
			CustomEndTime:      nil,
		})

		aifsClientRoster = append(aifsClientRoster, &pb.AIFSClientRoster{
			AifsClientRosterId: int64(i),
			Client: &pb.Client{
				ClientId:     int64(i),
				Name:         "test name" + strconv.Itoa(i),
				Abbreviation: "ABC" + strconv.Itoa(i),
				Email:        "email" + strconv.Itoa(i),
				Address:      "address" + strconv.Itoa(i),
				PostalCode:   434343,
				PhoneNumber:  "1232",
			},
		})
	}

	for i := 0; i < 3; i++ {
		roster := &pb.Roster{
			RosteringId:   int64(i),
			AifsId:        int64(i),
			StartTime:     time.Now().Format(common.DATETIME_FORMAT),
			EndTime:       time.Now().Add(1000000 * time.Second).Format(common.DATETIME_FORMAT),
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

func (s *Server) GetAvailableUsers(query *pb.AvailabilityQuery, stream pb.RosterServices_GetAvailableUsersServer) error {
	fmt.Println("GetAvailableUsers")

	res := pb.Response{Type: pb.Response_ACK}
	employeeEvalRes := pb.EmployeeEvaluationResponse{Response: &res}

	employees := make([]*pb.User, 0)

	for i := 0; i < 3; i++ {
		employees = append(employees, &pb.User{
			UserId:          int64(i),
			UserType:        pb.User_ISPECIALIST,
			Name:            "test name",
			Email:           "email",
			PhoneNumber:     "1232",
			TelegramHandle:  "sfds",
			UserSecurityImg: "dsfds",
			IsPartTimer:     false,
		})
	}

	for i := 0; i < 3; i++ {
		employeeEval := &pb.EmployeeEvaluation{
			Employee:      employees[i],
			IsAvailable:   true,
			EmployeeScore: float32((i + 1) * 4),
		}

		employeeEvalRes.Employee = employeeEval

		if err := stream.Send(&employeeEvalRes); err != nil {
			return err
		}
	}

	return nil
}
