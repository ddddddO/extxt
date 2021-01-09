package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBasicAuthenticated(t *testing.T) {
	tests := []struct {
		name     string
		userName string
		password string
		want     bool
	}{
		{
			name:     "Success",
			userName: "aaa",
			password: "pass",
			want:     true,
		},
		{
			name:     "invalid user name",
			userName: "bbb",
			password: "pass",
			want:     false,
		},
		{
			name:     "invalid password",
			userName: "aaa",
			password: "pppppp",
			want:     false,
		},
		{
			name:     "no set basic auth",
			userName: "",
			password: "",
			want:     false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
			if tt.userName != "" && tt.password != "" {
				req.SetBasicAuth(tt.userName, tt.password)
			}
			got := basicAuthenticated(req)
			if got != tt.want {
				t.Errorf("\nwant: %v\ngot: %v\n", tt.want, got)
			}
		})
	}
}
