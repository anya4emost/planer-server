package services

import (
	"database/sql"

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
	newTask := model.Task{
		Status:      inputTask.Status,
		Name:        inputTask.Name,
		Description: inputTask.Description,
		Icon:        inputTask.Icon,
		Color:       inputTask.Color,
		Type:        inputTask.Type,
		CreatorId:   inputTask.CreatorId,
		DoerId:      inputTask.DoerId,
		AimId:       sql.NullString{String: inputTask.AimId, Valid: len(inputTask.AimId) > 0},
	}

	rows, err := s.db.NamedQuery(
		`insert into tasks
		(status, name, description, icon, color, type, creator_id, doer_id, aim_id)
		values (:status, :name, :description, :icon, :color, :type, :creator_id, :doer_id, :aim_id)`,
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

	taskToUpdate := model.Task{
		Id:          inputTask.Id,
		Status:      inputTask.Status,
		Name:        inputTask.Name,
		Description: inputTask.Description,
		Icon:        inputTask.Icon,
		Color:       inputTask.Color,
		Type:        inputTask.Type,
		DoerId:      inputTask.DoerId,
		AimId:       sql.NullString{String: inputTask.AimId, Valid: len(inputTask.AimId) > 0},
	}

	rows, err := s.db.NamedQuery(
		`update tasks set 
		status=:status, name=:name, description=:description, icon=:icon, color=:color, type=:type, doer_id=:doer_id, aim_id=:aim_id
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
