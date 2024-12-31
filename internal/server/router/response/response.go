package response

import (
	"github.com/anya4emost/planer-server/internal/model"
	"github.com/gofiber/fiber/v2"
)

func Response(ctx *fiber.Ctx, code int, data interface{}) error {
	return ctx.Status(code).JSON(model.ApiResponse{
		Success: true,
		Data:    data,
	})
}

func Ok(ctx *fiber.Ctx, data interface{}) error {
	return Response(ctx, fiber.StatusOK, data)
}

func Created(ctx *fiber.Ctx, data interface{}) error {
	return Response(ctx, fiber.StatusCreated, data)
}
