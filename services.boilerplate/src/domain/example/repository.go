package example

import "context"

type Repository interface {
	Get(ctx context.Context, key string) (*Example, error)
	List(ctx context.Context, limit, offset int) ([]*Example, error)

	Create(ctx context.Context, example *Example) error
	Update(ctx context.Context, example *Example) error
}
