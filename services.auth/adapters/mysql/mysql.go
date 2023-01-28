package mysql

import (
	"github.com/jmoiron/sqlx"
	mysql_user "github.ssibrahimbas/mArchitecture/auth/adapters/mysql/user"
	"github.ssibrahimbas/mArchitecture/auth/domain/user"
)

type MySQL interface {
	NewUser(userFactory user.Factory, sql *sqlx.DB) user.Repository
}

type mySql struct{}

func New() MySQL {
	return &mySql{}
}

func (m *mySql) NewUser(userFactory user.Factory, sql *sqlx.DB) user.Repository {
	return mysql_user.New(userFactory, sql)
}
