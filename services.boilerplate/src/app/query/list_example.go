package query

import (
	"clean-boilerplate/boilerplate/src/domain/example"
	"clean-boilerplate/shared/decorator"
	"context"

	"github.com/sirupsen/logrus"
)

type ListExampleQuery struct {
	Limit  int
	Offset int
}

type ListExampleResult struct {
	Examples []example.Example
}

type ListExampleHandler decorator.QueryHandler[ListExampleQuery, ListExampleResult]

type listExampleHandler struct {
	exampleRepo example.Repository
}

func NewListExampleHandler(exampleRepo example.Repository, logger *logrus.Entry, metrics decorator.MetricsClient) ListExampleHandler {
	return decorator.ApplyQueryDecorators[ListExampleQuery, ListExampleResult](
		listExampleHandler{exampleRepo: exampleRepo},
		logger,
		metrics,
	)
}

func (h listExampleHandler) Handle(ctx context.Context, query ListExampleQuery) (ListExampleResult, error) {
	example, err := h.exampleRepo.List(ctx, query.Limit, query.Offset)
	if err != nil {
		return ListExampleResult{}, err
	}

	return ListExampleResult{Examples: example}, nil
}
