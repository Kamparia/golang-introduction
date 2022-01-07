package main

import (
	"errors"
	"fmt"
	"os"
)

func processFile(path string) error {
	if path == "" {
		return errors.New("empty path")
	}

	// open the file
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("unable to open file: %v", err)
	}

	// check if the file name
	if file.Name() != "info.log" {
		return fmt.Errorf("unexpected file name: %s", file.Name())
	}

	// close the file
	file.Close()

	return nil
}

func main() {
	err := processFile("./info.log")
	if err != nil {
		switch err.(type) {
		case *os.PathError:
			fmt.Printf("file not found: %+v\n", err)
		default:
			fmt.Printf("unknown error occurred: %+v\n", err)
		}
	}
}
