package controllers

import (
	"net/http"
	"undefeated-davout/echo-api-sample/entities"
	"undefeated-davout/echo-api-sample/usecases"

	"github.com/labstack/echo/v4"
)

type TaskController struct {
	ListTaskUsecase  usecases.ListTaskUsecase
	AddTaskUsecase   usecases.AddTaskUsecase
	GetUserIDUsecase usecases.GetUserIDUsecase
}

type task struct {
	ID     entities.TaskID     `json:"id"`
	Title  string              `json:"title"`
	Status entities.TaskStatus `json:"status"`
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
	rsp := []task{}
	for _, t := range tasks {
		rsp = append(rsp, task{
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
	title := c.FormValue("title")
	status := c.FormValue("status")

	userID, err := t.GetUserIDUsecase.GetUserID(c, ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	task, err := t.AddTaskUsecase.AddTask(ctx, userID, title, status)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, struct {
		ID entities.TaskID `json:"id"`
	}{ID: task.ID})
}
