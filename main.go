package main

import (
	"log"

	"github.com/damingerdai/tour/cmd"
)

var name string

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute error: %v", err)
	}
}
