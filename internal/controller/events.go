package controller

import (
	"slices"

	"github.com/anya4emost/planer-server/internal/model"
	"github.com/anya4emost/planer-server/internal/server/router/response"
	"github.com/anya4emost/planer-server/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type EventsController struct {
	s *services.EventService
}

func NewEventsController(s *services.EventService) *EventsController {
	return &EventsController{
		s: s,
	}
}

func (c *EventsController) GetEvents(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userid := claims["sub"].(string)

	events, err := c.s.GetAllByUserId(userid)

	if err != nil {
		return response.ErrorBadRequest(err)
	}

	outputRange := []fiber.Map{}

	for _, event := range events {
		outputRange = slices.Insert(outputRange, len(outputRange), fiber.Map{
			"id":               event.Id,
			"name":             event.Name,
			"description":      event.Description,
			"icon":             event.Icon,
			"color":            event.Color,
			"category":         event.Category.String,
			"date":             event.Date,
			"duration":         event.Duration,
			"timeZone":         event.TimeZone.String,
			"repit":            event.Repit,
			"remind":           event.Remind,
			"taskTracker":      event.TaskTracker,
			"customCategoryId": event.CustomCategoryId.String,
			"creatorId":        event.CreatorId,
		})
	}

	return response.Ok(ctx, outputRange)
}

func (c *EventsController) CreateEvent(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userid := claims["sub"].(string)

	input := model.EventInput{
		CreatorId: userid,
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

func (c *EventsController) UpdateEvent(ctx *fiber.Ctx) error {
	eventId := ctx.Params("eventid")
	input := model.EventInput{}

	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}

	input.Id = eventId

	event, err := c.s.Update(input)
	if err != nil {
		return response.ErrorBadRequest(err)
	}

	return response.Ok(ctx, event)
}

func (c *EventsController) DeleteEvent(ctx *fiber.Ctx) error {
	err := c.s.Delete(ctx.Params("eventid"))

	if err != nil {
		return response.ErrorBadRequest(err)
	}

	return response.Response(ctx, 204, fiber.Map{})
}
