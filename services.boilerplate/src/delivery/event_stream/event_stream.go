package event_stream

import (
	"clean-boilerplate/boilerplate/src/app"
	"clean-boilerplate/boilerplate/src/config"
	"clean-boilerplate/shared/events"
)

type Server struct {
	app    app.Application
	Topics config.ExampleTopics
	engine events.Engine
}

type Config struct {
	App    app.Application
	Topics config.ExampleTopics
	Engine events.Engine
}

func New(config Config) Server {
	return Server{
		app:    config.App,
		engine: config.Engine,
		Topics: config.Topics,
	}
}

func (s Server) Load() {
	//s.engine.Subscribe(s.Topics.AnotherCreated, s.ListenAnotherCreated)
}
