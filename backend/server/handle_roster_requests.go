// TODO: Add validation
package server

import (
	"context"
	"fmt"
	"sort"

	db_pck "capstone.operations_ecosystem/backend/database"
	pb "capstone.operations_ecosystem/backend/proto"
)

func (s *Server) AddRoster(cxt context.Context, rosters *pb.BulkRosters) (*pb.Response, error) {
	fmt.Println("AddRoster")
	res := &pb.Response{Type: pb.Response_ACK, PrimaryKey: 1}

	rosterIds := make([]int64, 0)

	for _, roster := range rosters.Rosters {
		// Fill up the blank values of the pb message
		err := s.insertDefaultRosterValues(roster)
		if err != nil {
			return res, err
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

		rosterIds = append(rosterIds, pk)
		res.PrimaryKey = pk
	}

	// Send rosters on telegram
	go s.sendNewRostersToTele(rosterIds)

	return res, nil
}

func (s *Server) UpdateRoster(cxt context.Context, rosters *pb.BulkRosters) (*pb.Response, error) {
	res := pb.Response{Type: pb.Response_ACK}
	for _, roster := range rosters.Rosters {

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

	// Ensure that the start time is found
	err := validateStartTime(query)

	if err != nil {
		rosterRes := pb.RosterResponse{Response: &res}
		res.Type = pb.Response_ERROR
		res.ErrorMessage = err.Error()
		stream.Send(&rosterRes)
		return nil
	}

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
		return nil
	}

	if len(foundRosters) > 0 {
		fmt.Println("FindRosters: Found rosters to return")
	} else {
		// Get default rosters for that particular day
		foundRosters, err = db_pck.GetDefaultRosters(
			s.db,
			query,
		)

		if err != nil {
			rosterRes := pb.RosterResponse{Response: &res}
			res.Type = pb.Response_ERROR
			res.ErrorMessage = err.Error()
			stream.Send(&rosterRes)
			return nil
		}
	}

	for _, roster := range foundRosters {
		rosterRes := pb.RosterResponse{Response: &res}
		rosterRes.Roster = roster
		if err := stream.Send(&rosterRes); err != nil {
			return err
		}
	}

	return nil
}

// Send back all the available users that are not yet assigned for the particular day
func (s *Server) GetAvailableUsers(query *pb.AvailabilityQuery, stream pb.RosterServices_GetAvailableUsersServer) error {
	fmt.Println("GetAvailableUsers")

	res := pb.Response{Type: pb.Response_ACK}
	employeeEvalRes := pb.EmployeeEvaluationResponse{Response: &res}

	err := getDefaultEndTime(query)
	if err != nil {
		return err
	}

	availEmployeeEvals, unavailEmployeeEvals, err := s.GetAvailableIspecialistsWithScore(query)

	if err != nil {
		return err
	}

	// Sort the available Ispecialists according to score desc
	sort.Slice(availEmployeeEvals, func(i, j int) bool {
		return availEmployeeEvals[i].EmployeeScore > availEmployeeEvals[j].EmployeeScore
	})

	sort.Slice(unavailEmployeeEvals, func(i, j int) bool {
		return unavailEmployeeEvals[i].EmployeeScore > unavailEmployeeEvals[j].EmployeeScore
	})

	// Send the available ones first
	for _, employeeEval := range availEmployeeEvals {
		employeeEvalRes.Employee = employeeEval

		if err := stream.Send(&employeeEvalRes); err != nil {
			return err
		}
	}

	for _, employeeEval := range unavailEmployeeEvals {
		employeeEvalRes.Employee = employeeEval

		if err := stream.Send(&employeeEvalRes); err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) UpdateRosterAssignment(cxt context.Context, RosterAssignement *pb.RosterAssignement) (*pb.Response, error) {
	res := pb.Response{Type: pb.Response_ACK}

	numAffected, err := db_pck.UpdateRosterAssignments(
		s.db,
		RosterAssignement,
	)

	if err != nil {
		res.Type = pb.Response_ERROR
		res.ErrorMessage = err.Error()
	} else {
		fmt.Println(numAffected, "roster assignments were updated.")
	}

	return &res, nil
}
