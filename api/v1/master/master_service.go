package master

import (
	"github.com/aavsss/programado/api/v1/schedule"
	"github.com/aavsss/programado/api/v1/schedule/repository"
)

type MasterService interface {
	viewRequests(masterId string) []repository.RequestInDB
	reviewRequests(masterId string, requestId string, accept bool)
}

type MasterServiceImpl struct {
	ScheduleService schedule.ScheduleService
}

func NewService() MasterService {
	return &MasterServiceImpl{}
}

func (ms *MasterServiceImpl) viewRequests(masterId string) []repository.RequestInDB {
	var userQueue []repository.RequestInDB
	for _, request := range repository.ScheduleQueue {
		if request.Requestee == masterId {
			userQueue = append(userQueue, request)
		}
	}
	return userQueue
}

func (ms *MasterServiceImpl) reviewRequests(masterId string, requestId string, accept bool) {
	for _, request := range repository.ScheduleQueue {
		if request.Id == requestId {
			if accept {
				// remove the item from the slice
			} else {
				// move it to temporary delete section
			}
		}
	}
}
