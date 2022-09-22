package controllers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"undefeated-davout/echo-api-sample/entities"
	customValidator "undefeated-davout/echo-api-sample/interface_adapters/gateways/custom_validator"
	"undefeated-davout/echo-api-sample/interface_adapters/gateways/repositories"
	"undefeated-davout/echo-api-sample/usecases"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestUserController_AddUser(t *testing.T) {
	e := echo.New()
	v := customValidator.NewValidator()
	e.Validator = v

	moqDB := &DBerMock{}
	moqRepo := &UserAdderMock{}
	moqRepo.AddUserFunc = func(ctx context.Context, db repositories.DBer, u *entities.User) error {
		u.ID = 1
		return nil
	}
	uc := usecases.AddUserUsecase{DB: moqDB, Repo: moqRepo}

	type fields struct {
		AddUserUsecase usecases.AddUserUsecase
		Validator      *customValidator.CustomValidator
	}
	type args struct {
		body string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		statusCode int
		response   string
		wantErr    bool
	}{
		{
			name:       "正常登録",
			fields:     fields{Validator: v, AddUserUsecase: uc},
			args:       args{body: `{"name":"hoge", "password":"fuga", "role":"normal"}`},
			statusCode: http.StatusOK,
			response:   "{\"id\":1}\n",
			wantErr:    false,
		},
		{
			name:       "エラー：name不足",
			fields:     fields{Validator: v, AddUserUsecase: uc},
			args:       args{body: `{"password":"fuga", "role":"normal"}`},
			statusCode: http.StatusBadRequest,
			response:   "",
			wantErr:    false,
		},
		{
			name:       "エラー：password不足",
			fields:     fields{Validator: v, AddUserUsecase: uc},
			args:       args{body: `{"name":"hoge", "role":"normal"}`},
			statusCode: http.StatusBadRequest,
			response:   "",
			wantErr:    false,
		},
		{
			name:       "エラー：role不足",
			fields:     fields{Validator: v, AddUserUsecase: uc},
			args:       args{body: `{"name":"hoge", "password":"fuga"`},
			statusCode: http.StatusBadRequest,
			response:   "",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.args.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			c := e.NewContext(req, rec)

			tr := &UserController{
				AddUserUsecase: tt.fields.AddUserUsecase,
				Validator:      tt.fields.Validator,
			}
			if err := tr.AddUser(c); (err != nil) != tt.wantErr {
				t.Errorf("UserController.AddUser() error = %v, wantErr %v", err, tt.wantErr)
			}

			assert.Equal(t, tt.statusCode, rec.Code)
			if tt.response != "" {
				assert.Equal(t, tt.response, rec.Body.String())
			}
		})
	}
}
