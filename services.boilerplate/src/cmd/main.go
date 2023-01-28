package main

import (
	"context"

	"github.ssibrahimbas/mArchitecture/shared/env"
	"github.ssibrahimbas/mArchitecture/shared/events/nats"
	"github.ssibrahimbas/mArchitecture/shared/i18n"
	"github.ssibrahimbas/mArchitecture/shared/logs"

	"github.ssibrahimbas/mArchitecture/boilerplate/src/config"
	"github.ssibrahimbas/mArchitecture/boilerplate/src/delivery"

	"github.ssibrahimbas/mArchitecture/boilerplate/src/service"
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
