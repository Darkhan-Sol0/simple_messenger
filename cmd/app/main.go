package main

import (
	"log"
	"simple_messenger/internal/server"
)

func main() {
	if err := server.New().Run(); err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}
