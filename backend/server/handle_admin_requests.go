package server

import (
	"fmt"
	"strconv"
	"time"

	"capstone.operations_ecosystem/backend/common"
	db_pck "capstone.operations_ecosystem/backend/database"
	pb "capstone.operations_ecosystem/backend/proto"
	"github.com/getsentry/sentry-go"

	"context"
)

// gRPC defined endpoint. Add a user in the DB.
func (s *Server) AddUser(cxt context.Context, user *pb.FullUser) (*pb.Response, error) {
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

// gRPC defined endpoint. Update a user in the DB.
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

// gRPC defined endpoint. Delete a user in the DB.
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

// gRPC defined endpoint. Find users in the DB. The users are filtered out based the query.
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

// gRPC defined endpoint. Add a client in the DB.
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

// gRPC defined endpoint. Update a client in the DB.
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

// gRPC defined endpoint. Delete a client in the DB.
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

// gRPC defined endpoint. Find clients in the DB. The clients are filtered out based the query.
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

// gRPC defined endpoint.
// Returns a cryptographic-secure nonce
func (s *Server) GetWANonce(cxt context.Context, user *pb.User) (*pb.ResponseNonce, error) {
	fmt.Println("GetWANonce")
	res := &pb.Response{Type: pb.Response_ACK, ErrorMessage: "No error"}

	// Create nonce
	nonce, err := getCryptographicallySecureString(64)
	if err != nil {
		fmt.Println("GetWANonce ERROR", err)
		res.ErrorMessage = err.Error()
		res.Type = pb.Response_ERROR
		return &pb.ResponseNonce{Response: res}, nil
	}

	// Put nonce in DB
	userQuery := &pb.UserQuery{}
	db_pck.AddUserFilter(userQuery, pb.UserFilter_TELEGRAM_USER_ID, pb.Filter_EQUAL, strconv.Itoa(int(user.TeleUserId)))
	numUpdated, err := db_pck.UpdateUserNonce(s.db, nonce, userQuery)

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
		Nonce:    nonce,
	}

	return &resNonce, nil
}

// gRPC defined endpoint.
// Get the user's security string from the DB
// Either the user's id or email must be filled to identify the user
// If there are no users to be identified, a random string not attached to any
// user will be returned.
func (s *Server) GetSecurityString(cxt context.Context, user *pb.User) (*pb.SecurityStringResponse, error) {
	defer sentry.Recover()

	res := pb.Response{Type: pb.Response_ACK}
	query := &pb.UserQuery{}

	// Identify the user by their DB id or their email
	if user.UserId > 0 {
		db_pck.AddUserFilter(query, pb.UserFilter_USER_ID, pb.Filter_EQUAL, strconv.Itoa(int(user.UserId)))
	} else if len(user.Email) > 0 {
		db_pck.AddUserFilter(query, pb.UserFilter_EMAIL, pb.Filter_EQUAL, user.Email)
	} else {
		// User is not identifiable, return a random string that is not tied to anyone
		randString, err := getCryptographicallySecureString(SECURITY_STRING_LENGTH)
		securityStringRes := pb.SecurityStringResponse{Response: &res}
		if err != nil {
			res.Type = pb.Response_ERROR
			res.ErrorMessage = err.Error()
			return &securityStringRes, nil
		} else {
			securityStringRes.SecurityString = randString
		}
		return &securityStringRes, nil
	}

	foundUsers, err := db_pck.GetUsers(
		s.db,
		query,
		false,
	)

	if err != nil {
		securityStringRes := pb.SecurityStringResponse{Response: &res}
		res.Type = pb.Response_ERROR
		res.ErrorMessage = err.Error()
		return &securityStringRes, nil
	}

	if len(foundUsers) < 1 {
		securityStringRes := pb.SecurityStringResponse{Response: &res}
		res.Type = pb.Response_ERROR
		res.ErrorMessage = "Unable to find user"
		return &securityStringRes, nil
	}

	securityStringRes := pb.SecurityStringResponse{Response: &res}
	securityStringRes.SecurityString = foundUsers[0].SecurityString

	return &securityStringRes, nil
}

// gRPC defined endpoint. User authentication.
// User should send a hashed password. The server will check if the password matches the user
// and if so, sends back a token.
func (s *Server) AuthenticateUser(cxt context.Context, loginRequest *pb.LoginRequest) (*pb.UserTokenResponse, error) {
	defer sentry.Recover()

	res := pb.Response{Type: pb.Response_ACK}
	query := &pb.UserQuery{}
	db_pck.AddUserFilter(query, pb.UserFilter_EMAIL, pb.Filter_EQUAL, loginRequest.UserEmail)
	foundUsers, err := db_pck.GetUsers(
		s.db,
		query,
		false,
	)

	if err != nil {
		userTokenResponse := pb.UserTokenResponse{Response: &res}
		res.Type = pb.Response_ERROR
		res.ErrorMessage = err.Error()
		return &userTokenResponse, nil
	}

	if len(foundUsers) < 1 {
		userTokenResponse := pb.UserTokenResponse{Response: &res}
		res.Type = pb.Response_ERROR
		res.ErrorMessage = "Unable to find user"
		return &userTokenResponse, nil
	}

	// Check if the passwords are ok
	err = s.validateUserPassword(foundUsers[0].HashedPassword, loginRequest.HashedPassword)
	if err != nil {
		userTokenResponse := pb.UserTokenResponse{Response: &res}
		res.Type = pb.Response_ERROR
		res.ErrorMessage = err.Error()
		return &userTokenResponse, nil
	}

	// Generate token
	token, err := s.generateUserToken(foundUsers[0].User)
	if err != nil {
		userTokenResponse := pb.UserTokenResponse{Response: &res}
		res.Type = pb.Response_ERROR
		res.ErrorMessage = err.Error()
		return &userTokenResponse, nil
	}

	// Send token
	userTokenResponse := pb.UserTokenResponse{Response: &res}
	userTokenResponse.UserToken = token
	return &userTokenResponse, nil
}

func (s *Server) GetRegistrationCode(ctx context.Context, req *pb.RegistrationCodeRequest) (*pb.RegistrationCodeResponse, error) {
	fmt.Println("GetRegistrationCode")
	fmt.Println(req)
	if req.User.UserType != pb.User_MANAGER {
		return &pb.RegistrationCodeResponse{}, fmt.Errorf("Unauthorised")
	}
	newToken, err := getCryptographicallySecureString(64)
	if err != nil {
		return &pb.RegistrationCodeResponse{
			Response: &pb.Response{
				Type:         pb.Response_ERROR,
				ErrorMessage: err.Error(),
			},
		}, nil
	}
	regOtp := &pb.RegistrationOTP{
		Token:            newToken,
		CreationDatetime: time.Now().Format(common.DATETIME_FORMAT),
		Creator:          req.User,
		UserType:         pb.User_UserType(req.Type),
	}
	_, err = db_pck.InsertRegOTP(s.db, regOtp, s.dbLock)
	if err != nil {
		return &pb.RegistrationCodeResponse{
			Response: &pb.Response{

				Type: pb.Response_ERROR,

				ErrorMessage: err.Error(),
			},
		}, nil

	}

	return &pb.RegistrationCodeResponse{
		Response: &pb.Response{Type: pb.Response_ACK},
		Code:     newToken,
	}, nil
}

func (s *Server) CheckRegistrationCode(ctx context.Context, regCode *pb.RegistrationCode) (*pb.SecurityStringResponse, error) {
	fmt.Println("CheckRegistrationCode")
	fmt.Println(regCode)

	res := &pb.SecurityStringResponse{
		Response: &pb.Response{Type: pb.Response_ERROR},
	}

	//Todo: Check registration code and return a loginstring (security string)
	_, err := s.validateRegistrationToken(regCode.Code)
	if err != nil {
		fmt.Println("CheckRegistrationCode ERROR:", err)
		res.Response.ErrorMessage = "Registration code is invalid."
		return res, nil
	}

	randString, err := getCryptographicallySecureString(SECURITY_STRING_LENGTH)

	if err != nil {
		fmt.Println("CheckRegistrationCode ERROR:", err)
		res.Response.ErrorMessage = "The server is unable to generate a security string."
		return res, nil
	}

	res.SecurityString = randString
	res.Response.Type = pb.Response_ACK

	return res, nil
}
