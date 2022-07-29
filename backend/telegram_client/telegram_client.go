package telegram_client

import (
	pb "capstone.operations_ecosystem/backend/proto"

	"google.golang.org/grpc"
)

// Define the functions a telegram client should have.
type TelegramClientInterface interface {
	CreateBroadcastClient(serverAddr *string, serverPort *int) (pb.BroadcastServicesClient, *grpc.ClientConn)
	InsertBroadcast(serverAddr *string, serverPort *int, broadcast *pb.Broadcast) (int64, error)
}

// The main telegram client being used to send information to the Telegram Bot gRPC server.
type TelegramClient struct {
	TelegramClientInterface
}
