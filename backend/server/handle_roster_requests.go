// TODO: Add validation
package server

// func (s *Server) AddRoster(cxt context.Context, roster *pb.Roster) (*pb.Response, error) {
// 	res := pb.Response{Type: pb.Response_ACK}

// 	getDefaultRecipients(broadcast)

// 	// Add creation datetime
// 	broadcast.CreationDate = timestamppb.Now()
// 	// TODO: define deadline
// 	broadcast.Deadline = timestamppb.Now()

// 	pk, err := db_pck.InsertBroadcast(
// 		s.db,
// 		broadcast,
// 		s.dbLock,
// 	)

// 	if err != nil {
// 		res.Type = pb.Response_ERROR
// 		res.ErrorMessage = err.Error()
// 	}

// 	res.PrimaryKey = int64(pk)

// 	return &res, nil
// }

// func (s *Server) UpdateRoster(cxt context.Context, broadcast *pb.Broadcast) (*pb.Response, error) {
// 	res := pb.Response{Type: pb.Response_ACK}
// 	numAffected, err := db_pck.UpdateBroadcast(
// 		s.db,
// 		broadcast,
// 		s.dbLock,
// 	)

// 	if err != nil {
// 		res.Type = pb.Response_ERROR
// 		res.ErrorMessage = err.Error()
// 	} else {
// 		fmt.Println(numAffected, "broadcasts were updated.")
// 	}

// 	return &res, nil
// }

// func (s *Server) DeleteRoster(cxt context.Context, broadcast *pb.Broadcast) (*pb.Response, error) {
// 	res := pb.Response{Type: pb.Response_ACK}
// 	numDel, err := db_pck.DeleteBroadcast(
// 		s.db,
// 		broadcast,
// 	)

// 	if err != nil {
// 		res.Type = pb.Response_ERROR
// 		res.ErrorMessage = err.Error()
// 	} else {
// 		fmt.Println(numDel, "broadcasts were deleted.")
// 	}

// 	return &res, nil
// }

// func (s *Server) FindRosters(query *pb.BroadcastQuery, stream pb.BroadcastServices_FindBroadcastsServer) error {
// 	res := pb.Response{Type: pb.Response_ACK}

// 	foundBroadcasts, err := db_pck.GetBroadcasts(
// 		s.db,
// 		query,
// 	)

// 	if err != nil {
// 		broadcastRes := pb.BroadcastResponse{Response: &res}
// 		res.Type = pb.Response_ERROR
// 		res.ErrorMessage = err.Error()
// 		stream.Send(&broadcastRes)

// 	} else {
// 		fmt.Println("FindBroadcasts: Found broadcasts to return")
// 		fmt.Println(foundBroadcasts)
// 		broadcastRes := pb.BroadcastResponse{Response: &res}

// 		for _, broadcast := range foundBroadcasts {
// 			broadcastRes.Broadcast = broadcast
// 			if err := stream.Send(&broadcastRes); err != nil {
// 				return err
// 			}
// 		}
// 	}

// 	return nil
// }

// func GetAvailableUsers() {

// }
