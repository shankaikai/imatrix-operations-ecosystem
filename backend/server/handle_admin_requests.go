// TODO: Add validation

package server

import (
	db_pck "capstone.operations_ecosystem/backend/database"
	pb "capstone.operations_ecosystem/backend/proto"

	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}

// TODO: Add field verification
func (s *Server) DeleteUser(cxt context.Context, user *pb.User) (*pb.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}

// TODO: Add field verification
func (s *Server) FindUsers(cxt context.Context, query *pb.UserQuery) (*pb.BulkUsers, error) {
	res := pb.Response{Type: pb.Response_ACK}
	Users := pb.BulkUsers{Response: &res}

	foundUsers, err := db_pck.GetUsers(
		s.db,
		query,
	)
	if err != nil {
		res.Type = pb.Response_ERROR
		res.ErrorMessage = err.Error()
	} else {
		Users.Users = foundUsers
	}

	return &Users, nil
}
