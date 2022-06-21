package schedule

import (
	"github.com/aavsss/programado/api/v1/schedule/models"
	"github.com/aavsss/programado/api/v1/schedule/repository"
)

type ScheduleService interface {
	AddToSchedule(period models.Period) bool
	RemoveFromSchedule(id float64) bool
	ViewSchedule() []repository.PeriodInDB
	UpdateSchedule(id float64, newPeriod models.Period) bool
}

type ScheduleServiceImpl struct {
	// Write dependencies here
}

func NewService() ScheduleService {
	return &ScheduleServiceImpl{}
}

func (s *ScheduleServiceImpl) AddToSchedule(period models.Period) bool {
	toBeAdded := repository.PeriodInDB{
		Id:          4,
		StartTime:   period.StartTime,
		EndTime:     period.EndTime,
		Description: period.Description,
	}
	repository.MockData = append(repository.MockData, toBeAdded)
	return true
}

func (s *ScheduleServiceImpl) RemoveFromSchedule(id float64) bool {
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
	return repository.MockData
}

func (s *ScheduleServiceImpl) UpdateSchedule(id float64, newPeriod models.Period) bool {
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
