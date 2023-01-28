package mapper

import "github.ssibrahimbas/mArchitecture/auth/delivery/http/mapper/user"

type Mapper struct {
	User user.Mapper
}

func New() *Mapper {
	return &Mapper{
		User: user.New(),
	}
}
