package util

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/raamaj/chat-app/internal/model"
)

func NewErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		var e *fiber.Error
		if errors.As(err, &e) {
			code = e.Code
		}

		if err.Error() == "missing or malformed JWT" {
			code = fiber.StatusUnauthorized
			err = fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
		}

		return ctx.Status(code).JSON(model.ErrorResponse{
			Errors: e.Error(),
		})
	}
}
