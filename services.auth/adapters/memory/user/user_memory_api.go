package user

import (
	"github.com/google/uuid"
	"github.ssibrahimbas/mArchitecture/auth/domain/user"
	"github.ssibrahimbas/mArchitecture/shared/i18n"
	"golang.org/x/net/context"
)

func (r *repo) Create(ctx context.Context, user *user.User) *i18n.I18nError {
	id := uuid.New().String()
	user.UUID = id
	r.users[id] = *user
	return nil
}

func (r *repo) Update(ctx context.Context, user *user.User) *i18n.I18nError {
	r.users[user.UUID] = *user
	return nil
}

func (r *repo) Get(ctx context.Context, email string) (*user.User, *i18n.I18nError) {
	for _, u := range r.users {
		if u.Email == email {
			return &u, nil
		}
	}
	return nil, r.userFactory.Errors.NotFound(email)
}
