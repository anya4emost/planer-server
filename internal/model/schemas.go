package model

import (
	"database/sql"
	"time"
)

type User struct {
	Id        string     `db:"id" json:"id"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
	Username  string     `db:"username" json:"username"`
	Password  string     `json:"-"`
}

type Aim struct {
	Id     string `json:"id"`
	UserId string `db:"user_id" json:"userId"`
	Name   string `json:"name"`
}

type CustomCategory struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Group struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UserGroup struct {
	Id      string `json:"id"`
	UserId  string `db:"user_id" json:"userId"`
	GroupId string `db:"group_id" json:"groupId"`
}

type Task struct {
	Id          string         `json:"id"`
	Name        string         `json:"name"`
	IsDone      bool           `db:"is_done" json:"isDone"`
	Description string         `json:"description"`
	Icon        string         `json:"icon"`
	Color       string         `json:"color"`
	Type        string         `json:"type"`
	Date        sql.NullString `json:"date"`
	TimeStart   sql.NullTime   `db:"time_start" json:"timeStart"`
	TimeEnd     sql.NullTime   `db:"time_end" json:"timeEnd"`
	TimeZone    sql.NullString `db:"time_zone" json:"timeZone"`
	CreatorId   string         `db:"creator_id" json:"creatorId"`
	DoerId      string         `db:"doer_id" json:"doerId"`
	AimId       sql.NullString `db:"aim_id" json:"aimId"`
}

type Event struct {
	Id               string         `json:"id"`
	Name             string         `json:"name"`
	Description      string         `json:"description"`
	Icon             string         `json:"icon"`
	Color            string         `json:"color"`
	Category         sql.NullString `json:"category"`
	Date             string         `json:"date"`
	Duration         int            `json:"duration"`
	TimeZone         sql.NullString `db:"time_zone" json:"timeZone"`
	Repit            string         `json:"repit"`
	Remind           string         `json:"remind"`
	TaskTracker      bool           `db:"task_tracker" json:"taskTracker"`
	CustomCategoryId sql.NullString `db:"custom_category_id" json:"customCategoryId"`
	CreatorId        string         `db:"creator_id" json:"creatorId"`
}

type Session struct {
	RefreshToken string `db:"refresh_token" json:"refreshToken"`
	UserId       string `db:"user_id" json:"userId"`
	CreatedAt    string `db:"created_at" json:"createdAt"`
	ExpiresAt    string `db:"expires_at" json:"expiresAt"`
	Family       string `db:"family" json:"family"`
	Revoked      bool   `db:"revoked" json:"revoked"`
}
