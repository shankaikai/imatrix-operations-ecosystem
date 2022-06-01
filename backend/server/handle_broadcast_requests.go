// TODO: Add validation
package server

import (
	"fmt"

	db_pck "capstone.operations_ecosystem/backend/database"
	pb "capstone.operations_ecosystem/backend/proto"

	"context"
)

func (s *Server) AddBroadcast(cxt context.Context, broadcast *pb.Broadcast) (*pb.Response, error) {
	res := pb.Response{Type: pb.Response_ACK}

	getDefaultRecipients(broadcast)

	pk, err := db_pck.InsertBroadcast(
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
		fmt.Println(foundBroadcasts)
		broadcastRes := pb.BroadcastResponse{Response: &res}

		for _, broadcast := range foundBroadcasts {
			broadcastRes.Broadcast = broadcast
			if err := stream.Send(&broadcastRes); err != nil {
				return err
			}
		}
	}

	return nil
}

// If the broadcast recipient is an AIFS,
// change the recipients to be actual users
// Modified the broadcast in place
func getDefaultRecipients(broadcast *pb.Broadcast) {
	for _, rec := range broadcast.Recipients {
		newRecipients := make([]*pb.BroadcastRecipient, 0)

		users := getFakeAIFSDuty(rec.AifsId)
		for _, user := range users {
			newRecipients = append(newRecipients, &pb.BroadcastRecipient{
				Recipient: user,
				AifsId:    rec.AifsId,
			})
		}

		rec.Recipient = newRecipients
	}
}

// TODO get actual roster for AIFS Groups
func getFakeAIFSDuty(aifsId int64) []*pb.User {
	users := make([]*pb.User, 0)

	for i := 1; i < 3; i++ {
		users = append(users, &pb.User{
			UserId: int64(i),
		})
	}

	return users
}
