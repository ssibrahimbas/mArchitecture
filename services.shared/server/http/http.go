package http

import (
	"encoding/json"
	"fmt"

	"clean-boilerplate/shared/i18n"
	"clean-boilerplate/shared/server/http/error_handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Host          string
	Port          int
	CreateHandler func(router fiber.Router) fiber.Router
	I18n          *i18n.I18n
}

func RunServer(cfg Config) {
	addr := fmt.Sprintf("%v:%v", cfg.Host, cfg.Port)
	RunServerOnAddr(addr, cfg)
}

func RunServerOnAddr(addr string, cfg Config) {
	app := fiber.New(fiber.Config{
		ErrorHandler: error_handler.New(error_handler.Config{
			DfMsgKey: "error_internal_server_error",
			I18n:     cfg.I18n,
		}),
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	group := app.Group("/api")
	cfg.CreateHandler(group)

	logrus.Infof("Starting server on %v", addr)
	if err := app.Listen(addr); err != nil {
		logrus.WithError(err).Panic("Unable to start HTTP server")
	}

	setGlobalMiddlewares(app)
}

func setGlobalMiddlewares(router fiber.Router) {
	router.Use(recover.New())
}
