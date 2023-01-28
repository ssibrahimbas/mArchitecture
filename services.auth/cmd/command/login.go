package command

import (
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	"github.ssibrahimbas/mArchitecture/auth/config"
	"github.ssibrahimbas/mArchitecture/auth/domain/user"
	"github.ssibrahimbas/mArchitecture/shared/cipher"
	"github.ssibrahimbas/mArchitecture/shared/decorator"
	"github.ssibrahimbas/mArchitecture/shared/events"
)

type LoginCommand struct {
	Email    string
	Password string
}

type LoginHandler decorator.CommandHandler[LoginCommand]

type loginHandler struct {
	userRepo   user.Repository
	authTopics config.AuthTopics
	publisher  events.Publisher
}

type LoginHandlerConfig struct {
	UserRepo      user.Repository
	AuthTopics    config.AuthTopics
	Publisher     events.Publisher
	Logger        *logrus.Entry
	MetricsClient decorator.MetricsClient
}

func NewLoginHandler(config LoginHandlerConfig) LoginHandler {
	return decorator.ApplyCommandDecorators[LoginCommand](
		loginHandler{
			userRepo:   config.UserRepo,
			authTopics: config.AuthTopics,
			publisher:  config.Publisher,
		},
		config.Logger,
		config.MetricsClient,
	)
}

func (h loginHandler) Handle(ctx context.Context, command LoginCommand) error {
	user, err := h.userRepo.GetByEmail(ctx, command.Email)
	if err != nil {
		return err
	}
	if err := cipher.Compare(user.Password, command.Password); err != nil {
		_ = h.publisher.Publish(h.authTopics.LoginFailed, user)
		return errors.New("invalid password")
	}
	_ = h.publisher.Publish(h.authTopics.LoggedIn, user)
	return nil
}
