package usecases

import (
	"context"
	"fmt"
	"undefeated-davout/echo-api-sample/entities"
	"undefeated-davout/echo-api-sample/interface_adapters/gateways/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AddUserUsecase struct {
	DB   repositories.DBer
	Repo UserAdder
}

func (a *AddUserUsecase) AddUser(ctx context.Context, name string, password string, role string) (*entities.User, error) {
	hasshedPassowrd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to register: %w", err)
	}

	user := &entities.User{
		Name:     name,
		Password: string(hasshedPassowrd),
		Role:     role,
	}

	if err := a.Repo.AddUser(ctx, a.DB, user); err != nil {
		return nil, fmt.Errorf("failed to register: %w", err)
	}
	return user, nil
}
