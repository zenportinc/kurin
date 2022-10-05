package engine

import (
	"github.com/zenportinc/kurin/example/domain"
)

type (
	UserRepository interface {
		Get(id string) *domain.User
		Create(user *domain.User)
		Delete(*domain.User)
		List() []*domain.User
	}

	ExampleProviderFactory interface {
		NewUserRepository() UserRepository
	}
)
