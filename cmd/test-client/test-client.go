package main

import (
	"log"
	http_client "testcontainers/internal/http-client"
)

func main() {
	client, err := http_client.NewClient("https://api.agify.io/")
	if err != nil {
		log.Fatalf("NewClient() failed: %v", err)
	}

	ageResponse, err := client.GetAge("Sig")
	if err != nil {
		log.Fatalf("GetAge() failed: %v", err)
	}

	log.Printf("AgeResponse: %+v", ageResponse)
}
