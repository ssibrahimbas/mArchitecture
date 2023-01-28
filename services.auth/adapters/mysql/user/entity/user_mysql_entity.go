package entity

import "github.ssibrahimbas/mArchitecture/auth/domain/user"

type MySQLUser struct {
	UUID      string `db:"uuid"`
	Email     string `db:"email"`
	Password  []byte `db:"password"`
	IsActive  bool   `db:"is_active"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

type fields struct {
	UUID      string
	Email     string
	Password  string
	IsActive  string
	CreatedAt string
	UpdatedAt string

	Table string
}

var Fields = fields{
	UUID:      "uuid",
	Password:  "password",
	Email:     "email",
	IsActive:  "is_active",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Table:     "users",
}

func (m *MySQLUser) ToUser() *user.User {
	return &user.User{
		UUID:     m.UUID,
		Email:    m.Email,
		Password: m.Password,
		IsActive: m.IsActive,
	}
}
