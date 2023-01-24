module clean-boilerplate/boilerplate

go 1.19

replace clean-boilerplate/shared => ../services.shared

require (
	clean-boilerplate/shared v0.0.0-00010101000000-000000000000
	github.com/jmoiron/sqlx v1.3.5
	github.com/pkg/errors v0.9.1
	gitlab.com/ssibrahimbas/sqb.go v0.0.0-20230123070951-1c15112715d6
)

require (
	github.com/go-sql-driver/mysql v1.7.0
	github.com/sirupsen/logrus v1.9.0 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
)
