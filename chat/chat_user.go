package chat

import (
	"context"
	"fmt"
	"net/http"

	"github.com/nasrul21/sendbird-go/errors"
)

func (c *ChatImpl) GetUserUnreadMessages(ctx context.Context, params UserUnreadMessagesParams) (resp UserUnreadMessages, err *errors.Error) {
	url := fmt.Sprintf("/v3/users/%s/unread_message_count", params.UserID)
	err = c.Client.Call(ctx, http.MethodGet, url, nil, nil, &resp)
	if err != nil {
		return
	}
	return
}
