// TODO: Add validation
package server

import (
	"context"
	"database/sql"
	"fmt"
	"io"

	db_pck "capstone.operations_ecosystem/backend/database"
	pb "capstone.operations_ecosystem/backend/proto"
)

func (s *Server) AddRoster(stream pb.RosterServices_AddRosterServer) error {
	fmt.Println("AddRoster")
	res := pb.Response{Type: pb.Response_ACK, PrimaryKey: 1}

	for {
		roster, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&res)
		}

		if err != nil {
			return err
		}

		// Fill up the blank values of the pb message
		err = s.insertDefaultRosterValues(roster)
		if err != nil {
			return err
		}

		pk, err := db_pck.InsertRoster(
			s.db,
			roster,
			s.dbLock,
		)

		if err != nil {
			res.Type = pb.Response_ERROR
			res.ErrorMessage = err.Error()
		}

		res.PrimaryKey = int64(pk)
	}
}

func (s *Server) UpdateRoster(cxt context.Context, roster *pb.Roster) (*pb.Response, error) {
	res := pb.Response{Type: pb.Response_ACK}
	numAffected, err := db_pck.UpdateRoster(
		s.db,
		roster,
		s.dbLock,
	)

	if err != nil {
		res.Type = pb.Response_ERROR
		res.ErrorMessage = err.Error()
	} else {
		fmt.Println(numAffected, "rosters were updated.")
	}

	return &res, nil
}

func (s *Server) DeleteRoster(cxt context.Context, roster *pb.Roster) (*pb.Response, error) {
	res := pb.Response{Type: pb.Response_ACK}
	numDel, err := db_pck.DeleteRoster(
		s.db,
		roster,
	)

	if err != nil {
		res.Type = pb.Response_ERROR
		res.ErrorMessage = err.Error()
	} else {
		fmt.Println(numDel, "rosters were deleted.")
	}

	return &res, nil
}

func (s *Server) FindRosters(query *pb.RosterQuery, stream pb.RosterServices_FindRostersServer) error {
	res := pb.Response{Type: pb.Response_ACK}

	foundRosters, err := db_pck.GetRosters(
		s.db,
		query,
	)

	if err != nil {
		rosterRes := pb.RosterResponse{Response: &res}
		res.Type = pb.Response_ERROR
		res.ErrorMessage = err.Error()
		stream.Send(&rosterRes)

	} else {
		fmt.Println("FindBroadcasts: Found broadcasts to return")

		rosterRes := pb.RosterResponse{Response: &res}

		for _, roster := range foundRosters {
			rosterRes.Roster = roster
			if err := stream.Send(&rosterRes); err != nil {
				return err
			}
		}
	}

	return nil
}

// TODO
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

// Adds the default values for the roster
func (s *Server) insertDefaultRosterValues(roster *pb.Roster) error {
	if roster.Clients == nil {
		err := fillDefaultClients(roster, s.db)
		if err != nil {
			return err
		}
	}

	// fill up custom time
	for _, assignment := range roster.GuardAssigned {
		if assignment.CustomStartTime == nil {
			assignment.CustomStartTime = roster.StartTime
		}
		if assignment.CustomEndTime == nil {
			assignment.CustomEndTime = roster.EndTime
		}
	}

	return nil
}

func fillDefaultClients(roster *pb.Roster, db *sql.DB) error {
	clientQuery := &pb.ClientQuery{}
	aifsClientRosters := make([]*pb.AIFSClientRoster, 0)
	var clients []*pb.Client
	var err error

	switch roster.AifsId {
	case 1:
		db_pck.AddClientFilter(clientQuery, pb.ClientFilter_CLIENT_ID,
			pb.Filter_EQUAL, "1")
		db_pck.AddClientFilter(clientQuery, pb.ClientFilter_CLIENT_ID,
			pb.Filter_EQUAL, "2")
		clients, err = db_pck.GetClients(db, clientQuery)
		if err != nil {
			return err
		}
	case 2:
		db_pck.AddClientFilter(clientQuery, pb.ClientFilter_CLIENT_ID,
			pb.Filter_EQUAL, "3")
		db_pck.AddClientFilter(clientQuery, pb.ClientFilter_CLIENT_ID,
			pb.Filter_EQUAL, "4")
		clients, err = db_pck.GetClients(db, clientQuery)
		if err != nil {
			return err
		}
	default:
		db_pck.AddClientFilter(clientQuery, pb.ClientFilter_CLIENT_ID,
			pb.Filter_EQUAL, "5")
		db_pck.AddClientFilter(clientQuery, pb.ClientFilter_CLIENT_ID,
			pb.Filter_EQUAL, "6")
		clients, err = db_pck.GetClients(db, clientQuery)
		if err != nil {
			return err
		}
	}

	for _, client := range clients {
		aifsClientRosters = append(aifsClientRosters, &pb.AIFSClientRoster{Client: client})
	}

	roster.Clients = aifsClientRosters

	return nil
}
