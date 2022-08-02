package server

import (
	"fmt"
	"strconv"
	"strings"

	db_pck "capstone.operations_ecosystem/backend/database"
	pb "capstone.operations_ecosystem/backend/proto"
	"github.com/getsentry/sentry-go"

	"context"
)

// gRPC HTTP defined endpoint.
func (s *Server) GetRosterAssignmentsForWebApp(cxt context.Context, getRequest *pb.HTTPAssignmentsGetRequest) (*pb.HTTPAssignmentResponse, error) {
	fmt.Println("GetRosterAssignmentsFromWebApp")
	fmt.Println(getRequest)

	res := &pb.HTTPAssignmentResponse{Response: &pb.HTTPMessage{Status: 1}}

	// Verify Nonce
	err := s.verifyNonce(getRequest.Twan, &pb.User{TeleUserId: getRequest.TeleUserId})
	if err != nil {
		res.Response.Value = err.Error()
		return res, nil
	}

	// Todo: look up DB and get all assignments for this person
	user, err := db_pck.IdUserByTelegramId(s.db, int(getRequest.TeleUserId), true)
	if err != nil {
		res.Response.Value = "No user by this ID exists."
		return res, err
	}
	res.Response.Value = user.User.Name

	rosterQuery := &pb.RosterQuery{}
	db_pck.AddRosterFilter(rosterQuery, pb.RosterFilter_GUARD_ASSIGNED_ID, pb.Filter_EQUAL, strconv.Itoa(int(user.User.UserId)))
	db_pck.AddRosterFilter(rosterQuery, pb.RosterFilter_IS_ASSIGNED, pb.Filter_EQUAL, "1")
	if getRequest.StartDate != "null" {
		db_pck.AddRosterFilter(rosterQuery, pb.RosterFilter_START_TIME, pb.Filter_GREATER_EQ, getRequest.StartDate)
	}
	if getRequest.EndDate != "null" {
		db_pck.AddRosterFilter(rosterQuery, pb.RosterFilter_END_TIME, pb.Filter_LESSER_EQ, getRequest.EndDate)

	}
	rosters, err := db_pck.GetRosters(s.db, rosterQuery)
	if err != nil {
		res.Response.Status = 2
		res.Response.Value = "This user has no upcoming assignments."
		return res, err
	}

	fmt.Println(rosters)
	for _, roster := range rosters {
		newHTTPResRoster := pb.HTTPRosterResponse{
			AifsId:        roster.AifsId,
			StartDatetime: roster.StartTime,
			EndDatetime:   roster.EndTime,
		}
		addresses := make([]string, 0)
		for _, client := range roster.Clients {
			addresses = append(addresses, client.Client.Address)
		}
		newHTTPResRoster.Addresses = addresses
		res.Rosters = append(res.Rosters, &newHTTPResRoster)
	}
	res.Response.Status = 0
	return res, nil
}

// gRPC HTTP defined endpoint.
// Insert a new incident report incoming from the telegram bot web
// Assumes the creator of the incident report is the owner of the nonce in the header
func (s *Server) PostWReportFromWebApp(ctx context.Context, postReq *pb.HTTPReportPostRequest) (*pb.HTTPMessage, error) {
	defer sentry.Recover()

	fmt.Println("PostWReportFromWebApp recieved", postReq)
	res := &pb.HTTPMessage{Status: 1}

	// Verify Nonce
	err := s.verifyNonce(postReq.Twan, &pb.User{TeleUserId: postReq.TeleUserId})
	if err != nil {
		res.Value = err.Error()
		return res, nil
	}

	incidentTime := postReq.Date + " " + postReq.Time + ":00"
	incidentTime = strings.Replace(incidentTime, "-", ":", -1)

	//Convert into incidentReport
	incidentReport := &pb.IncidentReport{
		Creator: &pb.User{UserId: int64(postReq.TeleUserId)},
		IncidentReportContent: &pb.IncidentReportContent{
			Title:                 postReq.Title,
			Address:               postReq.Address,
			IncidentTime:          incidentTime,
			IsPoliceNotified:      postReq.IsPoliceNotified,
			Description:           postReq.Details,
			HasActionTaken:        postReq.IsActionTaken,
			ActionTaken:           postReq.ActionDetails,
			HasInjury:             postReq.IsPeopleInjured,
			InjuryDescription:     postReq.InjuryDetails,
			HasStolenItem:         postReq.IsPropertyStolen,
			StolenItemDescription: postReq.PropertyStolenDetails,
		},
	}

	//Convert type
	submittedReportType, err := strconv.Atoi(postReq.ReportType)
	if err != nil || submittedReportType > 2 {
		res.Value = "Invalid report type."
		fmt.Println("Err: Invalid Report Type:", err)
		return res, err
	}
	switch submittedReportType {
	case 0:
		incidentReport.Type = pb.IncidentReport_FIRE_ALARM
	case 1:
		incidentReport.Type = pb.IncidentReport_INTRUDER
	case 2:
		incidentReport.Type = pb.IncidentReport_OTHERS
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

	res.Status = 0

	return res, nil
}

// gRPC HTTP defined endpoint.
func (s *Server) PostRegistrationFormFromWebApp(cxt context.Context, req *pb.HTTPRegistrationFormRequest) (*pb.HTTPMessage, error) {
	//Validate token
	//Todo: Memory fault, ask Hannah
	regOtp, err := s.validateRegistrationToken(req.Code)
	if err != nil {
		fmt.Println("PostRegistrationFormFromWebApp ERROR:", err)
		return &pb.HTTPMessage{Status: 1, Value: err.Error()}, nil
	}

	// TODO tele handle
	user := &pb.User{
		UserType:       regOtp.UserType,
		Name:           req.Name,
		Email:          req.Email,
		PhoneNumber:    req.PhoneNumber,
		TelegramHandle: req.TeleHandle,
		IsPartTimer:    req.IsPartTime,
		TeleUserId:     req.TeleUserId,
	}
	fullUser := &pb.FullUser{
		User:           user,
		HashedPassword: req.HasedLoginString,
		SecurityString: req.LoginString,
	}

	_, err = db_pck.InsertUser(s.db, fullUser, s.dbLock)
	if err != nil {
		fmt.Println("PostRegistrationFormFromWebApp ERROR:", err)
		return &pb.HTTPMessage{Status: 1, Value: err.Error()}, nil
	}

	// Mark registration token as used
	regOtp.IsUsed = true
	_, err = db_pck.UpdateRegOtp(s.db, regOtp, nil)

	if err != nil {
		fmt.Println("PostRegistrationFormFromWebApp ERROR:", err)
		return &pb.HTTPMessage{Status: 1, Value: err.Error()}, nil
	}
	return &pb.HTTPMessage{Status: 0}, nil
}
