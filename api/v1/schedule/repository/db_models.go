package repository

type PeriodInDB struct {
	Id          string
	StartTime   int
	EndTime     int
	Description string
	UserId      string
}

type RequestInDB struct {
	Id          string
	StartTime   int
	EndTime     int
	Description string
	Requester   string
	Requestee   string
}
