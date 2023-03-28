package main

import (
	"gRPC/internal/app"
	"log"
)

func main() {
	if err := app.RunGrpcServer(); err != nil {
		log.Fatal(err)
	}
}
