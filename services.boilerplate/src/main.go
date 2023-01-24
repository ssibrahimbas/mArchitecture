package main

import (
	"clean-boilerplate/boilerplate/src/app"
	"clean-boilerplate/boilerplate/src/config"
	"clean-boilerplate/boilerplate/src/protocols"
	"clean-boilerplate/boilerplate/src/service"
	"clean-boilerplate/shared/env"
	"clean-boilerplate/shared/genproto/example"
	"clean-boilerplate/shared/i18n"
	"clean-boilerplate/shared/logs"
	"clean-boilerplate/shared/server/http"
	"clean-boilerplate/shared/server/rpc"
	"clean-boilerplate/shared/validator"
	"context"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

type loader struct {
	app    app.Application
	config config.App
	i18n   *i18n.I18n
	ctx    context.Context
}

func main() {
	logs.Init()
	ctx := context.Background()
	config := config.App{}
	env.Load(&config)
	i18n := i18n.New(config.I18n.Fallback)
	i18n.Load(config.I18n.Dir, config.I18n.Locales...)
	app := service.NewApplication(ctx, config)
	l := loader{
		app:    app,
		config: config,
		i18n:   i18n,
		ctx:    ctx,
	}
	loadServer(l)

}

func loadRpc(l loader) {
	rpc.RunServer(l.config.Server.Port, func(server *grpc.Server) {
		svc := protocols.NewRpc(l.app)
		example.RegisterExampleServiceServer(server, svc)
	})
}

func loadHttp(l loader) {
	http.RunServer(http.Config{
		Host: l.config.Server.Host,
		Port: l.config.Server.Port,
		I18n: l.i18n,
		Cors: l.config.Cors,
		CreateHandler: func(router fiber.Router) fiber.Router {
			val := validator.New(l.i18n)
			val.ConnectCustom()
			val.RegisterTagName()
			return protocols.NewHttp(protocols.HttpConfig{
				App:       l.app,
				I18n:      *l.i18n,
				Validator: *val,
				Context:   l.ctx,
			}).Load(router)
		},
	})
}

func loadServer(l loader) {
	if l.config.Protocol == "rpc" {
		loadRpc(l)
		return
	}
	loadHttp(l)
}
