package command

import (
	"context"

	"github.ssibrahimbas/mArchitecture/shared/decorator"
	"github.ssibrahimbas/mArchitecture/shared/events"

	"github.ssibrahimbas/mArchitecture/boilerplate/src/config"
	"github.ssibrahimbas/mArchitecture/boilerplate/src/domain/example"

	"github.com/sirupsen/logrus"
)

type UpdateExampleCommand struct {
	Field   string
	Content string
}

type UpdateExampleHandler decorator.CommandHandler[UpdateExampleCommand]

type updateExampleHandler struct {
	exampleRepo   example.Repository
	exampleTopics config.ExampleTopics
	publisher     events.Publisher
}

type UpdateExampleHandlerConfig struct {
	ExampleRepo   example.Repository
	ExampleTopics config.ExampleTopics
	Publisher     events.Publisher
	Logger        *logrus.Entry
	MetricsClient decorator.MetricsClient
}

func NewUpdateExampleHandler(config UpdateExampleHandlerConfig) UpdateExampleHandler {
	return decorator.ApplyCommandDecorators[UpdateExampleCommand](
		updateExampleHandler{
			exampleRepo:   config.ExampleRepo,
			exampleTopics: config.ExampleTopics,
			publisher:     config.Publisher,
		},
		config.Logger,
		config.MetricsClient,
	)
}

func (h updateExampleHandler) Handle(ctx context.Context, command UpdateExampleCommand) error {
	example := &example.Example{
		Field:   command.Field,
		Content: command.Content,
	}

	err := h.exampleRepo.Update(ctx, example)
	if err != nil {
		return err
	}
	return h.publisher.Publish(h.exampleTopics.Updated, example)
}
