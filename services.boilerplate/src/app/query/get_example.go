package query

import (
	"clean-boilerplate/boilerplate/src/domain/example"
	"clean-boilerplate/shared/decorator"
	"context"

	"github.com/sirupsen/logrus"
)

type GetExampleQuery struct {
	Key string
}

type GetExampleResult struct {
	Value string
	Key   string
}

type GetExampleHandler decorator.QueryHandler[GetExampleQuery, GetExampleResult]

type getExampleHandler struct {
	exampleRepo example.Repository
}

func NewGetExampleHandler(exampleRepo example.Repository, logger *logrus.Entry, metrics decorator.MetricsClient) GetExampleHandler {
	return decorator.ApplyQueryDecorators[GetExampleQuery, GetExampleResult](
		getExampleHandler{exampleRepo: exampleRepo},
		logger,
		metrics,
	)
}

func (h getExampleHandler) Handle(ctx context.Context, query GetExampleQuery) (GetExampleResult, error) {
	example, err := h.exampleRepo.Get(ctx, query.Key)
	if err != nil {
		return GetExampleResult{}, err
	}

	return GetExampleResult{Value: example.Value, Key: example.Key}, nil
}
