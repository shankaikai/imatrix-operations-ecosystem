package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"time"

	"capstone.operations_ecosystem/backend/common"
	db_pck "capstone.operations_ecosystem/backend/database"
	pb "capstone.operations_ecosystem/backend/proto"

	"github.com/joho/godotenv"
)

const (
	IOT_POLLING_FREQUENCY = 10 * time.Second
	GATE_OPEN_KEYWORD     = "open"
	GATE_CLOSED_KEYWORD   = "closed"
)

type CameraIotStruct struct {
	// IoT Subscribers
	// Key: CameraIotId, Val: map[unique Id from Thread] channel
	// This channel is kept open by the subscriber.
	// The subscriber listens for any changes in the state via
	// the channel.
	GateSubscriptions      map[int64]map[string]chan *pb.CameraIot
	FireAlarmSubscriptions map[int64]map[string]chan *pb.CameraIot
	CpuTempSubscriptions   map[int64]map[string]chan *pb.CameraIot

	// Thingsboard credentials
	ThingsboardUsername string
	ThingsboardPassword string

	// Keep track of the states
	// Key: CameraIotId, Val: State
	GateStates      map[int64]pb.GateState_GatePosition
	FireAlarmStates map[int64]pb.FireAlarmState_AlarmState
	CpuTempStates   map[int64]float64
}

// Start go routines to all
func (s *Server) startAllIoTPolls() error {
	cameraIotAttributes, err := db_pck.GetCameraIotDetails(s.db, &pb.CameraIotQuery{})
	if err != nil {
		fmt.Println("startAllIoTPolls ERROR", err)
		return err
	}
	for _, attribute := range cameraIotAttributes {
		fmt.Println(attribute)
		go s.PollGateStatus(attribute.CameraIotId, "")
		// TODO poll other iot devices
	}

	return nil
}

// This function periodically checks the status of a particular gate
func (s *Server) PollGateStatus(cameraIotId int64, gateId string) {

	for range time.Tick(IOT_POLLING_FREQUENCY) {
		jwt, err := s.getJwtToken()

		if err != nil {
			fmt.Println("Unable to get JWT Token", err)
		}

		fmt.Println("Polling from gate", gateId)

		url := fmt.Sprintf("%s/%s", s.Config.ThingsboardUrl, s.Config.ThingsboardGetDeviceStateRelUrl)
		url = fmt.Sprintf(url, gateId)

		fmt.Println(url)
		resp, err := common.HttpGetWithJWT(url, jwt)

		if err != nil {
			fmt.Println("pollGateStatus ERROR", err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("pollGateStatus ERROR:", err)
			continue
		}

		if resp.StatusCode == 200 {
			// notify all subscribers
			stringBody := string(body)
			fmt.Println("string body", stringBody)

			// keyRegexp, err := regexp.Compile("\"key\":\"(.*?)\"")
			// if err != nil {
			// 	fmt.Println("pollGateStatus ERROR:", err)
			// 	continue
			// }

			valRegexp, err := regexp.Compile("\"value\":\"(.*?)\"")
			if err != nil {
				fmt.Println("pollGateStatus ERROR:", err)
				continue
			}

			// keyMatch := keyRegexp.FindStringSubmatch(stringBody)
			valMatch := valRegexp.FindStringSubmatch(stringBody)
			// fmt.Println("keyMatch[1]", keyMatch[1])
			fmt.Println("valMatch[1]", valMatch[1])

			// If there is no previous state, set it as init to prevent key error
			if _, ok := s.CameraIot.GateStates[cameraIotId]; !ok {
				s.CameraIot.GateStates[cameraIotId] = pb.GateState_INITIAL
			}

			go s.notifyGateSubscribers(cameraIotId, gateId, s.CameraIot.GateStates[cameraIotId], valMatch[1])
		}
	}
}

func (s *Server) setGateStatus(gateId string) error {
	postBody, err := json.Marshal(map[string]string{
		"label": "closed",
	})

	if err != nil {
		fmt.Println("setGateStatus ERROR", err)
		return err
	}

	jwt, err := s.getJwtToken()

	if err != nil {
		fmt.Println("Unable to get JWT Token", err)
	}

	fmt.Println("Setting state for gate", gateId)

	url := fmt.Sprintf("%s/%s", s.Config.ThingsboardUrl, s.Config.ThingsboardSetDeviceStateRelUrl)
	url = fmt.Sprintf(url, gateId)

	fmt.Println(url)
	resp, err := common.HttpPostWithJWT(url, jwt, string(postBody))

	if err != nil {
		fmt.Println("setGateStatus ERROR", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("setGateStatus ERROR:", err)
		return err
	}

	stringBody := string(body)
	fmt.Println("string body", stringBody)

	statusRegexp, err := regexp.Compile("\"status\"")
	// statusRegexp, err := regexp.Compile("\"status\":\"(.*?)\"")
	if err != nil {
		fmt.Println("getJwtToken ERROR:", err)
		return err
	}

	statusMatch := statusRegexp.FindStringSubmatch(stringBody)
	fmt.Println("statusMatch[1]", statusMatch[1])

	return nil
}

func (s *Server) getJwtToken() (string, error) {
	postBody, err := json.Marshal(map[string]string{
		"username": s.CameraIot.ThingsboardUsername,
		"password": s.CameraIot.ThingsboardPassword,
	})

	if err != nil {
		fmt.Println("getJwtToken ERROR", err)
		return "", err
	}

	responseBody := bytes.NewBuffer(postBody)
	// Make Auth Post Request
	resp, err := http.Post(s.Config.ThingsboardUrl+"/"+s.Config.ThingsboardAuthRelUrl, "application/json", responseBody)

	if err != nil {
		fmt.Println("getJwtToken ERROR:", err)
		return "", err
	}

	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("getJwtToken ERROR:", err)
		return "", err
	}
	stringBody := string(body)
	// fmt.Println(stringBody)

	tokenRegexp, err := regexp.Compile("\"token\":\"(.*?)\"")
	if err != nil {
		fmt.Println("getJwtToken ERROR:", err)
		return "", err
	}

	jwtMatch := tokenRegexp.FindStringSubmatch(stringBody)
	// fmt.Println("jwtMatch[1]", jwtMatch[1])
	return jwtMatch[1], nil
}

// If there is a state change, this function notifies all the subscribers of the gate
// that there is a new state change.
// If the new state is different from the old state, the global old state is changed.
func (s *Server) notifyGateSubscribers(cameraIotId int64, gateId string, oldState pb.GateState_GatePosition, newState string) {
	// Check if there is any change in state
	switch oldState {
	case pb.GateState_CLOSED:
		if GATE_CLOSED_KEYWORD == newState {
			return
		}
	case pb.GateState_OPEN:
		if GATE_OPEN_KEYWORD == newState {
			return
		}
	}

	// Create message to notify subscribers
	message := &pb.CameraIot{CameraIotId: cameraIotId, Gate: &pb.GateState{}}
	switch newState {
	case GATE_OPEN_KEYWORD:
		message.Gate.State = pb.GateState_OPEN
		s.CameraIot.GateStates[cameraIotId] = pb.GateState_OPEN
	case GATE_CLOSED_KEYWORD:
		message.Gate.State = pb.GateState_CLOSED
		s.CameraIot.GateStates[cameraIotId] = pb.GateState_CLOSED
	}

	for _, subscriberChannel := range s.CameraIot.GateSubscriptions[cameraIotId] {
		subscriberChannel <- message
	}
}

func (s *Server) subscribeToAllDevices(mainThreadChannel chan *pb.CameraIot, threadId string, cameraIotId int64) {
	// Subscribe to the gate device
	if _, ok := s.CameraIot.GateSubscriptions[cameraIotId]; !ok {
		s.CameraIot.GateSubscriptions[cameraIotId] = make(map[string]chan *pb.CameraIot)
	}
	s.CameraIot.GateSubscriptions[cameraIotId][threadId] = mainThreadChannel

	// Subscribe to the fire alarm device
	if _, ok := s.CameraIot.FireAlarmSubscriptions[cameraIotId]; !ok {
		s.CameraIot.FireAlarmSubscriptions[cameraIotId] = make(map[string]chan *pb.CameraIot)
	}
	s.CameraIot.GateSubscriptions[cameraIotId][threadId] = mainThreadChannel

	// Subscribe to the cpu temperatue
	if _, ok := s.CameraIot.CpuTempSubscriptions[cameraIotId]; !ok {
		s.CameraIot.CpuTempSubscriptions[cameraIotId] = make(map[string]chan *pb.CameraIot)
	}
	s.CameraIot.CpuTempSubscriptions[cameraIotId][threadId] = mainThreadChannel
}

func (s *Server) unsubscribeFromAllDevices(threadId string) {
	// Unsubscribe from the gate device
	for _, subs := range s.CameraIot.GateSubscriptions {
		delete(subs, threadId)
	}
	// Unsubscribe from the fire alarm
	for _, subs := range s.CameraIot.FireAlarmSubscriptions {
		delete(subs, threadId)
	}
	// Unsubscribe from the cpu temperatue
	for _, subs := range s.CameraIot.CpuTempSubscriptions {
		delete(subs, threadId)
	}

}

func (s *Server) getThingsBoardCreds() error {
	// load .env file
	envFilePath := ".env"
	err := godotenv.Load(envFilePath)

	if err != nil {
		fmt.Println(err)
	}

	s.CameraIot.ThingsboardUsername = os.Getenv("THINGSBOARD_USER")
	s.CameraIot.ThingsboardPassword = os.Getenv("THINGSBOARD_PASSWORD")

	return nil
}
