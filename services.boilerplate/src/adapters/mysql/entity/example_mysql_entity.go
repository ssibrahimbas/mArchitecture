package entity

type MySQLExampleEntity struct {
	UUID  string `db:"uuid"`
	Key   string `db:"key"`
	Value string `db:"value"`
}
