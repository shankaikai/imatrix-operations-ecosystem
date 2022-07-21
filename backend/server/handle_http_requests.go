package server

import (
	"fmt"
	"strconv"

	db_pck "capstone.operations_ecosystem/backend/database"
	pb "capstone.operations_ecosystem/backend/proto"
	"github.com/getsentry/sentry-go"

	"context"
)

func (s *Server) GetRosterAssignmentsForWebApp(cxt context.Context, HTTPRosterMessage *pb.HTTPRosterMessage) (*pb.HTTPMessage, error) {
	fmt.Println("GetRosterAssignmentsFromWebApp")
	fmt.Println(HTTPRosterMessage)
	return &pb.HTTPMessage{}, nil
}

func (s *Server) PostWReportFromWebApp(ctx context.Context, incidentReport *pb.IncidentReport) (*pb.HTTPMessage, error) {
	defer sentry.Recover()

	fmt.Println("PostWReportFromWebApp recieved", incidentReport)

	res := &pb.HTTPMessage{}

	fmt.Println("Report:\n", incidentReport)
	// Fill up the blank values of the pb message
	err := s.insertDefaultIncidentReportValues(incidentReport)
	if err != nil {
		return res, err
	}

	pk, err := db_pck.InsertIncidentReport(
		s.db,
		incidentReport,
		s.dbLock,
	)

	if err != nil {
		res.Value = err.Error()
	}

	res.Value = strconv.Itoa(int(pk))

	return res, nil
}
