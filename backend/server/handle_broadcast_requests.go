package server

import (
	db_pck "capstone.operations_ecosystem/backend/database"
	pb "capstone.operations_ecosystem/backend/proto"

	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) AddBroadcast(cxt context.Context, broadcast *pb.Broadcast) (*pb.Response, error) {
	res := pb.Response{Type: pb.Response_ACK}
	pk, err := db_pck.BroadcastInsert(
		s.db,
		broadcast,
		s.dbLock,
	)

	if err != nil {
		res.Type = pb.Response_ERROR
		res.ErrorMessage = err.Error()
	}

	res.PrimaryKey = int64(pk)

	return &res, nil
}

func (s *Server) UpdateBroadcast(cxt context.Context, broadcast *pb.Broadcast) (*pb.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBroadcast not implemented")
}

func (s *Server) DeleteBroadcast(cxt context.Context, broadcast *pb.Broadcast) (*pb.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBroadcast not implemented")
}

func (s *Server) FindBroadcasts(cxt context.Context, query *pb.BroadcastQuery) (*pb.BulkBroadcasts, error) {
	res := pb.Response{Type: pb.Response_ACK}
	broadcasts := pb.BulkBroadcasts{Response: &res}

	foundBroadcasts, err := db_pck.GetBroadcasts(
		s.db,
		query,
	)

	if err != nil {
		res.Type = pb.Response_ERROR
		res.ErrorMessage = err.Error()
	} else {
		broadcasts.Broadcasts = foundBroadcasts
	}

	return &broadcasts, nil
}
