package main

import (
	"log"
	"simple_message/internal/server"
)

func main() {
	if err := server.New().Run(); err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}
