package client

import (
	"context"
	"fmt"
	"io"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	pb "capstone.operations_ecosystem/backend/proto"
	telec "capstone.operations_ecosystem/backend/telegram_client"
)

func TestBroadcastClient(serverAddr *string, serverPort *int) {
	broadcast := CreateFakeBroadcast(1, true)
	pk := InsertBroadcast(serverAddr, serverPort, broadcast)
	broadcast.BroadcastId = pk
	InsertBroadcastAIFSID(serverAddr, serverPort)

	ConsolidatedFindBroadcastTest(serverAddr, serverPort)
	ConsolidatedUpdateBroadcastTest(serverAddr, serverPort, broadcast)
	DeleteBroadcast(serverAddr, serverPort, &pb.Broadcast{BroadcastId: 5})
}

func InsertBroadcast(serverAddr *string, serverPort *int, broadcast *pb.Broadcast) int64 {
	// Ensure that there are users first, if there are users already existing,
	// the returned users will be different, but its ok.
	// for i := 0; i < len(broadcast.Recipients); i++ {
	// 	for j := 0; j < len(broadcast.Recipients[i].Recipient); j++ {
	// 		InsertUser(serverAddr, serverPort, broadcast.Recipients[i].Recipient[j].Recipient)
	// 	}
	// }

	fmt.Println("Inserting Broadcast:", broadcast.BroadcastId)
	teleClient := telec.TelegramClient{}
	client, conn := teleClient.CreateBroadcastClient(serverAddr, serverPort)
	defer conn.Close()

	res, err := client.AddBroadcast(context.Background(), broadcast)
	if err != nil {
		fmt.Println("InsertBroadcast ERROR:", err)
		return -1
	}

	fmt.Println("Client received response:", res.Type)

	if res.Type == pb.Response_ERROR {
		fmt.Println("Client received error response:", res.ErrorMessage)
	}

	return res.PrimaryKey
}

func InsertBroadcastAIFSID(serverAddr *string, serverPort *int) int64 {
	broadcast := CreateFakeBroadcast(1, false)

	fmt.Println("Inserting Broadcast through AIFS id:", broadcast.BroadcastId)
	teleClient := telec.TelegramClient{}
	client, conn := teleClient.CreateBroadcastClient(serverAddr, serverPort)
	defer conn.Close()

	res, err := client.AddBroadcast(context.Background(), broadcast)
	if err != nil {
		fmt.Println("InsertBroadcast ERROR:", err)
		return -1
	}

	fmt.Println("Client received response:", res.Type)

	if res.Type == pb.Response_ERROR {
		fmt.Println("Client received error response:", res.ErrorMessage)
	}

	fmt.Println("Primary key of broadcast with AIFS ID:", res.PrimaryKey)
	return res.PrimaryKey
}

func ConsolidatedFindBroadcastTest(serverAddr *string, serverPort *int) {
	FindBroadcastsNoFilter(serverAddr, serverPort)
	// FindBroadcastsIdFilter(serverAddr, serverPort)
	// FindBroadcastsTypeFilter(serverAddr, serverPort)
	// FindBroadcastsContentFilter(serverAddr, serverPort)
	// FindBroadcastsCreateDateFilter(serverAddr, serverPort)
	// FindBroadcastsDeadlineFilter(serverAddr, serverPort)
	// FindBroadcastsCreatorIdFilter(serverAddr, serverPort)
	// FindBroadcastsRecipientIdFilter(serverAddr, serverPort)
	// FindBroadcastsNumRecFilter(serverAddr, serverPort)
	// FindBroadcastsUrgencyTypeFilter(serverAddr, serverPort)
	// FindBroadcastsAIFSIDFilter(serverAddr, serverPort)
	// FindBroadcastsMultipleFilters(serverAddr, serverPort)
}

func FindBroadcastsNoFilter(serverAddr *string, serverPort *int) {
	fmt.Println("Finding broadcasts without filter")
	FindBroadcastsTest(serverAddr, serverPort, &pb.BroadcastQuery{Limit: 4, Skip: 8})
}

func FindBroadcastsIdFilter(serverAddr *string, serverPort *int) {
	fmt.Println("Finding broadcasts id 15 filter")
	com := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: "15"}
	filter := &pb.BroadcastFilter{Comparisons: com, Field: pb.BroadcastFilter_BROADCAST_ID}

	query := &pb.BroadcastQuery{Limit: 4, Filters: []*pb.BroadcastFilter{filter}}

	FindBroadcastsTest(serverAddr, serverPort, query)
}

func FindBroadcastsTypeFilter(serverAddr *string, serverPort *int) {
	fmt.Println("Finding broadcasts type filter")

	com := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: pb.Broadcast_ASSIGNMENT.String()}
	filter := &pb.BroadcastFilter{Comparisons: com, Field: pb.BroadcastFilter_TYPE}

	query := &pb.BroadcastQuery{Limit: 4, Filters: []*pb.BroadcastFilter{filter}}

	FindBroadcastsTest(serverAddr, serverPort, query)
}

func FindBroadcastsContentFilter(serverAddr *string, serverPort *int) {
	fmt.Println("Finding broadcasts contents filter")

	com := &pb.Filter{Comparison: pb.Filter_CONTAINS, Value: "con"}
	filter := &pb.BroadcastFilter{Comparisons: com, Field: pb.BroadcastFilter_CONTENT}

	query := &pb.BroadcastQuery{Limit: 4, Filters: []*pb.BroadcastFilter{filter}}

	FindBroadcastsTest(serverAddr, serverPort, query)
}

func FindBroadcastsCreateDateFilter(serverAddr *string, serverPort *int) {
	creationDateMax := time.Now().AddDate(0, 0, -1)

	fmt.Println("Finding broadcasts creation date filter, max date:", creationDateMax)
	com := &pb.Filter{Comparison: pb.Filter_LESSER_EQ, Value: creationDateMax.Format("2006-01-02 15:04:05")}
	filter := &pb.BroadcastFilter{Comparisons: com, Field: pb.BroadcastFilter_CREATION_DATE}

	query := &pb.BroadcastQuery{Limit: 4, Filters: []*pb.BroadcastFilter{filter}}

	FindBroadcastsTest(serverAddr, serverPort, query)
}

func FindBroadcastsDeadlineFilter(serverAddr *string, serverPort *int) {
	deadlineMax := time.Now().AddDate(0, 0, 30)

	fmt.Println("Finding broadcasts deadline date filter, max date:", deadlineMax)
	com := &pb.Filter{Comparison: pb.Filter_LESSER_EQ, Value: deadlineMax.Format("2006-01-02 15:04:05")}
	filter := &pb.BroadcastFilter{Comparisons: com, Field: pb.BroadcastFilter_CREATION_DATE}

	query := &pb.BroadcastQuery{Limit: 4, Filters: []*pb.BroadcastFilter{filter}}

	FindBroadcastsTest(serverAddr, serverPort, query)
}

func FindBroadcastsCreatorIdFilter(serverAddr *string, serverPort *int) {
	fmt.Println("Finding broadcasts creator id filter")
	com := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: "1"}
	filter := &pb.BroadcastFilter{Comparisons: com, Field: pb.BroadcastFilter_CREATOR_ID}

	query := &pb.BroadcastQuery{Limit: 4, Filters: []*pb.BroadcastFilter{filter}}

	FindBroadcastsTest(serverAddr, serverPort, query)
}

func FindBroadcastsRecipientIdFilter(serverAddr *string, serverPort *int) {
	fmt.Println("Finding broadcasts recipient id filter")
	com := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: "2"}
	filter := &pb.BroadcastFilter{Comparisons: com, Field: pb.BroadcastFilter_RECEIPEIENT_ID}

	query := &pb.BroadcastQuery{Limit: 4, Filters: []*pb.BroadcastFilter{filter}}

	FindBroadcastsTest(serverAddr, serverPort, query)
}

func FindBroadcastsNumRecFilter(serverAddr *string, serverPort *int) {
	fmt.Println("Finding broadcasts number of recipients filter")
	com := &pb.Filter{Comparison: pb.Filter_GREATER_EQ, Value: "2"}
	filter := &pb.BroadcastFilter{Comparisons: com, Field: pb.BroadcastFilter_NUM_RECEIPIENTS}

	query := &pb.BroadcastQuery{Limit: 4, Filters: []*pb.BroadcastFilter{filter}}

	FindBroadcastsTest(serverAddr, serverPort, query)
}

func FindBroadcastsUrgencyTypeFilter(serverAddr *string, serverPort *int) {
	fmt.Println("Finding broadcasts medium urgency filter")
	com := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: "2"}
	filter := &pb.BroadcastFilter{Comparisons: com, Field: pb.BroadcastFilter_URGENCY}

	query := &pb.BroadcastQuery{Limit: 4, Filters: []*pb.BroadcastFilter{filter}}

	FindBroadcastsTest(serverAddr, serverPort, query)
}

func FindBroadcastsAIFSIDFilter(serverAddr *string, serverPort *int) {
	fmt.Println("Finding broadcasts AIFS ID 2 filter")
	com := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: "2"}
	filter := &pb.BroadcastFilter{Comparisons: com, Field: pb.BroadcastFilter_AIFS_ID}

	query := &pb.BroadcastQuery{Limit: 4, Filters: []*pb.BroadcastFilter{filter}}

	FindBroadcastsTest(serverAddr, serverPort, query)
}

func FindBroadcastsMultipleFilters(serverAddr *string, serverPort *int) {
	fmt.Println("Finding broadcasts mutiple filter")

	broadcastFilters := make([]*pb.BroadcastFilter, 0)

	// type
	typeCom := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: pb.Broadcast_ANNOUNCEMENT.String()}
	typeFilter := &pb.BroadcastFilter{Comparisons: typeCom, Field: pb.BroadcastFilter_TYPE}
	broadcastFilters = append(broadcastFilters, typeFilter)

	// creation date
	creationDateMax := time.Now().AddDate(0, 0, -1)
	cdCom := &pb.Filter{Comparison: pb.Filter_LESSER_EQ, Value: creationDateMax.Format("2006-01-02 15:04:05")}
	cdFilter := &pb.BroadcastFilter{Comparisons: cdCom, Field: pb.BroadcastFilter_CREATION_DATE}
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
	FindBroadcastsTest(serverAddr, serverPort, query)
}

func FindBroadcastsTest(serverAddr *string, serverPort *int, query *pb.BroadcastQuery) {
	fmt.Println("Finding broadcasts")
	teleClient := telec.TelegramClient{}
	client, conn := teleClient.CreateBroadcastClient(serverAddr, serverPort)
	defer conn.Close()

	stream, err := client.FindBroadcasts(context.Background(), query)
	if err != nil {
		fmt.Println("FindBroadcastsTest ERROR:", err)
		return
	}

	count := 0
	for {
		broadcastRes, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("FindBroadcastsTest ERROR:", err)
		}

		fmt.Println("Client received response:", broadcastRes.Response.Type)
		if broadcastRes.Response.Type == pb.Response_ERROR {
			continue
		}

		fmt.Println(count, ":", broadcastRes.Broadcast)
		fmt.Println(count, ":", broadcastRes.Broadcast.CreationDate.AsTime().String())
		count++
	}
}

func ConsolidatedUpdateBroadcastTest(serverAddr *string, serverPort *int, broadcast *pb.Broadcast) {
	UpdateBroadcastBasicFields(serverAddr, serverPort, broadcast)
	UpdateBroadcastRecipients(serverAddr, serverPort, broadcast)
}

func UpdateBroadcastBasicFields(serverAddr *string, serverPort *int, broadcast *pb.Broadcast) {
	creationTime := &timestamppb.Timestamp{Seconds: time.Now().Unix()}
	updateBroadcast := pb.Broadcast{
		BroadcastId:  broadcast.BroadcastId,
		Content:      "This is an updated broadcast content",
		CreationDate: creationTime,
	}

	UpdateBroadcastTest(serverAddr, serverPort, &updateBroadcast)
}

func UpdateBroadcastRecipients(serverAddr *string, serverPort *int, broadcast *pb.Broadcast) {
	updateBroadcast := pb.Broadcast{
		BroadcastId: broadcast.BroadcastId,
		Recipients:  broadcast.Recipients,
	}

	// replace one of the recipients with someone else
	updateBroadcast.Recipients[0].Recipient[0].Recipient.UserId = 6

	UpdateBroadcastTest(serverAddr, serverPort, &updateBroadcast)
}

func UpdateBroadcastTest(serverAddr *string, serverPort *int, broadcast *pb.Broadcast) {
	fmt.Println("Updating Broadcast...")
	teleClient := telec.TelegramClient{}
	client, conn := teleClient.CreateBroadcastClient(serverAddr, serverPort)
	defer conn.Close()

	res, err := client.UpdateBroadcast(context.Background(), broadcast)
	if err != nil {
		fmt.Println("UpdateBroadcastTest ERROR:", err)
		return
	}

	fmt.Println("Client received response:", res.Type)

	if res.Type == pb.Response_ERROR {
		fmt.Println("Client received error response:", res.ErrorMessage)
	}
}

func DeleteBroadcast(serverAddr *string, serverPort *int, broadcast *pb.Broadcast) {
	fmt.Println("Deleting Broadcast:", broadcast.BroadcastId)
	teleClient := telec.TelegramClient{}
	client, conn := teleClient.CreateBroadcastClient(serverAddr, serverPort)
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
