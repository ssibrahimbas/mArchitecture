package app

import (
	"clean-boilerplate/boilerplate/src/app/command"
	"clean-boilerplate/boilerplate/src/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateExample command.CreateExampleHandler
	UpdateExample command.UpdateExampleHandler
}

type Queries struct {
	ListExample query.ListExampleHandler
	GetExample  query.GetExampleHandler
}
