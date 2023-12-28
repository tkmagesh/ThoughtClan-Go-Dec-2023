package main

import (
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Println("application exiting with error :", err)
		os.Exit(1)
	}
}
