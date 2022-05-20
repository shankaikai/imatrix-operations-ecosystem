package client

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "capstone.operations_ecosystem/backend/proto"
)

func TestBroadcastClient(serverAddr *string, serverPort *int) {
	broadcast := createFakeBroadcast(1)
	InsertBroadcast(serverAddr, serverPort, broadcast)
	existingBroadcasts := FindBroadcastsNoFilter(serverAddr, serverPort, &pb.BroadcastQuery{Limit: 20})
	DeleteBroadcast(serverAddr, serverPort, existingBroadcasts.Broadcasts[0])
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

func FindBroadcastsNoFilter(serverAddr *string, serverPort *int, query *pb.BroadcastQuery) *pb.BulkBroadcasts {
	fmt.Println("Finding broadcasts without filter")
	client, conn := createBroadcastClient(serverAddr, serverPort)
	defer conn.Close()

	res, err := client.FindBroadcasts(context.Background(), query)
	if err != nil {
		fmt.Println("FindBroadcastsNoFilter ERROR:", err)
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
