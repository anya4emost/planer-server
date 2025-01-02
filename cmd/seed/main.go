package main

import (
	"fmt"
	"log"

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
		Status:      "Analysis",
		Description: "first task description",
		Icon:        "",
		Color:       "",
		Type:        "Important",
		CreatorId:   adminUser.Id,
		DoerId:      adminUser.Id,
		AimId:       firstAim.Id,
	}

	_, errt := db.NamedQuery(
		`insert into tasks
		(status, description, icon, color, type, creator_id, doer_id, aim_id)
		values (:status, :description, :icon, :color, :type, :creator_id, :doer_id, :aim_id)`,
		newTask,
	)

	if errt != nil {
		log.Fatalf("Error inserting task: %v\n", errA)
	}

	firstTask := model.Task{}
	db.Get(&firstTask, "select * from tasks where creator_id=$1", adminUser.Id)

	fmt.Printf("\n%#v\n", firstTask)

	log.Printf("\nSuccessfully inserted data in tables")
}
