package entity

import "clean-boilerplate/boilerplate/src/domain/example"

type MySQLExampleEntity struct {
	UUID  string `db:"uuid"`
	Key   string `db:"e_key"`
	Value string `db:"e_value"`
}

func (e *MySQLExampleEntity) ToExample() *example.Example {
	return &example.Example{
		UUID:  e.UUID,
		Key:   e.Key,
		Value: e.Value,
	}
}
