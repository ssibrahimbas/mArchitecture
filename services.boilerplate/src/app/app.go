package app

import (
	"clean-boilerplate/boilerplate/src/app/command"
	"clean-boilerplate/boilerplate/src/app/query"
)

type Application struct {
	Commands commands
	Queries  queries
}

type commands struct {
	CreateExample command.CreateExampleHandler
	UpdateExample command.UpdateExampleHandler
}

type queries struct {
	ListExample query.ListExampleHandler
	GetExample  query.GetExampleHandler
}
