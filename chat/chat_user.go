package chat

import (
	"context"
	"fmt"
	"net/http"

	"github.com/nasrul21/sendbird-go/errors"
)

func (c *ChatImpl) CreateUser(ctx context.Context, request CreateUserRequest) (resp CreateUserResponse, err *errors.Error) {
	url := "/v3/users"
	err = c.Client.Call(ctx, http.MethodPost, url, nil, request, &resp)
	if err != nil {
		return
	}
	return
}

func (c *ChatImpl) GetUserUnreadMessages(ctx context.Context, params UserUnreadMessagesParams) (resp UserUnreadMessages, err *errors.Error) {
	url := fmt.Sprintf("/v3/users/%s/unread_message_count", params.UserID)
	err = c.Client.Call(ctx, http.MethodGet, url, nil, nil, &resp)
	if err != nil {
		return
	}
	return
}
