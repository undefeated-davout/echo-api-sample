package entities

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserID int64

type User struct {
	ID        UserID `gorm:"primaryKey"`
	Name      string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) ComparePassword(pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pw))
}
