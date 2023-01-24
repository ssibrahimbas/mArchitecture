package mysql

import (
	"clean-boilerplate/boilerplate/src/config"
	"clean-boilerplate/boilerplate/src/domain/example"
	"context"
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	sqb_go "gitlab.com/ssibrahimbas/sqb.go"
)

type exampleRepo struct {
	db             *sqlx.DB
	exampleFactory example.Factory
}

func NewExampleRepo(db *sqlx.DB, exampleFactory example.Factory) example.Repository {
	if db == nil {
		panic("db is nil")
	}
	if exampleFactory.IsZero() {
		panic("exampleFactory is zero")
	}
	return &exampleRepo{
		db:             db,
		exampleFactory: exampleFactory,
	}
}

func (r *exampleRepo) Get(ctx context.Context, key string) (*example.Example, error) {
	e := example.Example{}
	query := sqb_go.QB.Table("example").Where("key", "=", key).Get()
	err := r.db.GetContext(ctx, &e, query)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, r.exampleFactory.NewNotFoundError(key)
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to get example")
	}
	return &e, nil
}

func (r *exampleRepo) Create(ctx context.Context, e *example.Example) error {
	query := sqb_go.QB.Table("example").Insert(&sqb_go.M{
		"key":   e.Key,
		"value": e.Value,
	})
	_, err := r.db.ExecContext(ctx, query)
	if err != nil {
		return errors.Wrap(err, "failed to create example")
	}
	return nil
}

func (r *exampleRepo) Update(ctx context.Context, e *example.Example) error {
	query := sqb_go.QB.Table("example").Where("key", "=", e.Key).Update(&sqb_go.M{
		"value": e.Value,
	})
	_, err := r.db.ExecContext(ctx, query)
	if err != nil {
		return errors.Wrap(err, "failed to update example")
	}
	return nil
}

func (r *exampleRepo) Delete(ctx context.Context, key string) error {
	query := sqb_go.QB.Table("example").Where("key", "=", key).Delete()
	_, err := r.db.ExecContext(ctx, query)
	if err != nil {
		return errors.Wrap(err, "failed to delete example")
	}
	return nil
}

func (r *exampleRepo) List(ctx context.Context, limit int, offset int) ([]*example.Example, int, error) {
	var examples []*example.Example
	query := sqb_go.QB.Table("example").Limit(limit).Offset(offset).Get()
	err := r.db.SelectContext(ctx, &examples, query)
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to list examples")
	}
	total, err := r.Count(ctx)
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to count examples")
	}
	return examples, total, nil
}

func (r *exampleRepo) Count(ctx context.Context) (int, error) {
	var count int
	query := sqb_go.QB.Table("example").Count("Id", "count").Get()
	err := r.db.GetContext(ctx, &count, query)
	if err != nil {
		return 0, errors.Wrap(err, "failed to count examples")
	}
	return count, nil
}

func New(cnf config.MySQL) (*sqlx.DB, error) {
	config := mysql.NewConfig()
	config.Net = "tcp"
	config.Addr = cnf.Address
	config.User = cnf.Username
	config.Passwd = cnf.Password
	config.DBName = cnf.Database
	config.ParseTime = true

	db, err := sqlx.Open("mysql", config.FormatDSN())
	if err != nil {
		return nil, errors.Wrap(err, "failed to open mysql connection")
	}
	return db, nil
}
