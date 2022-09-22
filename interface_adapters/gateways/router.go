package gateways

import (
	"context"
	"undefeated-davout/echo-api-sample/config"
	"undefeated-davout/echo-api-sample/entities"
	"undefeated-davout/echo-api-sample/interface_adapters/controllers"
	"undefeated-davout/echo-api-sample/interface_adapters/gateways/auth"
	customValidator "undefeated-davout/echo-api-sample/interface_adapters/gateways/custom_validator"
	"undefeated-davout/echo-api-sample/interface_adapters/gateways/repositories"
	"undefeated-davout/echo-api-sample/usecases"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"gorm.io/gorm"
)

func NewRouter(ctx context.Context, e *echo.Echo, db *gorm.DB, cfg *config.Config) error {
	validator := customValidator.NewValidator()
	e.Validator = validator
	clocker := entities.RealClocker{}
	repo := &repositories.Repository{Clocker: clocker}
	redisStore, err := repositories.NewKVS(ctx, cfg)
	if err != nil {
		return err
	}
	jwter := auth.NewJWTer(redisStore, clocker, cfg.JWTSecretKey)

	healthController := &controllers.HealthController{}
	e.GET("/health", healthController.CheckHealth)

	userController := &controllers.UserController{
		Validator:      validator,
		AddUserUsecase: usecases.AddUserUsecase{DB: db, Repo: repo},
	}
	e.POST("/users", userController.AddUser)

	authController := &controllers.AuthController{
		Validator:    validator,
		LoginUsecase: usecases.LoginUsecase{DB: db, Repo: repo, TokenGenerator: jwter},
	}
	e.POST("/login", authController.Login)

	// --- 認証あり ---
	config := middleware.JWTConfig{Claims: &auth.JWTCustomClaims{}, SigningKey: []byte(cfg.JWTSecretKey)}
	tg := e.Group("/tasks")
	tg.Use(middleware.JWTWithConfig(config))
	taskController := &controllers.TaskController{
		Validator:        validator,
		ListTaskUsecase:  usecases.ListTaskUsecase{DB: db, Repo: repo},
		AddTaskUsecase:   usecases.AddTaskUsecase{DB: db, Repo: repo},
		GetUserIDUsecase: usecases.GetUserIDUsecase{DB: db, Repo: repo, JWTer: jwter},
	}
	tg.GET("", taskController.ListTasks)
	tg.POST("", taskController.AddTask)

	return nil
}
