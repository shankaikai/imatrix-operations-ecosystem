package client

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "capstone.operations_ecosystem/backend/proto"
)

func TestBroadcastClient(serverAddr *string, serverPort *int) {
	// broadcast := createFakeBroadcast(1)
	// InsertBroadcast(serverAddr, serverPort, broadcast)
	ConsolidatedFindBroadcastTest(serverAddr, serverPort)
	DeleteBroadcast(serverAddr, serverPort, &pb.Broadcast{BroadcastId: 5})
}

func InsertBroadcast(serverAddr *string, serverPort *int, broadcast *pb.Broadcast) {
	// Ensure that there are users first, if there are users already existing,
	// the returned users will be different, but its ok.
	for i := 0; i < len(broadcast.Receipients); i++ {
		InsertUser(serverAddr, serverPort, broadcast.Receipients[0].Recipient)
	}

	fmt.Println("Inserting Broadcast:", broadcast.Title)
	client, conn := createBroadcastClient(serverAddr, serverPort)
	defer conn.Close()

	res, err := client.AddBroadcast(context.Background(), broadcast)
	if err != nil {
		fmt.Println("InsertBroadcast ERROR:", err)
		return
	}

	fmt.Println("Client received response:", res.Type)

	if res.Type == pb.Response_ERROR {
		fmt.Println("Client received error response:", res.ErrorMessage)
	}
}

func ConsolidatedFindBroadcastTest(serverAddr *string, serverPort *int) {
	FindBroadcastsNoFilter(serverAddr, serverPort)
	FindBroadcastsIdFilter(serverAddr, serverPort)
	FindBroadcastsTypeFilter(serverAddr, serverPort)
	FindBroadcastsTitleFilter(serverAddr, serverPort)
	FindBroadcastsContentFilter(serverAddr, serverPort)
	FindBroadcastsCreateDateFilter(serverAddr, serverPort)
	FindBroadcastsDeadlineFilter(serverAddr, serverPort)
	FindBroadcastsCreatorIdFilter(serverAddr, serverPort)
	FindBroadcastsRecipientIdFilter(serverAddr, serverPort)
	FindBroadcastsNumRecFilter(serverAddr, serverPort)
	FindBroadcastsMultipleFilters(serverAddr, serverPort)
}

func FindBroadcastsNoFilter(serverAddr *string, serverPort *int) *pb.BulkBroadcasts {
	fmt.Println("Finding broadcasts without filter")
	return FindBroadcastsTest(serverAddr, serverPort, &pb.BroadcastQuery{Limit: 4})
}

func FindBroadcastsIdFilter(serverAddr *string, serverPort *int) *pb.BulkBroadcasts {
	fmt.Println("Finding broadcasts id filter")
	com := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: "15"}
	filter := &pb.BroadcastFilter{Comparisons: com, Field: pb.BroadcastFilter_BROADCAST_ID}

	query := &pb.BroadcastQuery{Limit: 4, Filters: []*pb.BroadcastFilter{filter}}

	return FindBroadcastsTest(serverAddr, serverPort, query)
}

func FindBroadcastsTypeFilter(serverAddr *string, serverPort *int) *pb.BulkBroadcasts {
	fmt.Println("Finding broadcasts type filter")

	com := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: pb.Broadcast_ASSIGNMENT.String()}
	filter := &pb.BroadcastFilter{Comparisons: com, Field: pb.BroadcastFilter_TYPE}

	query := &pb.BroadcastQuery{Limit: 4, Filters: []*pb.BroadcastFilter{filter}}

	return FindBroadcastsTest(serverAddr, serverPort, query)
}

func FindBroadcastsTitleFilter(serverAddr *string, serverPort *int) *pb.BulkBroadcasts {
	fmt.Println("Finding broadcasts title filter")

	com := &pb.Filter{Comparison: pb.Filter_CONTAINS, Value: "name1"}
	filter := &pb.BroadcastFilter{Comparisons: com, Field: pb.BroadcastFilter_TITLE}

	query := &pb.BroadcastQuery{Limit: 4, Filters: []*pb.BroadcastFilter{filter}}

	return FindBroadcastsTest(serverAddr, serverPort, query)
}

func FindBroadcastsContentFilter(serverAddr *string, serverPort *int) *pb.BulkBroadcasts {
	fmt.Println("Finding broadcasts contents filter")

	com := &pb.Filter{Comparison: pb.Filter_CONTAINS, Value: "con"}
	filter := &pb.BroadcastFilter{Comparisons: com, Field: pb.BroadcastFilter_CONTENT}

	query := &pb.BroadcastQuery{Limit: 4, Filters: []*pb.BroadcastFilter{filter}}

	return FindBroadcastsTest(serverAddr, serverPort, query)
}

func FindBroadcastsCreateDateFilter(serverAddr *string, serverPort *int) *pb.BulkBroadcasts {
	creationDateMax := time.Now().AddDate(0, 0, -1)

	fmt.Println("Finding broadcasts creation date filter, max date:", creationDateMax)
	com := &pb.Filter{Comparison: pb.Filter_LESSER_EQ, Value: creationDateMax.Format("2006-01-02 15:04:05")}
	filter := &pb.BroadcastFilter{Comparisons: com, Field: pb.BroadcastFilter_CREATION_DATE}

	query := &pb.BroadcastQuery{Limit: 4, Filters: []*pb.BroadcastFilter{filter}}

	return FindBroadcastsTest(serverAddr, serverPort, query)
}

func FindBroadcastsDeadlineFilter(serverAddr *string, serverPort *int) *pb.BulkBroadcasts {
	deadlineMax := time.Now().AddDate(0, 0, 30)

	fmt.Println("Finding broadcasts deadline date filter, max date:", deadlineMax)
	com := &pb.Filter{Comparison: pb.Filter_LESSER_EQ, Value: deadlineMax.Format("2006-01-02 15:04:05")}
	filter := &pb.BroadcastFilter{Comparisons: com, Field: pb.BroadcastFilter_CREATION_DATE}

	query := &pb.BroadcastQuery{Limit: 4, Filters: []*pb.BroadcastFilter{filter}}

	return FindBroadcastsTest(serverAddr, serverPort, query)
}

func FindBroadcastsCreatorIdFilter(serverAddr *string, serverPort *int) *pb.BulkBroadcasts {
	fmt.Println("Finding broadcasts creator id filter")
	com := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: "1"}
	filter := &pb.BroadcastFilter{Comparisons: com, Field: pb.BroadcastFilter_CREATOR_ID}

	query := &pb.BroadcastQuery{Limit: 4, Filters: []*pb.BroadcastFilter{filter}}

	return FindBroadcastsTest(serverAddr, serverPort, query)
}

func FindBroadcastsRecipientIdFilter(serverAddr *string, serverPort *int) *pb.BulkBroadcasts {
	fmt.Println("Finding broadcasts recipient id filter")
	com := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: "2"}
	filter := &pb.BroadcastFilter{Comparisons: com, Field: pb.BroadcastFilter_RECEIPEIENT_ID}

	query := &pb.BroadcastQuery{Limit: 4, Filters: []*pb.BroadcastFilter{filter}}

	return FindBroadcastsTest(serverAddr, serverPort, query)
}

func FindBroadcastsNumRecFilter(serverAddr *string, serverPort *int) *pb.BulkBroadcasts {
	fmt.Println("Finding broadcasts number of recipients filter")
	com := &pb.Filter{Comparison: pb.Filter_GREATER_EQ, Value: "2"}
	filter := &pb.BroadcastFilter{Comparisons: com, Field: pb.BroadcastFilter_NUM_RECEIPIENTS}

	query := &pb.BroadcastQuery{Limit: 4, Filters: []*pb.BroadcastFilter{filter}}

	return FindBroadcastsTest(serverAddr, serverPort, query)
}

func FindBroadcastsMultipleFilters(serverAddr *string, serverPort *int) *pb.BulkBroadcasts {
	fmt.Println("Finding broadcasts mutiple filter")

	broadcastFilters := make([]*pb.BroadcastFilter, 0)

	// type
	typeCom := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: pb.Broadcast_ANNOUNCEMENT.String()}
	typeFilter := &pb.BroadcastFilter{Comparisons: typeCom, Field: pb.BroadcastFilter_TYPE}
	broadcastFilters = append(broadcastFilters, typeFilter)

	// creation date
	creationDateMax := time.Now().AddDate(0, 0, -1)
	cdFom := &pb.Filter{Comparison: pb.Filter_LESSER_EQ, Value: creationDateMax.Format("2006-01-02 15:04:05")}
	cdFilter := &pb.BroadcastFilter{Comparisons: cdFom, Field: pb.BroadcastFilter_CREATION_DATE}
	broadcastFilters = append(broadcastFilters, cdFilter)

	// creator id
	creatorCom := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: "1"}
	creatorFilter := &pb.BroadcastFilter{Comparisons: creatorCom, Field: pb.BroadcastFilter_CREATOR_ID}
	broadcastFilters = append(broadcastFilters, creatorFilter)

	// num recipients
	ncCom := &pb.Filter{Comparison: pb.Filter_GREATER_EQ, Value: "2"}
	ncFilter := &pb.BroadcastFilter{Comparisons: ncCom, Field: pb.BroadcastFilter_NUM_RECEIPIENTS}
	broadcastFilters = append(broadcastFilters, ncFilter)

	query := &pb.BroadcastQuery{Limit: 4, Filters: broadcastFilters}
	return FindBroadcastsTest(serverAddr, serverPort, query)
}

func FindBroadcastsTest(serverAddr *string, serverPort *int, query *pb.BroadcastQuery) *pb.BulkBroadcasts {
	fmt.Println("Finding broadcasts")
	client, conn := createBroadcastClient(serverAddr, serverPort)
	defer conn.Close()

	res, err := client.FindBroadcasts(context.Background(), query)
	if err != nil {
		fmt.Println("FindBroadcastsTest ERROR:", err)
		return res
	}

	fmt.Println("Client received response:", res.Response.Type)

	if res.Response.Type == pb.Response_ERROR {
		fmt.Println("Client received error response:", res.Response.ErrorMessage)
	} else {
		for i, broadcast := range res.Broadcasts {
			fmt.Println(i, ":", broadcast)
		}
	}

	return res
}

func DeleteBroadcast(serverAddr *string, serverPort *int, broadcast *pb.Broadcast) {
	fmt.Println("Deleting Broadcast:", broadcast.BroadcastId, broadcast.Title)
	client, conn := createBroadcastClient(serverAddr, serverPort)
	defer conn.Close()

	res, err := client.DeleteBroadcast(context.Background(), broadcast)
	if err != nil {
		fmt.Println("DeleteBroadcast ERROR:", err)
		return
	}

	fmt.Println("Client received response:", res.Type)

	if res.Type == pb.Response_ERROR {
		fmt.Println("Client received error response:", res.ErrorMessage)
	}
}

func createBroadcastClient(serverAddr *string, serverPort *int) (pb.BroadcastServicesClient, *grpc.ClientConn) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", *serverAddr, *serverPort), opts...)
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewBroadcastServicesClient(conn)

	return client, conn
}
