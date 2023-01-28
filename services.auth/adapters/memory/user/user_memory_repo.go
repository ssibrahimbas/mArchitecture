package user

import "github.ssibrahimbas/mArchitecture/auth/domain/user"

type repo struct {
	userFactory user.Factory
	users       map[string]user.User
}

func New(userFactory user.Factory) user.Repository {
	if userFactory.IsZero() {
		panic("exampleFactory is zero")
	}
	return &repo{
		userFactory: userFactory,
	}
}
