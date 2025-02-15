package httproutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guillospy92/crabi/bootstrap/container"
	httpmiddleware "github.com/guillospy92/crabi/internal/adapters/http/middleware"
)

// LoadRouters configure all routes
func LoadRouters(app *fiber.App) {
	routerAPI := app.Group("/api/v1")

	routerAPI.Post("login", container.AuthUserHandler.Handler)

	routerAPI.Post("user", container.CreateUserHandler.Handler)
	routerAPI.Get("user", httpmiddleware.ValidateTokenApp(), container.GetUserInfoHandler.Handler)
}
