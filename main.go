package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

var socket = ":8080"

type Response struct {
	TimeStamp time.Time `json:"time_stamp"`
	Hostname  string    `json:"hostname"`
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Create the data
	hn, _ := os.Hostname()
	resp := Response{TimeStamp: time.Now(), Hostname: hn}

	// 2. Set the header before writing the response
	w.Header().Set("Content-Type", "application/json")

	// 3. Encode and send the response
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/", jsonHandler)
	log.Printf("standing up server on %v\n", socket)
	if err := http.ListenAndServe(socket, nil); err != nil {
		log.Printf("failed to stand up server: %v\n", err)
	}
}
