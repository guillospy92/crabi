package bootstrap

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// ConfigTest config testing all application
type ConfigTest struct {
	Handler fiber.Handler
	Route   string
	Method  string
}

// LoadServerTestApplication  run the resources necessary to start the application test
func LoadServerTestApplication(config ConfigTest) *fiber.App {
	app := fiber.New()

	if config.Method == http.MethodGet {
		app.Get(config.Route, config.Handler)
	} else {
		app.Post(config.Route, config.Handler)
	}

	return app
}
