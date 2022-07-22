package client

import (
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "capstone.operations_ecosystem/backend/proto"
)

func TestCameraIotClientUser(serverAddr *string, serverPort *int) {
	SetGateStateTest(serverAddr, serverPort)
	GetIotStateTest(serverAddr, serverPort)
}

func SetGateStateTest(serverAddr *string, serverPort *int) {
	gateState := &pb.GateState{
		Id:    1,
		State: pb.GateState_CLOSED,
	}

	fmt.Println("Updating gate:", gateState.Id, gateState.State)

	client, conn := createIotClient(serverAddr, serverPort)
	defer conn.Close()

	res, err := client.SetGateState(context.Background(), gateState)
	if err != nil {
		fmt.Println("SetGateStateTest ERROR:", err)
		return
	}

	fmt.Println("Client received response:", res.Type)

	if res.Type == pb.Response_ERROR {
		fmt.Println("Client received error response:", res.ErrorMessage)
	}
}

func GetIotStateTest(serverAddr *string, serverPort *int) {
	fmt.Println("GetIotStateTest")
	client, conn := createIotClient(serverAddr, serverPort)
	defer conn.Close()

	stream, err := client.GetIotState(context.Background(), &emptypb.Empty{})
	if err != nil {
		fmt.Println("GetIotStateTest ERROR:", err)
		return
	}

	count := 0
	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("GetIotStateTest ERROR:", err)
		}

		fmt.Println("Client received response:", res.Response.Type)
		if res.Response.Type == pb.Response_ERROR {
			fmt.Println("Client received error response:", res.Response.ErrorMessage)
			continue
		}

		fmt.Println(count, ":", res.CameraIot)
		count++
	}
}

func createIotClient(serverAddr *string, serverPort *int) (pb.CameraIotServicesClient, *grpc.ClientConn) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", *serverAddr, *serverPort), opts...)
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewCameraIotServicesClient(conn)

	return client, conn
}
