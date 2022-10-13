package chat

import (
	"context"
	"fmt"
	"net/http"

	"github.com/nasrul21/sendbird-go/client"
	"github.com/nasrul21/sendbird-go/errors"
)

type Chat interface {
	GetUserUnreadMessages(ctx context.Context, userID string) (UnreadMessages, *errors.Error)
}

type ChatImpl struct {
	Client client.Client
}

func New(client client.Client) Chat {
	return &ChatImpl{
		Client: client,
	}
}

func (c *ChatImpl) GetUserUnreadMessages(ctx context.Context, userID string) (resp UnreadMessages, err *errors.Error) {
	url := fmt.Sprintf("/v3/users/%s/unread_message_count", userID)
	err = c.Client.Call(ctx, http.MethodGet, url, nil, nil, &resp)
	if err != nil {
		return
	}

	return
}
