package client

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "capstone.operations_ecosystem/backend/proto"
)

func TestClient(serverAddr *string, serverPort *int) {
	user := pb.User{
		UserId:          3,
		UserType:        pb.User_ISPECIALIST,
		Name:            "test name",
		Email:           "email",
		PhoneNumber:     "1232",
		TelegramHandle:  "sfds",
		UserSecurityImg: "dsfds",
		IsPartTimer:     false,
	}

	InsertUser(serverAddr, serverPort, &user)
}

func InsertUser(serverAddr *string, serverPort *int, user *pb.User) {
	fmt.Println("Inserting product:", user.Name)
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", *serverAddr, *serverPort), opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewAdminServicesClient(conn)
	res, err := client.AddUser(context.Background(), user)
	if err != nil {
		fmt.Println("InsertProduct ERROR:", err)
		return
	}

	fmt.Println("Client received response:", res.Type)

	if res.Type == pb.Response_ERROR {
		fmt.Println("Client received error response:", res.ErrorMessage)
	}
}
