package service

import (
	"clean-boilerplate/boilerplate/src/adapters/mysql"
	"clean-boilerplate/boilerplate/src/app"
	"clean-boilerplate/boilerplate/src/app/command"
	"clean-boilerplate/boilerplate/src/app/query"
	"clean-boilerplate/boilerplate/src/config"
	"clean-boilerplate/boilerplate/src/domain/example"
	"clean-boilerplate/shared/metrics"
	"context"

	"github.com/sirupsen/logrus"
)

func NewApplication(ctx context.Context, config config.App) app.Application {
	sqlDb, err := mysql.New(config.MySQL)
	if err != nil {
		panic(err)
	}

	factoryConfig := example.FactoryConfig{
		MaxKeyLength:   10,
		MaxValueLength: 10,
		MinKeyLength:   1,
		MinValueLength: 1,
	}

	exampleFactory, err := example.NewFactory(factoryConfig)
	if err != nil {
		panic(err)
	}

	exampleRepo := mysql.NewExampleRepo(sqlDb, exampleFactory)

	logger := logrus.NewEntry(logrus.StandardLogger())
	metricsClient := metrics.NoOp{}

	return app.Application{
		Commands: app.Commands{
			CreateExample: command.NewCreateExampleHandler(exampleRepo, logger, metricsClient),
			UpdateExample: command.NewUpdateExampleHandler(exampleRepo, logger, metricsClient),
		},
		Queries: app.Queries{
			ListExample: query.NewListExampleHandler(exampleRepo, logger, metricsClient),
			GetExample:  query.NewGetExampleHandler(exampleRepo, logger, metricsClient),
		},
	}
}
