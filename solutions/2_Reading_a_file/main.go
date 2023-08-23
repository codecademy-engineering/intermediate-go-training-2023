package main

import (
	"encoding/json"
	"net/http"
	"os"
)

type StatusResponse struct {
	Status string `json:"status"`
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fileContents, err := os.ReadFile("./greeting.txt")
	status := StatusResponse{Status: string(fileContents)}

	json, err := json.Marshal(status)
	if err != nil {
		panic("couldn't marshal")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(418)
	w.Write([]byte(json))
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(handleRoot))

	http.ListenAndServe("localhost:3000", mux)
}
