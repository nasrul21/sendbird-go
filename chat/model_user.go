package chat

type UserUnreadMessages struct {
	UnreadCount int `json:"unread_count"`
}

type UserUnreadMessagesParams struct {
	UserID string `json:"user_id"`
}

type UserJoinedChannelCountParams struct {
	UserID string `json:"user_id"`
}

type UserJoinedChannelCountResponse struct {
	GroupChannelCount int `json:"group_channel_count"`
}

type CreateUserRequest struct {
	UserID           string            `json:"user_id"`
	Nickname         string            `json:"nickname"`
	ProfileURL       string            `json:"profile_url"`
	IssueAccessToken bool              `json:"issue_access_token,omitempty"`
	Metadata         map[string]string `json:"metadata,omitempty"`
}

type UserResource struct {
	UserID                     string            `json:"user_id"`
	Nickname                   string            `json:"nickname"`
	ProfileURL                 string            `json:"profile_url"`
	AccessToken                string            `json:"access_token"`
	IsOnline                   bool              `json:"is_online"`
	IsActive                   bool              `json:"is_active"`
	IsCreated                  bool              `json:"is_created"`
	PhoneNumber                string            `json:"phone_number"`
	RequireAuthForProfileImage bool              `json:"require_auth_for_profile_image"`
	LastSeenAt                 int               `json:"last_seen_at"`
	DiscoveryKeys              []string          `json:"discovery_keys"`
	PreferredLanguages         []interface{}     `json:"preferred_languages"`
	HasEverLoggedIn            bool              `json:"has_ever_logged_in"`
	Metadata                   map[string]string `json:"metadata"`
}

type CreateUserResponse UserResource
