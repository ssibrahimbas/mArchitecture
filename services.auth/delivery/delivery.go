package delivery

import (
	"context"

	"github.ssibrahimbas/mArchitecture/auth/app"
	"github.ssibrahimbas/mArchitecture/auth/config"
	"github.ssibrahimbas/mArchitecture/shared/events"
	"github.ssibrahimbas/mArchitecture/shared/i18n"
)

type Delivery interface {
	Load()
}

type delivery struct {
	app         app.Application
	config      config.App
	i18n        *i18n.I18n
	ctx         context.Context
	eventEngine events.Engine
}

type Config struct {
	App         app.Application
	Config      config.App
	I18n        *i18n.I18n
	Ctx         context.Context
	EventEngine events.Engine
}

func New(config Config) Delivery {
	return &delivery{
		app:         config.App,
		config:      config.Config,
		i18n:        config.I18n,
		ctx:         config.Ctx,
		eventEngine: config.EventEngine,
	}
}

func (d *delivery) Load() {}
