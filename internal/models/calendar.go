package models

import "time"

type EventType int32

const (
	EventType_UNKNOWN_EVENT_TYPE EventType = 0
	EventType_TASK               EventType = 1
	EventType_MEETING            EventType = 2
)

type CalendarEvent struct {
	Id         int64
	Name       string
	EventType  EventType
	StartTime  time.Time
	EndTime    time.Time
	CreateTime time.Time
	UpdateTime time.Time
}
