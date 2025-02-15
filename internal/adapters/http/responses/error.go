package httpresponses

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	errorlogic "github.com/guillospy92/crabi/internal/core/errors"
)

const (
	// ErrorCodeBadFieldRequest errors validate request
	ErrorCodeBadFieldRequest = "ErrorCodeBadFieldRequest"

	// ErrorCodeUnexpected error unexpected internal
	ErrorCodeUnexpected = "ErrorCodeUnexpected"
)

// ResponseErrorWithAttribute responses errors with attribute
type ResponseErrorWithAttribute struct {
	ErrorCode     string   `json:"error_code"`
	Message       string   `json:"message"`
	MessageErrors []string `json:"message_errors"`
	StatusCode    int      `json:"status_code"`
}

// ResponseError responses errors general
type ResponseError struct {
	ErrorCode  string `json:"error_code"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

// EvaluateError evaluate error by type error application (ErrorGlobal)
func EvaluateError(c *fiber.Ctx, err error) error {
	var eg *errorlogic.ErrorLogic
	switch {
	case errors.As(err, &eg):
		return c.Status(eg.StatusCode).JSON(eg.GenerateErrorJSON())

	default:
		return c.Status(http.StatusInternalServerError).JSON(ResponseError{
			ErrorCode:  ErrorCodeUnexpected,
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal Server Error",
		})
	}
}
