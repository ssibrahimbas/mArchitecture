package mysql

import (
	mysql_example "clean-boilerplate/boilerplate/src/adapters/mysql/example"
	"clean-boilerplate/boilerplate/src/domain/example"

	"github.com/jmoiron/sqlx"
)

type MySQL interface {
	NewExample(db *sqlx.DB, exampleFactory example.Factory) example.Repository
}

type mysql struct{}

func New() MySQL {
	return &mysql{}
}

func (m *mysql) NewExample(db *sqlx.DB, exampleFactory example.Factory) example.Repository {
	return mysql_example.NewExampleRepo(db, exampleFactory)
}
