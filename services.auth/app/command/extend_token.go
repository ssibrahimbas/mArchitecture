package command

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.ssibrahimbas/mArchitecture/auth/config"
	"github.ssibrahimbas/mArchitecture/shared/auth/token"
	"github.ssibrahimbas/mArchitecture/shared/decorator"
	"github.ssibrahimbas/mArchitecture/shared/events"
	"github.ssibrahimbas/mArchitecture/shared/i18n"
)

type ExtendTokenCommand struct {
	Token string
}

type ExtendTokenResult struct{}

type ExtendTokenHandler decorator.CommandHandler[ExtendTokenCommand, ExtendTokenResult]

type extendTokenHandler struct {
	authTopics config.AuthTopics
	publisher  events.Publisher
	tokenSrv   token.Service
}

type ExtendTokenHandlerConfig struct {
	AuthTopics    config.AuthTopics
	Publisher     events.Publisher
	TokenSrv      token.Service
	Logger        *logrus.Entry
	MetricsClient decorator.MetricsClient
}

func NewExtendTokenHandler(config ExtendTokenHandlerConfig) ExtendTokenHandler {
	return decorator.ApplyCommandDecorators[ExtendTokenCommand, ExtendTokenResult](
		extendTokenHandler{
			authTopics: config.AuthTopics,
			publisher:  config.Publisher,
			tokenSrv:   config.TokenSrv,
		},
		config.Logger,
		config.MetricsClient,
	)
}

func (h extendTokenHandler) Handle(ctx context.Context, command ExtendTokenCommand) (ExtendTokenResult, *i18n.I18nError) {
	return ExtendTokenResult{}, nil
}
