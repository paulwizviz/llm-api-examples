package main

import (
	"fmt"
	"llma/internal/ollama"
	"log"
)

type reqbody struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type respbody struct {
	Model    string `json:"model"`
	Created  string `json:"created_at"`
	Response string `json:"response"`
}

func main() {

	timeoutDuration := 120 //seconds
	client := ollama.NewDefaultClient(timeoutDuration, ollama.DefaultLocalBaseURL, "llama2")
	resp, err := client.Generate("what is life?")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)

}
