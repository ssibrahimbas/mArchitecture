package http

import (
	"clean-boilerplate/boilerplate/src/app"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app app.Application
}

func New(app app.Application) Server {
	return Server{
		app: app,
	}
}

func (h Server) Load(router fiber.Router) fiber.Router {
	router.Get("/example/:key", h.GetExample)
	router.Get("/example", h.ListExample)
	router.Post("/example", h.CreateExample)
	router.Put("/example/:key", h.UpdateExample)
	return router
}
