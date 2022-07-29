package server

import (
	"context"
	"fmt"
	"time"

	"capstone.operations_ecosystem/backend/common"
	db_pck "capstone.operations_ecosystem/backend/database"
	pb "capstone.operations_ecosystem/backend/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	FULL_PATH_AUTH_USER        = "/operations_ecosys.AdminServices/AuthenticateUser"
	FULL_PATH_GET_SECURITY_STR = "/operations_ecosys.AdminServices/GetSecurityString"
	JWT_TOKEN_HEADER           = "authorization"
)

//TODO add more endpoint exceptions especially from http
// Authorization grpc unary interceptor function
func unaryServerInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	// Do not try to authorise for the following calls
	fmt.Println("info.FullMethod", info.FullMethod)

	if !(info.FullMethod == FULL_PATH_AUTH_USER ||
		info.FullMethod == FULL_PATH_GET_SECURITY_STR) {
		if err := authenticateUser(ctx); err != nil {
			return nil, err
		}
	}

	// Call handler
	return handler(ctx, req)
}

func streamServerInterceptor(srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler) error {
	fmt.Println("info.FullMethod", info.FullMethod)

	if err := authenticateUser(ss.Context()); err != nil {
		return err
	}

	// Call handler
	return handler(srv, ss)
}

func authenticateUser(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.InvalidArgument, "Could not get metadata")
	}
	fmt.Println("metadata:", md)

	authHeader, ok := md[JWT_TOKEN_HEADER]
	if !ok {
		return status.Errorf(codes.Unauthenticated, "Could not get JWT token header")
	}

	token := authHeader[0]

	err := validateJWT(token)

	return err
}

// Check if the JWT token is valid
// TODO join JWT to a user
func validateJWT(token string) error {
	dbInstance := db_pck.GetDB()
	userTokenQuery := &pb.UserTokenQuery{}
	db_pck.AddUserTokenFilter(userTokenQuery, pb.UserTokenFilter_TOKEN, pb.Filter_EQUAL, token)
	tokens, err := db_pck.GetUserTokens(dbInstance, userTokenQuery)

	if err != nil {
		return status.Errorf(codes.Unauthenticated, "token invalid - "+err.Error())
	}
	// check if token length is more than 0
	if len(tokens) < 1 {
		return status.Errorf(codes.Unauthenticated, "token invalid")
	}

	// check expiry
	expiry_time, err := time.Parse(common.DATETIME_FORMAT, tokens[0].ExpiryDatetime)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "token invalid - "+err.Error())
	}

	if expiry_time.Before(time.Now()) {
		// remove token from db
		db_pck.DeleteUserToken(dbInstance, tokens[0])
		return status.Errorf(codes.Unauthenticated, "token expired")
	}

	// token valid
	return nil
}
