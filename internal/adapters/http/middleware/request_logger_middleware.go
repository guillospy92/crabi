package httpmiddleware

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/guillospy92/logger"
)

// RequestLoggerMiddleware register and logger all responses
func RequestLoggerMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		start := time.Now()
		id := uuid.New().String()
		l := logger.Log()

		ctx.Locals("uuid", id)

		defer func() {
			code := ctx.Response().StatusCode()
			panicErr := recover()

			if panicErr != nil {
				code = http.StatusInternalServerError
				panic(panicErr)
			}

			l.With(
				slog.Any("code", code),
				slog.String("url", ctx.OriginalURL()),
				slog.String("User-Agent", ctx.Get("User-Agent")),
				slog.Any("elapsed_ms", time.Since(start)),
				slog.String("uuid", id),
			).Info("tracing request response")
		}()

		return ctx.Next()
	}
}
