package controller

import (
	"time"

	"github.com/anya4emost/planer-server/internal/model"
	"github.com/anya4emost/planer-server/internal/server/router"
	"github.com/anya4emost/planer-server/internal/services"
	"github.com/anya4emost/planer-server/pkg/util"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthController struct {
	s      *services.UserService
	secret string
}

func NewAuthController(s *services.UserService, secret string) *AuthController {
	return &AuthController{
		s:      s,
		secret: secret,
	}
}

func (c *AuthController) createToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"name": username,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
		"iat":  time.Now(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(c.secret))
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

	token, err := c.createToken(user.Username)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "registration error")
	}

	return router.Created(ctx, fiber.Map{
		"user":  user,
		"token": token,
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

	token, err := c.createToken(user.Username)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "login error")
	}

	return router.Ok(ctx, fiber.Map{
		"user":  user,
		"token": token,
	})
}

func (c *AuthController) Me(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["name"].(string)
	currentUser, err := c.s.GetByUsername(username)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid token")
	}

	token, err := c.createToken(currentUser.Username)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid token")
	}

	return router.Ok(ctx, fiber.Map{
		"user":  currentUser,
		"token": token,
	})
}
