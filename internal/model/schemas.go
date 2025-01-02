package model

import "time"

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

type Task struct {
	Id          string `json:"id"`
	Status      string `json:"status"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Color       string `json:"color"`
	Type        string `json:"type"`
	CreatorId   string `db:"creator_id" json:"creatorId"`
	DoerId      string `db:"doer_id" json:"doerId"`
	AimId       string `db:"aim_id" json:"aimId"`
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

type Event struct {
	Id               string `json:"id"`
	Category         string `json:"category"`
	Date             string `json:"date"`
	Time             string `json:"time"`
	Repit            string `json:"repit"`
	Remind           string `json:"remind"`
	CustomCategoryId string `db:"custom_category_id" json:"customCategoryId"`
	TaskId           string `db:"task_id" json:"taskId"`
}
