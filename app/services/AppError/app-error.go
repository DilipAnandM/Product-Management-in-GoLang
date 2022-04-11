package AppError

import (
	"net/http"
	"prodapi/app/services/ErrorCodes"

	"github.com/gofiber/fiber/v2"
)

type Error struct {
	Status     bool
	StatusCode int
	ErrorCode  string
}

func AppError(err_obj Error, c *fiber.Ctx) error {
	return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
		"status":       false,
		"statusCode":   406,
		"error_object": ErrorCodes.Error_codes[err_obj.ErrorCode],
	})
}
