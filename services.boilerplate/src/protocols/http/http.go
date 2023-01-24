package http

import (
	"clean-boilerplate/boilerplate/src/app"
	"clean-boilerplate/shared/i18n"
	"clean-boilerplate/shared/validator"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app       app.Application
	i18n      i18n.I18n
	validator validator.Validator
}

type Config struct {
	App       app.Application
	I18n      i18n.I18n
	Validator validator.Validator
}

func New(config Config) Server {
	return Server{
		app:       config.App,
		i18n:      config.I18n,
		validator: config.Validator,
	}
}

func (h Server) Load(router fiber.Router) fiber.Router {
	router.Get("/example/:key", h.GetExample)
	router.Get("/example", h.ListExample)
	router.Post("/example", h.CreateExample)
	router.Put("/example/:key", h.UpdateExample)
	return router
}
