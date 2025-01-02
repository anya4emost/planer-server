package services

import (
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
		Description: inputTask.Description,
		Icon:        inputTask.Icon,
		Color:       inputTask.Color,
		Type:        inputTask.Type,
		CreatorId:   inputTask.CreatorId,
		DoerId:      inputTask.DoerId,
		AimId:       inputTask.AimId,
	}

	rows, err := s.db.NamedQuery(
		`insert into tasks
		(status, description, icon, color, type, creator_id, doer_id, aim_id)
		values (:status, :description, :icon, :color, :type, :creator_id, :doer_id, :aim_id)`,
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

func (s *TaskService) Update(task model.TaskInput) {}

func (s *TaskService) Delete(task model.TaskInput) {}
