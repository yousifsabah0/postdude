package main

import (
	"log"
	"os"

	"github.com/yousifsabah0/postdude"
)

func main() {
	if err := postdude.New().Execute(); err != nil {
		log.Fatal("Unable to start the program. process exited.")
		os.Exit(1)
	}
}
