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

func (s *EventService) GetAllByUserId(id string) ([]model.Event, error) {
	events := []model.Event{}
	err := s.db.Select(&events, "select * from events where creator_id = $1", id)

	return events, err
}

func (s *EventService) Create(inputEvent model.EventInput) (*model.Event, error) {
	newEvent := model.Event{
		Name:             inputEvent.Name,
		Description:      inputEvent.Description,
		Icon:             inputEvent.Icon,
		Color:            inputEvent.Color,
		Category:         sql.NullString{String: inputEvent.Category, Valid: len(inputEvent.Category) > 0},
		Date:             inputEvent.Date,
		Duration:         inputEvent.Duration,
		TimeZone:         sql.NullString{String: inputEvent.TimeZone, Valid: len(inputEvent.TimeZone) > 0},
		Repit:            inputEvent.Repit,
		Remind:           inputEvent.Remind,
		TaskTracker:      inputEvent.TaskTracker,
		CustomCategoryId: sql.NullString{String: inputEvent.CustomCategoryId, Valid: len(inputEvent.CustomCategoryId) > 0},
		CreatorId:        inputEvent.CreatorId,
	}

	rows, err := s.db.NamedQuery(
		`insert into events
		(name, description, date, duration, time_zone, icon, color, creator_id, repit, remind, task_tracker)
		values (:name, :description, :date, :duration, :time_zone, :icon, :color, :creator_id, :repit, :remind, :task_tracker)`,
		newEvent)

	if err != nil {
		return nil, err
	}

	event := model.Event{}
	rows.Next()
	rows.StructScan(&event)

	return &event, err
}

func (s *EventService) Update(inputEvent model.EventInput) (*model.Event, error) {
	eventToUpdate := model.Event{
		Id:               inputEvent.Id,
		Name:             inputEvent.Name,
		Description:      inputEvent.Description,
		Icon:             inputEvent.Icon,
		Color:            inputEvent.Color,
		Category:         sql.NullString{String: inputEvent.Category, Valid: len(inputEvent.Category) > 0},
		Date:             inputEvent.Date,
		Duration:         inputEvent.Duration,
		TimeZone:         sql.NullString{String: inputEvent.TimeZone, Valid: len(inputEvent.TimeZone) > 0},
		Repit:            inputEvent.Repit,
		Remind:           inputEvent.Remind,
		TaskTracker:      inputEvent.TaskTracker,
		CustomCategoryId: sql.NullString{String: inputEvent.CustomCategoryId, Valid: len(inputEvent.CustomCategoryId) > 0},
		CreatorId:        inputEvent.CreatorId,
	}

	rows, err := s.db.NamedQuery(
		`update events set 
		name=:name, description=:description, icon=:icon, color=:color, date=:date, duration=:duration, time_zone=:time_zone, repit=:repit, remind=:remind, task_tracker=:task_tracker, custom_category_id=:custom_category_id
		where id=:id`,
		eventToUpdate,
	)

	if err != nil {
		return nil, err
	}
	event := model.Event{}
	rows.Next()
	rows.StructScan(&event)

	return &event, err
}

func (s *EventService) Delete(eventId string) error {
	_, err := s.db.Exec("delete from events where id=$1", eventId)

	return err
}
