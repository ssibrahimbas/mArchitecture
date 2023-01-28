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

type LogoutCommand struct {
	Token string
}

type LogoutResult struct{}

type LogoutHandler decorator.CommandHandler[LogoutCommand, LogoutResult]

type logoutHandler struct {
	authTopics config.AuthTopics
	publisher  events.Publisher
	tokenSrv   token.Service
}

type LogoutHandlerConfig struct {
	AuthTopics    config.AuthTopics
	Publisher     events.Publisher
	TokenSrv      token.Service
	Logger        *logrus.Entry
	MetricsClient decorator.MetricsClient
}

func NewLogoutHandler(config LogoutHandlerConfig) LogoutHandler {
	return decorator.ApplyCommandDecorators[LogoutCommand, LogoutResult](
		logoutHandler{
			authTopics: config.AuthTopics,
			publisher:  config.Publisher,
			tokenSrv:   config.TokenSrv,
		},
		config.Logger,
		config.MetricsClient,
	)
}

func (h logoutHandler) Handle(ctx context.Context, command LogoutCommand) (LogoutResult, *i18n.I18nError) {
	return LogoutResult{}, nil
}
