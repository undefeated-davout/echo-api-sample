package gateways

import (
	"context"
	"undefeated-davout/echo-api-sample/config"
	"undefeated-davout/echo-api-sample/entities"
	"undefeated-davout/echo-api-sample/interface_adapters/controllers"
	"undefeated-davout/echo-api-sample/interface_adapters/gateways/auth"
	"undefeated-davout/echo-api-sample/interface_adapters/gateways/repositories"
	"undefeated-davout/echo-api-sample/usecases"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewRouter(ctx context.Context, e *echo.Echo, db *gorm.DB, cfg *config.Config) error {
	clocker := entities.RealClocker{}
	repo := &repositories.Repository{Clocker: clocker}
	rcli, err := repositories.NewKVS(ctx, cfg)
	if err != nil {
		return err
	}
	jwter, err := auth.NewJWTer(rcli, clocker)
	if err != nil {
		return err
	}

	healthController := &controllers.HealthController{}
	e.GET("/health", healthController.CheckHealth)

	taskController := &controllers.TaskController{
		ListTaskUsecase: usecases.ListTaskUsecase{DB: db, Repo: repo},
		AddTaskUsecase:  usecases.AddTaskUsecase{DB: db, Repo: repo},
	}
	e.GET("/tasks", taskController.ListTasks)
	e.POST("/tasks", taskController.AddTask)

	userController := &controllers.UserController{
		AddUserUsecase: usecases.AddUserUsecase{DB: db, Repo: repo},
	}
	e.POST("/users", userController.AddUser)

	authController := &controllers.AuthController{
		LoginUsecase: usecases.LoginUsecase{DB: db, Repo: repo, TokenGenerator: jwter},
	}
	e.GET("/login", authController.Login)
	return nil
}
