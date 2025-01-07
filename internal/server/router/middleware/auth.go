package middleware

import (
	"github.com/anya4emost/planer-server/internal/server/router/response"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func AccessTokenErrorHandler(c *fiber.Ctx, err error) error {
	if err.Error() == jwtware.ErrJWTMissingOrMalformed.Error() {
		return response.ErrorBadRequest(err)
	}

	return response.InvalidTokenError(err)
}

func AccessTokenVerification(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(secret)},
		ErrorHandler: AccessTokenErrorHandler,
		TokenLookup:  "cookie:access-token",
	})
}
