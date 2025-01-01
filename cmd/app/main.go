package main

import (
	"fmt"
	"os"

	"github.com/idmaksim/img-augmentor/internal/augmentor"
)

func main() {
	aug := augmentor.New("data", 2)

	if err := aug.ProcessImages("archive.zip"); err != nil {
		fmt.Printf("Error processing images: %v\n", err)
		os.Exit(1)
	}
}
