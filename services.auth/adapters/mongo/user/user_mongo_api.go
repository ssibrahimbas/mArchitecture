package user

import (
	"context"

	"github.ssibrahimbas/mArchitecture/auth/adapters/mongo/user/entity"
	"github.ssibrahimbas/mArchitecture/auth/domain/user"
	"github.ssibrahimbas/mArchitecture/shared/i18n"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *repo) Create(ctx context.Context, user *user.User) *i18n.I18nError {
	e := r.checkExist(ctx, user.Email, false)
	if e != nil {
		return r.userFactory.Errors.AlreadyExists(user.Email)
	}
	u := &entity.MongoUser{}
	res, err := r.collection.InsertOne(ctx, u.FromUser(user))
	if err != nil {
		return r.userFactory.Errors.Failed("create")
	}
	user.UUID = res.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

func (r *repo) Update(ctx context.Context, user *user.User) *i18n.I18nError {
	u := &entity.MongoUser{}
	e := r.checkExist(ctx, user.Email, true)
	if e != nil {
		return e
	}
	res, err := r.collection.UpdateOne(ctx, bson.M{"email": user.Email}, bson.M{"$set": u.FromUser(user)})
	if err != nil {
		return r.userFactory.Errors.Failed("update")
	}
	if res.MatchedCount == 0 {
		return r.userFactory.Errors.NotFound(user.Email)
	}
	return nil
}

func (r *repo) Get(ctx context.Context, email string) (*user.User, *i18n.I18nError) {
	u := &entity.MongoUser{}
	res := r.collection.FindOne(ctx, bson.M{"email": email})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return nil, r.userFactory.Errors.NotFound(email)
		}
		return nil, r.userFactory.Errors.Failed("get")
	}
	err := res.Decode(u)
	if err != nil {
		return nil, r.userFactory.Errors.Failed("get")
	}
	return u.ToUser(), nil
}

func (r *repo) checkExist(ctx context.Context, email string, throwNotFound bool) *i18n.I18nError {
	res := r.collection.FindOne(ctx, bson.M{"email": email})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments && throwNotFound {
			return r.userFactory.Errors.NotFound(email)
		}
		return r.userFactory.Errors.Failed("get")
	}
	return nil
}