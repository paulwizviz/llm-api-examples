package main

import (
	"io"
	"net/http"
	"time"

	"log"
)

func main() {
	client := http.Client{
		Timeout: 120 * time.Second,
	}
	req, err := http.NewRequest(http.MethodGet, "http://localhost:1234/v1/models", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(body))
}
