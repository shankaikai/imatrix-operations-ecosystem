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

	fmt.Println("reportsss", incidentReport)
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

	// header := metadata.New(map[string]string{"Access-Control-Allow-Origin": "res-123"})
	// if err := grpc.SendHeader(ctx, header); err != nil {
	// 	return nil, status.Errorf(codes.Internal, "unable to send 'x-response-id' header")
	// }

	return res, nil
}
