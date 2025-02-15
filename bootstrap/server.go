package bootstrap

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
	"github.com/guillospy92/crabi/bootstrap/container"
	"github.com/guillospy92/crabi/internal/adapters/http/middleware"
	httproutes "github.com/guillospy92/crabi/internal/adapters/http/routes"

	loggermiddleware "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/guillospy92/crabi/resources"
	"github.com/guillospy92/logger"
)

// LoadApplication run the resources necessary to start the application
func LoadApplication(pathsProperties string) *fiber.App {
	loadEnvironmentVar(pathsProperties)
	container.InitializeServiceContainer()

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			return ctx.Status(code).JSON(fiber.Map{
				"error_code": "ErrorCodeUnexpected",
				"message":    err.Error(),
				"code":       code,
			})
		},
	})

	app.Static("/public", "./public")

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(_ *fiber.Ctx, e any) {
			logger.Log().Error("error panic not controller", slog.Any("err", e), slog.String("trace", string(debug.Stack())))
		},
	}))

	app.Use(httpmiddleware.RequestLoggerMiddleware())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		if err := app.Shutdown(); err != nil {
			logger.Log().Error("error shutdown process", slog.Any("err", err))
		}
	}()

	loadMiddleware(app)
	httproutes.LoadRouters(app)

	return app
}

// StartApplication init application server
func StartApplication(app *fiber.App) {
	if err := app.Listen(fmt.Sprintf(":%s", resources.ConfigurationEnv().APPPort)); err != nil {
		logger.Log().Error("error fatal init server", slog.Any("err", err))
		panic(fmt.Sprintf("error fatal init server %s", err.Error()))
	}
}

func loadMiddleware(app *fiber.App) {
	if resources.ConfigurationEnv().ShowLogRequest {
		app.Use(loggermiddleware.New(loggermiddleware.Config{
			Format:     "[${time}] ${status} - ${latency} ${method} ${path}\n",
			TimeFormat: "02-Jan-2006",
			TimeZone:   "utc",
		}))
	}

	if resources.ConfigurationEnv().ShowMonitor {
		app.Get("monitor", monitor.New())
	}
}

func loadEnvironmentVar(pathsProperties string) {
	resources.LoadConfig(pathsProperties)
}
