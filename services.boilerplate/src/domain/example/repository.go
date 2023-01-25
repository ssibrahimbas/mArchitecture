package example

import (
	"clean-boilerplate/shared/i18n"
	"context"
)

type Repository interface {
	Get(ctx context.Context, field string) (*Example, *i18n.I18nError)
	List(ctx context.Context, limit, offset int) ([]*Example, int, *i18n.I18nError)

	Create(ctx context.Context, example *Example) *i18n.I18nError
	Update(ctx context.Context, example *Example) *i18n.I18nError
}
