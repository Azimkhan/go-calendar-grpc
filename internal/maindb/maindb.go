package maindb

import (
	"context"
	"github.com/Azimkhan/go-calendar-grpc/internal/domain/errors"
	"github.com/Azimkhan/go-calendar-grpc/internal/domain/models"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

func NewPgEventRepository(dsn string) (*PgEventRepository, error) {
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &PgEventRepository{db: db}, nil

}

type PgEventRepository struct {
	db *sqlx.DB
}

func (p *PgEventRepository) Fetch(ctx context.Context) ([]*models.CalendarEvent, error) {
	rows, err := p.db.QueryxContext(ctx, "select id, name, event_type, start_time, end_time, create_time, update_time from "+
		"events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []*models.CalendarEvent
	for rows.Next() {
		event := &models.CalendarEvent{}
		scanErr := rows.StructScan(event)
		if scanErr != nil {
			return nil, scanErr
		}
		out = append(out, event)
	}
	return out, nil
}

func (p *PgEventRepository) GetByID(ctx context.Context, id int64) (*models.CalendarEvent, error) {
	row := p.db.QueryRowxContext(ctx, "select id, name, event_type, start_time, end_time, create_time, update_time from "+
		"events where id = $1", id)

	if row.Err() != nil {
		return nil, row.Err()
	}
	event := &models.CalendarEvent{}
	scanErr := row.StructScan(event)
	if scanErr != nil {
		return nil, scanErr
	}
	return event, nil

}

func (p *PgEventRepository) Update(ctx context.Context, e *models.CalendarEvent) error {
	result, err := p.db.ExecContext(ctx, "update events set name=$1, event_type=$2, start_time=$3, end_time=$4, create_time=$5, update_time=$6 "+
		"where id=$7",
		e.Name,
		e.EventType,
		e.StartTime,
		e.EndTime,
		e.CreateTime,
		e.UpdateTime,
		e.Id,
	)
	if err != nil {
		return err
	}
	if r, _ := result.RowsAffected(); r < 1 {
		return errors.EventNotFound
	}
	return nil
}

func (p *PgEventRepository) Store(ctx context.Context, e *models.CalendarEvent) error {
	row := p.db.QueryRowxContext(ctx, "insert into events(name, event_type, start_time, end_time, create_time, update_time) "+
		"values ($1, $2, $3, $4, $5, $6) returning id",
		e.Name,
		e.EventType,
		e.StartTime,
		e.EndTime,
		e.CreateTime,
		e.UpdateTime,
	)
	if row.Err() != nil {
		return row.Err()
	}
	var id int64
	scanErr := row.Scan(&id)
	if scanErr != nil {
		return scanErr
	}
	e.Id = id
	return nil
}

func (p *PgEventRepository) Delete(ctx context.Context, id int64) error {
	result, err := p.db.ExecContext(ctx, "delete from events where id=$1",
		id,
	)
	if err != nil {
		return err
	}
	if r, _ := result.RowsAffected(); r < 1 {
		return errors.EventNotFound
	}
	return nil
}
