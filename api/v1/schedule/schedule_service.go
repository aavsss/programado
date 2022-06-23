package schedule

import (
	"sort"

	"github.com/aavsss/programado/api/v1/schedule/models"
	"github.com/aavsss/programado/api/v1/schedule/repository"
	"github.com/google/uuid"
)

type ScheduleService interface {
	AddToSchedule(period models.Period) bool
	RemoveFromSchedule(id string) bool
	ViewSchedule() []repository.PeriodInDB
	UpdateSchedule(id string, newPeriod models.Period) bool
}

type ScheduleServiceImpl struct {
	// Write dependencies here
}

func NewService() ScheduleService {
	return &ScheduleServiceImpl{}
}

func (s *ScheduleServiceImpl) AddToSchedule(period models.Period) bool {
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
			return false
		}
	}
	repository.MockData = append(repository.MockData, toBeAdded)
	return true
}

func (s *ScheduleServiceImpl) RemoveFromSchedule(id string) bool {
	// Imperative side. Change it to declarative
	for i, period := range repository.MockData {
		if period.Id == id {
			repository.MockData = append(repository.MockData[:i], repository.MockData[i+1:]...)
			return true
		}
	}
	return false
}

func (s *ScheduleServiceImpl) ViewSchedule() []repository.PeriodInDB {
	sort.SliceStable(repository.MockData, func(i, j int) bool {
		return repository.MockData[i].StartTime < repository.MockData[j].StartTime
	})
	return repository.MockData
}

func (s *ScheduleServiceImpl) UpdateSchedule(id string, newPeriod models.Period) bool {
	// to check if it overlaps or not
	for _, period := range repository.MockData {
		if period.Id != id {
			if (newPeriod.StartTime >= period.StartTime || newPeriod.EndTime >= period.StartTime) &&
				(newPeriod.StartTime <= period.EndTime || newPeriod.EndTime <= period.EndTime) {
				return false
			}
		}
	}
	for i, period := range repository.MockData {
		if period.Id == id {
			period := &repository.MockData[i]
			(*period).StartTime = newPeriod.StartTime
			(*period).EndTime = newPeriod.EndTime
			(*period).Description = newPeriod.Description
			return true
		}
	}
	return false
}
