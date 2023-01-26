package service

import (
	"clean-boilerplate/boilerplate/src/adapters"
	mysql_example "clean-boilerplate/boilerplate/src/adapters/mysql/example"
	"clean-boilerplate/boilerplate/src/app"
	"clean-boilerplate/boilerplate/src/app/command"
	"clean-boilerplate/boilerplate/src/app/query"
	"clean-boilerplate/boilerplate/src/config"
	"clean-boilerplate/boilerplate/src/domain/example"
	"clean-boilerplate/shared/events"
	"clean-boilerplate/shared/metrics"

	"github.com/sirupsen/logrus"
)

func NewApplication(config config.App, eventEngine events.Engine) app.Application {
	sqlDb, err := mysql_example.New(config.MySQLExample)
	if err != nil {
		panic(err)
	}

	factoryConfig := example.FactoryConfig{
		MaxFieldLength:   10,
		MaxContentLength: 10,
		MinFieldLength:   1,
		MinContentLength: 1,
	}

	exampleFactory, err := example.NewFactory(factoryConfig)
	if err != nil {
		panic(err)
	}

	exampleRepo := adapters.MySQL.NewExample(sqlDb, exampleFactory)

	logger := logrus.NewEntry(logrus.StandardLogger())
	metricsClient := metrics.NoOp{}

	return app.Application{
		Commands: app.Commands{
			CreateExample: command.NewCreateExampleHandler(command.CreateExampleHandlerConfig{
				ExampleRepo:   exampleRepo,
				ExampleTopics: config.Topics.Example,
				Publisher:     eventEngine,
				Logger:        logger,
				MetricsClient: metricsClient,
			}),
			UpdateExample: command.NewUpdateExampleHandler(command.UpdateExampleHandlerConfig{
				ExampleRepo:   exampleRepo,
				ExampleTopics: config.Topics.Example,
				Publisher:     eventEngine,
				Logger:        logger,
				MetricsClient: metricsClient,
			}),
		},
		Queries: app.Queries{
			ListExample: query.NewListExampleHandler(exampleRepo, logger, metricsClient),
			GetExample:  query.NewGetExampleHandler(exampleRepo, logger, metricsClient),
		},
	}
}
