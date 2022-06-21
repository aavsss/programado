package schedule

import (
	"github.com/aavsss/programado/api/v1/schedule/models"
)

type ScheduleService interface {
	AddToSchedule() bool
	RemoveFromSchedule() bool
	ViewSchedule() []models.Period
	UpdateSchedule() bool
}

type ScheduleServiceImpl struct {
}

func NewService() ScheduleService {
	return &ScheduleServiceImpl{}
}

func (s *ScheduleServiceImpl) AddToSchedule() bool {
	return true
}

func (s *ScheduleServiceImpl) RemoveFromSchedule() bool {
	return true
}

func (s *ScheduleServiceImpl) ViewSchedule() []models.Period {
	return nil
}

func (s *ScheduleServiceImpl) UpdateSchedule() bool {
	return true
}
