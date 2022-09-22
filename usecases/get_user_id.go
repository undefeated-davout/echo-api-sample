package usecases

import (
	"context"
	"fmt"
	"undefeated-davout/echo-api-sample/entities"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type GetUserIDUsecase struct {
	DB    *gorm.DB
	Repo  UserGetter
	JWTer UserNameGetter
}

func (l *GetUserIDUsecase) GetUserID(c echo.Context, ctx context.Context) (entities.UserID, error) {
	userName := l.JWTer.GetUserName(c)

	user, err := l.Repo.GetUserByName(ctx, l.DB, userName)
	if err != nil {
		return 0, fmt.Errorf("failed to list: %w", err)
	}

	return user.ID, nil
}
