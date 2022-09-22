package controllers

import (
	"net/http"
	customValidator "undefeated-davout/echo-api-sample/interface_adapters/gateways/custom_validator"
	"undefeated-davout/echo-api-sample/interface_adapters/gateways/request"
	"undefeated-davout/echo-api-sample/interface_adapters/presenters/response"
	"undefeated-davout/echo-api-sample/usecases"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	Validator    *customValidator.CustomValidator
	LoginUsecase usecases.LoginUsecase
}

// ログイン処理
func (t *AuthController) Login(c echo.Context) error {
	req := new(request.LoginRequest)
	if err := t.Validator.GetValidatedRequest(c, req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	token, err := t.LoginUsecase.Login(c.Request().Context(), req.Name, req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, response.LoginResponse{AccessToken: token})
}
