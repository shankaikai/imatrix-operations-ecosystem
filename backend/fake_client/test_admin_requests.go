package client

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "capstone.operations_ecosystem/backend/proto"
)

func TestAdminClient(serverAddr *string, serverPort *int) {
	user := createFakeUser(1)
	InsertUser(serverAddr, serverPort, user)
	FindUsersNoFilter(serverAddr, serverPort, &pb.UserQuery{Limit: 20})
}

func InsertUser(serverAddr *string, serverPort *int, user *pb.User) {
	fmt.Println("Inserting product:", user.Name)
	client, conn := createAdminClient(serverAddr, serverPort)
	defer conn.Close()

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

func FindUsersNoFilter(serverAddr *string, serverPort *int, query *pb.UserQuery) {
	fmt.Println("Finding users without filter")
	client, conn := createAdminClient(serverAddr, serverPort)
	defer conn.Close()

	res, err := client.FindUsers(context.Background(), query)
	if err != nil {
		fmt.Println("InsertProduct ERROR:", err)
		return
	}

	fmt.Println("Client received response:", res.Response.Type)

	if res.Response.Type == pb.Response_ERROR {
		fmt.Println("Client received error response:", res.Response.ErrorMessage)
	} else {
		for i, user := range res.Users {
			fmt.Println(i, ":", user)
		}
	}
}

func createAdminClient(serverAddr *string, serverPort *int) (pb.AdminServicesClient, *grpc.ClientConn) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", *serverAddr, *serverPort), opts...)
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewAdminServicesClient(conn)

	return client, conn
}

func createFakeUser(id int) *pb.User {
	return &pb.User{
		UserId:          int64(id),
		UserType:        pb.User_ISPECIALIST,
		Name:            "test name" + strconv.Itoa(id),
		Email:           "email" + strconv.Itoa(id),
		PhoneNumber:     "1232",
		TelegramHandle:  "sfds",
		UserSecurityImg: "dsfds",
		IsPartTimer:     false,
	}
}
