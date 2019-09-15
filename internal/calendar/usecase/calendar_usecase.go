package usecase

import (
	"context"
	"github.com/Azimkhan/go-calendar-grpc/internal/calendar"
	"github.com/Azimkhan/go-calendar-grpc/internal/models"
	"time"
)

func NewCalendarUsecase(repository calendar.Repository, contextTimeout time.Duration) calendar.Usecase {
	return &CalendarUsecase{repository, contextTimeout}
}

type CalendarUsecase struct {
	repository     calendar.Repository
	contextTimeout time.Duration
}

// Fetch
func (c *CalendarUsecase) Fetch(ctx context.Context) ([]*models.CalendarEvent, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()
	return c.repository.Fetch(ctx)
}

// GetByID
func (c *CalendarUsecase) GetByID(ctx context.Context, id int64) (*models.CalendarEvent, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()
	return c.repository.GetByID(ctx, id)
}

// Update
func (c *CalendarUsecase) Update(ctx context.Context, record *models.CalendarEvent) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()
	record.UpdateTime = time.Now()
	return c.repository.Update(ctx, record)
}

// Store
func (c *CalendarUsecase) Store(ctx context.Context, record *models.CalendarEvent) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()
	record.CreateTime = time.Now()
	record.UpdateTime = record.CreateTime
	return c.repository.Store(ctx, record)
}

// Delete
func (c *CalendarUsecase) Delete(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()
	return c.repository.Delete(ctx, id)
}
