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

type mockGetUserUnreadMessages struct {
	mock.Mock
}

func (m *mockGetUserUnreadMessages) Call(ctx context.Context, method string, url string, header http.Header, body interface{}, result interface{}) *errors.Error {
	args := m.Called(ctx, method, url, header, body, result)
	if args.Get(0) != nil {
		return args.Get(0).(*errors.Error)
	}

	result.(*UserUnreadMessagesResponse).UnreadCount = 10

	return nil
}

func TestGetUserUnreadMessages(t *testing.T) {
	ctx := context.Background()

	type args struct {
		ctx    context.Context
		params UserUnreadMessagesParams
	}

	tests := []struct {
		name         string
		args         args
		setupMock    func(m *mockGetUserUnreadMessages)
		wantExpected UserUnreadMessagesResponse
		wantError    *errors.Error
	}{
		{
			name: "success",
			args: args{ctx: ctx, params: UserUnreadMessagesParams{UserID: "111001100"}},
			setupMock: func(m *mockGetUserUnreadMessages) {
				m.On(
					"Call",
					ctx,
					http.MethodGet,
					fmt.Sprintf("/v3/users/%s/unread_message_count", "111001100"),
					http.Header(nil),
					nil,
					&UserUnreadMessagesResponse{},
				).Return(nil)
			},
			wantExpected: UserUnreadMessagesResponse{UnreadCount: 10},
			wantError:    nil,
		},
		{
			name: "failed user not found",
			args: args{ctx: ctx, params: UserUnreadMessagesParams{UserID: "000000000"}},
			setupMock: func(m *mockGetUserUnreadMessages) {
				m.On(
					"Call",
					ctx,
					http.MethodGet,
					fmt.Sprintf("/v3/users/%s/unread_message_count", "000000000"),
					http.Header(nil),
					nil,
					&UserUnreadMessagesResponse{},
				).Return(errors.FromHTTPErr(400, []byte(`
					{
						"message": "\"User\" not found.",
						"code": 400201,
						"error": true
					}
				`)))
			},
			wantExpected: UserUnreadMessagesResponse{},
			wantError: errors.FromHTTPErr(400, []byte(`
			{
				"message": "\"User\" not found.",
				"code": 400201,
				"error": true
			}
		`)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := new(mockGetUserUnreadMessages)
			tt.setupMock(mockClient)

			chatService := New(mockClient)

			resp, err := chatService.GetUserUnreadMessages(tt.args.ctx, tt.args.params)
			assert.Equal(t, tt.wantExpected, resp)
			assert.Equal(t, tt.wantError, err)
		})
	}
}

type mockCreateUser struct {
	mock.Mock
}

func (m *mockCreateUser) Call(ctx context.Context, method string, url string, header http.Header, body interface{}, result interface{}) *errors.Error {
	args := m.Called(ctx, method, url, header, body, result)
	if args.Get(0) != nil {
		return args.Get(0).(*errors.Error)
	}

	result.(*CreateUserResponse).UserID = "testname"
	result.(*CreateUserResponse).Nickname = "Test Name"
	result.(*CreateUserResponse).ProfileURL = "http://profile_url.com/my_profile.png"

	return nil
}

func TestCreateUser(t *testing.T) {
	ctx := context.Background()

	type args struct {
		ctx     context.Context
		request CreateUserRequest
	}

	tests := []struct {
		name         string
		args         args
		setupMock    func(m *mockCreateUser)
		wantExpected CreateUserResponse
		wantError    *errors.Error
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				request: CreateUserRequest{
					UserID:     "testname",
					Nickname:   "Test Name",
					ProfileURL: "http://profile_url.com/my_profile.png",
				},
			},
			setupMock: func(m *mockCreateUser) {
				m.On(
					"Call",
					ctx,
					http.MethodPost,
					"/v3/users",
					http.Header(nil),
					CreateUserRequest{
						UserID:     "testname",
						Nickname:   "Test Name",
						ProfileURL: "http://profile_url.com/my_profile.png",
					},
					&CreateUserResponse{},
				).Return(nil)
			},
			wantExpected: CreateUserResponse{
				UserID:     "testname",
				Nickname:   "Test Name",
				ProfileURL: "http://profile_url.com/my_profile.png",
			},
			wantError: nil,
		},
		{
			name: "failed",
			args: args{
				ctx: ctx,
				request: CreateUserRequest{
					UserID:     "testduplicatename",
					Nickname:   "Test Duplicate Name",
					ProfileURL: "http://profile_url.com/my_profile.png",
				},
			},
			setupMock: func(m *mockCreateUser) {
				m.On(
					"Call",
					ctx,
					http.MethodPost,
					"/v3/users",
					http.Header(nil),
					CreateUserRequest{
						UserID:     "testduplicatename",
						Nickname:   "Test Duplicate Name",
						ProfileURL: "http://profile_url.com/my_profile.png",
					},
					&CreateUserResponse{},
				).Return(errors.FromHTTPErr(400, []byte(`
				  {
					"message":"\"user_id\" violates unique constraint.",
					"code":400202,
					"error":true
				  }
				`)))
			},
			wantExpected: CreateUserResponse{},
			wantError: errors.FromHTTPErr(400, []byte(`
				{
					"message":"\"user_id\" violates unique constraint.",
					"code":400202,
					"error":true
				}
			`)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := new(mockCreateUser)
			tt.setupMock(mockClient)

			chatService := New(mockClient)

			resp, err := chatService.CreateUser(tt.args.ctx, tt.args.request)
			assert.Equal(t, tt.wantExpected, resp)
			assert.Equal(t, tt.wantError, err)
		})
	}
}

type mockGetUserJoinedChannelCount struct {
	mock.Mock
}

func (m *mockGetUserJoinedChannelCount) Call(ctx context.Context, method string, url string, header http.Header, body interface{}, result interface{}) *errors.Error {
	args := m.Called(ctx, method, url, header, body, result)
	if args.Get(0) != nil {
		return args.Get(0).(*errors.Error)
	}
	result.(*UserJoinedChannelCountResponse).GroupChannelCount = 5
	return nil
}

func TestGetUserJoinedChannelCount(t *testing.T) {
	ctx := context.Background()

	type args struct {
		ctx    context.Context
		params UserJoinedChannelCountParams
	}

	tests := []struct {
		name         string
		args         args
		setupMock    func(m *mockGetUserJoinedChannelCount)
		wantExpected UserJoinedChannelCountResponse
		wantError    *errors.Error
	}{
		{
			name: "success",
			args: args{ctx: ctx, params: UserJoinedChannelCountParams{UserID: "111001100"}},
			setupMock: func(m *mockGetUserJoinedChannelCount) {
				m.On(
					"Call",
					ctx,
					http.MethodGet,
					fmt.Sprintf("/v3/users/%s/group_channel_count", "111001100"),
					http.Header(nil),
					nil,
					&UserJoinedChannelCountResponse{},
				).Return(nil)
			},
			wantExpected: UserJoinedChannelCountResponse{GroupChannelCount: 5},
			wantError:    nil,
		},
		{
			name: "failed user not found",
			args: args{ctx: ctx, params: UserJoinedChannelCountParams{UserID: "000000000"}},
			setupMock: func(m *mockGetUserJoinedChannelCount) {
				m.On(
					"Call",
					ctx,
					http.MethodGet,
					fmt.Sprintf("/v3/users/%s/group_channel_count", "000000000"),
					http.Header(nil),
					nil,
					&UserJoinedChannelCountResponse{},
				).Return(errors.FromHTTPErr(400, []byte(`
					{
						"message": "\"User\" not found.",
						"code": 400201,
						"error": true
					}
				`)))
			},
			wantExpected: UserJoinedChannelCountResponse{},
			wantError: errors.FromHTTPErr(400, []byte(`
			{
				"message": "\"User\" not found.",
				"code": 400201,
				"error": true
			}
		`)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := new(mockGetUserJoinedChannelCount)
			tt.setupMock(mockClient)

			chatService := New(mockClient)

			resp, err := chatService.GetUserJoinedChannelCount(tt.args.ctx, tt.args.params)
			assert.Equal(t, tt.wantExpected, resp)
			assert.Equal(t, tt.wantError, err)
		})
	}
}
