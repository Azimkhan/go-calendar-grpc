package repository

import (
	"context"
	"github.com/Azimkhan/go-calendar-grpc/internal/domain/errors"
	"github.com/Azimkhan/go-calendar-grpc/internal/domain/interfaces"
	"github.com/Azimkhan/go-calendar-grpc/internal/domain/models"
	"sync"
)

// Create new map calendar repository
func NewMapCalendarRepository() interfaces.Repository {
	return &MapCalendarRepository{db: make(map[int64]*models.CalendarEvent), idSeq: 1}
}

type MapCalendarRepository struct {
	wLock sync.Mutex
	idSeq int64
	db    map[int64]*models.CalendarEvent
}

// Fetch
func (r *MapCalendarRepository) Fetch(ctx context.Context) ([]*models.CalendarEvent, error) {
	out := make([]*models.CalendarEvent, len(r.db))
	i := 0
	for _, v := range r.db {
		out[i] = v
		i++
	}
	return out, nil
}

// GetByID
func (r *MapCalendarRepository) GetByID(ctx context.Context, id int64) (*models.CalendarEvent, error) {
	if e, ok := r.db[id]; ok {
		return e, nil
	}
	return nil, errors.EventNotFound
}

// Update
func (r *MapCalendarRepository) Update(ctx context.Context, record *models.CalendarEvent) error {
	r.wLock.Lock()
	defer r.wLock.Unlock()
	r.db[record.Id] = record
	return nil
}

// Store
func (r *MapCalendarRepository) Store(ctx context.Context, record *models.CalendarEvent) error {
	r.wLock.Lock()
	defer r.wLock.Unlock()
	record.Id = r.idSeq
	r.idSeq++
	r.db[record.Id] = record
	return nil
}

// Delete*//
func (r *MapCalendarRepository) Delete(ctx context.Context, id int64) error {
	r.wLock.Lock()
	defer r.wLock.Unlock()
	if _, ok := r.db[id]; ok {
		delete(r.db, id)
		return nil
	}
	return errors.EventNotFound

}
