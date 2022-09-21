package repositories

import (
	"context"
	"undefeated-davout/echo-api-sample/entities"

	"gorm.io/gorm"
)

func (r *Repository) GetUserByName(ctx context.Context, db *gorm.DB, name string) (*entities.User, error) {
	user := &entities.User{}
	if err := db.Where("name = ?", name).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repository) AddUser(ctx context.Context, db *gorm.DB, u *entities.User) error {
	if err := db.Create(u).Error; err != nil {
		return err
	}
	return nil
}
