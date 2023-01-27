package command

import (
	"context"
	"testing"

	"clean-boilerplate/boilerplate/src/config"
	"clean-boilerplate/boilerplate/src/domain/example"
	example_mocks "clean-boilerplate/boilerplate/src/domain/example/mocks"
	"clean-boilerplate/shared/events"
	event_mocks "clean-boilerplate/shared/events/mocks"
	"clean-boilerplate/shared/i18n"
)

func Test_updateExampleHandler_Handle(t *testing.T) {
	mock := preUpdatedTest()

	type fields struct {
		exampleRepo   example.Repository
		exampleTopics config.ExampleTopics
		publisher     events.Publisher
	}
	type args struct {
		ctx     context.Context
		command UpdateExampleCommand
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "update example",
			fields: fields{
				exampleRepo:   mock.repo,
				exampleTopics: mock.topics,
				publisher:     mock.publisher,
			},
			args: args{
				ctx: mock.ctx,
				command: UpdateExampleCommand{
					Field:   "field",
					Content: "content",
				},
			},
			wantErr: false,
		},
		{
			name: "update example error",
			fields: fields{
				exampleRepo:   mock.repo,
				exampleTopics: mock.topics,
				publisher:     mock.publisher,
			},
			args: args{
				ctx: mock.ctx,
				command: UpdateExampleCommand{
					Field:   "field2",
					Content: "content2",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := updateExampleHandler{
				exampleRepo:   tt.fields.exampleRepo,
				exampleTopics: tt.fields.exampleTopics,
				publisher:     tt.fields.publisher,
			}
			if err := h.Handle(tt.args.ctx, tt.args.command); (err != nil) != tt.wantErr {
				t.Errorf("updateExampleHandler.Handle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type updateTestMocks struct {
	repo      *example_mocks.Repository
	publisher *event_mocks.Publisher
	topics    config.ExampleTopics
	ctx       context.Context
}

func preUpdatedTest() updateTestMocks {
	ctx := context.Background()

	repo := &example_mocks.Repository{}
	publisher := &event_mocks.Publisher{}

	topics := config.ExampleTopics{
		Created: "Example.Updated",
	}

	publisher.On("Publish", topics.Updated, &example.Example{
		Field:   "field",
		Content: "content",
		UUID:    "",
	}).Return(nil)

	repo.On("Update", ctx, &example.Example{
		Field:   "field",
		Content: "content",
	}).Return(nil)

	repo.On("Update", ctx, &example.Example{
		Field:   "field2",
		Content: "content2",
	}).Return(i18n.NewError("error"))

	return updateTestMocks{
		repo:      repo,
		publisher: publisher,
		topics:    topics,
		ctx:       ctx,
	}
}
