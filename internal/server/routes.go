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

	// s.app.Use(encryptcookie.New(encryptcookie.Config{
	// 	Key: "secret-key-for-cookie-encription",
	// }))

	api := s.app.Group("/api")
	api.Get("/", healthCheck())

	authGroup := api.Group("auth")
	authGroup.Post("/login", authController.Login)
	authGroup.Post("/register", authController.Register)
	authGroup.Get("/me", middleware.AccessTokenVerification(s.jwtSecret), authController.Me)
	authGroup.Post("/logout", authController.Logout)
	authGroup.Post("/refresh", authController.Refresh)

	tasksApi := api.Group("/tasks")
	tasksApi.Use(middleware.AccessTokenVerification(s.jwtSecret))
	tasksApi.Get("/", taskController.GetTasks)
	tasksApi.Post("/", taskController.CreateTask)

	// POST "/api/tasks/${taskid}/event"
	tasksApi.Post("/:taskid/event", eventsController.CreateEvent)

	aimsApi := api.Group("/aims")
	aimsApi.Use(middleware.AccessTokenVerification(s.jwtSecret))
	aimsApi.Get("/", aimController.GetAims)
	aimsApi.Post("/", aimController.CreateAim)
}
