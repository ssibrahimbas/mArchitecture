package decorator

import (
	"context"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

type CommandHandler[C any] interface {
	Handle(ctx context.Context, cmd C) error
}

func ApplyCommandDecorators[H any](handler CommandHandler[H], logger *logrus.Entry, metricsClient MetricsClient) CommandHandler[H] {
	return &commandLoggingDecorator[H]{
		base: &commandMetricsDecorator[H]{
			base:   handler,
			client: metricsClient,
		},
		logger: logger,
	}
}

func generateActionName(handler any) string {
	return strings.Split(fmt.Sprintf("%T", handler), ".")[1]
}
