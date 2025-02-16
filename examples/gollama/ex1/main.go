package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

	b, err := json.Marshal(reqbody{
		Model:  "llama2",
		Prompt: "What is life?",
		Stream: false,
	})
	if err != nil {
		log.Fatal(err)
	}

	breader := bytes.NewReader(b)

	resp, err := http.Post("http://localhost:11434/api/generate", "application/json", breader)
	if err != nil {
		log.Fatalf("Unable to post: %v", err)
	}
	defer resp.Body.Close()

	rBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Unable to read body: %v", err)
	}

	var data respbody
	err = json.Unmarshal(rBody, &data)
	if err != nil {
		log.Fatalf("Unable to unmarshal response body: %v", err)
	}

	fmt.Printf("Response is: %v", data)
}
