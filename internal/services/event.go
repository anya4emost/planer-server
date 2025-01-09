package services

import (
	"database/sql"

	"github.com/anya4emost/planer-server/internal/model"
	"github.com/jmoiron/sqlx"
)

type EventService struct {
	db *sqlx.DB
}

func NewEventService(db *sqlx.DB) *EventService {
	return &EventService{
		db: db,
	}
}

func (s *EventService) Create(inputEvent model.EventInput) (*model.Event, error) {
	newEvent := model.Event{
		Category:         inputEvent.Category,
		Date:             inputEvent.Date,
		Time:             inputEvent.Time,
		Repit:            inputEvent.Repit,
		Remind:           inputEvent.Remind,
		CustomCategoryId: sql.NullString{String: inputEvent.CustomCategoryId, Valid: len(inputEvent.CustomCategoryId) > 0},
		TaskId:           inputEvent.TaskId,
	}

	rows, err := s.db.NamedQuery(
		`insert into events
		(category, date, time, repit, remind, task_id, custom_category_id)
		values (:category, :date, :time, :repit, :remind, :task_id, :custom_category_id)`,
		newEvent)

	if err != nil {
		return nil, err
	}

	event := model.Event{}
	rows.Next()
	rows.StructScan(&event)

	return &event, err
}
