package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gabrielroacueto/locc/api"
)

func main() {
	if len(os.Args) != 3 || os.Args[1] != "understand" {
		fmt.Println("Usage: locc understand <directory>")
		os.Exit(1)
	}

	directory := os.Args[2]

	callback := func(output string) {
		fmt.Print(output)
	}

	err := api.StreamDirectoryAnalysis(directory, callback)

	if err != nil {
		log.Fatalf("Error analyzing directory: %v", err)
	}
}
