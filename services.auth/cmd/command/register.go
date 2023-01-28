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

type RegisterCommand struct {
	Email    string
	Password string
}

type RegisterHandler decorator.CommandHandler[RegisterCommand, *user.User]

type registerHandler struct {
	userRepo   user.Repository
	authTopics config.AuthTopics
	publisher  events.Publisher
}

type RegisterHandlerConfig struct {
	UserRepo      user.Repository
	AuthTopics    config.AuthTopics
	Publisher     events.Publisher
	Logger        *logrus.Entry
	MetricsClient decorator.MetricsClient
}

func NewRegisterHandler(config RegisterHandlerConfig) RegisterHandler {
	return decorator.ApplyCommandDecorators[RegisterCommand, *user.User](
		registerHandler{
			userRepo:   config.UserRepo,
			authTopics: config.AuthTopics,
			publisher:  config.Publisher,
		},
		config.Logger,
		config.MetricsClient,
	)
}

func (h registerHandler) Handle(ctx context.Context, command RegisterCommand) (*user.User, *i18n.I18nError) {
	already, err := h.userRepo.GetByEmail(ctx, command.Email)
	if err != nil {
		return nil, err
	}
	if already != nil {
		return nil, i18n.NewError("user already exists", i18n.P{"Email": command.Email})
	}
	pw, error := cipher.Hash(command.Password)
	if error != nil {
		return nil, i18n.NewError("error hashing password", i18n.P{"Error": error.Error()})
	}
	user, err := h.userRepo.Create(ctx, command.Email, pw)
	if err != nil {
		return nil, err
	}
	_ = h.publisher.Publish(h.authTopics.Registered, user)
	return user, err
}
