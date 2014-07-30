package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	_ "github.com/joho/godotenv/autoload"
	"github.com/quamen/sup/go/broker"
	"github.com/quamen/sup/go/github"
)

var port = os.Getenv("PORT")

func main() {
	broker := broker.NewServer()
	github.NewEventFetcher(broker.Notifier)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleIndex)
	mux.HandleFunc("/events/", broker.ServeHTTP)

	n := negroni.Classic()
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
