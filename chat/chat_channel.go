package chat

import (
	"context"
	"fmt"
	"net/http"

	"github.com/nasrul21/sendbird-go/errors"
)

func (c *ChatImpl) UpdateGroupChannel(ctx context.Context, params UpdateGroupChannelParams, request UpdateGroupChannelRequest) (resp UpdateGroupChannelResponse, err *errors.Error) {
	url := fmt.Sprintf("/v3/group_channels/%s", params.ChannelURL)
	err = c.Client.Call(ctx, http.MethodPut, url, nil, request, &resp)
	if err != nil {
		return
	}
	return
}
