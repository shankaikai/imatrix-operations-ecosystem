package server

import (
	"context"
	"fmt"
	"strconv"
	"time"

	db_pck "capstone.operations_ecosystem/backend/database"
	pb "capstone.operations_ecosystem/backend/proto"
	"github.com/getsentry/sentry-go"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) SetGateState(cxt context.Context, gateState *pb.GateState) (*pb.Response, error) {
	defer sentry.Recover()
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
	err = s.setGateStatus(cameraIotAttributes[0].GateId)
	if err != nil {
		res := pb.Response{Type: pb.Response_ERROR}
		res.ErrorMessage = err.Error()
		return &res, nil
	}

	fmt.Println("Sending gate change command", gateState)
	return &res, nil
}

func (s *Server) GetIotState(emptypb *emptypb.Empty, stream pb.CameraIotServices_GetIotStateServer) error {
	defer sentry.Recover()

	fmt.Println("")
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
		}
		fmt.Println("attribute", attribute)

		// Subscribe to changes for all IoT devices
		s.subscribeToAllDevices(mainChannel, threadId, attribute.CameraIotId)

		// Get Gate Status
		// Get Fire Alarm Status
		// Get Cpu Temperature Status

		cameraIotRes.CameraIot = cameraIot
		err = stream.Send(&cameraIotRes)
		if err != nil {
			return err
		}
	}

	// Continuously wait for new states
	for {
		cameraIot := <-mainChannel
		res := pb.Response{Type: pb.Response_ACK}
		cameraIotRes := pb.CameraIotResponse{Response: &res, CameraIot: cameraIot}
		err = stream.Send(&cameraIotRes)
		if err != nil {
			return err
		}
	}

}