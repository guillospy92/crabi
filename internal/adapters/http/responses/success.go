package httpresponses

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// ResponseSuccess body response success
type ResponseSuccess struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

// ResponseSuccessWithData body response success with data
type ResponseSuccessWithData struct {
	Data       any    `json:"data"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

// HandleSuccess returns success response application in json
func HandleSuccess(c *fiber.Ctx, message string) error {
	return c.Status(http.StatusOK).JSON(ResponseSuccess{
		Message:    message,
		StatusCode: http.StatusOK,
	})
}

// HandleSuccessWithData returns success response application in json with fields or data additional
func HandleSuccessWithData(c *fiber.Ctx, message string, data any) error {
	return c.Status(http.StatusOK).JSON(ResponseSuccessWithData{
		Message:    message,
		StatusCode: http.StatusOK,
		Data:       data,
	})
}
