package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	_ "github.com/joho/godotenv/autoload"
	"github.com/nabeken/negroni-auth"
	"github.com/quamen/sup/go/broker"
	"github.com/quamen/sup/go/github"
	"github.com/quamen/sup/go/redis"
)

var port = os.Getenv("PORT")

func main() {
	broker := broker.NewServer()
	github.NewEventFetcher(broker.Notifier)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleIndex)
	mux.HandleFunc("/events/", handleEvents)
	mux.HandleFunc("/events/stream/", broker.ServeHTTP)

	n := negroni.Classic()
	n.Use(auth.Basic(os.Getenv("USERNAME"), os.Getenv("PASSWORD")))
	n.UseHandler(mux)
	n.Run(fmt.Sprintf(":%s", port))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("go/templates/index.html")

	if err != nil {
		log.Fatal("Error parsing templates/index.html")
	}

	t.Execute(w, "")
}

func handleEvents(w http.ResponseWriter, r *http.Request) {

	results, err := redis.FetchReversedCollectionFromSortedSet("EVENTS", 30)
	if err != nil {
		log.Fatal(err)
	}

	var events []github.Event

	w.Header().Set("Content-Type", "application/javascript")

	for _, event := range results {
		var a github.Event
		if err := json.Unmarshal(event.([]byte), &a); err != nil {
			panic(err)
		}
		events = append(events, a)
	}

	enc := json.NewEncoder(w)
	enc.Encode(events)

}
