package controller

import (
	"github.com/anya4emost/planer-server/internal/model"
	"github.com/anya4emost/planer-server/internal/server/router/response"
	"github.com/anya4emost/planer-server/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type TasksController struct {
	s *services.TaskService
}

func NewTasksController(s *services.TaskService) *TasksController {
	return &TasksController{
		s: s,
	}
}

func (c *TasksController) GetTasks(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userid := claims["sub"].(string)

	tasks, err := c.s.GetAllByUserId(userid)

	if err != nil {
		return response.ErrorBadRequest(err)
	}

	return response.Ok(ctx, tasks)
}

func (c *TasksController) CreateTask(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userid := claims["sub"].(string)

	input := model.TaskInput{
		CreatorId: userid,
	}

	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}

	if input.DoerId == "" {
		input.DoerId = userid
	}

	task, err := c.s.Create(input)
	if err != nil {
		return response.ErrorBadRequest(err)
	}

	return response.Ok(ctx, task)
}
