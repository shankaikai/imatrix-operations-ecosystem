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
	"sync"
	"time"

	"capstone.operations_ecosystem/backend/common"
	db_pck "capstone.operations_ecosystem/backend/database"
	pb "capstone.operations_ecosystem/backend/proto"

	"github.com/joho/godotenv"
)

const (
	// Frequency at which to update the server's state of the IoT devices
	// at each client.
	IOT_POLLING_FREQUENCY             = 10 * time.Second
	GATE_CONSISTENCY_UPDATE_FREQUENCY = 5 * time.Second

	// Frequency at which the server's Thingsboard JWT Token should be
	// refreshed
	JWT_TOKEN_REFRESH_FREQUENCY = 1*time.Hour + 30*time.Minute

	// Values expected from thingsboard
	GATE_SHARED_ATTRIBUTE_KEY   = "request"
	GATE_REQUST_OPEN_KEYWORD    = "open"
	GATE_REQUEST_CLOSED_KEYWORD = "close"
	GATE_STATUS_OPEN_KEYWORD    = "true"
	GATE_STATUS_CLOSED_KEYWORD  = "false"

	FIRE_ALARM_OFF_KEYWORD = "off"
	FIRE_ALARM_ON_KEYWORD  = "on"
)

// Used to store information about the IoT devices and their states.
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
	// This JWT Token is refreshed periodically
	// and used for all calls to the thingsboard server.
	JwtToken string

	// Keep track of the states
	// Key: CameraIotId, Val: State
	GateStates      map[int64]pb.GateState_GatePosition
	FireAlarmStates map[int64]pb.FireAlarmState_AlarmState
	CpuTempStates   map[int64]float64

	//Locks to write to the maps
	GateSubscriptionsLock *sync.RWMutex
	FireAlarmSubscriptionsLock *sync.RWMutex
	CpuTempSubscriptionsLock *sync.RWMutex
	GateStatesLock *sync.RWMutex
	FireAlarmStatesLock *sync.RWMutex
	CpuTempStatesLock *sync.RWMutex

}

// Initialise the camera and Iot service.
// This function will try to authenticate with Thingsboard first
// to get a JWT token. If it is succesfully, it start new goroutines
// to poll the states of all IoT devices periodically.
func (s *Server) initCameraIotService() error {
	// Set JWT token once
	err := s.refreshJwTToken()
	if err != nil {
		return err
	}
// Start polling the states of all IoT devices
	err = s.startAllIoTPolls()
	if err != nil {
		return err
	}

	// Auto refresh jwt token
	go func() {
		for range time.Tick(JWT_TOKEN_REFRESH_FREQUENCY) {
			err := s.refreshJwTToken()
			if err != nil {
				fmt.Println("AUTO REFRESH JWT TOKEN ERROR:", err)
			}
		}
	}()

	return nil
}

func (s *Server) refreshJwTToken() error {
	jwt, err := s.getJwtToken()

	if err != nil {
		fmt.Println("autoRefreshJwTToken ERROR", err)
		return err
	}

	s.CameraIot.JwtToken = jwt
	return nil
}

// Start go routines to all polls
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
	s.CameraIot.GateStatesLock.Lock()
	if _, ok := s.CameraIot.GateStates[cameraIotId]; !ok {
		s.CameraIot.GateStates[cameraIotId] = pb.GateState_CLOSED
	}
	s.CameraIot.GateStatesLock.Unlock()
	fmt.Println(gateId)
	for range time.Tick(GATE_CONSISTENCY_UPDATE_FREQUENCY) {

		fmt.Println("Polling from gate", gateId)

		url := fmt.Sprintf("%s/%s", s.Config.ThingsboardUrl, s.Config.ThingsboardGetDeviceStateRelUrl)
		url = fmt.Sprintf(url, gateId)

		fmt.Println(url)
		resp, err := common.HttpGetWithJWT(url, s.CameraIot.JwtToken)

		if err != nil {
			fmt.Println("pollGateStatus ERROR", err)
			continue
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

			valRegexp, err := regexp.Compile("\"key\":\"gate_opened\",\"value\":(.*?)}")
			if err != nil {
				fmt.Println("pollGateStatus ERROR:", err)
				continue
			}

			// keyMatch := keyRegexp.FindStringSubmatch(stringBody)
			valMatch := valRegexp.FindStringSubmatch(stringBody)
			if len(valMatch) == 0 {
				fmt.Println("pollGateStatus ERROR val 0", stringBody)
				continue
			}

			fmt.Println("valMatch[1]", valMatch[1])

			go s.notifyGateSubscribers(cameraIotId, gateId, valMatch[1])
		}
	}
}

func (s *Server) setGateStatus(gateId string, status pb.GateState_GatePosition) error {
	var statusJson map[string]string
	switch status {
	case pb.GateState_OPEN:
		statusJson = map[string]string{
			GATE_SHARED_ATTRIBUTE_KEY: GATE_REQUST_OPEN_KEYWORD,
		}
	case pb.GateState_CLOSED:
		statusJson = map[string]string{
			GATE_SHARED_ATTRIBUTE_KEY: GATE_REQUEST_CLOSED_KEYWORD,
		}
	}

	postBody, err := json.Marshal(statusJson)

	if err != nil {
		fmt.Println("setGateStatus ERROR", err)
		return err
	}

	fmt.Println("Setting state for gate", gateId)

	url := fmt.Sprintf("%s/%s", s.Config.ThingsboardUrl, s.Config.ThingsboardSetDeviceStateRelUrl)
	url = fmt.Sprintf(url, gateId)

	fmt.Println(url)
	resp, err := common.HttpPostWithJWT(url, s.CameraIot.JwtToken, string(postBody))

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
	s.CameraIot.FireAlarmStatesLock.Lock()
	if _, ok := s.CameraIot.FireAlarmStates[cameraIotId]; !ok {
		s.CameraIot.FireAlarmStates[cameraIotId] = pb.FireAlarmState_OFF
	}
	s.CameraIot.FireAlarmStatesLock.Unlock()

	fmt.Println(fireAlarmId)

	for range time.Tick(IOT_POLLING_FREQUENCY) {

		fmt.Println("Polling from Fire Alarm", fireAlarmId)

		url := fmt.Sprintf("%s/%s", s.Config.ThingsboardUrl, s.Config.ThingsboardGetDeviceStateRelUrl)
		url = fmt.Sprintf(url, fireAlarmId)

		fmt.Println(url)
		resp, err := common.HttpGetWithJWT(url, s.CameraIot.JwtToken)

		if err != nil {
			fmt.Println("PollFireAlarmStatus ERROR", err)
			continue
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
			if len(valMatch) == 0 {
				fmt.Println("PollFireAlarmStatus ERROR val 0", stringBody)
				continue
			}

			fmt.Println("valMatch[1]", valMatch[1])

			go s.notifyFireAlarmSubscribers(cameraIotId, fireAlarmId, s.CameraIot.FireAlarmStates[cameraIotId], valMatch[1])
		}
	}

}

// This function periodically checks the status of a particular cpu
func (s *Server) PollCpuTempStatus(cameraIotId int64, cpuId string) {
	// If there is no previous state, set it to a default -1 to prevent key error
	s.CameraIot.CpuTempStatesLock.Lock()
	if _, ok := s.CameraIot.CpuTempStates[cameraIotId]; !ok {
		s.CameraIot.CpuTempStates[cameraIotId] = -1
	}
	s.CameraIot.CpuTempStatesLock.Unlock()

	for range time.Tick(IOT_POLLING_FREQUENCY) {
		fmt.Println("Polling from cpu", cpuId)

		url := fmt.Sprintf("%s/%s", s.Config.ThingsboardUrl, s.Config.ThingsboardGetDeviceStateRelUrl)
		url = fmt.Sprintf(url, cpuId)

		fmt.Println(url)
		resp, err := common.HttpGetWithJWT(url, s.CameraIot.JwtToken)

		if err != nil {
			fmt.Println("PollCpuTempStatus ERROR", err)
			continue
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

			valMatch := valRegexp.FindStringSubmatch(stringBody)
			if len(valMatch) == 0 {
				fmt.Println("PollCpuTempStatus ERROR val 0", stringBody)
				continue
			}

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
	if len(jwtMatch) == 0 {
		fmt.Println("getJwtToken ERROR val 0", stringBody)
		return "", fmt.Errorf("getJwtToken ERROR: cannot get token")
	}

	return jwtMatch[1], nil
}

// No matter if there is any change in state, this function notifies all the subscribers of the gate
// that there is a new state change.
// If the new state is different from the old state, the global old state is changed.
func (s *Server) notifyGateSubscribers(cameraIotId int64, gateId string, newState string) {
	// Create message to notify subscribers
	message := &pb.CameraIot{
		CameraIotId: cameraIotId,
		Gate:        &pb.GateState{},
		Type:        pb.CameraIot_CHANGE_GATE,
	}

	switch newState {
	case GATE_STATUS_OPEN_KEYWORD:
		message.Gate.State = pb.GateState_OPEN
		
		s.CameraIot.GateStatesLock.Lock()
		s.CameraIot.GateStates[cameraIotId] = pb.GateState_OPEN
		s.CameraIot.GateStatesLock.Unlock()
	case GATE_STATUS_CLOSED_KEYWORD:
		message.Gate.State = pb.GateState_CLOSED

		s.CameraIot.GateStatesLock.Lock()
		s.CameraIot.GateStates[cameraIotId] = pb.GateState_CLOSED
		s.CameraIot.GateStatesLock.Unlock()
	}
	s.CameraIot.GateSubscriptionsLock.RLock()
	for _, subscriberChannel := range s.CameraIot.GateSubscriptions[cameraIotId] {
		subscriberChannel <- message
	}
	s.CameraIot.GateSubscriptionsLock.RUnlock()

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
	message := &pb.CameraIot{
		CameraIotId: cameraIotId,
		FireAlarm:   &pb.FireAlarmState{},
		Type:        pb.CameraIot_CHANGE_FIRE_ALARM,
	}
	switch newState {
	case FIRE_ALARM_OFF_KEYWORD:
		message.FireAlarm.State = pb.FireAlarmState_OFF
		s.CameraIot.FireAlarmStatesLock.Lock()
		s.CameraIot.FireAlarmStates[cameraIotId] = pb.FireAlarmState_OFF
		s.CameraIot.FireAlarmStatesLock.Unlock()
	case FIRE_ALARM_ON_KEYWORD:
		message.FireAlarm.State = pb.FireAlarmState_ON
		s.CameraIot.FireAlarmStatesLock.Lock()
		s.CameraIot.FireAlarmStates[cameraIotId] = pb.FireAlarmState_ON
		s.CameraIot.FireAlarmStatesLock.Unlock()
	}
	s.CameraIot.FireAlarmSubscriptionsLock.RLock()
	for _, subscriberChannel := range s.CameraIot.FireAlarmSubscriptions[cameraIotId] {
		subscriberChannel <- message
	}
	s.CameraIot.FireAlarmSubscriptionsLock.RUnlock()

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
	message := &pb.CameraIot{
		CameraIotId:    cameraIotId,
		CpuTemperature: &pb.CpuTempState{},
		Type:           pb.CameraIot_CHANGE_CPU_TEMP,
	}

	message.CpuTemperature.Temp = newState
	s.CameraIot.CpuTempStatesLock.Lock()
	s.CameraIot.CpuTempStates[cameraIotId] = newState
	s.CameraIot.CpuTempStatesLock.Unlock()

	s.CameraIot.CpuTempSubscriptionsLock.RLock()
	for _, subscriberChannel := range s.CameraIot.CpuTempSubscriptions[cameraIotId] {
		subscriberChannel <- message
	}
	s.CameraIot.CpuTempSubscriptionsLock.RUnlock()
}

func (s *Server) subscribeToAllDevices(mainThreadChannel chan *pb.CameraIot, threadId string, cameraIotId int64) {
	// Subscribe to the gate device
	s.CameraIot.GateSubscriptionsLock.Lock()
	fmt.Println("GATE DSFADSHAFIDSHFIUDSAHFDIUFHDISAFUHDASHFID")
	if _, ok := s.CameraIot.GateSubscriptions[cameraIotId]; !ok {
		s.CameraIot.GateSubscriptions[cameraIotId] = make(map[string]chan *pb.CameraIot)
	}
	s.CameraIot.GateSubscriptions[cameraIotId][threadId] = mainThreadChannel
	s.CameraIot.GateSubscriptionsLock.Unlock()

	// Subscribe to the fire alarm device
	s.CameraIot.FireAlarmSubscriptionsLock.Lock()
	fmt.Println("FIRE DSFADSHAFIDSHFIUDSAHFDIUFHDISAFUHDASHFIDU")
	if _, ok := s.CameraIot.FireAlarmSubscriptions[cameraIotId]; !ok {
		s.CameraIot.FireAlarmSubscriptions[cameraIotId] = make(map[string]chan *pb.CameraIot)
	}
	s.CameraIot.FireAlarmSubscriptions[cameraIotId][threadId] = mainThreadChannel
	s.CameraIot.FireAlarmSubscriptionsLock.Unlock()

	// Subscribe to the cpu temperatue
	s.CameraIot.CpuTempSubscriptionsLock.Lock()
	fmt.Println("CPU DSFADSHAFIDSHFIUDSAHFDIUFHDISAFUHDASHFIDU")

	if _, ok := s.CameraIot.CpuTempSubscriptions[cameraIotId]; !ok {
		s.CameraIot.CpuTempSubscriptions[cameraIotId] = make(map[string]chan *pb.CameraIot)
	}
	s.CameraIot.CpuTempSubscriptions[cameraIotId][threadId] = mainThreadChannel
	s.CameraIot.CpuTempSubscriptionsLock.Unlock()
}

func (s *Server) unsubscribeFromAllDevices(threadId string) {
	// Unsubscribe from the gate device
	s.CameraIot.GateSubscriptionsLock.Lock()
	for _, subs := range s.CameraIot.GateSubscriptions {
		delete(subs, threadId)
	}
	s.CameraIot.GateSubscriptionsLock.Unlock()

	// Unsubscribe from the fire alarm
	s.CameraIot.FireAlarmSubscriptionsLock.Lock()
	for _, subs := range s.CameraIot.FireAlarmSubscriptions {
		delete(subs, threadId)
	}
	s.CameraIot.FireAlarmSubscriptionsLock.Unlock()

	// Unsubscribe from the cpu temperatue
	s.CameraIot.CpuTempSubscriptionsLock.Lock()
	for _, subs := range s.CameraIot.CpuTempSubscriptions {
		delete(subs, threadId)
	}
	s.CameraIot.CpuTempSubscriptionsLock.Unlock()

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
