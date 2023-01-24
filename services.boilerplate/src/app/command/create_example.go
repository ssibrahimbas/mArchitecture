package command

import (
	"clean-boilerplate/boilerplate/src/domain/example"
	"clean-boilerplate/shared/decorator"
	"context"

	"github.com/sirupsen/logrus"
)

type CreateExampleCommand struct {
	Key   string
	Value string
}

type CreateExampleHandler decorator.CommandHandler[CreateExampleCommand]

type createExampleHandler struct {
	exampleRepo example.Repository
}

func NewCreateExampleHandler(exampleRepo example.Repository, logger *logrus.Entry, metrics decorator.MetricsClient) CreateExampleHandler {
	return decorator.ApplyCommandDecorators[CreateExampleCommand](
		createExampleHandler{exampleRepo: exampleRepo},
		logger,
		metrics,
	)
}

func (h createExampleHandler) Handle(ctx context.Context, command CreateExampleCommand) error {
	example := &example.Example{
		Key:   command.Key,
		Value: command.Value,
	}

	return h.exampleRepo.Create(ctx, example)
}
