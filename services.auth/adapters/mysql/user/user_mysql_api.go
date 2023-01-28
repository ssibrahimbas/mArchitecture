package user

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.ssibrahimbas/mArchitecture/auth/adapters/mysql/user/entity"
	"github.ssibrahimbas/mArchitecture/auth/domain/user"
	"github.ssibrahimbas/mArchitecture/shared/formats"
	"github.ssibrahimbas/mArchitecture/shared/i18n"
	sqb_go "gitlab.com/ssibrahimbas/sqb.go"
)

func (r *repo) Create(ctx context.Context, user *user.User) *i18n.I18nError {
	e := r.checkExist(ctx, user.Email, false)
	if e != nil {
		return r.userFactory.Errors.AlreadyExists(user.Email)
	}
	t := time.Now().Format(formats.DateYYYYMMDDHHMMSS)
	query := sqb_go.QB.Table(entity.Fields.Table).Insert(&sqb_go.M{
		entity.Fields.UUID:      user.UUID,
		entity.Fields.Email:     user.Email,
		entity.Fields.IsActive:  true,
		entity.Fields.CreatedAt: t,
		entity.Fields.UpdatedAt: t,
	})
	_, err := r.db.ExecContext(ctx, query)
	if err != nil {
		return r.userFactory.Errors.Failed("create")
	}
	return nil
}

func (r *repo) Update(ctx context.Context, user *user.User) *i18n.I18nError {
	e := r.checkExist(ctx, user.Email, true)
	if e != nil {
		return e
	}
	query := sqb_go.QB.Table(entity.Fields.Table).Where(entity.Fields.UUID, "=", user.UUID).Update(&sqb_go.M{
		entity.Fields.Email:     user.Email,
		entity.Fields.IsActive:  user.IsActive,
		entity.Fields.UpdatedAt: time.Now().Format(formats.DateYYYYMMDDHHMMSS),
	})
	_, err := r.db.ExecContext(ctx, query)
	if err != nil {
		return r.userFactory.Errors.Failed("update")
	}
	return nil
}

func (r *repo) Get(ctx context.Context, email string) (*user.User, *i18n.I18nError) {
	e := &entity.MySQLUser{}
	query := sqb_go.QB.Table(entity.Fields.Table).Where(entity.Fields.Email, "=", email).Get()
	err := r.db.GetContext(ctx, e, query)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, r.userFactory.Errors.NotFound(email)
	}
	if err != nil {
		return nil, r.userFactory.Errors.Failed("get")
	}
	return e.ToUser(), nil
}

func (r *repo) checkExist(ctx context.Context, email string, throwNotFound bool) *i18n.I18nError {
	e := &entity.MySQLUser{}
	query := sqb_go.QB.Table(entity.Fields.Table).Where(entity.Fields.Email, "=", email).Get()
	err := r.db.GetContext(ctx, e, query)
	if errors.Is(err, sql.ErrNoRows) && throwNotFound {
		return r.userFactory.Errors.NotFound(email)
	}
	if err != nil {
		return r.userFactory.Errors.Failed("checkExist")
	}
	return nil
}
