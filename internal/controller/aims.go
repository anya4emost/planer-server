package controller

import (
	"github.com/anya4emost/planer-server/internal/model"
	"github.com/anya4emost/planer-server/internal/server/router/response"
	"github.com/anya4emost/planer-server/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AimsController struct {
	s *services.AimService
}

func NewAimsController(s *services.AimService) *AimsController {
	return &AimsController{
		s: s,
	}
}

func (c *AimsController) GetAims(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userid := claims["sub"].(string)

	aims, err := c.s.GetAllByUserId(userid)

	if err != nil {
		return response.ErrorBadRequest(err)
	}

	return response.Ok(ctx, aims)
}

func (c *AimsController) CreateAim(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userid := claims["sub"].(string)

	input := model.AimInput{
		UserId: userid,
	}
	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}

	aim, err := c.s.Create(input)
	if err != nil {
		return response.ErrorBadRequest(err)
	}

	return response.Ok(ctx, aim)
}
