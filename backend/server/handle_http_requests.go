package server

import (
	"fmt"

	db_pck "capstone.operations_ecosystem/backend/database"
	pb "capstone.operations_ecosystem/backend/proto"
	"github.com/getsentry/sentry-go"

	"context"
)

// gRPC HTTP defined endpoint.
func (s *Server) GetRosterAssignmentsForWebApp(cxt context.Context, HTTPRosterMessage *pb.HTTPRosterMessage) (*pb.HTTPMessage, error) {
	fmt.Println("GetRosterAssignmentsFromWebApp")
	fmt.Println(HTTPRosterMessage)
	return &pb.HTTPMessage{}, nil
}

// gRPC HTTP defined endpoint.
// Insert a new incident report incoming from the telegram bot web
// Assumes the creator of the incident report is the owner of the nonce in the header
func (s *Server) PostWReportFromWebApp(ctx context.Context, incidentReport *pb.IncidentReport) (*pb.HTTPMessage, error) {
	defer sentry.Recover()

	fmt.Println("PostWReportFromWebApp recieved", incidentReport)
	res := &pb.HTTPMessage{}

	// Verify Nonce
	err := s.verifyNonceFromHeaders(ctx, incidentReport.Creator)
	if err != nil {
		res.Value = err.Error()
		return res, nil
	}

	// get the user Id for this particular telegram user
	fullUser, err := db_pck.IdUserByTelegramId(s.db, int(incidentReport.Creator.TeleChatId), true)
	if err != nil {
		return res, err
	}
	// put the user id into the incident report
	incidentReport.Creator = fullUser.User

	// Fill up the blank values of the pb message
	err = s.insertDefaultIncidentReportValues(incidentReport)
	if err != nil {
		return res, err
	}

	_, err = db_pck.InsertIncidentReport(
		s.db,
		incidentReport,
		s.dbLock,
	)

	if err != nil {
		res.Value = err.Error()
	}

	res.Value = "1"

	return res, nil
}
