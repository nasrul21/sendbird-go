package chat

type UpdateGroupChannelParams struct {
	ChannelURL string
}

type UpdateGroupChannelRequest struct {
	Name        *string   `json:"name,omitempty"`
	CoverURL    *string   `json:"cover_url,omitempty"`
	CoverFile   *string   `json:"cover_file,omitempty"`
	CostumType  *string   `json:"custom_type,omitempty"`
	Data        *string   `json:"data,omitempty"`
	IsDistinct  *bool     `json:"is_distinct,omitempty"`
	IsPublic    *bool     `json:"is_public,omitempty"`
	AccessCode  *string   `json:"access_code,omitempty"`
	OperatorIDs *[]string `json:"operator_ids,omitempty"`
}

type ChannelResource struct {
	Name                 string `json:"name"`
	ChannelURL           string `json:"channel_url"`
	CoverURL             string `json:"cover_url"`
	CustomType           string `json:"custom_type"`
	UnreadMessageCount   int    `json:"unread_message_count"`
	Data                 string `json:"data"`
	IsDistinct           bool   `json:"is_distinct"`
	IsPublic             bool   `json:"is_public"`
	IsSuper              bool   `json:"is_super"`
	IsEphemeral          bool   `json:"is_ephemeral"`
	IsAccessCodeRequired bool   `json:"is_access_code_required"`
	HiddenState          string `json:"hidden_state"`
	MemberCount          int    `json:"member_count"`
	JoinedMemberCount    int    `json:"joined_member_count"`
	Members              []struct {
		UserID             string   `json:"user_id"`
		Nickname           string   `json:"nickname"`
		ProfileURL         string   `json:"profile_url"`
		IsActive           bool     `json:"is_active"`
		IsOnline           bool     `json:"is_online"`
		FriendDiscoveryKey []string `json:"friend_discovery_key"`
		LastSeenAt         int64    `json:"last_seen_at"`
		State              string   `json:"state"`
		Role               string   `json:"role"`
		Metadata           struct {
			Location string `json:"location"`
			Marriage string `json:"marriage"`
		} `json:"metadata"`
	} `json:"members"`
	Operators []struct {
		UserID     string `json:"user_id"`
		Nickname   string `json:"nickname"`
		ProfileURL string `json:"profile_url"`
		Metadata   struct {
			Location string `json:"location"`
			Marriage string `json:"marriage"`
		} `json:"metadata"`
	} `json:"operators"`
	MaxLengthMessage int         `json:"max_length_message"`
	LastMessage      interface{} `json:"last_message"`
	CreatedAt        int         `json:"created_at"`
	Freeze           bool        `json:"freeze"`
}

type UpdateGroupChannelResponse ChannelResource
