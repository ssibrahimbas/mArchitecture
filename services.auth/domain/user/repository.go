package user

import (
	"context"

	"github.ssibrahimbas/mArchitecture/shared/i18n"
)

type Repository interface {
	Get(ctx context.Context, email string) (*User, *i18n.I18nError)
	Create(ctx context.Context, user *User) *i18n.I18nError
	Update(ctx context.Context, user *User) *i18n.I18nError
}
