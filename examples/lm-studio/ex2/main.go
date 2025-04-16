package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"log"
)

func main() {
	client := http.Client{
		Timeout: 120 * time.Second,
	}
	reqBody := make(map[string]any)
	reqBody["model"] = "gemma-2-2b-it"
	reqBody["prompt"] = []string{"Give me a list of 3 cryptocurrencies", "Please only present the list in HTML only."}
	reqBody["stream"] = false

	b, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatal(err)
	}
	reader := bytes.NewReader(b)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:1234/v1/completions", reader)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
