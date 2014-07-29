package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"code.google.com/p/goauth2/oauth"
)

// EventFetcher manages access to Githubs event list API for an organisation.
// https://developer.github.com/v3/activity/events/#list-events-for-an-organization
type EventFetcher struct {
	previousHeaders responseHeaders
}

type responseHeaders struct {
	ETag   string
	Status string
}

// NewEventFetcher returns an instance of EventFetcher that is
// polling for new events.
func NewEventFetcher(notifier chan Event) (eventFetcher *EventFetcher) {
	eventFetcher = &EventFetcher{}

	go eventFetcher.poll(notifier)

	return
}

// poll loops infinitely, fetching new events every minute.
func (eventFetcher *EventFetcher) poll(notifier chan Event) {
	for {

		for _, event := range eventFetcher.events() {
			notifier <- event
		}

		time.Sleep(time.Minute * 1)
	}
}

func (eventFetcher *EventFetcher) events() (events [30]Event) {

	// Set OAuth access token
	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: os.Getenv("GITHUB_ACCESS_TOKEN")},
	}

	uri := fmt.Sprintf("https://api.github.com/users/%s/events/orgs/%s", os.Getenv("GITHUB_USER"), os.Getenv("GITHUB_ORG"))

	req, err := http.NewRequest("GET", uri, nil)
	if eventFetcher.previousHeaders.ETag != "" {
		req.Header.Set("If-None-Match", eventFetcher.previousHeaders.ETag)
	}

	resp, err := t.Client().Do(req)

	if err != nil {
		log.Printf("Could not fetch %s", uri)
		return events
	}

	status := resp.Header.Get("Status")

	switch status {
	case "304 Not Modified":
		log.Printf("No new events received for %s", os.Getenv("GITHUB_ORG"))
		eventFetcher.previousHeaders.Status = status
		return
	default:
		eventFetcher.previousHeaders.Status = status
		eventFetcher.previousHeaders.ETag = resp.Header.Get("ETag")
		log.Printf("New events received for %s", os.Getenv("GITHUB_ORG"))
	}

	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Could not read response body for %s", uri)
		return events
	}

	if err := json.Unmarshal(contents, &events); err != nil {
		panic(err)
	}

	event := events[0]

	log.Printf("ID is %s", *event.ID)
	log.Printf("Type is %s", *event.Type)
	log.Printf("Actor is %s", *event.Actor)
	log.Printf("Repo is %s", *event.Repo)
	log.Printf("Payload is %s", *event.RawPayload)
	log.Printf("Public is %s", *event.Public)
	log.Printf("Created At is %s", *event.CreatedAt)
	log.Printf("Org is %s", *event.Org)

	return events
}
