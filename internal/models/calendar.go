package models

import "time"

type CalendarEvent struct {
	Id        int64
	Name      string
	StartTime time.Time
	EndTime   time.Time
}
