package services

import (
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
		CustomCategoryId: inputEvent.CustomCategoryId,
		TaskId:           inputEvent.TaskId,
	}

	rows, err := s.db.NamedQuery(
		`insert into events
		(category, date, time, repit, remind, custom_category_id, task_id)
		values (:category, :date, :time, :repit, :remind, :custom_category_id, :task_id)`,
		newEvent,
	)

	if err != nil {
		return nil, err
	}

	event := model.Event{}
	rows.Next()
	rows.StructScan(&event)

	return &event, err
}
