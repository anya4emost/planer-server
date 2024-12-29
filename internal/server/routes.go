package server

import (
	"github.com/anya4emost/planer-server/internal/server/router"
	"github.com/gofiber/fiber/v2"
)

func healthCheck() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return router.Ok(ctx, fiber.Map{})
	}
}

func (s *Server) SetupRoutes() {
	api := s.app.Group("/api")
	api.Get("/", healthCheck())
}
