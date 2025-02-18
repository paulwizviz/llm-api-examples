package main

import (
	"encoding/json"
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
	respReader, err := client.GenerateAPI(reqBody)
	if err != nil {
		log.Fatal(err)
	}
	defer respReader.Close()

	decoder := json.NewDecoder(respReader)

	var resp ollama.ResponseGenerate
	err = decoder.Decode(&resp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)

}
