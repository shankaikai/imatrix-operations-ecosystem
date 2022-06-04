// TODO: Add validation
package server

import (
	"context"
	"fmt"
	"io"
	"sort"

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

	// Only find rosters assignments that are still assigned
	db_pck.AddRosterFilter(query, pb.RosterFilter_IS_ASSIGNED, pb.Filter_EQUAL, "1")

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

// Send back all the available users that are not yet assigned for the particular day
func (s *Server) GetAvailableUsers(query *pb.AvailabilityQuery, stream pb.RosterServices_GetAvailableUsersServer) error {
	fmt.Println("GetAvailableUsers")

	res := pb.Response{Type: pb.Response_ACK}
	employeeEvalRes := pb.EmployeeEvaluationResponse{Response: &res}

	employeeEvals, err := s.GetAvailableIspecialistsWithScore(query)

	if err != nil {
		return err
	}

	// Sort the available Ispecialists according to score desc
	sort.Slice(employeeEvals, func(i, j int) bool {
		return employeeEvals[i].EmployeeScore < employeeEvals[j].EmployeeScore
	})

	for _, employeeEval := range employeeEvals {
		employeeEvalRes.Employee = employeeEval

		if err := stream.Send(&employeeEvalRes); err != nil {
			return err
		}
	}

	return nil
}
