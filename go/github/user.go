package github

// User represents a User returned by the github
// API. Currently implements just enough to capture the
// 'actor' field returned by the Event API.
//
// https://developer.github.com/v3/users/
type User struct {
	ID         *int    `json:"id,omitempty"`
	Login      *string `json:"login,omitempty"`
	GravatarID *string `json:"gravatar_id,omitempty"`
	URL        *string `json:"url,omitempty"`
	AvatarURL  *string `json:"avatar_url,omitempty"`
}
