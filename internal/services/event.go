package services

import (
	"fmt"
	"slices"
	"strings"

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

	comand := "insert into events"

	colsToInsert := []string{
		"category",
		"date",
		"time",
		"repit",
		"remind",
		"task_id",
	}

	if len(newEvent.CustomCategoryId) > 0 {
		colsToInsert = slices.Insert(colsToInsert, len(colsToInsert), "custom_category_id")
	}

	params := "(" + strings.Join(colsToInsert, ", ") + ")"
	values := "values (:" + strings.Join(colsToInsert, ", :") + ")"

	fmt.Println(comand + params + values)

	rows, err := s.db.NamedQuery(comand+params+values, newEvent)

	if err != nil {
		return nil, err
	}

	event := model.Event{}
	rows.Next()
	rows.StructScan(&event)

	return &event, err
}
