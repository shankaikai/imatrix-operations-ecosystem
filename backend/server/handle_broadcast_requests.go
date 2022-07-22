package server

import (
	"fmt"

	db_pck "capstone.operations_ecosystem/backend/database"
	"github.com/getsentry/sentry-go"

	pb "capstone.operations_ecosystem/backend/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"context"
)

// gRPC defined endpoint. Insert a broadcast into the DB.
func (s *Server) AddBroadcast(cxt context.Context, broadcast *pb.Broadcast) (*pb.Response, error) {
	defer sentry.Recover()

	res := pb.Response{Type: pb.Response_ACK}

	s.getDefaultBroadcastFields(broadcast)
	// Add creation datetime
	broadcast.CreationDate = timestamppb.Now()
	// Set a deadline, currently not in use
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
	go s.sendNewBroadcastsOut(pk)

	return &res, nil
}

// gRPC defined endpoint. Update a broadcast in the DB.
func (s *Server) UpdateBroadcast(cxt context.Context, broadcast *pb.Broadcast) (*pb.Response, error) {
	defer sentry.Recover()

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

// gRPC defined endpoint. Delete a broadcast in the DB.
func (s *Server) DeleteBroadcast(cxt context.Context, broadcast *pb.Broadcast) (*pb.Response, error) {
	defer sentry.Recover()

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

// gRPC defined endpoint. Find broadcasts in the DB. The broadcasts are filtered out based the query.
func (s *Server) FindBroadcasts(query *pb.BroadcastQuery, stream pb.BroadcastServices_FindBroadcastsServer) error {
	defer sentry.Recover()

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

// gRPC defined endpoint. Update a broadcast recipient in the DB.
func (s *Server) UpdateBroadcastRecipient(cxt context.Context, broadcastRecipient *pb.BroadcastRecipient) (*pb.Response, error) {
	defer sentry.Recover()

	fmt.Println("UpdateBroadcastRecipient", broadcastRecipient)
	res := pb.Response{Type: pb.Response_ACK}
	numAffected, err := db_pck.UpdateBroadcastRecipients(
		s.db,
		broadcastRecipient,
	)

	if err != nil {
		res.Type = pb.Response_ERROR
		res.ErrorMessage = err.Error()
	} else {
		fmt.Println(numAffected, "broadcast recipients were updated.")
	}

	// Notify aifs led lights to turn back to default
	if numAffected > 0 && (broadcastRecipient.Acknowledged || broadcastRecipient.Rejected) {
		s.notifyAIFSofBroadcastAck(broadcastRecipient)
	}

	return &res, nil
}
