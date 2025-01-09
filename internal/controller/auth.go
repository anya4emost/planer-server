package controller

import (
	"strconv"
	"time"

	"github.com/anya4emost/planer-server/internal/model"
	"github.com/anya4emost/planer-server/internal/server/router/response"
	"github.com/anya4emost/planer-server/internal/services"
	"github.com/anya4emost/planer-server/pkg/util"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type AuthController struct {
	userService    *services.UserService
	sessionService *services.SessionService
	secret         string
}

func NewAuthController(userService *services.UserService, sessionService *services.SessionService, secret string) *AuthController {
	return &AuthController{
		userService:    userService,
		sessionService: sessionService,
		secret:         secret,
	}
}

func (c *AuthController) createToken(userId string, expiration int64) (string, error) {
	claims := jwt.MapClaims{
		"sub": userId,
		"exp": expiration,
		"iat": time.Now(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(c.secret))
}

func getAccessTokenTime() int64 {
	// 15 Minutes
	return time.Now().Add(time.Minute * 15).Unix()
}

func getRefreshTokenTime() int64 {
	// week
	return time.Now().Add(time.Hour * 24 * 7).Unix()
}

func ClearCookies(c *fiber.Ctx, key ...string) {
	for i := range key {
		c.Cookie(&fiber.Cookie{
			Name:    key[i],
			Expires: time.Now().Add(-time.Hour * 24),
			Value:   "",
		})
	}
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

	user, err := c.userService.Create(input.Username, password)
	if err != nil {
		return response.ErrorUnauthorized(err, "Registration error")
	}

	accessToken, err := c.createToken(user.Id, getAccessTokenTime())
	if err != nil {
		return response.ErrorUnauthorized(err, "Registration error")
	}

	refreshTokenTime := getRefreshTokenTime()
	refreshToken, err := c.createToken(user.Id, refreshTokenTime)
	if err != nil {
		return response.ErrorUnauthorized(err, "Registration error")
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "access-token",
		Value:    accessToken,
		HTTPOnly: true,
	})

	ctx.Cookie(&fiber.Cookie{
		Name:     "refresh-token",
		Value:    refreshToken,
		HTTPOnly: true,
	})

	familyId, err := gonanoid.New()
	if err != nil {
		return response.ErrorBadRequest(err)
	}

	session := model.Session{
		RefreshToken: refreshToken,
		UserId:       user.Id,
		CreatedAt:    time.Now().String(),
		ExpiresAt:    time.Unix(refreshTokenTime, 0).Local().String(),
		Family:       familyId,
		Revoked:      false,
	}

	c.sessionService.Create(session)

	return response.Created(ctx, fiber.Map{
		"user": user,
	})
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {

	input := model.AuthInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}

	user, err := c.userService.GetByUsername(input.Username)
	if err != nil {
		return response.ErrorUnauthorized(err, "Login error")
	}
	if !util.CheckPassword(input.Pasword, user.Password) {
		return response.ErrorUnauthorized(err, "Login error")
	}

	accessToken, err := c.createToken(user.Id, getAccessTokenTime())
	if err != nil {
		return response.ErrorUnauthorized(err, "Registration error")
	}

	refreshTokenTime := getRefreshTokenTime()
	refreshToken, err := c.createToken(user.Id, refreshTokenTime)
	if err != nil {
		return response.ErrorUnauthorized(err, "Registration error")
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "access-token",
		Value:    accessToken,
		HTTPOnly: true,
	})

	ctx.Cookie(&fiber.Cookie{
		Name:     "refresh-token",
		Value:    refreshToken,
		HTTPOnly: true,
	})

	familyId, err2 := gonanoid.New()
	if err2 != nil {
		return response.ErrorBadRequest(err2)
	}

	session := model.Session{
		RefreshToken: refreshToken,
		UserId:       user.Id,
		CreatedAt:    strconv.FormatInt(time.Now().Unix(), 10),
		ExpiresAt:    strconv.FormatInt(refreshTokenTime, 10),
		Family:       familyId,
		Revoked:      false,
	}

	err = c.sessionService.Create(session)

	return response.Ok(ctx, fiber.Map{
		"user": user,
	})
}

func (c *AuthController) Me(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userid := claims["sub"].(string)

	currentUser, err := c.userService.GetById(userid)

	if err != nil {
		return response.ErrorUnauthorized(err, "Invalid credentials")
	}

	return response.Ok(ctx, fiber.Map{
		"user": currentUser,
	})
}

func (c *AuthController) Logout(ctx *fiber.Ctx) error {
	refreshToken := ctx.Cookies("refresh-token")

	session, err := c.sessionService.GetByName(refreshToken)
	if err != nil || session.RefreshToken == "" {
		return response.ErrorUnauthorized(err, "User is not loged in")
	}

	err = c.sessionService.DeleteAllFamily(session.Family)

	if err != nil {
		return response.DefaultErrorHandler(ctx, err)
	}

	ClearCookies(ctx, "access-token", "refresh-token")

	return response.Ok(ctx, fiber.Map{})
}

func (c *AuthController) Refresh(ctx *fiber.Ctx) error {
	refreshToken := ctx.Cookies("refresh-token")

	session, err := c.sessionService.GetByName(refreshToken)
	if err != nil || session.RefreshToken == "" {
		return response.ErrorUnauthorized(err, "User is not loged in")
	}

	if session.Revoked {
		err = c.sessionService.DeleteAllFamily(session.Family)
		if err != nil {
			return response.DefaultErrorHandler(ctx, err)
		}

		return response.ErrorUnauthorized(err, "User is not loged in")
	}

	c.sessionService.MakeRevoked(session)

	accessToken, err := c.createToken(session.UserId, getAccessTokenTime())
	if err != nil {
		return response.ErrorUnauthorized(err, "Registration error")
	}

	refreshTokenTime := getRefreshTokenTime()
	newRefreshToken, err := c.createToken(session.UserId, refreshTokenTime)
	if err != nil {
		return response.ErrorUnauthorized(err, "Registration error")
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "access-token",
		Value:    accessToken,
		HTTPOnly: true,
	})

	ctx.Cookie(&fiber.Cookie{
		Name:     "refresh-token",
		Value:    newRefreshToken,
		HTTPOnly: true,
	})

	newSession := model.Session{
		RefreshToken: newRefreshToken,
		UserId:       session.UserId,
		CreatedAt:    time.Now().String(),
		ExpiresAt:    time.Unix(refreshTokenTime, 0).Local().String(),
		Family:       session.Family,
		Revoked:      false,
	}

	c.sessionService.Create(newSession)

	return response.Ok(ctx, fiber.Map{})
}
