package usecases

import (
	"context"
	"undefeated-davout/echo-api-sample/entities"

	"gorm.io/gorm"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . TaskLister TaskAdder UserGetter UserAdder TokenGenerator

type TaskLister interface {
	ListTasks(ctx context.Context, db *gorm.DB, id entities.UserID) ([]entities.Task, error)
}
type TaskAdder interface {
	AddTask(ctx context.Context, db *gorm.DB, t *entities.Task) error
}

type UserGetter interface {
	GetUserByName(ctx context.Context, db *gorm.DB, name string) (*entities.User, error)
}
type UserAdder interface {
	AddUser(ctx context.Context, db *gorm.DB, u *entities.User) error
}

type TokenGenerator interface {
	GenerateToken(ctx context.Context, u entities.User) ([]byte, error)
}
