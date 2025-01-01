package services

import (
	"github.com/anya4emost/planer-server/internal/model"
	"github.com/jmoiron/sqlx"
)

type AimService struct {
	db *sqlx.DB
}

func NewAimService(db *sqlx.DB) *AimService {
	return &AimService{
		db: db,
	}
}

func (s *AimService) GetAllByUserId(id string) ([]model.Aim, error) {
	aims := []model.Aim{}
	err := s.db.Select(&aims, "select * from aims where user_id = $1", id)

	return aims, err
}

func (s *AimService) GetById(id string) (*model.Aim, error) {
	aim := model.Aim{}
	err := s.db.Get(&aim, "select * from aims where id = $1", id)

	return &aim, err
}

func (s *AimService) Create(inputAim model.AimInput) (*model.Aim, error) {
	newAim := model.Aim{
		UserId: inputAim.UserId,
		Name:   inputAim.Name,
	}

	rows, err := s.db.NamedQuery(
		`insert into aims (user_id, name)
		values (:user_id, :name)`,
		newAim,
	)

	if err != nil {
		return nil, err
	}
	aim := model.Aim{}
	rows.Next()
	rows.StructScan(&aim)

	return &aim, err
}
