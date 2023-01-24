package protocols

import (
	"clean-boilerplate/boilerplate/src/app"
	"clean-boilerplate/boilerplate/src/protocols/http"
	"clean-boilerplate/boilerplate/src/protocols/rpc"
	"clean-boilerplate/shared/genproto/example"
)

func NewHttp(app app.Application) http.Server {
	return http.New(app)
}

func NewRpc(app app.Application) example.ExampleServiceServer {
	return rpc.New(app)
}
