package fake_server

import (
	"fmt"

	pb "capstone.operations_ecosystem/backend/proto"

	"context"
)

func (s *Server) GetRosterAssignmentsForWebApp(cxt context.Context, getRequest *pb.HTTPAssignmentsGetRequest) (*pb.HTTPAssignmentResponse, error) {
	fmt.Println("GetRosterAssignmentsFromWebApp")
	fmt.Println(getRequest)
	return &pb.HTTPAssignmentResponse{}, nil
}

func (s *Server) PostWReportFromWebApp(cxt context.Context, incidentReport *pb.HTTPReportPostRequest) (*pb.HTTPMessage, error) {
	fmt.Println("PostWReportFromWebApp")
	fmt.Println(incidentReport)
	return &pb.HTTPMessage{Value: "OK"}, nil
}
