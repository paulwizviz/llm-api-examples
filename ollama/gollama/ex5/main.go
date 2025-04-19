package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/ollama/ollama/api"
)

func main() {

	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}
	format := json.RawMessage([]byte(`{
		"type": "object",
		"properties": {
	  		"height": {
	   			"type": "integer"
	  		},
	  		"width": {
	    		"type": "integer"
	  		},
			"area":{
				"type": "integer"
			}
		},
		"required": [
      		"height",
      		"width",
			"area"
    	]
	}`))
	isStream := false
	req := &api.GenerateRequest{
		Model:  "llama3.2:1b",
		Prompt: "What is the area of a rectangle with height 2m and width 1m? Respond in JSON",
		Format: format,
		Stream: &isStream,
	}

	err = client.Generate(context.TODO(), req, func(resp api.GenerateResponse) error {
		fmt.Println(resp)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	req = &api.GenerateRequest{
		Model:  "llama3.2:1b",
		Prompt: "What is the area of a circle with radis 1m? Respond in JSON",
		Format: json.RawMessage([]byte(`"json"`)),
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
