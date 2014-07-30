package github

// PullRequest represents a PullRequest returned by the github
// API. Currently implements just enough to capture the
// 'pull_request' field returned by the Event API in a
// PullRequestReviewCommentEvent.
//
// https://developer.github.com/v3/pull_request/
type PullRequest struct {
	ID     *int    `json:"id,omitempty"`
	URL    *string `json:"url,omitempty"`
	Number *int    `json:"number,omitempty"`
	State  *string `json:"state,omitempty"`
	Title  *string `json:"title,omitempty"`
}
