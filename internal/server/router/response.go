package router

import "github.com/gofiber/fiber/v2"

func Response(ctx *fiber.Ctx, code int, data interface{}) error {
	return ctx.Status(code).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func Ok(ctx *fiber.Ctx, data interface{}) error {
	return Response(ctx, fiber.StatusOK, data)
}

func Created(ctx *fiber.Ctx, data interface{}) error {
	return Response(ctx, fiber.StatusCreated, data)
}
