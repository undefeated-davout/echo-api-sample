package customvalidator

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"undefeated-davout/echo-api-sample/interface_adapters/gateways/request"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func TestCustomValidator_GetValidatedRequest(t *testing.T) {
	e := echo.New()
	v := NewValidator()
	e.Validator = v

	createContext := func(body string) echo.Context {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		return c
	}

	type fields struct {
		validator *validator.Validate
	}
	type args struct {
		c   echo.Context
		req interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "正常",
			fields:  fields{validator: v.validator},
			args:    args{c: createContext(`{"name":"hoge","password":"fuga"}`), req: &request.LoginRequest{}},
			wantErr: false,
		},
		{
			name:    "エラー：パラメータ不足",
			fields:  fields{validator: v.validator},
			args:    args{c: createContext(`{"name":"hoge"}`), req: &request.LoginRequest{}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cv := &CustomValidator{
				validator: tt.fields.validator,
			}
			if err := cv.GetValidatedRequest(tt.args.c, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("CustomValidator.GetValidatedRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
