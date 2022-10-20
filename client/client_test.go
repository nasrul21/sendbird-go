package client_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/nasrul21/sendbird-go/client"
	"github.com/nasrul21/sendbird-go/errors"
	"github.com/stretchr/testify/assert"
)

func TestCall(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name         string
		setupHandler func(w http.ResponseWriter, r *http.Request)
		wantExpected map[string]interface{}
		wantErr      *errors.Error
	}{
		{
			name: "success",
			setupHandler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{"hello": "world"}`))
			},
			wantExpected: map[string]interface{}{
				"hello": "world",
			},
			wantErr: nil,
		},
		{
			name: "error failed decode",
			setupHandler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(``))
			},
			wantExpected: nil,
			wantErr:      errors.FromGoErr(fmt.Errorf("unexpected end of JSON input")),
		},
		{
			name: "error http response",
			setupHandler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{"error": "true", "message": "internal server error", "code": "500221"}`))
			},
			wantExpected: nil,
			wantErr:      errors.FromHTTPErr(500, []byte(`{"error": "true", "message": "internal server error", "code": "500221"}`)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(tt.setupHandler))
			defer server.Close()

			var result map[string]interface{}

			client := ClientImpl{
				HttpClient: &http.Client{},
			}
			err := client.Call(ctx, http.MethodGet, server.URL, http.Header{}, nil, &result)

			assert.Equal(t, tt.wantExpected, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
