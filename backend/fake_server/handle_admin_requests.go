package fake_server

import (
	"fmt"

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

func (s *Server) FindUsers(cxt context.Context, query *pb.UserQuery) (*pb.BulkUsers, error) {
	fmt.Println("FindUsers")

	res := pb.Response{Type: pb.Response_ACK}
	users := pb.BulkUsers{Response: &res}

	foundUsers := make([]*pb.User, 0)

	for i := 0; i < 3; i++ {
		foundUsers = append(foundUsers, &pb.User{
			UserId:          3,
			UserType:        pb.User_ISPECIALIST,
			Name:            "test name",
			Email:           "email",
			PhoneNumber:     "1232",
			TelegramHandle:  "sfds",
			UserSecurityImg: "dsfds",
			IsPartTimer:     false,
		})
	}

	users.Users = foundUsers

	return &users, nil
}
