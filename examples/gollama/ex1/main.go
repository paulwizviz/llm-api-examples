package main

import (
	"fmt"
	"llma/internal/ollama"
	"log"
)

func main() {

	timeoutDuration := 120 //seconds
	client := ollama.NewDefaultClient(timeoutDuration, ollama.DefaultLocalBaseURL)
	reqBody := ollama.RequestGenerate{
		Model:  "llama3.2:1b",
		Prompt: "What is the formula for calculating the area of a rectangle?",
	}
	resp, err := client.Generate(reqBody)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*resp)

}
