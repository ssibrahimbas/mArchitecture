package query

import (
	"context"

	"github.ssibrahimbas/mArchitecture/shared/decorator"
	"github.ssibrahimbas/mArchitecture/shared/i18n"

	"github.ssibrahimbas/mArchitecture/boilerplate/src/domain/example"

	"github.com/sirupsen/logrus"
)

type ListExampleQuery struct {
	Limit  int
	Offset int
}

type ListExampleResult struct {
	Examples []*example.Example
	Total    int
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

func (h listExampleHandler) Handle(ctx context.Context, query ListExampleQuery) (ListExampleResult, *i18n.I18nError) {
	example, total, err := h.exampleRepo.List(ctx, query.Limit, query.Offset)
	if err != nil {
		return ListExampleResult{}, err
	}

	return ListExampleResult{Examples: example, Total: total}, nil
}
