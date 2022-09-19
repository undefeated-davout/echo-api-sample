package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	wantPort := 1234
	t.Setenv("PORT", fmt.Sprint(wantPort))

	tests := []struct {
		name    string
		want    *Config
		wantErr bool
	}{
		{
			name: "正常",
			want: &Config{Env: "dev", Port: 1234, DBHost: "db", DBName: "sample", DBPort: 3306, DBUser: "sample",
				DBPassword: "sample"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New()
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
