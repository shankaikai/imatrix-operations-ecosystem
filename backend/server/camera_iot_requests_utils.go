package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	"capstone.operations_ecosystem/backend/common"
	db_pck "capstone.operations_ecosystem/backend/database"
	pb "capstone.operations_ecosystem/backend/proto"

	"github.com/joho/godotenv"
)

const (
	// IOT_POLLING_FREQUENCY = 10 * time.Second
	IOT_POLLING_FREQUENCY = 2 * time.Second

	// Values expected from thingsboard
	GATE_OPEN_KEYWORD   = "open"
	GATE_CLOSED_KEYWORD = "closed"

	FIRE_ALARM_OFF_KEYWORD = "off"
	FIRE_ALARM_ON_KEYWORD  = "on"
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
		go s.PollGateStatus(attribute.CameraIotId, attribute.GateId)
		go s.PollFireAlarmStatus(attribute.CameraIotId, attribute.FireAlarmId)
		go s.PollCpuTempStatus(attribute.CameraIotId, attribute.CpuId)
	}

	return nil
}

// This function periodically checks the status of a particular gate
func (s *Server) PollGateStatus(cameraIotId int64, gateId string) {
	// If there is no previous state, set it to a default closed to prevent key error
	if _, ok := s.CameraIot.GateStates[cameraIotId]; !ok {
		s.CameraIot.GateStates[cameraIotId] = pb.GateState_CLOSED
	}

	fmt.Println(gateId)
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

			go s.notifyGateSubscribers(cameraIotId, gateId, s.CameraIot.GateStates[cameraIotId], valMatch[1])
		}
	}
}

func (s *Server) setGateStatus(gateId string, status pb.GateState_GatePosition) error {
	var statusJson map[string]string
	switch status {
	case pb.GateState_OPEN:
		statusJson = map[string]string{
			"status": GATE_OPEN_KEYWORD,
		}
	case pb.GateState_CLOSED:
		statusJson = map[string]string{
			"status": GATE_CLOSED_KEYWORD,
		}
	}

	postBody, err := json.Marshal(statusJson)

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

	if resp.StatusCode != 200 {
		fmt.Println("setGateStatus ERROR resp.StatusCode:", resp.StatusCode)
		return fmt.Errorf("setGateStatus STATUS CODE ERROR %d", resp.StatusCode)
	}

	return nil
}

// This function periodically checks the status of a particular fire alarm
func (s *Server) PollFireAlarmStatus(cameraIotId int64, fireAlarmId string) {
	// If there is no previous state, set it to a default off to prevent key error
	if _, ok := s.CameraIot.FireAlarmStates[cameraIotId]; !ok {
		s.CameraIot.FireAlarmStates[cameraIotId] = pb.FireAlarmState_OFF
	}
	fmt.Println(fireAlarmId)

	for range time.Tick(IOT_POLLING_FREQUENCY) {
		jwt, err := s.getJwtToken()

		if err != nil {
			fmt.Println("Unable to get JWT Token", err)
		}

		fmt.Println("Polling from Fire Alarm", fireAlarmId)

		url := fmt.Sprintf("%s/%s", s.Config.ThingsboardUrl, s.Config.ThingsboardGetDeviceStateRelUrl)
		url = fmt.Sprintf(url, fireAlarmId)

		fmt.Println(url)
		resp, err := common.HttpGetWithJWT(url, jwt)

		if err != nil {
			fmt.Println("PollFireAlarmStatus ERROR", err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("PollFireAlarmStatus ERROR:", err)
			continue
		}

		if resp.StatusCode == 200 {
			// notify all subscribers
			stringBody := string(body)
			fmt.Println("fire alarm string body", stringBody)

			valRegexp, err := regexp.Compile("\"value\":\"(.*?)\"")
			if err != nil {
				fmt.Println("PollFireAlarmStatus ERROR:", err)
				continue
			}

			// keyMatch := keyRegexp.FindStringSubmatch(stringBody)
			valMatch := valRegexp.FindStringSubmatch(stringBody)
			// fmt.Println("keyMatch[1]", keyMatch[1])
			fmt.Println("valMatch[1]", valMatch[1])

			go s.notifyFireAlarmSubscribers(cameraIotId, fireAlarmId, s.CameraIot.FireAlarmStates[cameraIotId], valMatch[1])
		}
	}

}

// This function periodically checks the status of a particular cpu
func (s *Server) PollCpuTempStatus(cameraIotId int64, cpuId string) {
	// If there is no previous state, set it to a default -1 to prevent key error
	if _, ok := s.CameraIot.CpuTempStates[cameraIotId]; !ok {
		s.CameraIot.CpuTempStates[cameraIotId] = -1
	}

	for range time.Tick(IOT_POLLING_FREQUENCY) {
		jwt, err := s.getJwtToken()

		if err != nil {
			fmt.Println("Unable to get JWT Token", err)
		}

		fmt.Println("Polling from cpu", cpuId)

		url := fmt.Sprintf("%s/%s", s.Config.ThingsboardUrl, s.Config.ThingsboardGetDeviceStateRelUrl)
		url = fmt.Sprintf(url, cpuId)

		fmt.Println(url)
		resp, err := common.HttpGetWithJWT(url, jwt)

		if err != nil {
			fmt.Println("PollCpuTempStatus ERROR", err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("PollCpuTempStatus ERROR:", err)
			continue
		}

		if resp.StatusCode == 200 {
			// notify all subscribers
			stringBody := string(body)
			fmt.Println("string body", stringBody)

			valRegexp, err := regexp.Compile("\"value\":(.*?)\\}")
			if err != nil {
				fmt.Println("PollCpuTempStatus ERROR:", err)
				continue
			}

			// keyMatch := keyRegexp.FindStringSubmatch(stringBody)
			valMatch := valRegexp.FindStringSubmatch(stringBody)
			if len(valMatch) == 0 {
				fmt.Println("PollCpuTempStatus ERROR val 0", stringBody)
				continue
			}
			// fmt.Println("keyMatch[1]", keyMatch[1])
			fmt.Println("valMatch[1]", valMatch[1])

			newState, err := strconv.ParseFloat(valMatch[1], 64)
			if err != nil {
				fmt.Println("PollCpuTempStatus ERROR:", err)
				continue
			}

			go s.notifyCpuTempSubscribers(cameraIotId, cpuId, s.CameraIot.CpuTempStates[cameraIotId], newState)
		}
	}

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

	if resp.StatusCode != 200 {
		fmt.Println("getJwtToken ERROR STATUS CODE", resp.StatusCode)
		return "", fmt.Errorf("getJwtToken STATUS CODE ERROR %d", resp.StatusCode)

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

// If there is a state change, this function notifies all the subscribers of the gate
// that there is a new state change.
// If the new state is different from the old state, the global old state is changed.
func (s *Server) notifyFireAlarmSubscribers(cameraIotId int64, fireAlarmId string, oldState pb.FireAlarmState_AlarmState, newState string) {
	// Check if there is any change in state
	switch oldState {
	case pb.FireAlarmState_OFF:
		if FIRE_ALARM_OFF_KEYWORD == newState {
			return
		}
	case pb.FireAlarmState_ON:
		if FIRE_ALARM_ON_KEYWORD == newState {
			return
		}
	}

	// Create message to notify subscribers
	message := &pb.CameraIot{CameraIotId: cameraIotId, FireAlarm: &pb.FireAlarmState{}}
	switch newState {
	case FIRE_ALARM_OFF_KEYWORD:
		message.FireAlarm.State = pb.FireAlarmState_OFF
		s.CameraIot.FireAlarmStates[cameraIotId] = pb.FireAlarmState_OFF
	case FIRE_ALARM_ON_KEYWORD:
		message.FireAlarm.State = pb.FireAlarmState_ON
		s.CameraIot.FireAlarmStates[cameraIotId] = pb.FireAlarmState_ON
	}

	for _, subscriberChannel := range s.CameraIot.GateSubscriptions[cameraIotId] {
		subscriberChannel <- message
	}
}

// If there is a state change, this function notifies all the subscribers of the gate
// that there is a new state change.
// If the new state is different from the old state, the global old state is changed.
func (s *Server) notifyCpuTempSubscribers(cameraIotId int64, cpuId string, oldState float64, newState float64) {
	// Check if there is any change in state
	if oldState == newState {
		return
	}

	// Create message to notify subscribers
	message := &pb.CameraIot{CameraIotId: cameraIotId, CpuTemperature: &pb.CpuTempState{}}

	message.CpuTemperature.Temp = newState
	s.CameraIot.CpuTempStates[cameraIotId] = newState

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
	s.CameraIot.FireAlarmSubscriptions[cameraIotId][threadId] = mainThreadChannel

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

	fmt.Println("SUBSCRIPTIONS UPDATE")
	fmt.Println(s.CameraIot.GateSubscriptions)
	fmt.Println(s.CameraIot.FireAlarmSubscriptions)
	fmt.Println(s.CameraIot.CpuTempSubscriptions)
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
