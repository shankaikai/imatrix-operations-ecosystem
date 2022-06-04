package client

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "capstone.operations_ecosystem/backend/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestRosteringClient(serverAddr *string, serverPort *int) {
	// rosters := make([]*pb.Roster, 0)
	// for i := 1; i < 4; i++ {
	// 	rosters = append(rosters, createFakeRoster(i))
	// }

	// // pk := InsertRoster(serverAddr, serverPort, rosters)
	// rosters[2].RosteringId = 4 //pk

	// ConsolidatedFindRosterTest(serverAddr, serverPort)
	// ConsolidatedUpdateRosterTest(serverAddr, serverPort, rosters[2])

	// DeleteRosterTest(serverAddr, serverPort, &pb.Roster{RosteringId: 9})
	// FindRosterIdFilter(serverAddr, serverPort)
	ConsolidatedGetAvailableUsersTest(serverAddr, serverPort)
}

func InsertRoster(serverAddr *string, serverPort *int, rosters []*pb.Roster) int64 {
	grpcClient, conn := createRosterClient(serverAddr, serverPort)
	defer conn.Close()

	stream, err := grpcClient.AddRoster(context.Background())
	if err != nil {
		fmt.Println("InsertRoster ERROR:", err)
		return -1
	}

	for _, roster := range rosters {
		fmt.Println("Inserting roster:", roster.RosteringId)
		if err := stream.Send(roster); err != nil {
			fmt.Println("InsertRoster ERROR:", err)
			return -1
		}
	}

	res, err := stream.CloseAndRecv()

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
	FindRostersNoFilter(serverAddr, serverPort)
	FindRosterIdFilter(serverAddr, serverPort)
	FindRosterAssignmentFilter(serverAddr, serverPort)
	FindRosterAifsClientIdFilter(serverAddr, serverPort)
	FindRosterAifsIdFilter(serverAddr, serverPort)
	FindRosterGuardIdFilter(serverAddr, serverPort)
	FindRosterClientIdFilter(serverAddr, serverPort)
	FindRosterGuardConfirmationFilter(serverAddr, serverPort)
	FindRosterGuardAttendedFilter(serverAddr, serverPort)
	FindRostersMultipleFilters(serverAddr, serverPort)
}

func FindRostersNoFilter(serverAddr *string, serverPort *int) {
	fmt.Println("Finding rosters without filter")
	FindRostersTest(serverAddr, serverPort, &pb.RosterQuery{Limit: 5})
}

func FindRosterIdFilter(serverAddr *string, serverPort *int) {
	fmt.Println("Finding roster id filter")
	com := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: "18"}
	filter := &pb.RosterFilter{Comparisons: com, Field: pb.RosterFilter_ROSTER_ID}

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
	fmt.Println("Roster Update time id:", roster.RosteringId, "Before time: ", roster.StartTime.String())

	updatedRoster := &pb.Roster{
		RosteringId: roster.RosteringId,
		StartTime:   &timestamppb.Timestamp{Seconds: roster.StartTime.Seconds - 100},
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
	// getAvailableUsersTestWithRoster(serverAddr, serverPort)
}

func getAvailableUsersTestNoRoster(serverAddr *string, serverPort *int) {
	// Give a fictional time in 2022 Jan
	startTimeTime := time.Date(2022, 1, 25, 18, 0, 0, 0, time.UTC)
	endTimeTime := time.Date(2022, 1, 26, 6, 0, 0, 0, time.UTC)
	startTime := &timestamppb.Timestamp{Seconds: startTimeTime.Unix()}
	endTime := &timestamppb.Timestamp{Seconds: endTimeTime.Unix()}

	availQuery := &pb.AvailabilityQuery{StartTime: startTime, EndTime: endTime}

	fmt.Println(availQuery)
	year, week := startTimeTime.ISOWeek()
	fmt.Println("start year, week, day", year, week, startTimeTime.Weekday())
	year, week = endTimeTime.ISOWeek()
	fmt.Println("end year, week, day", year, week, endTimeTime.Weekday())

	GetAvailableUsersTest(serverAddr, serverPort, availQuery)
}

func getAvailableUsersTestWithRoster(serverAddr *string, serverPort *int) {
	// create a roster for Aug 1 2022 from 6pm to 6am the next day
	startTimeTime := time.Date(2022, 8, 1, 18, 0, 0, 0, time.UTC)
	endTimeTime := time.Date(2022, 8, 2, 6, 0, 0, 0, time.UTC)
	startTime := &timestamppb.Timestamp{Seconds: startTimeTime.Unix()}
	endTime := &timestamppb.Timestamp{Seconds: endTimeTime.Unix()}

	year, week := startTimeTime.ISOWeek()
	fmt.Println("start year, week, day", year, week, startTimeTime.Weekday())
	year, week = endTimeTime.ISOWeek()
	fmt.Println("end year, week, day", year, week, endTimeTime.Weekday())

	rosters := make([]*pb.Roster, 0)
	for i := 1; i < 4; i++ {
		roster := createFakeRoster(i)
		roster.StartTime = startTime
		roster.EndTime = endTime
		for _, assignment := range roster.GuardAssigned {
			assignment.CustomStartTime = startTime
			assignment.CustomEndTime = endTime
		}
		rosters = append(rosters, roster)
	}

	InsertRoster(serverAddr, serverPort, rosters)

	availQuery := &pb.AvailabilityQuery{StartTime: startTime, EndTime: endTime}
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
