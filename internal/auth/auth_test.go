package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name         string
		setupRequest func(r *http.Request)
		want         string
		wantErr      bool
	}{
		{
			name: "no auth header",
			setupRequest: func(r *http.Request) {
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "malformed auth header",
			setupRequest: func(r *http.Request) {
				r.Header.Set("Authorization", "Bearer")
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "valid auth header",
			setupRequest: func(r *http.Request) {
				r.Header.Set("Authorization", "Bearer test-api-key")
			},
			want:    "test-api-key",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/", nil)
			tt.setupRequest(req)

			got, err := GetAPIKey(req.Header)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
