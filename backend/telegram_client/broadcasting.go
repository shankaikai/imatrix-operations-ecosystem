package telegram_client

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "capstone.operations_ecosystem/backend/proto"
)

func (t *TelegramClient) InsertBroadcast(serverAddr *string, serverPort *int, broadcast *pb.Broadcast) (int64, error) {
	fmt.Println("Sending Broadcast to Telegram Server:", broadcast.BroadcastId)
	client, conn := t.CreateBroadcastClient(serverAddr, serverPort)
	defer conn.Close()

	res, err := client.AddBroadcast(context.Background(), broadcast)
	if err != nil {
		fmt.Println("InsertBroadcast ERROR:", err)
		return -1, err
	}

	fmt.Println("Client received response:", res.Type)

	if res.Type == pb.Response_ERROR {
		fmt.Println("Client received error response:", res.ErrorMessage)
	}

	return res.PrimaryKey, err
}

func (t *TelegramClient) CreateBroadcastClient(serverAddr *string, serverPort *int) (pb.BroadcastServicesClient, *grpc.ClientConn) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", *serverAddr, *serverPort), opts...)
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewBroadcastServicesClient(conn)

	return client, conn
}
