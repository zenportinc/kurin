package engine

import (
	"github.com/zenportinc/kurin/example/domain"
)

type (
	CreateUserRequest struct {
		Username string `json:"username"`
		Email    string `json:"email" valid:"email"`
	}
)

func (engine *exampleEngine) GetUser(id string) (*domain.User, error) {
	user := engine.userRepository.Get(id)
	if user == nil {
		return nil, NewNotFound(id, "user")
	}

	return user, nil
}

func (engine *exampleEngine) CreateUser(r *CreateUserRequest) (*domain.User, error) {
	valid, _, err := engine.validator.Validate(r)
	if err != nil {
		return nil, err
	}

	if !valid {
		return nil, NewInvalid(r, err)
	}

	user := &domain.User{
		Email:    r.Email,
		Username: r.Username,
	}
	user.GenerateId()

	engine.userRepository.Create(user)

	return user, nil
}

func (engine *exampleEngine) DeleteUser(id string) error {
	user, err := engine.GetUser(id)
	if err != nil {
		return err
	}

	engine.userRepository.Delete(user)

	return nil
}

func (engine *exampleEngine) ListUsers() []*domain.User {
	return engine.userRepository.List()
}
