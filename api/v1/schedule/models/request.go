package models

type Request struct {
	StartTime   int
	EndTime     int
	Description string
	Requester   string
	Requestee   string
}
