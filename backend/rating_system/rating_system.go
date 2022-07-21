package rating_system

import (
	pb "capstone.operations_ecosystem/backend/proto"
)

const (
	ERROR_SCORE = -99999
)

type ScoreStruct struct {
	Score     float32
	Err       error
	ChannelId int
}

// TODO: Ask Emily for the calculation
func GetUserScore(user *pb.User) (float32, error) {
	return float32(100 - user.UserId), nil
}

func GetUserScoreFromChan(user *pb.User, userChan chan ScoreStruct, channelId int) {
	score, err := GetUserScore(user)

	userChan <- ScoreStruct{
		Score:     score,
		Err:       err,
		ChannelId: channelId,
	}
}
