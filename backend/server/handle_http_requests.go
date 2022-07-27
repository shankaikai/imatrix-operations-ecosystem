package server

import (
	"fmt"

	db_pck "capstone.operations_ecosystem/backend/database"
	pb "capstone.operations_ecosystem/backend/proto"
	"github.com/getsentry/sentry-go"

	"context"
)

// gRPC HTTP defined endpoint.
func (s *Server) GetRosterAssignmentsForWebApp(cxt context.Context, getRequest *pb.HTTPAssignmentsGetRequest) (*pb.HTTPMessage, error) {
	fmt.Println("GetRosterAssignmentsFromWebApp")
	fmt.Println(getRequest)
	// Todo: look up DB and get all assignments for this person

	//This is a random format just for testing; feel free to change
	assignments := []string{"100 ABC Road; 1000hrs 10 May 2023; 1200hrs 11 May 2024",
					 "101 ABC Road; 1000hrs 11 May 2023; 1200hrs 12 May 2024"}

	return &pb.HTTPMessage{Status: 0, ValueArr: assignments}, nil
}

// gRPC HTTP defined endpoint.
// Insert a new incident report incoming from the telegram bot web
// Assumes the creator of the incident report is the owner of the nonce in the header
func (s *Server) PostWReportFromWebApp(ctx context.Context, postReq *pb.HTTPReportPostRequest) (*pb.HTTPMessage, error) {
	defer sentry.Recover()

	fmt.Println("PostWReportFromWebApp recieved", postReq)
	res := &pb.HTTPMessage{}

	// Verify Nonce
	err := s.verifyNonce(postReq.Twan, &pb.User{UserId: postReq.TeleUserId})
	if err != nil {
		res.Value = err.Error()
		return res, nil
	}

	//Convert into incidentReport
	incidentReport := &pb.IncidentReport{
		//Type: postReq.ReportType,
		Creator: &pb.User{UserId: int64(postReq.TeleUserId)},
		IncidentReportContent: &pb.IncidentReportContent{
			Address: postReq.Address,
			IncidentTime: postReq.Time,
			IsPoliceNotified: postReq.IsPoliceNotified,
			Title: postReq.Title,
			Description: postReq.Details,
		},
	}

	
	// get the user Id for this particular telegram user
	fullUser, err := db_pck.IdUserByTelegramId(s.db, int(incidentReport.Creator.TeleUserId), true)
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

// gRPC HTTP defined endpoint.
func (s *Server) CheckRegistrationCode(cxt context.Context, req *pb.HTTPMessage) (*pb.HTTPMessage, error) {
	//Todo: Check registration code and return a loginstring (security string)

	return &pb.HTTPMessage{Status: 0, Value: "fake_security_string"}, nil
}

// gRPC HTTP defined endpoint.
func (s *Server) PostRegistrationFormFromWebApp(cxt context.Context, req *pb.HTTPRegistrationFormRequest) (*pb.HTTPMessage, error) {
	//Todo

	return &pb.HTTPMessage{Status: 0}, nil
}
