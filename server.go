package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	_ "github.com/joho/godotenv/autoload"
	"github.com/quamen/sup/go/github"
)

var port = os.Getenv("PORT")

func main() {

	go func() {
		// Loop for ever
		eventFetcher := github.NewEventFetcher()
		for {
			eventFetcher = eventFetcher.Events()
			// Wait 1 minute before looping again
			time.Sleep(time.Minute * 1)
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleIndex)

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(fmt.Sprintf(":%s", port))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")

	if err != nil {
		log.Fatal("Error parsing templates/index.html")
	}

	t.Execute(w, "")
}
