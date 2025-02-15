package httpmiddleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/guillospy92/crabi/internal/adapters/http/context"
	"github.com/guillospy92/crabi/internal/adapters/http/responses"
	"github.com/guillospy92/crabi/internal/core/domain"
	"github.com/guillospy92/crabi/pkg/jwt"
)

// ValidateTokenApp validate token app
func ValidateTokenApp() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("token")
		if tokenString == "" {
			return c.Status(http.StatusBadRequest).JSON(httpresponses.ResponseError{
				ErrorCode:  httpresponses.ErrorCodeBadFieldRequest,
				StatusCode: http.StatusBadRequest,
				Message:    "token not found",
			})
		}

		user, err := pkgjwt.ValidateToken[domain.UserEntity](tokenString)

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(httpresponses.ResponseError{
				ErrorCode:  httpresponses.ErrorCodeUnexpected,
				StatusCode: http.StatusInternalServerError,
				Message:    "error deserializing token",
			})
		}

		c.Locals(httpcontext.UserKeyContextKey, *user)

		return c.Next()
	}
}
