package fake_server

import (
	"fmt"
	"time"

	"capstone.operations_ecosystem/backend/common"
	pb "capstone.operations_ecosystem/backend/proto"

	"context"
)

func (s *Server) AddIncidentReport(cxt context.Context, incidentReport *pb.IncidentReport) (*pb.Response, error) {
	fmt.Println("AddIncidentReport")
	res := pb.Response{Type: pb.Response_ACK, PrimaryKey: 1}
	return &res, nil
}

func (s *Server) UpdateIncidentReport(cxt context.Context, incidentReport *pb.IncidentReport) (*pb.Response, error) {
	fmt.Println("UpdateIncidentReport")
	res := pb.Response{Type: pb.Response_ACK, PrimaryKey: 1}
	return &res, nil
}

func (s *Server) DeleteIncidentReport(cxt context.Context, incidentReport *pb.IncidentReport) (*pb.Response, error) {
	fmt.Println("DeleteIncidentReport")
	res := pb.Response{Type: pb.Response_ACK, PrimaryKey: 1}
	return &res, nil
}

func (s *Server) FindIncidentReports(query *pb.IncidentReportQuery, stream pb.IncidentReportServices_FindIncidentReportsServer) error {
	fmt.Println("FindIncidentReports")

	res := pb.Response{Type: pb.Response_ACK}
	incidentReportRes := pb.IncidentReportResponse{Response: &res}

	user := &pb.User{
		UserId:          1,
		UserType:        pb.User_ISPECIALIST,
		Name:            "test name",
		Email:           "email",
		PhoneNumber:     "1232",
		TelegramHandle:  "sfds",
		UserSecurityImg: "dsfds",
		IsPartTimer:     false,
	}

	content := &pb.IncidentReportContent{
		ReportContentId:       1,
		LastModifiedDate:      time.Now().Format(common.DATETIME_FORMAT),
		LastModifedUser:       user,
		Address:               "address",
		IncidentTime:          time.Now().Format(common.DATETIME_FORMAT),
		IsPoliceNotified:      false,
		Title:                 "Title",
		Description:           "Description",
		HasActionTaken:        false,
		ActionTaken:           "",
		HasInjury:             true,
		InjuryDescription:     "InjuryDescription",
		HasStolenItem:         true,
		StolenItemDescription: "StolenItemDescription",
		ReportImageLink:       "",
	}

	for i := 0; i < 3; i++ {
		incidentReport := &pb.IncidentReport{
			IncidentReportId:      int64(i),
			Type:                  pb.IncidentReport_INTRUDER,
			Creator:               user,
			CreationDate:          time.Now().Format(common.DATETIME_FORMAT),
			LastModifiedDate:      time.Now().Format(common.DATETIME_FORMAT),
			LastModifedUser:       user,
			IsOriginal:            i%2 == 0,
			IsApproved:            i%2 != 0,
			Signature:             nil,
			ApprovalDate:          "",
			IncidentReportContent: content,
		}

		incidentReportRes.IncidentReport = incidentReport

		if err := stream.Send(&incidentReportRes); err != nil {
			return err
		}
	}

	return nil
}
