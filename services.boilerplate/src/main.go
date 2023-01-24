package main

import (
	"clean-boilerplate/boilerplate/src/app"
	"clean-boilerplate/boilerplate/src/protocols"
	"clean-boilerplate/boilerplate/src/service"
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
	app := service.NewApplication(ctx)
	loadHttp(app)

}

func loadRpc(app app.Application) {
	rpc.RunServer(3001, func(server *grpc.Server) {
		svc := protocols.NewRpc(app)
		example.RegisterExampleServiceServer(server, svc)
	})

}

func loadHttp(app app.Application) {
	http.RunServer("localhost", 3000, func(router fiber.Router) fiber.Router {
		return protocols.NewHttp(app).Load(router)
	})
}
