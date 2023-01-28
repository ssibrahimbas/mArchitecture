package mongo

import (
	"context"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestDB_TransformId(t *testing.T) {
	type fields struct {
		c   *mongo.Client
		db  *mongo.Database
		ctx context.Context
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    primitive.ObjectID
		wantErr bool
	}{
		{
			name: "Test TransformId with valid id",
			fields: fields{
				c:   nil,
				db:  nil,
				ctx: context.TODO(),
			},
			args: args{
				id: "000000000000000000000000",
			},
			want:    primitive.NilObjectID,
			wantErr: false,
		},
		{
			name: "Test TransformId with invalid id",
			fields: fields{
				c:   nil,
				db:  nil,
				ctx: context.TODO(),
			},
			args: args{
				id: "x",
			},
			want:    primitive.NilObjectID,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DB{
				c:   tt.fields.c,
				db:  tt.fields.db,
				ctx: tt.fields.ctx,
			}
			got, err := m.TransformId(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DB.TransformId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.TransformId() = %v, want %v", got, tt.want)
			}
		})
	}
}
