package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page is working!")
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var metadata Metadata
	error := decoder.Decode(&metadata)

	if error != nil {
		fmt.Fprintf(w, "Error: %v", error)
		return
	}

	fmt.Fprintf(w, "Payload %v\n", metadata)
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
