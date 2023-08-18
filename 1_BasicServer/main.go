package main

import (
	"encoding/json"
	"net/http"
)

type MyStruct struct {
	A int `json:"NotA"`
}

func main() {

	myStruct := MyStruct{A: 5}

	json, err := json.Marshal(myStruct)
	if err != nil {
		panic("couldn't unmarshal")
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(json))
	}))

	http.ListenAndServe("localhost:3000", mux)
}
