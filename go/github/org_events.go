package github

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

// NewEventFetcher returns an instance of EventFetcher
func NewEventFetcher() (eventFetcher *EventFetcher) {
	eventFetcher = &EventFetcher{}

	return
}

// Events returns a struct representing events from the github API
//
//
// It is stateful, only returning new events as it finds them.
func (eventFetcher *EventFetcher) Events() (events *EventFetcher) {

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

	if err == nil {
		status := resp.Header.Get("Status")

		switch status {
		case "304 Not Modified":
			log.Printf("No new events received for %s", os.Getenv("GITHUB_ORG"))
			eventFetcher.previousHeaders.Status = status
		default:
			eventFetcher.previousHeaders.Status = status
			eventFetcher.previousHeaders.ETag = resp.Header.Get("ETag")
			log.Printf("New events received for %s", os.Getenv("GITHUB_ORG"))

		}

	} else {
		log.Printf("Could not fetch %s", uri)
		return eventFetcher
	}

	return eventFetcher

	/*
	   defer resp.Body.Close()

	   contents, err := ioutil.ReadAll(resp.Body)
	   if err != nil {
	     log.Printf("Could not read response body for %s", uri)
	     return
	   }

	   log.Printf("%s", contents)
	*/
}
