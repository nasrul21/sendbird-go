package sendbird

import (
	"net/http"
	"testing"

	"github.com/nasrul21/sendbird-go/client"
	"github.com/stretchr/testify/assert"
)

func TestSendbird(t *testing.T) {
	apiToken := "XXX_XXX_XXX"
	baseURL := "https://sendbird_dummy_url.com"

	sendbirdClient := New(&Option{
		APIToken: apiToken,
		BaseURL:  baseURL,
	})

	wantExpected := &Sendbird{
		option: &Option{
			APIToken: apiToken,
			BaseURL:  baseURL,
		},
		httpClient: client.New(
			&http.Client{},
			sendbirdClient.option.BaseURL,
			map[string]string{
				"Content-Type": "application/json",
				"Api-Token":    apiToken,
			},
		),
	}

	wantExpected.init()

	assert.Equal(t, wantExpected, sendbirdClient)
}
