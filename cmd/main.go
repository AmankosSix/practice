package main

import (
	"fmt"
	"log"
	"practice/internal/config"
)

func main() {
	status, err := config.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Project status:", status)
}
