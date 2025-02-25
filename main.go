package main

import (
	"boids/app"
	"log"
	"os"
)

func main() {
	if err := app.Run(); err != nil {
		log.Printf("exiting with 0: %v", err)
		os.Exit(1)
	}
}
