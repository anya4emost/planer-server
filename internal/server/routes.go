package server

import (
	"github.com/anya4emost/planer-server/internal/controller"
	"github.com/anya4emost/planer-server/internal/server/router/middleware"
	"github.com/anya4emost/planer-server/internal/server/router/response"
	"github.com/gofiber/fiber/v2"
)

func healthCheck() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return response.Ok(ctx, fiber.Map{})
	}
}

func (s *Server) SetupRoutes(
	authController *controller.AuthController,
	taskController *controller.TasksController,
	aimController *controller.AimsController,
	eventsController *controller.EventsController) {

	api := s.app.Group("/api")
	api.Get("/", healthCheck())

	api.Post("/login", authController.Login)
	api.Post("/register", authController.Register)
	api.Get("/me", middleware.Authenticate(s.jwtSecret), authController.Me)

	tasksApi := api.Group("/tasks")
	tasksApi.Use(middleware.Authenticate(s.jwtSecret))
	tasksApi.Get("/", taskController.GetTasks)
	tasksApi.Post("/", taskController.CreateTask)

	// POST "/api/tasks/${taskid}/event"
	tasksApi.Post("/:taskid/event", eventsController.CreateEvent)

	aimsApi := api.Group("/aims")
	aimsApi.Use(middleware.Authenticate(s.jwtSecret))
	aimsApi.Get("/", aimController.GetAims)
	aimsApi.Post("/", aimController.CreateAim)
}
