package main

import (
	"bizdaysapi/bizdays"
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
	fmt.Fprintf(w, "%d", bizdays.Between(fromDate, toDate))
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
