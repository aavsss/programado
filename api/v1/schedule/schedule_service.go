package schedule

import (
	"sort"

	"github.com/aavsss/programado/api/v1/schedule/models"
	"github.com/aavsss/programado/api/v1/schedule/repository"
	"github.com/google/uuid"
)

type ScheduleService interface {
	AddToSchedule(period models.Period, toReturn chan bool)
	RemoveFromSchedule(id string, toReturn chan bool)
	ViewSchedule(toReturn chan []repository.PeriodInDB)
	UpdateSchedule(id string, newPeriod models.Period, toReturn chan bool)
}

type ScheduleServiceImpl struct {
	// Write dependencies here
}

func NewService() ScheduleService {
	return &ScheduleServiceImpl{}
}

func (s *ScheduleServiceImpl) AddToSchedule(period models.Period, toReturn chan bool) {
	toBeAdded := repository.PeriodInDB{
		Id:          uuid.New().String(),
		StartTime:   period.StartTime,
		EndTime:     period.EndTime,
		Description: period.Description,
	}
	// checking to see if the schedule overlaps or not
	for _, period := range repository.MockData {
		if (toBeAdded.StartTime >= period.StartTime || toBeAdded.EndTime >= period.StartTime) &&
			(toBeAdded.StartTime <= period.EndTime || toBeAdded.EndTime <= period.EndTime) {
			toReturn <- false
		}
	}
	repository.MockData = append(repository.MockData, toBeAdded)
	toReturn <- true
}

func (s *ScheduleServiceImpl) RemoveFromSchedule(id string, toReturn chan bool) {
	// Imperative side. Change it to declarative
	for i, period := range repository.MockData {
		if period.Id == id {
			repository.MockData = append(repository.MockData[:i], repository.MockData[i+1:]...)
			toReturn <- true
		}
	}
	toReturn <- false
}

func (s *ScheduleServiceImpl) ViewSchedule(toReturn chan []repository.PeriodInDB) {
	sort.SliceStable(repository.MockData, func(i, j int) bool {
		return repository.MockData[i].StartTime < repository.MockData[j].StartTime
	})
	toReturn <- repository.MockData
}

func (s *ScheduleServiceImpl) UpdateSchedule(id string, newPeriod models.Period, toReturn chan bool) {
	// to check if it overlaps or not
	for _, period := range repository.MockData {
		if period.Id != id {
			if (newPeriod.StartTime >= period.StartTime || newPeriod.EndTime >= period.StartTime) &&
				(newPeriod.StartTime <= period.EndTime || newPeriod.EndTime <= period.EndTime) {
				toReturn <- false
			}
		}
	}
	for i, period := range repository.MockData {
		if period.Id == id {
			period := &repository.MockData[i]
			(*period).StartTime = newPeriod.StartTime
			(*period).EndTime = newPeriod.EndTime
			(*period).Description = newPeriod.Description
			toReturn <- true
		}
	}
	toReturn <- false
}
