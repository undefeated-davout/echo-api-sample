package entities

import (
	"time"
)

type TaskID int64
type TaskStatus string

const (
	TaskStatusTodo  TaskStatus = "todo"
	TaskStatusDoing TaskStatus = "doing"
	TaskStatusDone  TaskStatus = "done"
)

type Task struct {
	ID        TaskID `gorm:"primaryKey"`
	UserID    UserID
	Title     string
	Status    TaskStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}
