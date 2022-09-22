package usecases

import (
	"context"
	"undefeated-davout/echo-api-sample/entities"
	"undefeated-davout/echo-api-sample/interface_adapters/gateways/repositories"

	"github.com/labstack/echo/v4"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . UserGetter UserAdder TaskLister TaskAdder TokenGenerator UserNameGetter

// ユーザ
type UserGetter interface {
	GetUserByName(ctx context.Context, db repositories.DBer, name string) (*entities.User, error)
}
type UserAdder interface {
	AddUser(ctx context.Context, db repositories.DBer, u *entities.User) error
}

// タスク
type TaskLister interface {
	ListTasks(ctx context.Context, db repositories.DBer, id entities.UserID) ([]entities.Task, error)
}
type TaskAdder interface {
	AddTask(ctx context.Context, db repositories.DBer, t *entities.Task) error
}

// 認証
type TokenGenerator interface {
	GenerateToken(ctx context.Context, u entities.User) (string, error)
}
type UserNameGetter interface {
	GetUserName(c echo.Context) string
}
