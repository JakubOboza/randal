package server

import (
	"errors"
	"fmt"
	"strings"

	"github.com/JakubOboza/randal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var (
	ErrPortOutOfBounds = errors.New("server port should be between 1024 and 65535")
)

type App struct {
	port   int
	engine *fiber.App
	conf   *config.RandalConfig
}

func New(port int, conf *config.RandalConfig) *App {
	return &App{port: port, conf: conf}
}

func (app *App) Setup() error {
	app.engine = fiber.New()
	app.engine.Use(logger.New())

	for path, endPoint := range app.conf.EndPoints {

		if endPoint != nil && endPoint.IsValid() {
			fmt.Println("registering path: ", path)
			app.engine.Get(path, app.HandleRequest(endPoint))
		} else {
			fmt.Println("endpoint invalid: ", path)
		}

	}

	// handle root if configured
	if strings.TrimSpace(app.conf.RootUrl) != "" {
		app.engine.Get("/", func(c *fiber.Ctx) error {
			return c.Redirect(app.conf.RootUrl)
		})
	}

	return nil
}

func (app *App) Start() error {
	if app.port < 1024 || app.port > 65535 {
		return ErrPortOutOfBounds
	}
	return app.engine.Listen(fmt.Sprintf(":%d", app.port))
}

func (app *App) HandleRequest(endPoint *config.RandalEndpoint) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Redirect(endPoint.Next())
	}
}
