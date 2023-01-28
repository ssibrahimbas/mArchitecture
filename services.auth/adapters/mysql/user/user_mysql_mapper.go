package user

import (
	"github.ssibrahimbas/mArchitecture/auth/adapters/mysql/user/entity"
	"github.ssibrahimbas/mArchitecture/auth/domain/user"
	sqb_go "gitlab.com/ssibrahimbas/sqb.go"
)

type Mapper interface {
	ToEntityMap(u *user.User) *sqb_go.M
}

type mapper struct{}

func NewMapper() Mapper {
	return &mapper{}
}

func (m *mapper) ToEntityMap(u *user.User) *sqb_go.M {
	return &sqb_go.M{
		entity.Fields.UUID:      u.UUID,
		entity.Fields.Email:     u.Email,
		entity.Fields.Password:  u.Password,
		entity.Fields.IsActive:  u.IsActive,
		entity.Fields.CreatedAt: u.CreatedAt,
		entity.Fields.UpdatedAt: u.UpdatedAt,
	}
}
