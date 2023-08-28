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

type furColor struct {
	Primary    string
	Highlights string
}

type coordinates struct {
	Lat float64
	Lng float64
}

type SquirrelSighting struct {
	SquirrelID             string   `json:"Squirrel_ID"`
	FurColor               furColor `json:"Fur_Color"`
	Park                   string
	InteractionsWithHumans string `json:"Interactions_With_Humans"`
	Activities             []string
	Coorditantes           coordinates
}

func handleSquirrels(w http.ResponseWriter, r *http.Request) {
	fileContents, err := os.ReadFile("../../static/data.json")
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(fileContents)
}

func handleSquirrelByID(w http.ResponseWriter, r *http.Request) {
	fileContents, err := os.ReadFile("../../static/data.json")
	if err != nil {
		panic(err)
	}
	list := []SquirrelSighting{}
	err = json.Unmarshal(fileContents, list)
	if err != nil {
		panic(err)
	}

	first := list[0]
	resp, err := json.Marshal(first)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(handleRoot))
	mux.Handle("/squirrels", http.HandlerFunc(handleSquirrels))
	// TODO: we can't handle this in standard lib alone:
	// maybe we switch to using gin framework?
	mux.Handle("/squirrels/:id", http.HandlerFunc(handleSquirrels))

	http.ListenAndServe("localhost:4321", mux)
}
