package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"path/filepath"
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

	db     *sql.DB
	dbLock *sync.Mutex

	teleServerAddr *string
	teleServerPort *int

	// AIFS LED lights for broadcasting
	testLEDAddr *string

	Config *ServerConfig
}

type ServerConfig struct {
	Aifs1Id int `json:"AIFS1_USER_ID"`
	Aifs2Id int `json:"AIFS2_USER_ID"`
	Aifs3Id int `json:"AIFS3_USER_ID"`
}

// Initialises the Server with all the necessary configurations.
// Server starts and listens at the given address and port.
func InitServer(serverAddr *string, serverPort *int, teleServerAddr *string, teleServerPort *int, testLEDAddr *string) {
	fmt.Println("Starting gRPC server...")
	server := Server{
		dbLock:         &sync.Mutex{},
		teleServerAddr: teleServerAddr,
		teleServerPort: teleServerPort,
		testLEDAddr:    testLEDAddr,
	}
	server.db = db_pck.GetDB()
	server.getServerConfigs()

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

	grpcServer.Serve(lis)
}

func (s *Server) getServerConfigs() {
	configFilePath := filepath.Join("config.json")
	fmt.Println("configFilePath", configFilePath)
	configFile, err := ioutil.ReadFile(configFilePath)

	if err != nil {
		fmt.Println("getServerConfigs", err, configFilePath)
		s.Config = &ServerConfig{}
		return
	}

	// we initialize our Users array
	config := &ServerConfig{}

	err = json.Unmarshal([]byte(configFile), config)
	if err != nil {
		fmt.Println("getServerConfigs", err, configFilePath)
		s.Config = &ServerConfig{}
		return
	}

	s.Config = config
}
