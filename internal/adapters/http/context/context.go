package httpcontext

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/guillospy92/crabi/internal/core/domain"
	"github.com/guillospy92/logger"
)

// ContextTextKey type string of context
type ContextTextKey string

// UserKeyContextKey key context user
const UserKeyContextKey ContextTextKey = "user"

// GetContextApplication transform ctx fiber to native ctx
func GetContextApplication(c *fiber.Ctx) context.Context {
	var ctx = context.TODO()
	loggerAdditional, ok := c.Locals("uuid").(string)

	if ok {
		ctx = logger.AppendCtx(ctx, slog.String("uuid", loggerAdditional))
	}

	user, ok := c.Locals(UserKeyContextKey).(domain.UserEntity)
	if ok {
		ctx = context.WithValue(ctx, UserKeyContextKey, user)
	} else {
		ctx = context.WithValue(ctx, UserKeyContextKey, domain.UserEntity{})
	}

	return ctx
}

// GetContextUser get information user context
func GetContextUser(c *fiber.Ctx) domain.UserEntity {
	user, ok := c.Locals(UserKeyContextKey).(domain.UserEntity)
	if !ok {
		panic(fmt.Errorf("user not found in context"))
	}

	return user
}
