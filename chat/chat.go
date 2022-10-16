package chat

import (
	"context"

	"github.com/nasrul21/sendbird-go/client"
	"github.com/nasrul21/sendbird-go/errors"
)

type Chat interface {
	GetUserUnreadMessages(ctx context.Context, params UserUnreadMessagesParams) (UserUnreadMessages, *errors.Error)
}

type ChatImpl struct {
	Client client.Client
}

func New(client client.Client) Chat {
	return &ChatImpl{
		Client: client,
	}
}
