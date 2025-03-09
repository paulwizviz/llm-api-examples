package main

import (
	"fmt"
	"llma/internal/golm"
	"log"
)

func main() {
	client := golm.NewClient("http://localhost:1234", 120)
	req := golm.CompletionRequest{
		Model:  "gemma-2-2b-it",
		Prompt: []string{"Give me a list of 3 cryptocurrencies", "Please only present the list in HTML only."},
		Stream: true,
	}
	resp, err := client.Completion(req)
	if err != nil {
		log.Fatal(err)
	}
	for _, r := range resp {
		fmt.Println(r)
	}
}
