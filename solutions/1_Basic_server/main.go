package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type StatusResponse struct {
	Status string
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	status := StatusResponse{Status: "I'm a teapot 🫖"}

	jsonBytes, err := json.Marshal(status)
	if err != nil {
		panic("couldn't marshal")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(418)
	w.Write(jsonBytes)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(handleRoot))

	fmt.Println("Started sever on port 4321")

	http.ListenAndServe("localhost:4321", mux)
}
