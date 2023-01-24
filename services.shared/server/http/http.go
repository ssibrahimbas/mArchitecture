package http

import (
	"encoding/json"
	"fmt"

	"clean-boilerplate/shared/server/http/error_handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirupsen/logrus"
)

func RunServer(host string, port int, createHandler func(router fiber.Router) fiber.Router) {
	RunServerOnAddr(fmt.Sprintf("%v:%v", host, port), createHandler)
}

func RunServerOnAddr(addr string, createHandler func(router fiber.Router) fiber.Router) {
	app := fiber.New(fiber.Config{
		ErrorHandler: error_handler.New(),
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	})

	group := app.Group("/api")
	createHandler(group)

	logrus.Infof("Starting server on %v", addr)
	if err := app.Listen(addr); err != nil {
		logrus.WithError(err).Panic("Unable to start HTTP server")
	}

	setGlobalMiddlewares(app)
}

func setGlobalMiddlewares(router fiber.Router) {
	router.Use(recover.New())
}
