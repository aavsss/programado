package master

import (
	"github.com/aavsss/programado/api/v1/schedule"
	"github.com/aavsss/programado/api/v1/schedule/repository"
	"github.com/google/uuid"
)

type MasterService interface {
	viewRequests(masterId string, toReturn chan []repository.RequestInDB)
	reviewRequest(masterId string, requestId string, accept bool)
}

type MasterServiceImpl struct {
	ScheduleService schedule.ScheduleService
}

func NewService() MasterService {
	return &MasterServiceImpl{}
}

func (ms *MasterServiceImpl) viewRequests(masterId string, toReturn chan []repository.RequestInDB) {
	var userQueue []repository.RequestInDB
	for _, request := range repository.ScheduleQueue {
		if request.Requestee == masterId {
			userQueue = append(userQueue, request)
		}
	}
	toReturn <- userQueue
}

func (ms *MasterServiceImpl) reviewRequest(masterId string, requestId string, accept bool) {
	for i, request := range repository.ScheduleQueue {
		if request.Id == requestId && request.Requestee == masterId {
			if accept {
				// todo: check to see if the schedule fits into period
				// or if there are alternatives to it
				// maybe the user wants to add or maybe they wanna replace it. depends on frontend

				// add the item to the schedule
				toBeAdded := repository.PeriodInDB{
					Id:          uuid.New().String(),
					StartTime:   request.StartTime,
					EndTime:     request.EndTime,
					Description: request.Description,
					UserId:      request.Requestee,
				}
				repository.MockData = append(repository.MockData, toBeAdded)
				// remove the item from the slice
				repository.ScheduleQueue = append(repository.ScheduleQueue[:i], repository.ScheduleQueue[i+1:]...)
			} else {
				// copy the item to removedScheduleRequests
				repository.RemovedScheduleRequests = append(repository.RemovedScheduleRequests, request)
				// remove the item form the slice
				repository.ScheduleQueue = append(repository.ScheduleQueue[:i], repository.ScheduleQueue[i+1:]...)
			}
		}
		return
	}
}
