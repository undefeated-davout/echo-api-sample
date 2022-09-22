package controllers

import (
	"net/http"
	customValidator "undefeated-davout/echo-api-sample/interface_adapters/gateways/custom_validator"
	"undefeated-davout/echo-api-sample/interface_adapters/gateways/request"
	"undefeated-davout/echo-api-sample/interface_adapters/presenters/response"
	"undefeated-davout/echo-api-sample/usecases"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	AddUserUsecase usecases.AddUserUsecase
	Validator      *customValidator.CustomValidator
}

// ユーザ登録
func (t *UserController) AddUser(c echo.Context) error {
	req := new(request.AddUserRequest)
	if err := t.Validator.GetValidatedRequest(c, req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user, err := t.AddUserUsecase.AddUser(c.Request().Context(), req.Name, req.Password, req.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, response.AddUserResponse{ID: user.ID})
}
