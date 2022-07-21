package server

import (
	"context"
	"fmt"
	"strconv"

	db_pck "capstone.operations_ecosystem/backend/database"
	pb "capstone.operations_ecosystem/backend/proto"
	"google.golang.org/grpc/metadata"
)

const (
	NONCE_HEADER = "Authorization"
)

// Removes a Nonce that has been used from the DB for a particular user.
// Needs the user's db user id
func (s *Server) removeUserNonce(user *pb.User) error {
	// Put nonce in DB
	userQuery := &pb.UserQuery{}
	db_pck.AddUserFilter(userQuery, pb.UserFilter_USER_ID, pb.Filter_EQUAL, strconv.Itoa(int(user.UserId)))
	numUpdated, err := db_pck.UpdateUserNonce(s.db, "", userQuery)

	// Check if successfully updated
	if err != nil {
		fmt.Println("removeUserNonce ERROR:", err)
		return err
	}
	if numUpdated < 1 {
		return fmt.Errorf("No user's nonce was updated for user with id " + strconv.Itoa(int(user.TeleChatId)))
	}

	return nil
}

// Checks if the nonce from the header is valid, and if so,
// removes the nonce from the DB
func (s *Server) verifyNonceFromHeaders(ctx context.Context, user *pb.User) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		errMsg := "unable to get nonce from header"
		fmt.Println("verifyNonceFromHeaders ERROR no metadata:", errMsg)
		return fmt.Errorf(errMsg)
	}

	vals := md.Get(NONCE_HEADER)
	if len(vals) < 1 {
		errMsg := "unable to get nonce from header"
		fmt.Println("verifyNonceFromHeaders ERROR has metadata:", md, errMsg)
		return fmt.Errorf(errMsg)
	}

	nonce := vals[0]

	// Find the user
	fullUser, err := db_pck.IdUserByTelegramId(s.db, int(user.TeleChatId), false)
	if err != nil {
		return err
	}

	// Check the nonce
	if nonce != fullUser.Nonce {
		errMsg := "nonce does not match"
		fmt.Println("verifyNonceFromHeaders ERROR:", errMsg)
		return fmt.Errorf(errMsg)
	}

	fmt.Println("Removing Nonce...")
	// Remove nonce from db
	err = s.removeUserNonce(fullUser.User)
	if err != nil {
		return err
	}

	// All good
	return nil
}
