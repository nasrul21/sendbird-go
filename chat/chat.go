package chat

import (
	"context"

	"github.com/nasrul21/sendbird-go/client"
	"github.com/nasrul21/sendbird-go/errors"
)

type Chat interface {
	// user
	GetUserUnreadMessages(ctx context.Context, params UserUnreadMessagesParams) (UserUnreadMessages, *errors.Error)
	GetUserJoinedChannelCount(ctx context.Context, params UserJoinedChannelCountParams) (resp UserJoinedChannelCountResponse, err *errors.Error)
	CreateUser(ctx context.Context, request CreateUserRequest) (resp CreateUserResponse, err *errors.Error)

	// channel
	UpdateGroupChannel(ctx context.Context, params UpdateGroupChannelParams, request UpdateGroupChannelRequest) (resp UpdateGroupChannelResponse, err *errors.Error)
}

type ChatImpl struct {
	Client client.Client
}

func New(client client.Client) Chat {
	return &ChatImpl{
		Client: client,
	}
}
