package protocols

import (
	"clean-boilerplate/boilerplate/src/app"
	"clean-boilerplate/boilerplate/src/protocols/http"
	"clean-boilerplate/boilerplate/src/protocols/rpc"
	"clean-boilerplate/shared/genproto/example"
)

type HttpConfig http.Config

func NewHttp(config HttpConfig) http.Server {
	return http.New(http.Config{
		App:       config.App,
		I18n:      config.I18n,
		Validator: config.Validator,
	})
}

func NewRpc(app app.Application) example.ExampleServiceServer {
	return rpc.New(app)
}
