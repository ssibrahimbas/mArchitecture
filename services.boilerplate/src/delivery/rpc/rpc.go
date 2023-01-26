package rpc

import (
	"clean-boilerplate/boilerplate/src/app"
	"clean-boilerplate/shared/genproto/example"
)

type Server struct {
	app app.Application
	example.UnsafeExampleServiceServer
}

func New(app app.Application) example.ExampleServiceServer {
	return Server{
		app: app,
	}
}