package server

import (
	"context"
	"crypto/rand"
	"fmt"
	"strconv"
	"time"

	"capstone.operations_ecosystem/backend/common"
	db_pck "capstone.operations_ecosystem/backend/database"
	pb "capstone.operations_ecosystem/backend/proto"
	"google.golang.org/grpc/metadata"
)

const (
	NONCE_HEADER          = "Authorization"
	TOKEN_EXPIRY_DURATION = time.Hour * 24

	// Cryptographic string lengths
	SECURITY_STRING_LENGTH = 64
	TOKEN_STRING_LENGTH    = 128
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
		return fmt.Errorf("No user's nonce was updated for user with id " + strconv.Itoa(int(user.TeleUserId)))
	}

	return nil
}

// Retrives a nonce from the header and verifies it with verifyNonce
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
	return s.verifyNonce(vals[0], user)
}

// Checks if a nonce is valid, and if so, removes the nonce from the DB
func (s *Server) verifyNonce(nonce string, user *pb.User) error {
	// Find the user
	fullUser, err := db_pck.IdUserByTelegramId(s.db, int(user.TeleUserId), false)
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

// Validates the password of the user
func (s *Server) validateUserPassword(actualPassword string, inputPassword string) error {
	if actualPassword == inputPassword {
		return nil
	}
	fmt.Println("validateUserPassword ERROR:", "password does not match")
	return fmt.Errorf("password does not match")
}

// Generates and inserts user tokens into the DB
func (s *Server) generateUserToken(user *pb.User) (*pb.UserToken, error) {
	fmt.Println("Generating user token")
	// Generate token
	token, err := getCryptographicallySecureString(TOKEN_STRING_LENGTH)
	if err != nil {
		return &pb.UserToken{}, err
	}

	userToken := &pb.UserToken{
		User:             user,
		Token:            token,
		CreationDatetime: time.Now().Format(common.DATETIME_FORMAT),
		ExpiryDatetime:   time.Now().Add(TOKEN_EXPIRY_DURATION).Format(common.DATETIME_FORMAT),
	}
	// Insert into DB
	_, err = db_pck.InsertUserToken(s.db, userToken, s.dbLock)
	if err != nil {
		return &pb.UserToken{}, err
	}

	return userToken, err
}

// Generates a cryptographically secure string of a certain length.
func getCryptographicallySecureString(length int) (string, error) {
	const chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		fmt.Println("getCryptographicallySecureString ERROR", err)
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = chars[b%byte(len(chars))]
	}

	return string(bytes), nil
}
