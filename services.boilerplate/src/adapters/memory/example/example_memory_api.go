package memory_example

import (
	"clean-boilerplate/boilerplate/src/domain/example"
	"clean-boilerplate/shared/i18n"
	"context"

	"github.com/google/uuid"
)

func (r *exampleRepo) Create(ctx context.Context, example *example.Example) *i18n.I18nError {
	id := uuid.New().String()
	example.UUID = id
	r.examples[id] = *example
	return nil
}
func (r *exampleRepo) Update(ctx context.Context, example *example.Example) *i18n.I18nError {
	r.examples[example.UUID] = *example
	return nil
}

func (r *exampleRepo) Get(ctx context.Context, field string) (*example.Example, *i18n.I18nError) {
	for _, e := range r.examples {
		if e.Field == field {
			return &e, nil
		}
	}
	return nil, r.exampleFactory.NewNotFoundError(field)
}

func (r *exampleRepo) List(ctx context.Context, limit int, offset int) ([]*example.Example, int, *i18n.I18nError) {
	var examples []*example.Example
	for _, e := range r.examples {
		examples = append(examples, &e)
	}
	return examples, len(examples), nil
}
