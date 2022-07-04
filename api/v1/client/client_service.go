package client

import (
	"github.com/aavsss/programado/api/v1/schedule"
	"github.com/aavsss/programado/api/v1/schedule/repository"

	"github.com/google/uuid"
)

type ClientService interface {
	RequestNewMeeting(toReturn chan bool)
}

type ClientServiceImpl struct {
	ScheduleService schedule.ScheduleService
}

func NewService(scheduleService schedule.ScheduleService) ClientService {
	return &ClientServiceImpl{
		ScheduleService: scheduleService,
	}
}

func (cs *ClientServiceImpl) RequestNewMeeting(toReturn chan bool) {
	tempData := repository.RequestInDB{
		Id:          uuid.New().String(),
		StartTime:   500,
		EndTime:     600,
		Description: "First request",
		Requester:   "1",
		Requestee:   "2",
	}
	repository.ScheduleQueue = append(repository.ScheduleQueue, tempData)
	toReturn <- true
}
