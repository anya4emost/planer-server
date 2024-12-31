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

func (s *Server) SetupRoutes(uc *controller.AuthController) {
	api := s.app.Group("/api")
	api.Get("/", healthCheck())

	api.Post("/login", uc.Login)
	api.Post("/register", uc.Register)
	api.Get("/me", middleware.Authenticate(s.jwtSecret), uc.Me)
}
