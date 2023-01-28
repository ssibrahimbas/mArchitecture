package rpc

import (
	"github.ssibrahimbas/mArchitecture/shared/genproto/example"

	"github.ssibrahimbas/mArchitecture/boilerplate/src/app"
)

type Server struct {
	app app.Application
	example.ExampleServiceServer
}

func New(app app.Application) example.ExampleServiceServer {
	return Server{
		app: app,
	}
}
