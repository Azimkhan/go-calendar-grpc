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
	EventType  EventType `db:"event_type"`
	StartTime  time.Time `db:"start_time"`
	EndTime    time.Time `db:"end_time"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
}
