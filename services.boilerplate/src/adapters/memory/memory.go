package memory

import (
	memory_example "github.ssibrahimbas/mArchitecture/boilerplate/src/adapters/memory/example"
	"github.ssibrahimbas/mArchitecture/boilerplate/src/domain/example"
)

type Memory interface {
	NewExample(exampleFactory example.Factory) example.Repository
}

type memory struct{}

func New() Memory {
	return &memory{}
}

func (m *memory) NewExample(exampleFactory example.Factory) example.Repository {
	return memory_example.NewExampleRepo(exampleFactory)
}
