package fake_server

import (
	"fmt"
	"strconv"

	pb "capstone.operations_ecosystem/backend/proto"

	"context"
)

func (s *Server) AddUser(cxt context.Context, user *pb.User) (*pb.Response, error) {
	fmt.Println("AddUser")
	res := pb.Response{Type: pb.Response_ACK, PrimaryKey: 1}
	return &res, nil
}

func (s *Server) UpdateUser(cxt context.Context, user *pb.User) (*pb.Response, error) {
	fmt.Println("UpdateUser")
	res := pb.Response{Type: pb.Response_ACK, PrimaryKey: 1}
	return &res, nil
}

func (s *Server) DeleteUser(cxt context.Context, user *pb.User) (*pb.Response, error) {
	fmt.Println("DeleteUser")
	res := pb.Response{Type: pb.Response_ACK, PrimaryKey: 1}
	return &res, nil
}

func (s *Server) FindUsers(query *pb.UserQuery, stream pb.AdminServices_FindUsersServer) error {
	fmt.Println("FindUsers")

	res := pb.Response{Type: pb.Response_ACK}
	userRes := pb.UsersResponse{Response: &res}

	for i := 0; i < 3; i++ {
		user := &pb.User{
			UserId:          3,
			UserType:        pb.User_ISPECIALIST,
			Name:            "test name" + strconv.Itoa(i),
			Email:           "email",
			PhoneNumber:     "1232",
			TelegramHandle:  "sfds",
			UserSecurityImg: "dsfds",
			IsPartTimer:     false,
		}
		userRes.User = user

		if err := stream.Send(&userRes); err != nil {
			return err
		}
	}

	return nil
}

// Client
func (s *Server) AddClient(cxt context.Context, client *pb.Client) (*pb.Response, error) {
	fmt.Println("AddUser")
	res := pb.Response{Type: pb.Response_ACK, PrimaryKey: 1}
	return &res, nil
}

func (s *Server) UpdateClient(cxt context.Context, client *pb.Client) (*pb.Response, error) {
	fmt.Println("UpdateUser")
	res := pb.Response{Type: pb.Response_ACK, PrimaryKey: 1}
	return &res, nil
}

func (s *Server) DeleteClient(cxt context.Context, client *pb.Client) (*pb.Response, error) {
	fmt.Println("DeleteUser")
	res := pb.Response{Type: pb.Response_ACK, PrimaryKey: 1}
	return &res, nil
}

func (s *Server) FindClients(query *pb.ClientQuery, stream pb.AdminServices_FindClientsServer) error {
	fmt.Println("FindUsers")

	res := pb.Response{Type: pb.Response_ACK}
	clientRes := pb.ClientResponse{Response: &res}

	for i := 0; i < 3; i++ {
		client := &pb.Client{
			ClientId:             int64(i),
			Name:                 "test name" + strconv.Itoa(i),
			Email:                "email" + strconv.Itoa(i),
			Address:              "address" + strconv.Itoa(i),
			PhoneNumber:          "1232",
			NumberOfGuardsNeeded: 3,
		}
		clientRes.Client = client

		if err := stream.Send(&clientRes); err != nil {
			return err
		}
	}

	return nil
}
