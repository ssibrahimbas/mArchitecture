package entity

import (
	mysql_migration "clean-boilerplate/shared/migration/mysql"

	"github.com/jmoiron/sqlx"
)

type MySQLExampleMigration struct {
	Database string
}

func NewExampleMigration() mysql_migration.MigrationItem {
	return &MySQLExampleMigration{
		Database: "boilerplate.example",
	}
}

func (m *MySQLExampleMigration) Up(db *sqlx.DB) error {
	query := `CREATE TABLE IF NOT EXISTS boilerplate.example (
		uuid VARCHAR(255) NOT NULL PRIMARY KEY,
		e_key VARCHAR(255) NOT NULL,
		e_value VARCHAR(255) NOT NULL
	)`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (m *MySQLExampleMigration) Down(db *sqlx.DB) error {
	query := `DROP TABLE boilerplate.example`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (m *MySQLExampleMigration) GetDatabase() string {
	return m.Database
}
