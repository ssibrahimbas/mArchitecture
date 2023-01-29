package command

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.ssibrahimbas/mArchitecture/auth/config"
	"github.ssibrahimbas/mArchitecture/auth/domain/user"
	"github.ssibrahimbas/mArchitecture/shared/auth/token"
	"github.ssibrahimbas/mArchitecture/shared/decorator"
	"github.ssibrahimbas/mArchitecture/shared/events"
	"github.ssibrahimbas/mArchitecture/shared/i18n"
	"github.ssibrahimbas/mArchitecture/shared/jwt"
)

type RefreshTokenCommand struct {
	Token string
	Claim *jwt.UserClaim
}

type RefreshTokenResult struct {
	Token string
}

type RefreshTokenHandler decorator.CommandHandler[RefreshTokenCommand, *RefreshTokenResult]

type refreshTokenHandler struct {
	authTopics config.AuthTopics
	publisher  events.Publisher
	tokenSrv   token.Service
	errors     user.Errors
}

type RefreshTokenHandlerConfig struct {
	AuthTopics    config.AuthTopics
	Publisher     events.Publisher
	TokenSrv      token.Service
	Errors        user.Errors
	Logger        *logrus.Entry
	MetricsClient decorator.MetricsClient
}

func NewRefreshTokenHandler(config RefreshTokenHandlerConfig) RefreshTokenHandler {
	return decorator.ApplyCommandDecorators[RefreshTokenCommand, *RefreshTokenResult](
		refreshTokenHandler{
			authTopics: config.AuthTopics,
			publisher:  config.Publisher,
			tokenSrv:   config.TokenSrv,
			errors:     config.Errors,
		},
		config.Logger,
		config.MetricsClient,
	)
}

func (h refreshTokenHandler) Handle(ctx context.Context, command RefreshTokenCommand) (*RefreshTokenResult, *i18n.I18nError) {
	err := h.tokenSrv.Expire(command.Token)
	if err != nil {
		return nil, h.errors.Failed("refresh token")
	}
	tkn, err := h.tokenSrv.Generate(command.Claim)
	if err != nil {
		return nil, h.errors.Failed("refresh token")
	}
	return &RefreshTokenResult{
		Token: tkn,
	}, nil
}
