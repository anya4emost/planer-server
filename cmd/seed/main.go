package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/anya4emost/planer-server/internal/config"
	"github.com/anya4emost/planer-server/internal/database"
	"github.com/anya4emost/planer-server/internal/model"
	"github.com/anya4emost/planer-server/pkg/util"
)

func main() {
	cfg := config.Load()
	db := database.Connect(cfg.DatabaseUrl)
	password, err := util.HashPassword("password")

	if err != nil {
		log.Fatalf("Error generating password: %v\n", err)
	}

	user := model.User{
		Username: "admin",
		Password: password,
	}

	_, errU := db.NamedQuery(
		`insert into users (username, password)
		values (:username, :password)`,
		user,
	)

	adminUser := model.User{}
	db.Get(&adminUser, "select * from users where username=$1", "admin")

	fmt.Printf("\n%#v\n", adminUser)

	if errU != nil {
		log.Fatalf("Error inserting users: %v\n", errU)
	}

	aims := []model.Aim{
		{
			Name:   "Переехать в америку",
			UserId: adminUser.Id,
		},
		{
			Name:   "Дойти до зарплаты 350т",
			UserId: adminUser.Id,
		},
		{
			Name:   "Сделать ремонт в доме",
			UserId: adminUser.Id,
		},
	}

	_, errA := db.NamedQuery(
		`insert into aims (user_id, name)
		values(:user_id, :name)`,
		aims,
	)
	if errA != nil {
		log.Fatalf("Error inserting aims: %v\n", errA)
	}

	firstAim := model.Aim{}
	db.Get(&firstAim, "select * from aims where name=$1", "Переехать в америку")

	fmt.Printf("\n%#v\n", firstAim)

	newTask := model.Task{
		IsDone:      false,
		Name:        "task name",
		Description: "first task description",
		Date:        sql.NullString{String: "02-20-2025", Valid: true},
		TimeStart:   sql.NullTime{Time: time.Now(), Valid: true},
		TimeEnd:     sql.NullTime{Time: time.Now().Add(time.Hour), Valid: true},
		Icon:        "",
		Color:       "",
		Type:        "Important",
		CreatorId:   adminUser.Id,
		DoerId:      adminUser.Id,
		AimId:       sql.NullString{String: firstAim.Id, Valid: true},
	}

	_, errt := db.NamedQuery(
		`insert into tasks
		(is_done, name, description, date, time_start, time_end, icon, color, type, creator_id, doer_id, aim_id)
		values (:is_done, :name, :description, :date, :time_start, :time_end, :icon, :color, :type, :creator_id, :doer_id, :aim_id)`,
		newTask,
	)

	if errt != nil {
		log.Fatalf("Error inserting task: %v\n", errt)
	}

	firstTask := model.Task{}
	db.Get(&firstTask, "select * from tasks where creator_id=$1", adminUser.Id)

	fmt.Printf("\n%#v\n", firstTask)

	newEvent := model.Event{
		Name:        "event name",
		Description: "first event description",
		Date:        "02-20-2025",
		Duration:    120,
		Icon:        "",
		Color:       "",
		Category:    sql.NullString{String: "", Valid: true},
		Repit:       "EveryDay",
		Remind:      "TenMinBefore",
		TaskTracker: false,
		CreatorId:   adminUser.Id,
	}

	_, erre := db.NamedQuery(
		`insert into events
		(name, description, date, duration, icon, color, creator_id, repit, remind, task_tracker)
		values (:name, :description, :date, :duration, :icon, :color, :creator_id, :repit, :remind, :task_tracker)`,
		newEvent,
	)

	if erre != nil {
		log.Fatalf("Error inserting event: %v\n", erre)
	}

	firstEvent := model.Event{}
	db.Get(&firstEvent, "select * from events where creator_id=$1", adminUser.Id)

	fmt.Printf("\n%#v\n", firstEvent)

	log.Printf("\nSuccessfully inserted data in tables")
}
