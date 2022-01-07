package main

import (
	"log"
	"os"
)

func consoleLogger() {
	log.Print("Logging in Go!")
}

func fileLogger() {
	// open a log file, create if not exists
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Print("Logging to a file in Go!")
}

func main() {
	// Logging to the console
	consoleLogger()

	// Logging to a file
	fileLogger()
}
