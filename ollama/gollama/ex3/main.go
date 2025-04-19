package main

import (
	"fmt"
	"llma/internal/gollama"
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
	b64Img := gollama.EncodeBase64(img)
	if err != nil {
		log.Fatal(err)
	}
	timeoutDuration := 120 //seconds
	client := gollama.NewDefaultClient(timeoutDuration, gollama.DefaultLocalBaseURL)
	reqBody := gollama.RequestGenerate{
		Model:  "llava",
		Prompt: "Is this a picture of a cat?",
		Images: []string{b64Img},
	}
	resp, err := client.GenerateAPI(reqBody)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Response:", resp)
}
