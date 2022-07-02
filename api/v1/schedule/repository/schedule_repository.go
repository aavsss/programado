package repository

var MockData = []PeriodInDB{
	{Id: "1-id", StartTime: 1, EndTime: 3, Description: "First meeting", UserId: "1"},
	{Id: "2-id", StartTime: 10, EndTime: 14, Description: "Stand up", UserId: "1"},
	{Id: "3-id", StartTime: 5, EndTime: 8, Description: "One on one", UserId: "2"},
}

var ScheduleQueue = []RequestInDB{}
