package usecases

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type LoginUsecase struct {
	DB             *gorm.DB
	Repo           UserGetter
	TokenGenerator TokenGenerator
}

func (l *LoginUsecase) Login(ctx context.Context, name, pw string) (string, error) {
	u, err := l.Repo.GetUserByName(ctx, l.DB, name)
	if err != nil {
		return "", fmt.Errorf("failed to list: %w", err)
	}

	if err = u.ComparePassword(pw); err != nil {
		return "", fmt.Errorf("wrong password: %w", err)
	}

	token, err := l.TokenGenerator.GenerateToken(ctx, *u)
	if err != nil {
		return "", fmt.Errorf("failed to generate JWT: %w", err)
	}

	return token, nil
}
