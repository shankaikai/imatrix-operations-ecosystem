package fake_server

import (
	"fmt"

	pb "capstone.operations_ecosystem/backend/proto"

	"context"
)

func (s *Server) GetRosterAssignmentsForWebApp(cxt context.Context, HTTPRosterMessage *pb.HTTPRosterMessage) (*pb.HTTPMessage, error) {
	fmt.Println("GetRosterAssignmentsFromWebApp")
	fmt.Println(HTTPRosterMessage)
	return &pb.HTTPMessage{}, nil
}

func (s *Server) PostWReportFromWebApp(cxt context.Context, incidentReport *pb.IncidentReport) (*pb.HTTPMessage, error) {
	fmt.Println("PostWReportFromWebApp")
	fmt.Println(incidentReport)
	return &pb.HTTPMessage{Value: "OK"}, nil
}
