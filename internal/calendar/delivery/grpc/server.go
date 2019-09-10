//go:generate protoc --go_out=plugins=grpc:. --proto_path=../../../../api/protobuf calendar.proto
package grpc

import (
	"context"
	"log"
)

type CalendarServer struct {
	logger *log.Logger
}

func (c CalendarServer) CreateEvent(context.Context, *CreateCalendarEventRequest) (*CalendarEventObject, error) {
	panic("implement me")
}

func (c CalendarServer) ListEvents(context.Context, *ListCalendarEventRequest) (*ListCalendarEventResponse, error) {
	panic("implement me")
}

func (c CalendarServer) GetEvent(context.Context, *GetCalendarEventRequest) (*CalendarEventObject, error) {
	panic("implement me")
}

func (c CalendarServer) DeleteEvent(context.Context, *DeleteCalendarEventRequest) (*DeleteCalendarEventResponse, error) {
	panic("implement me")
}
