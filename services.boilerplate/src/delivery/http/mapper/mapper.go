package mapper

import example_mapper "github.ssibrahimbas/mArchitecture/boilerplate/src/delivery/http/mapper/example"

type Mapper struct {
	Example example_mapper.Mapper
}

func New() *Mapper {
	return &Mapper{
		Example: example_mapper.New(),
	}
}
