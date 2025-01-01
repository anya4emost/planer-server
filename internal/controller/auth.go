package controller

import (
	"time"

	"github.com/anya4emost/planer-server/internal/model"
	"github.com/anya4emost/planer-server/internal/server/router/response"
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

func (c *AuthController) createToken(userId string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
		"iat": time.Now(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(c.secret))
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	input := model.AuthInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}

	password, err := util.HashPassword(input.Pasword)
	if err != nil {
		return response.ErrorBadRequest(err)
	}

	user, err := c.s.Create(input.Username, password)
	if err != nil {
		return response.ErrorUnauthorized(err, "Registration error")
	}

	token, err := c.createToken(user.Id)
	if err != nil {
		return response.ErrorUnauthorized(err, "Registration error")
	}

	return response.Created(ctx, fiber.Map{
		"user":  user,
		"token": token,
	})
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	input := model.AuthInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}

	user, err := c.s.GetByUsername(input.Username)
	if err != nil {
		return response.ErrorUnauthorized(err, "Login error")
	}
	if !util.CheckPassword(input.Pasword, user.Password) {
		return response.ErrorUnauthorized(err, "Login error")
	}

	token, err := c.createToken(user.Id)
	if err != nil {
		return response.ErrorUnauthorized(err, "Login error")
	}

	return response.Ok(ctx, fiber.Map{
		"user":  user,
		"token": token,
	})
}

func (c *AuthController) Me(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userid := claims["sub"].(string)

	currentUser, err := c.s.GetById(userid)

	if err != nil {
		return response.ErrorUnauthorized(err, "Invalid credentials")
	}

	token, err := c.createToken(userid)
	if err != nil {
		return response.ErrorUnauthorized(err, "Invalid credentials")
	}

	return response.Ok(ctx, fiber.Map{
		"user":  currentUser,
		"token": token,
	})
}
