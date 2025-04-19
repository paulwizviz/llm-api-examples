package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ollama/ollama/api"
)

func main() {

	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}
	isStream := false
	req := &api.GenerateRequest{
		Model:  "llama3.2:1b",
		Prompt: "What is the formula for calculating the area of a rectangle?",
		Stream: &isStream,
	}

	err = client.Generate(context.TODO(), req, func(resp api.GenerateResponse) error {
		fmt.Println(resp)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
