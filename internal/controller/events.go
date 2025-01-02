package controller

import (
	"github.com/anya4emost/planer-server/internal/model"
	"github.com/anya4emost/planer-server/internal/server/router/response"
	"github.com/anya4emost/planer-server/internal/services"
	"github.com/gofiber/fiber/v2"
)

type EventsController struct {
	s *services.EventService
}

func NewEventsController(s *services.EventService) *EventsController {
	return &EventsController{
		s: s,
	}
}

func (c *EventsController) CreateEvent(ctx *fiber.Ctx) error {
	input := model.EventInput{
		TaskId: ctx.Params("taskid"),
	}

	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}

	event, err := c.s.Create(input)
	if err != nil {
		return response.ErrorBadRequest(err)
	}

	return response.Ok(ctx, event)
}
