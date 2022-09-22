package usecases

import (
	"context"
	"fmt"
	"undefeated-davout/echo-api-sample/entities"

	"gorm.io/gorm"
)

type ListTaskUsecase struct {
	DB   *gorm.DB
	Repo TaskLister
}

func (l *ListTaskUsecase) ListTasks(ctx context.Context, userID entities.UserID) ([]entities.Task, error) {
	tasks, err := l.Repo.ListTasks(ctx, l.DB, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}
	return tasks, nil
}
