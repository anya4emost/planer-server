package model

type AuthInput struct {
	Username string `json:"username"`
	Pasword  string `json:"password"`
}

type TaskInput struct {
	Id          string `json:"id"`
	Status      string `json:"status"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Color       string `json:"color"`
	Type        string `json:"type"`
	CreatorId   string `json:"creatorId"`
	DoerId      string `json:"doerId"`
	AimId       string `json:"aimId"`
	EventId     string `json:"eventId"`
}

type EventInput struct {
	Id               string `json:"id"`
	Category         string `json:"category"`
	Date             string `json:"date"`
	Time             string `json:"time"`
	Repit            string `json:"repit"`
	Remind           string `json:"remind"`
	CustomCategoryId string `json:"customCategoryId"`
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
