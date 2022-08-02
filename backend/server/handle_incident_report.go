package server

import (
	"context"
	"fmt"

	db_pck "capstone.operations_ecosystem/backend/database"
	pb "capstone.operations_ecosystem/backend/proto"
	"github.com/getsentry/sentry-go"
)

// gRPC defined endpoint. Add an incident report to the DB.
func (s *Server) AddIncidentReport(cxt context.Context, incidentReport *pb.IncidentReport) (*pb.Response, error) {
	defer sentry.Recover()

	fmt.Println("AddIncidentReport")
	res := &pb.Response{Type: pb.Response_ACK, PrimaryKey: 1}

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
		res.Type = pb.Response_ERROR
		res.ErrorMessage = err.Error()
	}

	res.PrimaryKey = pk

	return res, nil
}

// gRPC defined endpoint. Update an incident report in the DB.
// Two copies of the incident report are stored in the DB, the original report
// and the most updated copy of the report.
func (s *Server) UpdateIncidentReport(cxt context.Context, incidentReport *pb.IncidentReport) (*pb.Response, error) {
	defer sentry.Recover()

	fmt.Println("UPDATE INCIDENT REPORT ", incidentReport.IncidentReportId)
	fmt.Println("hdsfdsfds", incidentReport)
	res := pb.Response{Type: pb.Response_ACK}

	numAffected, err := db_pck.UpdateIncidentReport(
		s.db,
		incidentReport,
		s.dbLock,
	)

	if err != nil {
		res.Type = pb.Response_ERROR
		res.ErrorMessage = err.Error()
	} else {
		fmt.Println(numAffected, "incidentReports were updated.")
	}

	return &res, nil
}

// gRPC defined endpoint. Delete an incident report in the DB.
func (s *Server) DeleteIncidentReport(cxt context.Context, incidentReport *pb.IncidentReport) (*pb.Response, error) {
	defer sentry.Recover()

	res := pb.Response{Type: pb.Response_ACK}
	numDel, err := db_pck.DeleteIncidentReport(
		s.db,
		incidentReport,
	)

	if err != nil {
		res.Type = pb.Response_ERROR
		res.ErrorMessage = err.Error()
	} else {
		fmt.Println(numDel, "incidentReports were deleted.")
	}

	return &res, nil
}

// gRPC defined endpoint. Find incident reports in the DB. The reports are filtered out based the query.
func (s *Server) FindIncidentReports(query *pb.IncidentReportQuery, stream pb.IncidentReportServices_FindIncidentReportsServer) error {
	defer sentry.Recover()

	res := pb.Response{Type: pb.Response_ACK}

	foundIncidentReports, err := db_pck.GetIncidentReports(
		s.db,
		query,
	)

	if err != nil {
		incidentReportRes := pb.IncidentReportResponse{Response: &res}
		res.Type = pb.Response_ERROR
		res.ErrorMessage = err.Error()
		stream.Send(&incidentReportRes)
		return nil
	}

	if len(foundIncidentReports) > 0 {
		fmt.Println("FindIncidentReports: Found incidentReports to return")
	}

	for _, incidentReport := range foundIncidentReports {
		incidentReportRes := pb.IncidentReportResponse{Response: &res}
		incidentReportRes.IncidentReport = incidentReport
		// fmt.Println(incidentReport)
		if err := stream.Send(&incidentReportRes); err != nil {
			return err
		}
	}

	return nil
}
