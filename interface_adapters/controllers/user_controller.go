package controllers

import (
	"net/http"
	"undefeated-davout/echo-api-sample/entities"
	"undefeated-davout/echo-api-sample/usecases"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	AddUserUsecase usecases.AddUserUsecase
}

// ユーザ登録
func (t *UserController) AddUser(c echo.Context) error {
	name := c.FormValue("name")
	password := c.FormValue("password")
	role := c.FormValue("role")

	user, err := t.AddUserUsecase.AddUser(c.Request().Context(), name, password, role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, struct {
		ID entities.UserID `json:"id"`
	}{ID: user.ID})
}
