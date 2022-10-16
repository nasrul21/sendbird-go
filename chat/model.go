package chat

type UserUnreadMessages struct {
	UnreadCount int `json:"unread_count"`
}

type UserUnreadMessagesParams struct {
	UserID string `json:"user_id"`
}
