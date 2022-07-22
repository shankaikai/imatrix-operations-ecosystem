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
	pb.CameraIotServicesServer
	pb.WebAppServicesServer

	db     *sql.DB
	dbLock *sync.Mutex

	teleServerAddr *string
	teleServerPort *int

	// AIFS LED lights for broadcasting
	testLEDAddr *string

	Config *ServerConfig

	CameraIot *CameraIotStruct
}

// Configurations for the server as defined in the config.json file.
type ServerConfig struct {
	// The IDs of the AIFS official user accounts in the DB.
	Aifs1Id int `json:"AIFS1_USER_ID"`
	Aifs2Id int `json:"AIFS2_USER_ID"`
	Aifs3Id int `json:"AIFS3_USER_ID"`

	// Thingsboard Server URLs
	ThingsboardUrl                  string `json:"THINGSBOARD_URL"`
	ThingsboardAuthRelUrl           string `json:"THINGSBOARD_AUTH_RELATIVE_URL"`
	ThingsboardGetDeviceStateRelUrl string `json:"THINGSBOARD_GET_DEVICE_STATE_RELATIVE_URL"`
	ThingsboardSetDeviceStateRelUrl string `json:"THINGSBOARD_SET_DEVICE_STATE_RELATIVE_URL"`
}

// Initialise the gRPC server with all the necessary information.
// The server will listen to requests incoming to serverAddr:serverPort.
func InitServer(serverAddr *string, serverPort *int, teleServerAddr *string, teleServerPort *int,
	testLEDAddr *string, webProxyAddr *string, webProxyPort *int) {
	fmt.Println("Starting gRPC server...")
	server := Server{
		dbLock:         &sync.Mutex{},
		teleServerAddr: teleServerAddr,
		teleServerPort: teleServerPort,
		testLEDAddr:    testLEDAddr,
		CameraIot: &CameraIotStruct{
			GateSubscriptions:      make(map[int64]map[string]chan *pb.CameraIot),
			FireAlarmSubscriptions: make(map[int64]map[string]chan *pb.CameraIot),
			CpuTempSubscriptions:   make(map[int64]map[string]chan *pb.CameraIot),
			GateStates:             make(map[int64]pb.GateState_GatePosition),
			FireAlarmStates:        make(map[int64]pb.FireAlarmState_AlarmState),
			CpuTempStates:          make(map[int64]float64),
		},
	}

	server.db = db_pck.GetDB()
	server.getServerConfigs()
	server.getThingsBoardCreds()

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
	pb.RegisterWebAppServicesServer(grpcServer, &server)

	server.initCameraIotService()

	go Proxy_main(serverAddr, serverPort, webProxyAddr, webProxyPort)
	grpcServer.Serve(lis)
}

// Get all server configurations from the config file.
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
