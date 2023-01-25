package command

import (
	"clean-boilerplate/boilerplate/src/domain/example"
	"clean-boilerplate/shared/decorator"
	"context"

	"github.com/sirupsen/logrus"
)

type UpdateExampleCommand struct {
	Field   string
	Content string
}

type UpdateExampleHandler decorator.CommandHandler[UpdateExampleCommand]

type updateExampleHandler struct {
	exampleRepo example.Repository
}

func NewUpdateExampleHandler(exampleRepo example.Repository, logger *logrus.Entry, metrics decorator.MetricsClient) UpdateExampleHandler {
	return decorator.ApplyCommandDecorators[UpdateExampleCommand](
		updateExampleHandler{exampleRepo: exampleRepo},
		logger,
		metrics,
	)
}

func (h updateExampleHandler) Handle(ctx context.Context, command UpdateExampleCommand) error {
	example := &example.Example{
		Field:   command.Field,
		Content: command.Content,
	}

	return h.exampleRepo.Update(ctx, example)
}
