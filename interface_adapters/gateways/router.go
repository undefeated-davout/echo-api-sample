package gateways

import (
	"undefeated-davout/echo-api-sample/entities"
	"undefeated-davout/echo-api-sample/interface_adapters/controllers"
	"undefeated-davout/echo-api-sample/interface_adapters/gateways/repositories"
	taskusecases "undefeated-davout/echo-api-sample/usecases/task_usecases"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewRouter(e *echo.Echo, db *gorm.DB) {
	repository := &repositories.Repository{Clocker: entities.RealClocker{}}

	healthController := &controllers.HealthController{}
	e.GET("/health", healthController.CheckHealth)

	taskController := &controllers.TaskController{
		ListTaskUsecase: taskusecases.ListTaskUsecase{DB: db, Repo: repository},
		AddTaskUsecase:  taskusecases.AddTaskUsecase{DB: db, Repo: repository},
	}
	e.GET("/tasks", taskController.ListTasks)
	e.POST("/tasks", taskController.AddTask)
}
