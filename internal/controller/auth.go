package controller

import (
	"github.com/anya4emost/planer-server/internal/model"
	"github.com/anya4emost/planer-server/internal/server/router"
	"github.com/anya4emost/planer-server/internal/services"
	"github.com/anya4emost/planer-server/pkg/util"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	s *services.UserService
}

func NewAuthController(s *services.UserService) *AuthController {
	return &AuthController{
		s: s,
	}
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	input := model.AuthInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	password, err := util.HashPassword(input.Pasword)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid password")
	}

	user, err := c.s.Create(input.Username, password)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "registration error")
	}

	return router.Created(ctx, fiber.Map{
		"user": user,
	})
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	input := model.AuthInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	user, err := c.s.GetByUsername(input.Username)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "login error")
	}
	if !util.CheckPassword(input.Pasword, user.Password) {
		return fiber.NewError(fiber.StatusUnauthorized, "login error")
	}

	return router.Ok(ctx, fiber.Map{
		"user": user,
	})
}
