package command

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.ssibrahimbas/mArchitecture/auth/config"
	"github.ssibrahimbas/mArchitecture/auth/domain/user"
	"github.ssibrahimbas/mArchitecture/shared/cipher"
	"github.ssibrahimbas/mArchitecture/shared/decorator"
	"github.ssibrahimbas/mArchitecture/shared/events"
	"github.ssibrahimbas/mArchitecture/shared/i18n"
)

type LoginCommand struct {
	Email    string
	Password string
}

type LoginHandler decorator.CommandHandler[LoginCommand, *user.User]

type loginHandler struct {
	userRepo   user.Repository
	authTopics config.AuthTopics
	publisher  events.Publisher
	errors     user.Errors
}

type LoginHandlerConfig struct {
	UserRepo      user.Repository
	AuthTopics    config.AuthTopics
	Publisher     events.Publisher
	Logger        *logrus.Entry
	MetricsClient decorator.MetricsClient
	Errors        user.Errors
}

func NewLoginHandler(config LoginHandlerConfig) LoginHandler {
	return decorator.ApplyCommandDecorators[LoginCommand, *user.User](
		loginHandler{
			userRepo:   config.UserRepo,
			authTopics: config.AuthTopics,
			publisher:  config.Publisher,
			errors:     config.Errors,
		},
		config.Logger,
		config.MetricsClient,
	)
}

func (h loginHandler) Handle(ctx context.Context, command LoginCommand) (*user.User, *i18n.I18nError) {
	user, err := h.userRepo.GetByEmail(ctx, command.Email)
	if err != nil {
		return nil, err
	}
	if err := cipher.Compare(user.Password, command.Password); err != nil {
		_ = h.publisher.Publish(h.authTopics.LoginFailed, user)
		return nil, h.errors.InvalidPassword()
	}
	_ = h.publisher.Publish(h.authTopics.LoggedIn, user)
	return nil, err
}
