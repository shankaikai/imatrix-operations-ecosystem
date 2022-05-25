// TODO: Add validation

package server

import (
	"fmt"

	db_pck "capstone.operations_ecosystem/backend/database"
	pb "capstone.operations_ecosystem/backend/proto"

	"context"
)

// TODO: Add field verification
func (s *Server) AddUser(cxt context.Context, user *pb.User) (*pb.Response, error) {
	res := pb.Response{Type: pb.Response_ACK}
	pk, err := db_pck.UserInsert(
		s.db,
		user,
		s.dbLock,
	)

	if err != nil {
		res.Type = pb.Response_ERROR
		res.ErrorMessage = err.Error()
	}

	res.PrimaryKey = int64(pk)

	return &res, nil
}

// TODO: Add field verification
func (s *Server) UpdateUser(cxt context.Context, user *pb.User) (*pb.Response, error) {
	res := pb.Response{Type: pb.Response_ACK}
	pk, err := db_pck.UpdateUser(
		s.db,
		user,
	)

	if err != nil {
		res.Type = pb.Response_ERROR
		res.ErrorMessage = err.Error()
	}

	res.PrimaryKey = int64(pk)

	return &res, nil
}

// TODO: Add field verification
func (s *Server) DeleteUser(cxt context.Context, user *pb.User) (*pb.Response, error) {
	res := pb.Response{Type: pb.Response_ACK}
	numDel, err := db_pck.DeleteUser(
		s.db,
		user,
	)

	if err != nil {
		res.Type = pb.Response_ERROR
		res.ErrorMessage = err.Error()
	}

	fmt.Println("Deleted", numDel, "users")

	return &res, nil
}

// TODO: Add field verification
func (s *Server) FindUsers(query *pb.UserQuery, stream pb.AdminServices_FindUsersServer) error {
	res := pb.Response{Type: pb.Response_ACK}
	foundUsers, err := db_pck.GetUsers(
		s.db,
		query,
	)

	if err != nil {
		userRes := pb.UsersResponse{Response: &res}
		res.Type = pb.Response_ERROR
		res.ErrorMessage = err.Error()
		stream.Send(&userRes)

	} else {
		userRes := pb.UsersResponse{Response: &res}
		for _, user := range foundUsers {
			userRes.User = user
			if err := stream.Send(&userRes); err != nil {
				return err
			}
		}
	}

	return nil
}
