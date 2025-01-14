package server

import (
	"github.com/anya4emost/planer-server/internal/config"
	"github.com/anya4emost/planer-server/internal/controller"
	"github.com/anya4emost/planer-server/internal/database"
	"github.com/anya4emost/planer-server/internal/server/router/response"
	"github.com/anya4emost/planer-server/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	app       *fiber.App
	port      string
	jwtSecret string
	db        *sqlx.DB
}

func (s *Server) Start() error {
	sessionService := services.NewSessionService(s.db)
	userService := services.NewUserService(s.db)
	eventService := services.NewEventService(s.db)
	taskService := services.NewTaskService(s.db)
	aimService := services.NewAimService(s.db)

	authController := controller.NewAuthController(userService, sessionService, s.jwtSecret)
	taskController := controller.NewTasksController(taskService)
	aimController := controller.NewAimsController(aimService)
	eventsController := controller.NewEventsController(eventService)

	s.SetupRoutes(authController, taskController, aimController, eventsController)
	return s.app.Listen(s.port)
}

func (s *Server) Stop() error {
	s.db.Close()
	return s.app.Shutdown()
}

func NewServer(cfg *config.Config) *Server {
	app := fiber.New(fiber.Config{
		ErrorHandler: response.DefaultErrorHandler,
	})

	port := ":" + cfg.Port
	db := database.Connect(cfg.DatabaseUrl)

	return &Server{
		app:       app,
		port:      port,
		jwtSecret: cfg.JwtSecret,
		db:        db,
	}
}
