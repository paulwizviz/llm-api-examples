package main

import (
	"fmt"
	"llma/internal/golm"
	"log"
)

func main() {
	client := golm.NewClient("http://localhost:1234", 120)
	models, err := client.ListModels()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(models)
}
