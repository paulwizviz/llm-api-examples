package main

import (
	"encoding/json"
	"fmt"
	"llma/internal/ollama"
	"log"
	"os"
	"path/filepath"
)

func readImage(fname string) ([]byte, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	imageFile := filepath.Join(pwd, "testdata", fname)
	data, err := os.ReadFile(imageFile)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func main() {

	img, err := readImage("cat.jpeg")
	b64Img := ollama.EncodeBase64(img)
	if err != nil {
		log.Fatal(err)
	}
	timeoutDuration := 120 //seconds
	client := ollama.NewDefaultClient(timeoutDuration, ollama.DefaultLocalBaseURL)
	reqBody := ollama.RequestGenerate{
		Model:  "llava",
		Prompt: "Is this a picture of a cat?",
		Images: []string{b64Img},
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
