package mapper

import example_mapper "clean-boilerplate/boilerplate/src/protocols/http/mapper/example"

type Mapper struct {
	Example example_mapper.Mapper
}

func New() *Mapper {
	return &Mapper{
		Example: example_mapper.New(),
	}
}
