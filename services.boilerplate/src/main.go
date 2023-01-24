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

func main() {
	logs.Init()
	ctx := context.Background()
	config := config.App{}
	env.Load(&config)
	i18n := i18n.New(config.I18n.Fallback)
	i18n.Load(config.I18n.Dir, config.I18n.Locales)
	app := service.NewApplication(ctx, config)
	loadServer(app, config, i18n)

}

func loadRpc(app app.Application, config config.App) {
	rpc.RunServer(config.Server.Port, func(server *grpc.Server) {
		svc := protocols.NewRpc(app)
		example.RegisterExampleServiceServer(server, svc)
	})
}

func loadHttp(app app.Application, config config.App, i18n *i18n.I18n) {
	http.RunServer(http.Config{
		Host: config.Server.Host,
		Port: config.Server.Port,
		I18n: i18n,
		CreateHandler: func(router fiber.Router) fiber.Router {
			val := validator.New(i18n)
			val.ConnectCustom()
			val.RegisterTagName()
			return protocols.NewHttp(protocols.HttpConfig{
				App:       app,
				I18n:      *i18n,
				Validator: *val,
			}).Load(router)
		},
	})
}

func loadServer(app app.Application, config config.App, i18n *i18n.I18n) {
	if config.Protocol == "rpc" {
		loadRpc(app, config)
		return
	}
	loadHttp(app, config, i18n)
}
