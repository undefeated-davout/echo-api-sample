package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHealthController_CheckHealth(t *testing.T) {
	healthJSON := "{\"status\":\"ok\"}\n"

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(healthJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		h       *HealthController
		args    args
		wantErr bool
	}{
		{name: "正常",
			args:    args{c: c},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HealthController{}
			if err := h.CheckHealth(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("HealthController.CheckHealth() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, healthJSON, rec.Body.String())
		})
	}
}
