package github

// Organisation represents an Organisation returned by the github
// API. Currently implements just enough to capture the
// 'repo' field returned by the Event API.
//
// https://developer.github.com/v3/orgs/
type Organisation struct {
	ID         *int    `json:"id,omitempty"`
	Login      *string `json:"login,omitempty"`
	GravatarID *string `json:"gravatar_id,omitempty"`
	URL        *string `json:"url,omitempty"`
	AvatarURL  *string `json:"avatar_url,omitemtpy"`
}
