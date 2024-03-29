package server

import (
	"context"
	"fmt"
	"sort"
	"strconv"

	db_pck "capstone.operations_ecosystem/backend/database"
	pb "capstone.operations_ecosystem/backend/proto"
)

// gRPC defined endpoint. Add a roster to the DB.
func (s *Server) AddRoster(cxt context.Context, rosters *pb.BulkRosters) (*pb.Response, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered", r)
		}
	}()

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

// gRPC defined endpoint. Update multiple rosters in the DB.
func (s *Server) UpdateRoster(cxt context.Context, rosters *pb.BulkRosters) (*pb.Response, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered", r)
		}
	}()

	fmt.Println("UPDATE ROSTER SEARCH HERE", rosters.Rosters[0].RosteringId)
	res := pb.Response{Type: pb.Response_ACK}
	for _, roster := range rosters.Rosters {

		numAffected, newRosterAssignmentsPk, err := db_pck.UpdateRoster(
			s.db,
			roster,
			s.dbLock,
		)

		if err != nil {
			res.Type = pb.Response_ERROR
			res.ErrorMessage = err.Error()
		} else {
			fmt.Println(numAffected, "rosters were updated.")
			s.sendUpdatedRostersToTele(newRosterAssignmentsPk)
		}
	}
	return &res, nil
}

// gRPC defined endpoint. Delete a roster in the DB.
func (s *Server) DeleteRoster(cxt context.Context, roster *pb.Roster) (*pb.Response, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered", r)
		}
	}()

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

// gRPC defined endpoint. Find rosters in the DB. The rosters are filtered out based the query.
func (s *Server) FindRosters(query *pb.RosterQuery, stream pb.RosterServices_FindRostersServer) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered", r)
		}
	}()

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

// gRPC defined endpoint.
// Send back all the available users that are not yet assigned for the particular day
func (s *Server) GetAvailableUsers(query *pb.AvailabilityQuery, stream pb.RosterServices_GetAvailableUsersServer) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered", r)
		}
	}()

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

// gRPC defined endpoint. Update a particular roster assignment in the DB.
// This is mostly used to save an acknowledgement or a rejection of a roster assignment.
func (s *Server) UpdateRosterAssignment(cxt context.Context, RosterAssignement *pb.RosterAssignement) (*pb.Response, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered", r)
		}
	}()

	res := pb.Response{Type: pb.Response_ACK}
	query := &pb.RosterQuery{}
	db_pck.AddRosterFilter(query, pb.RosterFilter_ROSTER_ASSIGNMENT_ID, pb.Filter_EQUAL, strconv.Itoa(int(RosterAssignement.RosterAssignmentId)))
	numAffected, err := db_pck.UpdateRosterAssignments(
		s.db,
		RosterAssignement,
		query,
	)

	if err != nil {
		res.Type = pb.Response_ERROR
		res.ErrorMessage = err.Error()
	} else {
		fmt.Println(numAffected, "roster assignments were updated.")
	}

	return &res, nil
}

// gRPC defined endpoint. Find roster assignments in the DB. The roster assignments are filtered out based the query.
func (s *Server) FindRosterAssignments(query *pb.RosterQuery, stream pb.RosterServices_FindRosterAssignmentsServer) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered", r)
		}
	}()

	res := pb.Response{Type: pb.Response_ACK}

	foundRosterAssignments, err := db_pck.GetRosterAssingments(
		s.db,
		query,
		-1,
	)

	if err != nil {
		rosterRes := pb.RosterAssignmentResponse{Response: &res}
		res.Type = pb.Response_ERROR
		res.ErrorMessage = err.Error()
		stream.Send(&rosterRes)
		return nil
	}

	fmt.Println("FindRosters: Found rosters to return")

	for _, rosterAsgn := range foundRosterAssignments {
		rosterRes := pb.RosterAssignmentResponse{Response: &res}
		rosterRes.RosterAssignment = rosterAsgn
		if err := stream.Send(&rosterRes); err != nil {
			return err
		}
	}

	return nil
}
