package service

import (
	"github.ssibrahimbas/mArchitecture/shared/events"
	"github.ssibrahimbas/mArchitecture/shared/metrics"

	"github.ssibrahimbas/mArchitecture/boilerplate/src/adapters"
	mysql_example "github.ssibrahimbas/mArchitecture/boilerplate/src/adapters/mysql/example"
	"github.ssibrahimbas/mArchitecture/boilerplate/src/app"
	"github.ssibrahimbas/mArchitecture/boilerplate/src/app/command"
	"github.ssibrahimbas/mArchitecture/boilerplate/src/app/query"
	"github.ssibrahimbas/mArchitecture/boilerplate/src/config"
	"github.ssibrahimbas/mArchitecture/boilerplate/src/domain/example"

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
