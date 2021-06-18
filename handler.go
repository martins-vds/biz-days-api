package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func bizdaysHandler(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	if from == "" || to == "" {
		http.Error(w, "'from' and 'to' dates must be provided", http.StatusBadRequest)
		return
	}

	fromDate, err := time.Parse(time.RFC3339, from)
	toDate, err := time.Parse(time.RFC3339, to)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	fmt.Fprintf(w, "%f", diff(fromDate, toDate))
}

func diff(a, b time.Time) (days float64) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}

	days = b.Sub(a).Hours() / 24

	return
}

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	http.HandleFunc("/api/bizdays", bizdaysHandler)
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
