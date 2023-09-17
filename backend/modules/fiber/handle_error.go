package fiber

import (
	"densomap-backend/types/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	// Case of *fiber.Error
	if e, ok := err.(*fiber.Error); ok {
		return c.Status(e.Code).JSON(response.ErrorResponse{
			Success: false,
			Code:    strings.ReplaceAll(strings.ToUpper(e.Error()), " ", "_"),
			Message: e.Error(),
			Error:   e.Error(),
		})
	}

	// Case of ErrorInstance
	if e, ok := err.(*response.ErrorInstance); ok {
		if e.Err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&response.ErrorResponse{
				Success: false,
				Code:    e.Code,
				Message: e.Message,
				Error:   e.Err.Error(),
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(&response.ErrorResponse{
			Success: false,
			Message: e.Message,
			Code:    e.Code,
		})
	}

	// Case of validator.ValidationErrors
	if e, ok := err.(validator.ValidationErrors); ok {
		var lists []string
		for _, err := range e {
			lists = append(lists, err.Field()+" ("+err.Tag()+")")
		}

		message := strings.Join(lists[:], ", ")

		return c.Status(fiber.StatusBadRequest).JSON(&response.ErrorResponse{
			Success: false,
			Code:    "VALIDATION_FAILED",
			Message: "VALIDATION failed on field " + message,
			Error:   e.Error(),
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(&response.ErrorResponse{
		Success: false,
		Code:    "UNKNOWN_SERVER_SIDE_ERROR",
		Message: "Unknown server side error",
		Error:   err.Error(),
	})
}
