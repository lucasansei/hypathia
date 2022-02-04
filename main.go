package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"database/sql"

	_ "github.com/lib/pq"
)

type httpResponse struct {
	Message string
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	healthCheckResponse := httpResponse{"health OK"}

	data, err := json.Marshal(healthCheckResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func messengerWebhookHandler(w http.ResponseWriter, r *http.Request) {
	res := httpResponse{"OK"}

	data, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func twitterWebhookHandler(w http.ResponseWriter, r *http.Request) {
	res := httpResponse{"OK"}

	data, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func main() {
	connStr := "user=postgres dbname=chatbot-foo sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	row := db.QueryRow("SELECT 1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(row)

	handler := http.NewServeMux()
	handler.HandleFunc("/health-check", healthCheckHandler)
	handler.HandleFunc("/webhooks/facebook", messengerWebhookHandler)
	handler.HandleFunc("/webhooks/twitter", twitterWebhookHandler)

	log.Fatal(http.ListenAndServe(":2049", handler))
}
