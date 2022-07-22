package client

import (
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "capstone.operations_ecosystem/backend/proto"
)

func TestAdminClientUser(serverAddr *string, serverPort *int) {
	// user := createFakeUser(1)
	// pk := InsertUser(serverAddr, serverPort, user)
	// user.UserId = pk
	// ConsolidatedFindUserTest(serverAddr, serverPort)
	// UpdateUserTest(serverAddr, serverPort, user)
	// DeleteUserTest(serverAddr, serverPort, &pb.User{UserId: -1})
	// GetNonceTest(serverAddr, serverPort, user)
	GetUserRandomStringTest(serverAddr, serverPort)
	AuthenticateUserTest(serverAddr, serverPort)
}

func InsertUser(serverAddr *string, serverPort *int, user *pb.User) int64 {
	fmt.Println("Inserting user:", user.Name)
	client, conn := createAdminClient(serverAddr, serverPort)
	defer conn.Close()

	res, err := client.AddUser(context.Background(), user)
	if err != nil {
		fmt.Println("InsertUser ERROR:", err)
		return -1
	}

	fmt.Println("Client received response:", res.Type)

	if res.Type == pb.Response_ERROR {
		fmt.Println("Client received error response:", res.ErrorMessage)
	}

	return res.PrimaryKey
}

func ConsolidatedFindUserTest(serverAddr *string, serverPort *int) {
	FindUsersNoFilter(serverAddr, serverPort)
	FindUserIdFilter(serverAddr, serverPort)
	FindUserTypeFilter(serverAddr, serverPort)
	FindUserNameFilter(serverAddr, serverPort)
	FindUserEmaililter(serverAddr, serverPort)
	FindUserPNUMilter(serverAddr, serverPort)
	FindUserTeleFilter(serverAddr, serverPort)
	FindUserPartTimerFilter(serverAddr, serverPort)
	FindUsersMultipleFilters(serverAddr, serverPort)
}

func FindUsersNoFilter(serverAddr *string, serverPort *int) {
	fmt.Println("Finding users without filter")
	FindUsersTest(serverAddr, serverPort, &pb.UserQuery{Limit: 5})
}

func FindUserIdFilter(serverAddr *string, serverPort *int) {
	fmt.Println("Finding user id filter")
	com := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: "1"}
	filter := &pb.UserFilter{Comparisons: com, Field: pb.UserFilter_USER_ID}

	query := &pb.UserQuery{Limit: 4, Filters: []*pb.UserFilter{filter}}

	FindUsersTest(serverAddr, serverPort, query)
}

func FindUserTypeFilter(serverAddr *string, serverPort *int) {
	fmt.Println("Finding user type filter")
	com := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: pb.User_ISPECIALIST.String()}
	filter := &pb.UserFilter{Comparisons: com, Field: pb.UserFilter_TYPE}

	query := &pb.UserQuery{Limit: 4, Filters: []*pb.UserFilter{filter}}

	FindUsersTest(serverAddr, serverPort, query)
}

func FindUserNameFilter(serverAddr *string, serverPort *int) {
	fmt.Println("Finding user name filter")
	com := &pb.Filter{Comparison: pb.Filter_CONTAINS, Value: "name2"}
	filter := &pb.UserFilter{Comparisons: com, Field: pb.UserFilter_NAME}

	query := &pb.UserQuery{Limit: 4, Filters: []*pb.UserFilter{filter}}

	FindUsersTest(serverAddr, serverPort, query)
}

func FindUserEmaililter(serverAddr *string, serverPort *int) {
	fmt.Println("Finding user email filter")
	com := &pb.Filter{Comparison: pb.Filter_CONTAINS, Value: "email2"}
	filter := &pb.UserFilter{Comparisons: com, Field: pb.UserFilter_EMAIL}

	query := &pb.UserQuery{Limit: 4, Filters: []*pb.UserFilter{filter}}

	FindUsersTest(serverAddr, serverPort, query)
}

func FindUserPNUMilter(serverAddr *string, serverPort *int) {
	fmt.Println("Finding user phone num filter")
	com := &pb.Filter{Comparison: pb.Filter_CONTAINS, Value: "12"}
	filter := &pb.UserFilter{Comparisons: com, Field: pb.UserFilter_PHONE_NUMBER}

	query := &pb.UserQuery{Limit: 4, Filters: []*pb.UserFilter{filter}}

	FindUsersTest(serverAddr, serverPort, query)
}

func FindUserTeleFilter(serverAddr *string, serverPort *int) {
	fmt.Println("Finding user tele handle filter")
	com := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: "sfds"}
	filter := &pb.UserFilter{Comparisons: com, Field: pb.UserFilter_TELEGRAM_HANDLE}

	query := &pb.UserQuery{Limit: 4, Filters: []*pb.UserFilter{filter}}

	FindUsersTest(serverAddr, serverPort, query)
}

func FindUserPartTimerFilter(serverAddr *string, serverPort *int) {
	fmt.Println("Finding user part timer filter")
	com := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: "1"}
	filter := &pb.UserFilter{Comparisons: com, Field: pb.UserFilter_IS_PART_TIMER}

	query := &pb.UserQuery{Limit: 4, Filters: []*pb.UserFilter{filter}}

	FindUsersTest(serverAddr, serverPort, query)
}

func FindUsersMultipleFilters(serverAddr *string, serverPort *int) {
	fmt.Println("Finding user mutiple filter")

	userFilters := make([]*pb.UserFilter, 0)

	// type
	typeCom := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: pb.User_ISPECIALIST.String()}
	typeFilter := &pb.UserFilter{Comparisons: typeCom, Field: pb.UserFilter_TYPE}
	userFilters = append(userFilters, typeFilter)

	// tele handle
	teleCom := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: "sfds"}
	teleFilter := &pb.UserFilter{Comparisons: teleCom, Field: pb.UserFilter_TELEGRAM_HANDLE}
	userFilters = append(userFilters, teleFilter)

	// part timer
	ptCom := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: "1"}
	ptFilter := &pb.UserFilter{Comparisons: ptCom, Field: pb.UserFilter_IS_PART_TIMER}
	userFilters = append(userFilters, ptFilter)

	query := &pb.UserQuery{Limit: 4, Filters: userFilters}
	FindUsersTest(serverAddr, serverPort, query)
}

func FindUsersTest(serverAddr *string, serverPort *int, query *pb.UserQuery) {
	fmt.Println("Finding users...")
	client, conn := createAdminClient(serverAddr, serverPort)
	defer conn.Close()

	stream, err := client.FindUsers(context.Background(), query)

	if err != nil {
		fmt.Println("FindUsersNoFilter ERROR:", err)
		return
	}

	count := 0
	for {
		userRes, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("FindUsersNoFilter ERROR:", err)
		}

		fmt.Println("Client received response:", userRes.Response.Type)
		if userRes.Response.Type == pb.Response_ERROR {
			continue
		}

		fmt.Println(count, ":", userRes.User)
		count++
	}
}

func UpdateUserTest(serverAddr *string, serverPort *int, user *pb.User) {
	updatedUser := &pb.User{
		UserId:         user.UserId,
		UserType:       pb.User_CONTROLLER,
		TelegramHandle: "smthnew",
		IsPartTimer:    true,
	}

	fmt.Println("Updating user:", user.UserId)

	client, conn := createAdminClient(serverAddr, serverPort)
	defer conn.Close()

	res, err := client.UpdateUser(context.Background(), updatedUser)
	if err != nil {
		fmt.Println("UpdateUserTest ERROR:", err)
		return
	}

	fmt.Println("Client received response:", res.Type)

	if res.Type == pb.Response_ERROR {
		fmt.Println("Client received error response:", res.ErrorMessage)
	}
}

func DeleteUserTest(serverAddr *string, serverPort *int, user *pb.User) {
	fmt.Println("Deleting user:", user.Name)
	client, conn := createAdminClient(serverAddr, serverPort)
	defer conn.Close()

	res, err := client.DeleteUser(context.Background(), user)
	if err != nil {
		fmt.Println("DeleteUserTest ERROR:", err)
		return
	}

	fmt.Println("Client received response:", res.Type)

	if res.Type == pb.Response_ERROR {
		fmt.Println("Client received error response:", res.ErrorMessage)
	}
}

func GetNonceTest(serverAddr *string, serverPort *int, user *pb.User) {
	fmt.Println("GetNonceTest user:", user.UserId)
	client, conn := createAdminClient(serverAddr, serverPort)
	defer conn.Close()

	res, err := client.GetWANonce(context.Background(), user)
	if err != nil {
		fmt.Println("GetNonceTest ERROR:", err)
		return
	}

	fmt.Println("Client received response:", res.Response.Type)

	if res.Response.Type == pb.Response_ERROR {
		fmt.Println("Client received error response:", res.Response.ErrorMessage)
	} else {
		fmt.Println("Nonce:", string(res.Nonce), "")
	}
}

func GetUserRandomStringTest(serverAddr *string, serverPort *int) {
	fmt.Println("GetUserRandomStringTest")
	user := &pb.User{
		Email: "testemail@gmail.com",
	}

	client, conn := createAdminClient(serverAddr, serverPort)
	defer conn.Close()

	res, err := client.GetSecurityString(context.Background(), user)
	if err != nil {
		fmt.Println("GetUserRandomStringTest ERROR:", err)
		return
	}

	fmt.Println("Client received response:", res.Response.Type)

	if res.Response.Type == pb.Response_ERROR {
		fmt.Println("Client received error response:", res.Response.ErrorMessage)
	} else {
		fmt.Println("Ran string:", string(res.SecurityString))
	}
}

func AuthenticateUserTest(serverAddr *string, serverPort *int) {
	fmt.Println("GetUserRandomStringTest")

	loginRequest := &pb.LoginRequest{
		UserEmail:      "testemail@gmail.com",
		HashedPassword: "fsdkvvhcxnddsofdhc",
	}

	client, conn := createAdminClient(serverAddr, serverPort)
	defer conn.Close()

	res, err := client.AuthenticateUser(context.Background(), loginRequest)
	if err != nil {
		fmt.Println("GetUserRandomStringTest ERROR:", err)
		return
	}

	fmt.Println("Client received response:", res.Response.Type)

	if res.Response.Type == pb.Response_ERROR {
		fmt.Println("Client received error response:", res.Response.ErrorMessage)
	} else {
		fmt.Println("Token:", res.UserToken)
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
