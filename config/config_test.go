package config

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	t.Setenv("ENV", "dev")
	t.Setenv("PORT", "1234")
	t.Setenv("DB_HOST", "testhost")
	t.Setenv("DB_NAME", "testname")
	t.Setenv("DB_PORT", "3307")
	t.Setenv("DB_USER", "testdbuser")
	t.Setenv("DB_PASSWORD", "testdbpassword")

	tests := []struct {
		name    string
		want    *Config
		wantErr bool
	}{
		{
			name: "正常",
			want: &Config{Env: "dev", Port: 1234, DBHost: "testhost", DBName: "testname", DBPort: 3307, DBUser: "testdbuser",
				DBPassword: "testdbpassword"},
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
