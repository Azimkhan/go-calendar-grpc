package service

import (
	"context"
	"github.com/Azimkhan/go-calendar-grpc/internal/domain/interfaces"
	"github.com/Azimkhan/go-calendar-grpc/internal/domain/models"
	"time"
)

func NewCalendarService(repository interfaces.Repository, contextTimeout time.Duration) *CalendarService {
	return &CalendarService{repository, contextTimeout}
}

type CalendarService struct {
	repository     interfaces.Repository
	contextTimeout time.Duration
}

// Fetch
func (c *CalendarService) Fetch(ctx context.Context) ([]*models.CalendarEvent, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()
	return c.repository.Fetch(ctx)
}

// GetByID
func (c *CalendarService) GetByID(ctx context.Context, id int64) (*models.CalendarEvent, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()
	return c.repository.GetByID(ctx, id)
}

// Update
func (c *CalendarService) Update(ctx context.Context, record *models.CalendarEvent) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()
	record.UpdateTime = time.Now().UTC()
	return c.repository.Update(ctx, record)
}

// Store
func (c *CalendarService) Store(ctx context.Context, record *models.CalendarEvent) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()
	record.CreateTime = time.Now().UTC()
	record.UpdateTime = record.CreateTime
	return c.repository.Store(ctx, record)
}

// Delete
func (c *CalendarService) Delete(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()
	return c.repository.Delete(ctx, id)
}
