package server

import (
	"fmt"
	"time"

	"capstone.operations_ecosystem/backend/common"
	db_pck "capstone.operations_ecosystem/backend/database"
	pb "capstone.operations_ecosystem/backend/proto"
)

const (
	// If the I-Specialist submitted the report after duty
	// we do not know who submitted the report
	// Allow them to submit the report up till 1h after.
	INCIDENT_REPORTING_END_HOUR_BUFFER = 1
)

func (s *Server) insertDefaultIncidentReportValues(report *pb.IncidentReport) error {
	if report.Creator.UserId == int64(s.Config.Aifs1Id) ||
		report.Creator.UserId == int64(s.Config.Aifs2Id) ||
		report.Creator.UserId == int64(s.Config.Aifs3Id) {

		// Get the AIFS ID
		switch report.Creator.UserId {
		case int64(s.Config.Aifs1Id):
			report.AifsId = 1
		case int64(s.Config.Aifs2Id):
			report.AifsId = 2
		case int64(s.Config.Aifs3Id):
			report.AifsId = 3
		}

		// Get actual user from roster
		rosterQuery := &pb.RosterQuery{}
		// Start time
		db_pck.AddRosterFilter(rosterQuery, pb.RosterFilter_START_TIME,
			pb.Filter_LESSER_EQ, time.Now().Format(common.DATETIME_FORMAT))
		// End time + buffer of 1h
		db_pck.AddRosterFilter(rosterQuery, pb.RosterFilter_END_TIME,
			pb.Filter_GREATER_EQ, time.Now().Add(time.Hour*INCIDENT_REPORTING_END_HOUR_BUFFER).
				Format(common.DATETIME_FORMAT))
		db_pck.AddRosterFilter(rosterQuery, pb.RosterFilter_IS_ASSIGNED,
			pb.Filter_EQUAL, "1")

		rosterAssignment, err := db_pck.GetRosterAssingments(s.db, rosterQuery, -1)

		if err != nil {
			fmt.Println("insertDefaultIncidentReportValues", err)
			return err
		}

		if len(rosterAssignment) > 0 {
			report.Creator = rosterAssignment[0].GuardAssigned.Employee
		}

	} else {
		report.AifsId = -1
	}

	report.CreationDate = time.Now().Format(common.DATETIME_FORMAT)
	report.IncidentReportContent.LastModifedUser = report.Creator
	report.IncidentReportContent.LastModifiedDate = report.CreationDate
	return nil
}
