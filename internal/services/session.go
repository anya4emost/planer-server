package services

import (
	"github.com/anya4emost/planer-server/internal/model"
	"github.com/jmoiron/sqlx"
)

type SessionService struct {
	db *sqlx.DB
}

func NewSessionService(db *sqlx.DB) *SessionService {
	return &SessionService{
		db: db,
	}
}

func (s *SessionService) GetByName(refreshToken string) (model.Session, error) {
	session := model.Session{}
	err := s.db.Get(&session, "select * from sessions where refresh_token=$1", refreshToken)
	return session, err
}

func (s *SessionService) DeleteAllFamily(family string) error {
	_, err := s.db.Exec("delete from sessions where family=$1", family)

	return err
}

func (s *SessionService) Create(session model.Session) error {
	_, err := s.db.Exec(
		`insert into sessions (refresh_token, user_id, created_at, expires_at, family, revoked)
		values($1, $2, $3, $4, $5, $6)`,
		session.RefreshToken,
		session.UserId,
		session.CreatedAt,
		session.ExpiresAt,
		session.Family,
		session.Revoked,
	)

	return err
}

func (s *SessionService) MakeRevoked(session model.Session) error {
	_, err := s.db.Exec(
		`update sessions set revoked=true where refresh_token=$1`,
		session.RefreshToken,
	)

	return err
}
