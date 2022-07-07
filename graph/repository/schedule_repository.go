package repository

import (
	"github.com/aavsss/programado/graph/model"
)

var MockData = []model.Period{
	{StartTime: 1, EndTime: 3, Description: "First meeting"},
	{StartTime: 10, EndTime: 14, Description: "Stand up"},
	{StartTime: 5, EndTime: 8, Description: "One on one"},
}

var ScheduleQueue = []model.Request{}

var RemovedScheduleRequests = []model.Request{}

var MockUsers = []model.User{
	{ID: "1", Role: "ADMIN"},
	{ID: "2", Role: "ADMIN"},
	{ID: "3", Role: "CLIENT"},
}
