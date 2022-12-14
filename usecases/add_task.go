package usecases

import (
	"context"
	"fmt"
	"undefeated-davout/echo-api-sample/entities"
	"undefeated-davout/echo-api-sample/interface_adapters/gateways/repositories"
)

type AddTaskUsecase struct {
	DB   repositories.DBer
	Repo TaskAdder
}

func (a *AddTaskUsecase) AddTask(ctx context.Context, userID entities.UserID, title string, status entities.TaskStatus) (*entities.Task, error) {
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
