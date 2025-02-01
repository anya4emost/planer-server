package model

type AuthInput struct {
	Username string `json:"username"`
	Pasword  string `json:"password"`
}

type TaskInput struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	IsDone      bool   `db:"is_done" json:"isDone"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Color       string `json:"color"`
	Type        string `json:"type"`
	Date        string `json:"date"`
	TimeStart   string `db:"time_start" json:"timeStart"`
	TimeEnd     string `db:"time_end" json:"timeEnd"`
	TimeZone    string `db:"time_zone" json:"timeZone"`
	CreatorId   string `json:"creatorId"`
	DoerId      string `json:"doerId"`
	AimId       string `json:"aimId"`
}

type EventInput struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	Icon             string `json:"icon"`
	Color            string `json:"color"`
	Category         string `json:"category"`
	Date             string `json:"date"`
	Duration         int    `json:"duration"`
	TimeZone         string `db:"time_zone" json:"timeZone"`
	Repit            string `json:"repit"`
	Remind           string `json:"remind"`
	TaskTracker      bool   `db:"task_tracker" json:"taskTracker"`
	CustomCategoryId string `json:"customCategoryId"`
	CreatorId        string `json:"creatorId"`
}

type AimInput struct {
	Id     string `json:"id"`
	UserId string `json:"userId"`
	Name   string `json:"name"`
}

type CustomCategoryInput struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type GroupInput struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UserGroupInput struct {
	Id      string `json:"id"`
	UserId  string `json:"userId"`
	GroupId string `json:"groupId"`
}
