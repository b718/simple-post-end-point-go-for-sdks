package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

type PostRequestResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(PostRequestResponse{
			StatusCode: http.StatusMethodNotAllowed,
			Message:    "Method not allowed",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(PostRequestResponse{
		StatusCode: http.StatusOK,
		Message:    "OK",
	})
}

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	http.HandleFunc("/", postHandler)

	go func() {
		if err := http.ListenAndServe("0.0.0.0:4040", nil); err != nil {
			log.Fatal(err)
		}
	}()

	<-signals
	fmt.Println("Shutting down...")
}
