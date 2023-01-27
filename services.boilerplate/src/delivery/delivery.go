package delivery

import (
	"clean-boilerplate/boilerplate/src/app"
	"clean-boilerplate/boilerplate/src/config"
	"clean-boilerplate/boilerplate/src/delivery/event_stream"
	"clean-boilerplate/boilerplate/src/delivery/http"
	"clean-boilerplate/boilerplate/src/delivery/rpc"
	"clean-boilerplate/shared/events"
	"clean-boilerplate/shared/genproto/example"
	"clean-boilerplate/shared/i18n"
	"clean-boilerplate/shared/validator"
	"context"

	sharedHttp "clean-boilerplate/shared/server/http"
	sharedRpc "clean-boilerplate/shared/server/rpc"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

type Loader struct {
}

type Delivery interface {
	Load()
}

type delivery struct {
	app         app.Application
	config      config.App
	i18n        *i18n.I18n
	ctx         context.Context
	eventEngine events.Engine
}

type Config struct {
	App         app.Application
	Config      config.App
	I18n        *i18n.I18n
	Ctx         context.Context
	EventEngine events.Engine
}

func New(config Config) Delivery {
	return &delivery{
		app:         config.App,
		config:      config.Config,
		i18n:        config.I18n,
		ctx:         config.Ctx,
		eventEngine: config.EventEngine,
	}
}

func (d *delivery) Load() {
	d.loadEventStream()
	if d.config.Protocol == "grpc" {
		d.loadRPC()
		return
	}
	d.loadHTTP()
}

func (d *delivery) loadHTTP() *delivery {
	sharedHttp.RunServer(sharedHttp.Config{
		Host: d.config.Server.Host,
		Port: d.config.Server.Port,
		I18n: d.i18n,
		Cors: d.config.Cors,
		CreateHandler: func(router fiber.Router) fiber.Router {
			val := validator.New(d.i18n)
			val.ConnectCustom()
			val.RegisterTagName()
			return http.New(http.Config{
				App:       d.app,
				I18n:      *d.i18n,
				Validator: *val,
				Context:   d.ctx,
			}).Load(router)
		},
	})
	return d
}

func (d *delivery) loadRPC() *delivery {
	sharedRpc.RunServer(d.config.Server.Port, func(server *grpc.Server) {
		svc := rpc.New(d.app)
		example.RegisterExampleServiceServer(server, svc)
	})
	return d
}

func (d *delivery) loadEventStream() *delivery {
	eventStream := event_stream.New(event_stream.Config{
		App:    d.app,
		Topics: d.config.Topics.Example,
		Engine: d.eventEngine,
	})
	err := d.eventEngine.Open()
	if err != nil {
		panic(err)
	}
	eventStream.Load()
	return d
}
