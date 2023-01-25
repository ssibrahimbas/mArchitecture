package main

import (
	"clean-boilerplate/boilerplate/src/config"
	"clean-boilerplate/boilerplate/src/delivery"

	"clean-boilerplate/boilerplate/src/service"
	"clean-boilerplate/shared/env"
	"clean-boilerplate/shared/events/nats"
	"clean-boilerplate/shared/i18n"
	"clean-boilerplate/shared/logs"
	"context"
)

func main() {
	logs.Init()
	ctx := context.Background()
	config := config.App{}
	env.Load(&config)
	i18n := i18n.New(config.I18n.Fallback)
	i18n.Load(config.I18n.Dir, config.I18n.Locales...)
	eventEngine := nats.New(nats.Config{
		Url:     config.Nats.Url,
		Streams: config.Nats.Streams,
	})
	app := service.NewApplication(config, eventEngine)
	delivery := delivery.New(delivery.Config{
		App:         app,
		Config:      config,
		I18n:        i18n,
		Ctx:         ctx,
		EventEngine: eventEngine,
	})
	delivery.Load()
}
