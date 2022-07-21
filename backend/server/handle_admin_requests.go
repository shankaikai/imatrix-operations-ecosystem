package server

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"

	db_pck "capstone.operations_ecosystem/backend/database"
	pb "capstone.operations_ecosystem/backend/proto"
	"github.com/getsentry/sentry-go"

	"context"
)

func (s *Server) AddUser(cxt context.Context, user *pb.User) (*pb.Response, error) {
	defer sentry.Recover()

	res := pb.Response{Type: pb.Response_ACK}
	pk, err := db_pck.InsertUser(
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

func (s *Server) UpdateUser(cxt context.Context, user *pb.User) (*pb.Response, error) {
	defer sentry.Recover()

	res := pb.Response{Type: pb.Response_ACK}
	num_affected, err := db_pck.UpdateUser(
		s.db,
		user,
		&pb.UserQuery{},
	)

	if err != nil {
		res.Type = pb.Response_ERROR
		res.ErrorMessage = err.Error()
	}

	res.PrimaryKey = int64(num_affected)
	fmt.Println("Updated", num_affected, "users")

	return &res, nil
}

func (s *Server) DeleteUser(cxt context.Context, user *pb.User) (*pb.Response, error) {
	defer sentry.Recover()

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

func (s *Server) FindUsers(query *pb.UserQuery, stream pb.AdminServices_FindUsersServer) error {
	defer sentry.Recover()

	fmt.Println("FindUsers", query)
	res := pb.Response{Type: pb.Response_ACK}
	foundUsers, err := db_pck.GetUsers(
		s.db,
		query,
		true,
	)

	if err != nil {
		userRes := pb.UsersResponse{Response: &res}
		res.Type = pb.Response_ERROR
		res.ErrorMessage = err.Error()
		stream.Send(&userRes)

	} else {
		userRes := pb.UsersResponse{Response: &res}
		for _, fullUser := range foundUsers {
			userRes.User = fullUser.User
			if err := stream.Send(&userRes); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *Server) AddClient(cxt context.Context, client *pb.Client) (*pb.Response, error) {
	defer sentry.Recover()

	res := pb.Response{Type: pb.Response_ACK}
	pk, err := db_pck.InsertClient(
		s.db,
		client,
		s.dbLock,
	)

	if err != nil {
		res.Type = pb.Response_ERROR
		res.ErrorMessage = err.Error()
	}

	res.PrimaryKey = int64(pk)

	return &res, nil
}

func (s *Server) UpdateClient(cxt context.Context, client *pb.Client) (*pb.Response, error) {
	defer sentry.Recover()

	res := pb.Response{Type: pb.Response_ACK}
	num_affected, err := db_pck.UpdateClients(
		s.db,
		client,
	)

	if err != nil {
		res.Type = pb.Response_ERROR
		res.ErrorMessage = err.Error()
	}

	res.PrimaryKey = int64(num_affected)
	fmt.Println("Updated", num_affected, "clients")
	return &res, nil
}

func (s *Server) DeleteClient(cxt context.Context, client *pb.Client) (*pb.Response, error) {
	defer sentry.Recover()

	res := pb.Response{Type: pb.Response_ACK}
	numDel, err := db_pck.DeleteClient(
		s.db,
		client,
	)

	if err != nil {
		res.Type = pb.Response_ERROR
		res.ErrorMessage = err.Error()
	}

	fmt.Println("Deleted", numDel, "clients")

	return &res, nil
}

func (s *Server) FindClients(query *pb.ClientQuery, stream pb.AdminServices_FindClientsServer) error {
	defer sentry.Recover()

	res := pb.Response{Type: pb.Response_ACK}
	foundClients, err := db_pck.GetClients(
		s.db,
		query,
	)

	if err != nil {
		clientRes := pb.ClientResponse{Response: &res}
		res.Type = pb.Response_ERROR
		res.ErrorMessage = err.Error()
		stream.Send(&clientRes)

	} else {
		clientRes := pb.ClientResponse{Response: &res}
		for _, clients := range foundClients {
			clientRes.Client = clients
			if err := stream.Send(&clientRes); err != nil {
				return err
			}
		}
	}

	return nil
}

// Return a cryptographic-secure nonce
func (s *Server) GetWANonce(cxt context.Context, user *pb.User) (*pb.ResponseNonce, error) {
	fmt.Println("GetWANonce")
	res := &pb.Response{Type: pb.Response_ACK, ErrorMessage: "No error"}

	// Create nonce
	nonce := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println("GetWANonce ERROR", err)
		res.ErrorMessage = err.Error()
		res.Type = pb.Response_ERROR
		return &pb.ResponseNonce{Response: res}, nil
	}

	nonceString := hex.EncodeToString(nonce)
	// Put nonce in DB
	userQuery := &pb.UserQuery{}
	db_pck.AddUserFilter(userQuery, pb.UserFilter_TELEGRAM_USER_ID, pb.Filter_EQUAL, strconv.Itoa(int(user.TeleChatId)))
	numUpdated, err := db_pck.UpdateUserNonce(s.db, nonceString, userQuery)

	// Check if successfully updated
	if err != nil {
		fmt.Println("GetWANonce ERROR", err)
		res.ErrorMessage = err.Error()
		res.Type = pb.Response_ERROR
		return &pb.ResponseNonce{Response: res}, nil
	}
	if numUpdated < 1 {
		fmt.Println("GetWANonce ERROR No user's nonce updated", numUpdated)
		res.ErrorMessage = "ERROR No user's nonce updated"
		res.Type = pb.Response_ERROR
		return &pb.ResponseNonce{Response: res}, nil
	}

	// Return Nonce
	resNonce := pb.ResponseNonce{
		Response: res,
		Nonce:    nonceString,
	}

	return &resNonce, nil
}
