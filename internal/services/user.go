package services

import (
	"github.com/anya4emost/planer-server/internal/model"
	"github.com/jmoiron/sqlx"
)

type UserService struct {
	db *sqlx.DB
}

func NewUserService(db *sqlx.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (s *UserService) GetByUsername(username string) (model.User, error) {
	user := model.User{}
	err := s.db.Get(&user, "select * from users where username = $1", username)
	return user, err
}

func (s *UserService) GetById(id string) (model.User, error) {
	user := model.User{}
	err := s.db.Get(&user, "select * from users where id = $1", id)
	return user, err
}

func (s *UserService) Create(username, password string) (*model.User, error) {
	rows, err := s.db.Queryx(
		`insert into users (username, password)
		values ($1, $2)
		returning *`,
		username,
		password,
	)

	if err != nil {
		return nil, err
	}

	user := model.User{}
	rows.Next()
	err = rows.StructScan(&user)
	return &user, err
}
