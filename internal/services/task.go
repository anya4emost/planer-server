package services

import (
	"database/sql"
	"time"

	"github.com/anya4emost/planer-server/internal/model"
	"github.com/jmoiron/sqlx"
)

type TaskService struct {
	db *sqlx.DB
}

func NewTaskService(db *sqlx.DB) *TaskService {
	return &TaskService{
		db: db,
	}
}

func (s *TaskService) GetAllByUserId(id string) ([]model.Task, error) {
	tasks := []model.Task{}
	err := s.db.Select(&tasks, "select * from tasks where creator_id = $1", id)

	return tasks, err
}

func (s *TaskService) GetById(id string) (*model.Task, error) {
	task := model.Task{}
	err := s.db.Get(&task, "select * from tasks where id = $1", id)

	return &task, err
}

func (s *TaskService) Create(inputTask model.TaskInput) (*model.Task, error) {
	timeStart, tserr := time.Parse("15:04", inputTask.TimeStart)
	timeEnd, teerr := time.Parse("15:04", inputTask.TimeEnd)

	newTask := model.Task{
		IsDone:      inputTask.IsDone,
		Name:        inputTask.Name,
		Description: inputTask.Description,
		Date:        sql.NullString{String: inputTask.Date, Valid: len(inputTask.Date) > 0},
		TimeStart:   sql.NullTime{Time: timeStart, Valid: tserr == nil},
		TimeEnd:     sql.NullTime{Time: timeEnd, Valid: teerr == nil},
		TimeZone:    sql.NullString{String: inputTask.TimeZone, Valid: len(inputTask.TimeZone) > 0},
		Icon:        inputTask.Icon,
		Color:       inputTask.Color,
		Type:        inputTask.Type,
		CreatorId:   inputTask.CreatorId,
		DoerId:      inputTask.DoerId,
		AimId:       sql.NullString{String: inputTask.AimId, Valid: len(inputTask.AimId) > 0},
	}

	rows, err := s.db.NamedQuery(
		`insert into tasks
		(is_done, name, description, date, time_start, time_end, time_zone, icon, color, type, creator_id, doer_id, aim_id)
		values (:is_done, :name, :description, :date, :time_start, :time_end, :time_zone, :icon, :color, :type, :creator_id, :doer_id, :aim_id)`,
		newTask,
	)

	if err != nil {
		return nil, err
	}
	task := model.Task{}
	rows.Next()
	rows.StructScan(&task)

	return &task, err
}

func (s *TaskService) Update(inputTask model.TaskInput) (*model.Task, error) {
	timeStart, tserr := time.Parse("15:04", inputTask.TimeStart)
	timeEnd, teerr := time.Parse("15:04", inputTask.TimeEnd)

	taskToUpdate := model.Task{
		Id:          inputTask.Id,
		IsDone:      inputTask.IsDone,
		Name:        inputTask.Name,
		Description: inputTask.Description,
		Date:        sql.NullString{String: inputTask.Date, Valid: len(inputTask.Date) > 0},
		TimeStart:   sql.NullTime{Time: timeStart, Valid: tserr == nil},
		TimeEnd:     sql.NullTime{Time: timeEnd, Valid: teerr == nil},
		TimeZone:    sql.NullString{String: inputTask.TimeZone, Valid: len(inputTask.TimeZone) > 0},
		Icon:        inputTask.Icon,
		Color:       inputTask.Color,
		Type:        inputTask.Type,
		DoerId:      inputTask.DoerId,
		AimId:       sql.NullString{String: inputTask.AimId, Valid: len(inputTask.AimId) > 0},
	}

	rows, err := s.db.NamedQuery(
		`update tasks set 
		is_done=:is_done, name=:name, description=:description, icon=:icon, date=:date, time_start=:time_start, time_end=:time_end, time_zone=:time_zone, color=:color, type=:type, doer_id=:doer_id, aim_id=:aim_id
		where id=:id`,
		taskToUpdate,
	)

	if err != nil {
		return nil, err
	}
	task := model.Task{}
	rows.Next()
	rows.StructScan(&task)

	return &task, err
}

func (s *TaskService) Delete(taskId string) error {
	_, err := s.db.Exec("delete from tasks where id=$1", taskId)

	return err
}
