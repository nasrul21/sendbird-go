package sendbird

import (
	"net/http"

	"github.com/nasrul21/sendbird-go/chat"
	"github.com/nasrul21/sendbird-go/client"
)

type Sendbird struct {
	option     *Option
	httpClient client.Client
	Chat       chat.Chat
}

type Option struct {
	APIToken string
	BaseURL  string
}

func (s *Sendbird) init() {
	s.Chat = chat.New(s.httpClient)
}

func New(option *Option) *Sendbird {
	headers := map[string]string{
		"Content-Type": "application/json",
		"Api-Token":    option.APIToken,
	}
	s := &Sendbird{
		option:     option,
		httpClient: client.New(&http.Client{}, option.BaseURL, headers),
	}

	s.init()

	return s
}

func (s *Sendbird) WithHttpClient(httpClient client.Client) *Sendbird {
	s.httpClient = httpClient
	s.init()
	return s
}
