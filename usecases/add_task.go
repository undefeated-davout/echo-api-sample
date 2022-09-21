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

func (a *AddTaskUsecase) AddTask(ctx context.Context, title string, status string) (*entities.Task, error) {
	// id, ok := auth.GetUserID(ctx)
	// if !ok {
	// 	return nil, fmt.Errorf("user_id not found")
	// }

	var userID entities.UserID = 1

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
