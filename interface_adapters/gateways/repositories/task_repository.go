package repositories

import (
	"context"
	"undefeated-davout/echo-api-sample/entities"

	"gorm.io/gorm"
)

func (r *Repository) ListTasks(ctx context.Context, db *gorm.DB, userID entities.UserID) ([]entities.Task, error) {
	tasks := []entities.Task{}
	if err := db.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		return tasks, err
	}
	return tasks, nil
}

func (r *Repository) AddTask(ctx context.Context, db *gorm.DB, t *entities.Task) error {
	if err := db.Create(t).Error; err != nil {
		return err
	}
	return nil
}
