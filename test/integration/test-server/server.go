package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type AgeResponse struct {
	Count uint64
	Name  string
	Age   int
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	ageResponse := AgeResponse{
		Count: 1000,
		Name:  name,
		Age:   62,
	}
	jsonResponse, err := json.Marshal(ageResponse)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	_, err = io.WriteString(w, string(jsonResponse))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func main() {
	http.HandleFunc("/", getRoot)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
