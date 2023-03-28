package main

import (
	"gRPC/internal/app"
	"log"
)

func main() {
	if err := app.RunHttpServer(); err != nil {
		log.Fatal(err)
	}
}
