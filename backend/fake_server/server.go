package fake_server

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"sync"

	db_pck "capstone.operations_ecosystem/backend/database"
	pb "capstone.operations_ecosystem/backend/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.AdminServicesServer
	pb.BroadcastServicesServer
	pb.RosterServicesServer
	pb.IncidentReportServicesServer
	pb.CameraIotServicesServer
	db     *sql.DB
	dbLock *sync.Mutex
}

func InitServer(serverAddr *string, serverPort *int) {
	fmt.Println("Starting gRPC fake server...")
	server := Server{dbLock: &sync.Mutex{}}
	server.db = db_pck.GetDB()

	if server.db == nil {
		log.Fatalf("InitServer: Failed to connect to DB")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *serverAddr, *serverPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterAdminServicesServer(grpcServer, &server)
	pb.RegisterBroadcastServicesServer(grpcServer, &server)
	pb.RegisterRosterServicesServer(grpcServer, &server)
	pb.RegisterIncidentReportServicesServer(grpcServer, &server)
	pb.RegisterCameraIotServicesServer(grpcServer, &server)

	grpcServer.Serve(lis)
}
