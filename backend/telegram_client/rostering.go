package telegram_client

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "capstone.operations_ecosystem/backend/proto"
)

// Send a gRPC client request to the telegram bot gRPC server, letting them know that a new roster had been created.
func InsertRoster(serverAddr *string, serverPort *int, rosters []*pb.Roster) int64 {
	fmt.Println("insert roster telegram bot")
	grpcClient, conn := createRosterClient(serverAddr, serverPort)
	defer conn.Close()

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

// Creates and returns a gRPC client for rostering services.
// Note, it is the caller's responsibility to close the connection afterwards.
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
