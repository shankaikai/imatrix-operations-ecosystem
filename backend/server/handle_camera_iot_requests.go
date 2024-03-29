package server

import (
	"context"
	"fmt"
	"strconv"
	"time"

	db_pck "capstone.operations_ecosystem/backend/database"
	pb "capstone.operations_ecosystem/backend/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

// gRPC defined endpoint. Opens or closes a particular gate.
func (s *Server) SetGateState(cxt context.Context, gateState *pb.GateState) (*pb.Response, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered", r)
		}
	}()
	res := pb.Response{Type: pb.Response_ACK}
	// Send gate command to IoT Server
	// Get attributes of the gate needed
	query := &pb.CameraIotQuery{}
	db_pck.AddCameraIotFilter(query, pb.CameraIotFilter_CAMERA_IOT_ID, pb.Filter_EQUAL, strconv.Itoa(int(gateState.Id)))
	cameraIotAttributes, err := db_pck.GetCameraIotDetails(s.db, query)

	if err != nil {
		res := pb.Response{Type: pb.Response_ERROR}
		res.ErrorMessage = err.Error()
		return &res, nil
	}

	if len(cameraIotAttributes) < 1 {
		res := pb.Response{Type: pb.Response_ERROR}
		res.ErrorMessage = "SetGateState: could not find id " + strconv.Itoa(int(gateState.Id)) + "in DB"
		return &res, nil
	}

	// Send the gate state request to the correct device
	err = s.setGateStatus(cameraIotAttributes[0].GateId, gateState.State)
	if err != nil {
		res := pb.Response{Type: pb.Response_ERROR}
		res.ErrorMessage = err.Error()
		return &res, nil
	}

	fmt.Println("Sending gate change command", gateState)
	return &res, nil
}

// gRPC defined endpoint. This opens a stream that will not close. Periodically, new updates of the IoT device states
// will be sent to the client.
func (s *Server) GetIotState(emptypb *emptypb.Empty, stream pb.CameraIotServices_GetIotStateServer) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered", r)
		}
	}()
	fmt.Println("STATES SUB DSFADSHAFIDSHFIUDSAHFDIUFHDISAFUHDASHFIDU")

	fmt.Println("GetIotState")
	cameraIotAttributes, err := db_pck.GetCameraIotDetails(s.db, &pb.CameraIotQuery{})

	if err != nil {
		res := pb.Response{Type: pb.Response_ERROR}
		cameraIotRes := pb.CameraIotResponse{Response: &res}
		res.ErrorMessage = err.Error()
		stream.Send(&cameraIotRes)
		return nil
	}

	// Create a unique ID for this thread
	threadId := strconv.Itoa(int(time.Now().Unix()))
	// Create a channel to listen to all the subscriptions with
	mainChannel := make(chan *pb.CameraIot, 10000)
	// Defer the unsubscriptions
	defer s.unsubscribeFromAllDevices(threadId)

	// Return camera end point, and initial states of all IoT devices
	for _, attribute := range cameraIotAttributes {
		res := pb.Response{Type: pb.Response_ACK}
		cameraIotRes := pb.CameraIotResponse{Response: &res, CameraIot: &pb.CameraIot{}}
		cameraIot := &pb.CameraIot{
			CameraIotId: attribute.CameraIotId,
			Name:        attribute.Name,
			Camera:      &pb.Camera{Url: attribute.CameraUrl},
			Type:        pb.CameraIot_INITIAL,
		}
		fmt.Println("attribute", attribute)

		// Subscribe to changes for all IoT devices
		s.subscribeToAllDevices(mainChannel, threadId, attribute.CameraIotId)

		// Get Gate Status
		if _, ok := s.CameraIot.GateStates[attribute.CameraIotId]; !ok {
			s.CameraIot.GateStates[attribute.CameraIotId] = pb.GateState_CLOSED
		}
		s.CameraIot.GateStatesLock.RLock()
		cameraIot.Gate = &pb.GateState{State: s.CameraIot.GateStates[attribute.CameraIotId]}
		s.CameraIot.GateStatesLock.RUnlock()

		// Get Fire Alarm Status
		if _, ok := s.CameraIot.FireAlarmStates[attribute.CameraIotId]; !ok {
			s.CameraIot.FireAlarmStates[attribute.CameraIotId] = pb.FireAlarmState_OFF
		}
		s.CameraIot.FireAlarmStatesLock.RLock()
		cameraIot.FireAlarm = &pb.FireAlarmState{State: s.CameraIot.FireAlarmStates[attribute.CameraIotId]}
		s.CameraIot.FireAlarmStatesLock.RUnlock()

		// Get Cpu Temperature Status
		if _, ok := s.CameraIot.CpuTempStates[attribute.CameraIotId]; !ok {
			s.CameraIot.CpuTempStates[attribute.CameraIotId] = -1
		}
		s.CameraIot.CpuTempStatesLock.RLock()
		cameraIot.CpuTemperature = &pb.CpuTempState{Temp: s.CameraIot.CpuTempStates[attribute.CameraIotId]}
		s.CameraIot.CpuTempStatesLock.RUnlock()

		cameraIotRes.CameraIot = cameraIot
		err = stream.Send(&cameraIotRes)
		if err != nil {
			fmt.Println("GetIotState ERROR:", err)
			return err
		}
	}

	// Continuously wait for new states
	for {
		res := pb.Response{Type: pb.Response_ACK}

		cameraIot := <-mainChannel
		cameraIotRes := pb.CameraIotResponse{Response: &res, CameraIot: cameraIot}
		err = stream.Send(&cameraIotRes)
		if err != nil {
			fmt.Println("GetIotState ERROR:", err)
			return err
		}
	}

}
