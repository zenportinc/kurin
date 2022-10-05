package engine

import "github.com/zenportinc/kensho"

type (
	Factory interface {
		NewEngine() Engine
	}

	engineFactory struct {
		ExampleProviderFactory
	}
)

func NewFactory(e ExampleProviderFactory) Factory {
	return &engineFactory{e}
}

func (f *engineFactory) NewEngine() Engine {
	validator := kensho.NewValidator()
	return &exampleEngine{
		userRepository: f.NewUserRepository(),
		validator:      validator,
	}
}
