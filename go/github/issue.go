package github

// Issue represents an Issue returned by the github
// API. Currently implements just enough to capture the
// 'issue' field returned by the Event API in a
// IssueCommentEvent.
//
// https://developer.github.com/v3/issue/
type Issue struct {
	ID     *int    `json:"id,omitempty"`
	URL    *string `json:"url,omitempty"`
	Number *int    `json:"number,omitempty"`
	State  *string `json:"state,omitempty"`
	Title  *string `json:"title,omitempty"`
}
