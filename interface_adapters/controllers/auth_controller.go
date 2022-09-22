package controllers

import (
	"net/http"
	"undefeated-davout/echo-api-sample/usecases"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	LoginUsecase usecases.LoginUsecase
}

// ログイン処理
func (t *AuthController) Login(c echo.Context) error {
	name := c.FormValue("name")
	password := c.FormValue("password")

	token, err := t.LoginUsecase.Login(c.Request().Context(), name, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: token,
	})
}
