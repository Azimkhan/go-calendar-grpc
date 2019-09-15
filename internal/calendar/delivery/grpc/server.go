//go:generate protoc --go_out=plugins=grpc:. --proto_path=../../../../api/protobuf calendar_server.proto
package grpc

import (
	"context"
	"github.com/Azimkhan/go-calendar-grpc/internal/calendar"
	"github.com/Azimkhan/go-calendar-grpc/internal/models"
	"github.com/golang/protobuf/ptypes"
	"go.uber.org/zap"
	"time"
)

type CalendarServer struct {
	usecase calendar.Usecase
	logger  *zap.Logger
}

// Create server
func NewCalendarServer(logger *zap.Logger, usecase calendar.Usecase) *CalendarServer {
	return &CalendarServer{logger: logger, usecase: usecase}
}

// CreateEvent
func (s *CalendarServer) CreateEvent(c context.Context, r *CreateCalendarEventRequest) (*CalendarEventObject, error) {
	s.logger.Debug("CreateEvent()", zap.Any("event", r))
	ctx := context.Background()
	model := convertCreateCalendarEventRequestToModel(r)
	err := s.usecase.Store(ctx, model)
	if err != nil {
		return nil, err
	}
	return convertModelToCalendarEventObject(model), nil

}

// ListEvents
func (s *CalendarServer) ListEvents(c context.Context, r *ListCalendarEventRequest) (*ListCalendarEventResponse, error) {
	s.logger.Debug("ListEvents()")
	ctx := context.Background()
	events, err := s.usecase.Fetch(ctx)
	if err != nil {
		return nil, err
	}
	result := &ListCalendarEventResponse{
		Events: convertFromModelList(events),
	}
	return result, nil

}

// GetEvent
func (s *CalendarServer) GetEvent(c context.Context, r *GetCalendarEventRequest) (*CalendarEventObject, error) {
	s.logger.Debug("GetEvent()", zap.Int64("id", r.Id))
	ctx := context.Background()
	event, err := s.usecase.GetByID(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return convertModelToCalendarEventObject(event), nil
}

// DeleteEvent
func (s *CalendarServer) DeleteEvent(c context.Context, r *DeleteCalendarEventRequest) (*DeleteCalendarEventResponse, error) {
	s.logger.Debug("DeleteEvent()", zap.Int64("id", r.Id))
	ctx := context.Background()
	err := s.usecase.Delete(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &DeleteCalendarEventResponse{}, nil
}

func convertFromModelList(events []*models.CalendarEvent) (out []*CalendarEventObject) {
	out = make([]*CalendarEventObject, len(events))
	for i := 0; i < len(events); i++ {
		out[i] = convertModelToCalendarEventObject(events[i])
	}
	return
}

func convertFromModelEventType(eventType models.EventType) EventType {
	return EventType(eventType)
}
func convertToModelEventType(eventType EventType) models.EventType {
	return models.EventType(eventType)
}
func convertModelToCalendarEventObject(model *models.CalendarEvent) *CalendarEventObject {
	startTime, _ := ptypes.TimestampProto(model.StartTime)
	endTime, _ := ptypes.TimestampProto(model.EndTime)
	return &CalendarEventObject{
		Id:        model.Id,
		Name:      model.Name,
		Type:      convertFromModelEventType(model.EventType),
		StartTime: startTime,
		EndTime:   endTime,
	}
}
func convertCreateCalendarEventRequestToModel(r *CreateCalendarEventRequest) *models.CalendarEvent {
	startTime, _ := ptypes.Timestamp(r.StartTime)
	endTime, _ := ptypes.Timestamp(r.EndTime)
	return &models.CalendarEvent{
		Id:         0,
		Name:       r.Name,
		EventType:  convertToModelEventType(r.Type),
		StartTime:  startTime,
		EndTime:    endTime,
		CreateTime: time.Time{},
		UpdateTime: time.Time{},
	}
}
