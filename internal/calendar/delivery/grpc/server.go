//go:generate protoc --go_out=plugins=grpc:. --proto_path=../../../../api/protobuf calendar_server.proto
package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"go.uber.org/zap"
	"time"
)

type CalendarServer struct {
	Logger *zap.Logger
}

func (s *CalendarServer) CreateEvent(context.Context, *CreateCalendarEventRequest) (*CalendarEventObject, error) {
	panic("implement me")
}

func (s *CalendarServer) ListEvents(context.Context, *ListCalendarEventRequest) (*ListCalendarEventResponse, error) {
	s.Logger.Info("ListEvents()")
	start := time.Now()
	end := start.Add(time.Hour)
	startProto, _ := ptypes.TimestampProto(start)
	endProto, _ := ptypes.TimestampProto(end)

	objects := []*CalendarEventObject{
		{Id: 1, Name: "Clean the room", Type: EventType_TASK, StartTime: startProto, EndTime: endProto},
		{Id: 2, Name: "Meeting with Antony", Type: EventType_MEETING, StartTime: startProto, EndTime: endProto},
	}
	result := &ListCalendarEventResponse{Events: objects}
	return result, nil

}

func (s *CalendarServer) GetEvent(context.Context, *GetCalendarEventRequest) (*CalendarEventObject, error) {
	panic("implement me")
}

func (s *CalendarServer) DeleteEvent(context.Context, *DeleteCalendarEventRequest) (*DeleteCalendarEventResponse, error) {
	panic("implement me")
}
