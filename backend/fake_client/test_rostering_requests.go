package client

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"capstone.operations_ecosystem/backend/common"
	pb "capstone.operations_ecosystem/backend/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestRosteringClient(serverAddr *string, serverPort *int) {
	rosters := make([]*pb.Roster, 0)
	for i := 1; i < 4; i++ {
		rosters = append(rosters, CreateFakeRoster(i))
	}

	pk := InsertRoster(serverAddr, serverPort, rosters)
	rosters[2].RosteringId = pk

	ConsolidatedFindRosterTest(serverAddr, serverPort)
	// ConsolidatedUpdateRosterTest(serverAddr, serverPort, rosters[2])

	// DeleteRosterTest(serverAddr, serverPort, &pb.Roster{RosteringId: 9})
	// ConsolidatedGetAvailableUsersTest(serverAddr, serverPort)
}

func InsertRoster(serverAddr *string, serverPort *int, rosters []*pb.Roster) int64 {
	fmt.Println("insert roster test")
	grpcClient, conn := createRosterClient(serverAddr, serverPort)
	defer conn.Close()
	rosters[0].Clients = nil
	res, err := grpcClient.AddRoster(context.Background(), &pb.BulkRosters{Rosters: rosters})
	if err != nil {
		fmt.Println("InsertRoster ERROR:", err)
		return -1
	}

	if err != nil {
		fmt.Println("InsertRoster ERROR:", err)
		return -1
	}

	fmt.Println("Roster received response:", res.Type)

	if res.Type == pb.Response_ERROR {
		fmt.Println("Roster received error response:", res.ErrorMessage)
	}

	return res.PrimaryKey
}

func ConsolidatedFindRosterTest(serverAddr *string, serverPort *int) {
	FindRosterStartTimeFilter(serverAddr, serverPort)

	// FindRosterAssignmentFilter(serverAddr, serverPort)
	// FindRosterAifsClientIdFilter(serverAddr, serverPort)
	// FindRosterAifsIdFilter(serverAddr, serverPort)
	// FindRosterGuardIdFilter(serverAddr, serverPort)
	// FindRosterClientIdFilter(serverAddr, serverPort)
	// FindRosterGuardConfirmationFilter(serverAddr, serverPort)
	// FindRosterGuardAttendedFilter(serverAddr, serverPort)
	// FindRostersMultipleFilters(serverAddr, serverPort)
}

func FindRosterStartTimeFilter(serverAddr *string, serverPort *int) {
	fmt.Println("Finding Roster Start time filter")
	// startTime := time.Date(2022, 6, 21, 18, 0, 0, 0, time.UTC)
	startTime := time.Date(2022, 6, 8, 18, 0, 0, 0, time.UTC)
	com := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: startTime.Format(common.DATETIME_FORMAT)}
	filter := &pb.RosterFilter{Comparisons: com, Field: pb.RosterFilter_START_TIME}

	query := &pb.RosterQuery{Limit: 4, Filters: []*pb.RosterFilter{filter}}

	FindRostersTest(serverAddr, serverPort, query)
}

func FindRosterAssignmentFilter(serverAddr *string, serverPort *int) {
	fmt.Println("Finding roster assignment id filter")
	com := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: "1"}
	filter := &pb.RosterFilter{Comparisons: com, Field: pb.RosterFilter_ROSTER_ASSIGNMENT_ID}

	query := &pb.RosterQuery{Limit: 4, Filters: []*pb.RosterFilter{filter}}

	FindRostersTest(serverAddr, serverPort, query)
}

func FindRosterAifsClientIdFilter(serverAddr *string, serverPort *int) {
	fmt.Println("Finding roster aifs client id filter")
	com := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: "1"}
	filter := &pb.RosterFilter{Comparisons: com, Field: pb.RosterFilter_ROSTER_AIFS_CLIENT_ID}

	query := &pb.RosterQuery{Limit: 4, Filters: []*pb.RosterFilter{filter}}

	FindRostersTest(serverAddr, serverPort, query)
}

func FindRosterAifsIdFilter(serverAddr *string, serverPort *int) {
	fmt.Println("Finding roster aifs id filter")
	com := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: "1"}
	filter := &pb.RosterFilter{Comparisons: com, Field: pb.RosterFilter_AIFS_ID}

	query := &pb.RosterQuery{Limit: 4, Filters: []*pb.RosterFilter{filter}}

	FindRostersTest(serverAddr, serverPort, query)
}

func FindRosterGuardIdFilter(serverAddr *string, serverPort *int) {
	fmt.Println("Finding roster guard id filter")
	com := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: "1"}
	filter := &pb.RosterFilter{Comparisons: com, Field: pb.RosterFilter_GUARD_ASSIGNED_ID}

	query := &pb.RosterQuery{Limit: 4, Filters: []*pb.RosterFilter{filter}}

	FindRostersTest(serverAddr, serverPort, query)
}

func FindRosterClientIdFilter(serverAddr *string, serverPort *int) {
	fmt.Println("Finding roster client id filter")
	com := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: "1"}
	filter := &pb.RosterFilter{Comparisons: com, Field: pb.RosterFilter_CLIENT_ID}

	query := &pb.RosterQuery{Limit: 4, Filters: []*pb.RosterFilter{filter}}

	FindRostersTest(serverAddr, serverPort, query)
}

func FindRosterGuardConfirmationFilter(serverAddr *string, serverPort *int) {
	fmt.Println("Finding roster guard confirmation filter")
	com := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: "1"}
	filter := &pb.RosterFilter{Comparisons: com, Field: pb.RosterFilter_GUARD_ASSIGNMENT_CONFIRMATION}

	query := &pb.RosterQuery{Limit: 4, Filters: []*pb.RosterFilter{filter}}

	FindRostersTest(serverAddr, serverPort, query)
}

func FindRosterGuardAttendedFilter(serverAddr *string, serverPort *int) {
	fmt.Println("Finding roster guard attended filter")
	com := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: "1"}
	filter := &pb.RosterFilter{Comparisons: com, Field: pb.RosterFilter_GUARD_ASSIGNMENT_ATTENDED}

	query := &pb.RosterQuery{Limit: 4, Filters: []*pb.RosterFilter{filter}}

	FindRostersTest(serverAddr, serverPort, query)
}

func FindRostersMultipleFilters(serverAddr *string, serverPort *int) {
	fmt.Println("Finding roster mutiple filter")

	rosterFilters := make([]*pb.RosterFilter, 0)

	// Add filters here when there are more filters
	startTime := time.Date(2022, 6, 21, 18, 0, 0, 0, time.UTC)

	startTimecom := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: startTime.Format(common.DATETIME_FORMAT)}
	startTimeFilter := &pb.RosterFilter{Comparisons: startTimecom, Field: pb.RosterFilter_START_TIME}
	rosterFilters = append(rosterFilters, startTimeFilter)

	// client id
	cidCom := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: "1"}
	cidFilter := &pb.RosterFilter{Comparisons: cidCom, Field: pb.RosterFilter_CLIENT_ID}
	rosterFilters = append(rosterFilters, cidFilter)

	// guard id
	gcidCom := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: "1"}
	gidFilter := &pb.RosterFilter{Comparisons: gcidCom, Field: pb.RosterFilter_GUARD_ASSIGNED_ID}
	rosterFilters = append(rosterFilters, gidFilter)

	query := &pb.RosterQuery{Limit: 4, Filters: rosterFilters}
	FindRostersTest(serverAddr, serverPort, query)
}

func FindRostersTest(serverAddr *string, serverPort *int, query *pb.RosterQuery) {
	fmt.Println("Finding rosters...")
	grpcRoster, conn := createRosterClient(serverAddr, serverPort)
	defer conn.Close()

	stream, err := grpcRoster.FindRosters(context.Background(), query)

	if err != nil {
		fmt.Println("FindRostersNoFilter ERROR:", err)
		return
	}

	count := 0
	for {
		rosterRes, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("FindRostersNoFilter ERROR:", err)
		}

		fmt.Println("Roster received response:", rosterRes.Response.Type)
		if rosterRes.Response.Type == pb.Response_ERROR {
			continue
		}

		fmt.Println(count, ":", rosterRes.Roster)
		count++
	}
}

func ConsolidatedUpdateRosterTest(serverAddr *string, serverPort *int, roster *pb.Roster) {
	UpdateRosterBasicFields(serverAddr, serverPort, roster)
	UpdateRosterAssignments(serverAddr, serverPort, roster)
	UpdateRosterAifsClients(serverAddr, serverPort, roster)
}

func UpdateRosterBasicFields(serverAddr *string, serverPort *int, roster *pb.Roster) {
	fmt.Println("Roster Update time id:", roster.RosteringId, "Before time: ", roster.StartTime)

	updatedRoster := &pb.Roster{
		RosteringId: roster.RosteringId,
		StartTime:   roster.StartTime,
	}

	UpdateRosterTest(serverAddr, serverPort, updatedRoster)
}

func UpdateRosterAssignments(serverAddr *string, serverPort *int, roster *pb.Roster) {
	fmt.Println("Test Update roster assignments")

	updatedRoster := &pb.Roster{
		RosteringId:   roster.RosteringId,
		GuardAssigned: roster.GuardAssigned,
	}

	// replace one of the guards with someone else
	updatedRoster.GuardAssigned[0].GuardAssigned.Employee.UserId = 6

	UpdateRosterTest(serverAddr, serverPort, updatedRoster)
}

func UpdateRosterAifsClients(serverAddr *string, serverPort *int, roster *pb.Roster) {
	fmt.Println("Test Update roster aifs clients")

	updatedRoster := &pb.Roster{
		RosteringId: roster.RosteringId,
		Clients:     roster.Clients,
	}

	// replace one of the clients with another
	updatedRoster.Clients[0].Client.ClientId = 4

	UpdateRosterTest(serverAddr, serverPort, updatedRoster)
}

func UpdateRosterTest(serverAddr *string, serverPort *int, roster *pb.Roster) {
	fmt.Println("Updating roster:", roster.RosteringId)

	grpcRoster, conn := createRosterClient(serverAddr, serverPort)
	defer conn.Close()

	res, err := grpcRoster.UpdateRoster(context.Background(), roster)
	if err != nil {
		fmt.Println("UpdateRosterTest ERROR:", err)
		return
	}

	fmt.Println("Roster received response:", res.Type)

	if res.Type == pb.Response_ERROR {
		fmt.Println("Roster received error response:", res.ErrorMessage)
	}
}

func DeleteRosterTest(serverAddr *string, serverPort *int, roster *pb.Roster) {
	fmt.Println("Deleting roster:", roster.RosteringId)
	grpcRoster, conn := createRosterClient(serverAddr, serverPort)
	defer conn.Close()

	res, err := grpcRoster.DeleteRoster(context.Background(), roster)
	if err != nil {
		fmt.Println("DeleteRosterTest ERROR:", err)
		return
	}

	fmt.Println("Roster received response:", res.Type)

	if res.Type == pb.Response_ERROR {
		fmt.Println("Roster received error response:", res.ErrorMessage)
	}
}

func ConsolidatedGetAvailableUsersTest(serverAddr *string, serverPort *int) {
	getAvailableUsersTestNoRoster(serverAddr, serverPort)
	getAvailableUsersTestWithRoster(serverAddr, serverPort)
}

func getAvailableUsersTestNoRoster(serverAddr *string, serverPort *int) {
	// Give a fictional time in 2022 Jan
	startTime := time.Date(2022, 1, 25, 18, 0, 0, 0, time.UTC)
	endTime := time.Date(2022, 1, 26, 6, 0, 0, 0, time.UTC)

	availQuery := &pb.AvailabilityQuery{
		StartTime: startTime.Format(common.DATETIME_FORMAT),
		EndTime:   endTime.Format(common.DATETIME_FORMAT),
	}

	fmt.Println(availQuery)
	year, week := startTime.ISOWeek()
	fmt.Println("start year, week, day", year, week, startTime.Weekday())
	year, week = endTime.ISOWeek()
	fmt.Println("end year, week, day", year, week, endTime.Weekday())

	GetAvailableUsersTest(serverAddr, serverPort, availQuery)
}

func getAvailableUsersTestWithRoster(serverAddr *string, serverPort *int) {
	// create a roster for Aug 1 2022 from 6pm to 6am the next day
	startTimeTime := time.Date(2022, 4, 22, 18, 0, 0, 0, time.UTC)
	endTimeTime := time.Date(2022, 4, 22, 6, 0, 0, 0, time.UTC)
	startTime := &timestamppb.Timestamp{Seconds: startTimeTime.Unix()}
	endTime := &timestamppb.Timestamp{Seconds: endTimeTime.Unix()}

	year, week := startTimeTime.ISOWeek()
	fmt.Println("start year, week, day", year, week, startTimeTime.Weekday())
	year, week = endTimeTime.ISOWeek()
	fmt.Println("end year, week, day", year, week, endTimeTime.Weekday())

	rosters := make([]*pb.Roster, 0)
	for i := 1; i < 4; i++ {
		roster := CreateFakeRoster(i)
		roster.StartTime = startTimeTime.Format(common.DATETIME_FORMAT)
		roster.EndTime = endTimeTime.Format(common.DATETIME_FORMAT)
		for _, assignment := range roster.GuardAssigned {
			assignment.CustomStartTime = startTime
			assignment.CustomEndTime = endTime
		}
		rosters = append(rosters, roster)
	}

	InsertRoster(serverAddr, serverPort, rosters)

	availQuery := &pb.AvailabilityQuery{
		StartTime: startTimeTime.Format(common.DATETIME_FORMAT),
		EndTime:   endTimeTime.Format(common.DATETIME_FORMAT),
	}
	GetAvailableUsersTest(serverAddr, serverPort, availQuery)
}

func GetAvailableUsersTest(serverAddr *string, serverPort *int, query *pb.AvailabilityQuery) {
	fmt.Println("Finding Available Users ...")
	grpcRoster, conn := createRosterClient(serverAddr, serverPort)
	defer conn.Close()

	stream, err := grpcRoster.GetAvailableUsers(context.Background(), query)

	if err != nil {
		fmt.Println("GetAvailableUsersTest ERROR:", err)
		return
	}

	count := 0
	for {
		availabilityRes, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("GetAvailableUsersTest ERROR:", err)
			continue
		}

		fmt.Println("Availability check received response:", availabilityRes.Response.Type)
		if availabilityRes.Response.Type == pb.Response_ERROR {
			continue
		}

		fmt.Println(count, ":", availabilityRes.Employee)
		count++
	}
}

func createRosterClient(serverAddr *string, serverPort *int) (pb.RosterServicesClient, *grpc.ClientConn) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", *serverAddr, *serverPort), opts...)
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewRosterServicesClient(conn)

	return client, conn
}
