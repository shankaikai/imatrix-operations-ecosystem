// TODO: Add validation
package server

import (
	"fmt"

	db_pck "capstone.operations_ecosystem/backend/database"

	pb "capstone.operations_ecosystem/backend/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"context"
)

func (s *Server) AddBroadcast(cxt context.Context, broadcast *pb.Broadcast) (*pb.Response, error) {
	res := pb.Response{Type: pb.Response_ACK}

	getDefaultRecipients(broadcast)

	// Add creation datetime
	broadcast.CreationDate = timestamppb.Now()
	// TODO: define deadline
	broadcast.Deadline = timestamppb.Now()

	pk, err := db_pck.InsertBroadcast(
		s.db,
		broadcast,
		s.dbLock,
	)

	if err != nil {
		res.Type = pb.Response_ERROR
		res.ErrorMessage = err.Error()
	}

	res.PrimaryKey = pk

	// Send to Telegram Bot
	go s.sendNewBroadcastToTele(pk)

	return &res, nil
}

func (s *Server) UpdateBroadcast(cxt context.Context, broadcast *pb.Broadcast) (*pb.Response, error) {
	res := pb.Response{Type: pb.Response_ACK}
	numAffected, err := db_pck.UpdateBroadcast(
		s.db,
		broadcast,
		s.dbLock,
	)

	if err != nil {
		res.Type = pb.Response_ERROR
		res.ErrorMessage = err.Error()
	} else {
		fmt.Println(numAffected, "broadcasts were updated.")
	}

	return &res, nil
}

func (s *Server) DeleteBroadcast(cxt context.Context, broadcast *pb.Broadcast) (*pb.Response, error) {
	res := pb.Response{Type: pb.Response_ACK}
	numDel, err := db_pck.DeleteBroadcast(
		s.db,
		broadcast,
	)

	if err != nil {
		res.Type = pb.Response_ERROR
		res.ErrorMessage = err.Error()
	} else {
		fmt.Println(numDel, "broadcasts were deleted.")
	}

	return &res, nil
}

func (s *Server) FindBroadcasts(query *pb.BroadcastQuery, stream pb.BroadcastServices_FindBroadcastsServer) error {
	fmt.Println("FindBroadcasts Start")
	res := pb.Response{Type: pb.Response_ACK}

	foundBroadcasts, err := db_pck.GetBroadcasts(
		s.db,
		query,
	)

	if err != nil {
		broadcastRes := pb.BroadcastResponse{Response: &res}
		res.Type = pb.Response_ERROR
		res.ErrorMessage = err.Error()
		stream.Send(&broadcastRes)

	} else {
		fmt.Println("FindBroadcasts: Found broadcasts to return")
		// fmt.Println(foundBroadcasts)
		broadcastRes := pb.BroadcastResponse{Response: &res}

		for _, broadcast := range foundBroadcasts {
			broadcastRes.Broadcast = broadcast
			fmt.Println(broadcast)
			if err := stream.Send(&broadcastRes); err != nil {
				return err
			}
		}
	}

	return nil
}
