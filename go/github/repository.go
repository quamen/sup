package github

// Repository represents a Repository returned by the github
// API. Currently implements just enough to capture the
// 'repo' field returned by the Event API.
//
// https://developer.github.com/v3/repos/
type Repository struct {
	ID   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	URL  *string `json:"url,omitempty"`
}
