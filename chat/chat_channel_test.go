package chat

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/nasrul21/sendbird-go/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockUpdateGroupChannel struct {
	mock.Mock
}

func (m *mockUpdateGroupChannel) Call(ctx context.Context, method string, url string, header http.Header, body interface{}, result interface{}) *errors.Error {
	args := m.Called(ctx, method, url, header, body, result)
	if args.Get(0) != nil {
		return args.Get(0).(*errors.Error)
	}

	result.(*UpdateGroupChannelResponse).Name = "Test Channel"
	result.(*UpdateGroupChannelResponse).ChannelURL = "sendbird_group_channel_xxxx"
	result.(*UpdateGroupChannelResponse).IsPublic = true

	return nil
}

func TestUpdateGroupChannel(t *testing.T) {
	ctx := context.Background()

	type args struct {
		ctx     context.Context
		params  UpdateGroupChannelParams
		request UpdateGroupChannelRequest
	}

	channelName := "Test Channel"
	isPublic := true

	tests := []struct {
		name         string
		args         args
		setupMock    func(m *mockUpdateGroupChannel)
		wantExpected UpdateGroupChannelResponse
		wantError    *errors.Error
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				params: UpdateGroupChannelParams{
					ChannelURL: "sendbird_group_channel_xxxx",
				},
				request: UpdateGroupChannelRequest{
					Name:     &channelName,
					IsPublic: &isPublic,
				},
			},
			setupMock: func(m *mockUpdateGroupChannel) {
				m.On(
					"Call",
					ctx,
					http.MethodPut,
					fmt.Sprintf("/v3/group_channels/%s", "sendbird_group_channel_xxxx"),
					http.Header(nil),
					UpdateGroupChannelRequest{
						Name:     &channelName,
						IsPublic: &isPublic,
					},
					&UpdateGroupChannelResponse{},
				).Return(nil)
			},
			wantExpected: UpdateGroupChannelResponse{
				Name:       "Test Channel",
				ChannelURL: "sendbird_group_channel_xxxx",
				IsPublic:   true,
			},
			wantError: nil,
		},
		{
			name: "failed",
			args: args{
				ctx: ctx,
				params: UpdateGroupChannelParams{
					ChannelURL: "invalid_sendbird_group_channel_xxxx",
				},
				request: UpdateGroupChannelRequest{
					Name:     &channelName,
					IsPublic: &isPublic,
				},
			},
			setupMock: func(m *mockUpdateGroupChannel) {
				m.On(
					"Call",
					ctx,
					http.MethodPut,
					fmt.Sprintf("/v3/group_channels/%s", "invalid_sendbird_group_channel_xxxx"),
					http.Header(nil),
					UpdateGroupChannelRequest{
						Name:     &channelName,
						IsPublic: &isPublic,
					},
					&UpdateGroupChannelResponse{},
				).Return(errors.FromHTTPErr(400, []byte(`
					{
						"message": "\"Channel\" not found.",
						"code": 400201,
						"error": true
					}
				`)))
			},
			wantExpected: UpdateGroupChannelResponse{},
			wantError: errors.FromHTTPErr(400, []byte(`
				{
					"message": "\"Channel\" not found.",
					"code": 400201,
					"error": true
				}
			`)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := new(mockUpdateGroupChannel)
			tt.setupMock(mockClient)

			chatService := New(mockClient)

			resp, err := chatService.UpdateGroupChannel(tt.args.ctx, tt.args.params, tt.args.request)
			assert.Equal(t, tt.wantExpected, resp)
			assert.Equal(t, tt.wantError, err)
		})
	}
}
