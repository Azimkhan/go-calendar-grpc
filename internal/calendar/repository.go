package calendar

import (
	"context"
	"github.com/Azimkhan/go-calendar-grpc/internal/models"
)

type Repository interface {
	Fetch(ctx context.Context) ([]*models.CalendarEvent, error)
	GetByID(ctx context.Context, id int64) (*models.CalendarEvent, error)
	Update(ctx context.Context, ar *models.CalendarEvent) error
	Store(context.Context, *models.CalendarEvent) error
	Delete(ctx context.Context, id int64) error
}
