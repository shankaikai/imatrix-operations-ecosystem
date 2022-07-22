package telegram_client

import (
	pb "capstone.operations_ecosystem/backend/proto"

	"google.golang.org/grpc"
)

type TelegramClientInterface interface {
	CreateBroadcastClient(serverAddr *string, serverPort *int) (pb.BroadcastServicesClient, *grpc.ClientConn)
	InsertBroadcast(serverAddr *string, serverPort *int, broadcast *pb.Broadcast) (int64, error)
}

type TelegramClient struct {
	TelegramClientInterface
}
