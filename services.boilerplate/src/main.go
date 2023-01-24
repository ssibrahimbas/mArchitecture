package main

import (
	"clean-boilerplate/boilerplate/src/app"
	"clean-boilerplate/boilerplate/src/config"
	"clean-boilerplate/boilerplate/src/protocols"
	"clean-boilerplate/boilerplate/src/service"
	"clean-boilerplate/shared/env"
	"clean-boilerplate/shared/genproto/example"
	"clean-boilerplate/shared/logs"
	"clean-boilerplate/shared/server/http"
	"clean-boilerplate/shared/server/rpc"
	"context"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

func main() {
	logs.Init()
	ctx := context.Background()
	config := config.App{}
	env.Load(&config)
	app := service.NewApplication(ctx, config)
	loadServer(app, config)

}

func loadRpc(app app.Application, config config.App) {
	rpc.RunServer(config.Server.Port, func(server *grpc.Server) {
		svc := protocols.NewRpc(app)
		example.RegisterExampleServiceServer(server, svc)
	})
}

func loadHttp(app app.Application, config config.App) {
	http.RunServer(config.Server.Host, config.Server.Port, func(router fiber.Router) fiber.Router {
		return protocols.NewHttp(app).Load(router)
	})
}

func loadServer(app app.Application, config config.App) {
	if config.Protocol == "rpc" {
		loadRpc(app, config)
		return
	}
	loadHttp(app, config)
}
