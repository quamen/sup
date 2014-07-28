package github

import (
	"encoding/json"
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
