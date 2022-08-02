package client

import (
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "capstone.operations_ecosystem/backend/proto"
	"capstone.operations_ecosystem/backend/server"
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

	client, conn, cxt := createIotClient(serverAddr, serverPort)
	defer conn.Close()

	res, err := client.SetGateState(cxt, gateState)
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
	client, conn, cxt := createIotClient(serverAddr, serverPort)
	defer conn.Close()

	stream, err := client.GetIotState(cxt, &emptypb.Empty{})
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

func createIotClient(serverAddr *string, serverPort *int) (pb.CameraIotServicesClient, *grpc.ClientConn, context.Context) {
	var opts []grpc.DialOption
	ctx := context.Background()

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", *serverAddr, *serverPort), opts...)
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewCameraIotServicesClient(conn)
	// md := metadata.New(map[string]string{server.JWT_TOKEN_HEADER: "0lWtzj9wOQevbserbFutaxEiVHy9Sj5ZVaaOJxIWTmaVI01beYObd3Curq4xsBZnseMoil6LaN2lFwN1aWBiXSmOF9sa2Yoj4Wdy89gfsQr3JSl5064Q8NozaRf3dpC6"})
	md := metadata.New(map[string]string{server.JWT_TOKEN_HEADER: "U6VFo936AGzQ47QQkkTURNVWAbeuaoavx2O6y8WamPgSHd9XN4EniK8TfIEbMCM7iOsCcSWb83JZawa8npZlHCMfOPV1G4ymiCz1zhtkz4ZJLtFFJkptuVTdNwv2D0M2"})
	ctx = metadata.NewOutgoingContext(ctx, md)

	return client, conn, ctx
}
