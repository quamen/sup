package github

import (
	"encoding/json"
	"log"
	"time"
)

// Event represents the generic EventType returned by the github
// API, https://developer.github.com/v3/activity/events/types/
type Event struct {
	ID         *string          `json:"id,omitempty"`
	Type       *string          `json:"type,omitempty"`
	Actor      *User            `json:"actor,omitempty"`
	Repo       *Repository      `json:"repo,omitempty"`
	RawPayload *json.RawMessage `json:"payload,omitempty"`
	Public     *bool            `json:"public"`
	CreatedAt  *time.Time       `json:"created_at,omitempty"`
	Org        *Organisation    `json:"org,omitempty"`
}

// Payload returns the parsed event payload.
func (event *Event) Payload() (payload interface{}) {
	switch *event.Type {
	case "PullRequestReviewCommentEvent":
		payload = &PullRequestReviewCommentEvent{}
	case "IssueCommentEvent":
		payload = &IssueCommentEvent{}
	default:
		return
	}
	if err := json.Unmarshal(*event.RawPayload, &payload); err != nil {
		panic(err.Error())
	}
	return payload
}

// SupportedPayload returns true if the RawPayload can be converted into a
// concrete Payload
func (event *Event) SupportedPayload() (supported bool) {
	switch *event.Type {
	case "PullRequestReviewCommentEvent":
		return true
	case "IssueCommentEvent":
		return true
	default:
		log.Printf("Unsupported Payload: %s", *event.Type)
		return false
	}
}

// PullRequestReviewCommentEvent represents a comment on a pull request.
// Only a subset of the total payload is parsed.
type PullRequestReviewCommentEvent struct {
	ID          *int         `json:"id,omitempty"`
	PullRequest *PullRequest `json:"pull_request,omitemtpy"`
}

// IssueCommentEvent represents a comment on an issue.
// Only a subset of the total payload is parsed.
type IssueCommentEvent struct {
	ID    *int   `json:"id,omitempty"`
	Issue *Issue `json:"issue,omitemtpy"`
}
