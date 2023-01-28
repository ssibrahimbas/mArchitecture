package entity

import "github.ssibrahimbas/mArchitecture/auth/domain/user"

type MongoUser struct {
	UUID      string `bson:"uuid"`
	Email     string `bson:"email"`
	IsActive  bool   `bson:"is_active"`
	CreatedAt string `bson:"created_at"`
	UpdatedAt string `bson:"updated_at"`
}

func (m *MongoUser) ToUser() *user.User {
	return &user.User{
		UUID:     m.UUID,
		Email:    m.Email,
		IsActive: m.IsActive,
	}
}

func (m *MongoUser) FromUser(user *user.User) *MongoUser {
	m.UUID = user.UUID
	m.Email = user.Email
	m.IsActive = user.IsActive
	return m
}
