package controllers

import (
	"net/http"
	"undefeated-davout/echo-api-sample/entities"
	customValidator "undefeated-davout/echo-api-sample/interface_adapters/gateways/custom_validator"
	"undefeated-davout/echo-api-sample/interface_adapters/gateways/request"
	"undefeated-davout/echo-api-sample/interface_adapters/presenters/response"
	"undefeated-davout/echo-api-sample/usecases"

	"github.com/labstack/echo/v4"
)

type TaskController struct {
	Validator        *customValidator.CustomValidator
	ListTaskUsecase  usecases.ListTaskUsecase
	AddTaskUsecase   usecases.AddTaskUsecase
	GetUserIDUsecase usecases.GetUserIDUsecase
}

// タスク取得
func (t *TaskController) ListTasks(c echo.Context) error {
	ctx := c.Request().Context()

	userID, err := t.GetUserIDUsecase.GetUserID(c, ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	tasks, err := t.ListTaskUsecase.ListTasks(c.Request().Context(), userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	rsp := response.ListTasksResponse{}
	for _, t := range tasks {
		rsp = append(rsp, response.TaskResponse{
			ID:     t.ID,
			Title:  t.Title,
			Status: t.Status,
		})
	}
	return c.JSON(http.StatusOK, rsp)
}

// タスク登録
func (t *TaskController) AddTask(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(request.AddTaskRequest)
	if err := t.Validator.GetValidatedRequest(c, req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	userID, err := t.GetUserIDUsecase.GetUserID(c, ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	task, err := t.AddTaskUsecase.AddTask(ctx, userID, req.Title, entities.TaskStatus(req.Status))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, response.AddTaskResponse{ID: task.ID})
}
