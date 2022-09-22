package usecases

import (
	"context"
	"fmt"
	"undefeated-davout/echo-api-sample/entities"

	"gorm.io/gorm"
)

type AddTaskUsecase struct {
	DB   *gorm.DB
	Repo TaskAdder
}

func (a *AddTaskUsecase) AddTask(ctx context.Context, userID entities.UserID, title string, status string) (*entities.Task, error) {
	task := &entities.Task{
		UserID: userID,
		Title:  title,
		Status: entities.TaskStatusTodo,
	}

	err := a.Repo.AddTask(ctx, a.DB, task)
	if err != nil {
		return nil, fmt.Errorf("failed to register: %w", err)
	}
	return task, nil
}
