package main

import (
	"fmt"
	"log"
	"nanogengo/genio"
	"os"
)

func main() {
	textFilesDir := os.Args[1]
	fmt.Printf("Read files from: %v\n", textFilesDir)

	linesProvider := genio.DirectoryLinesProvider{DirLocation: textFilesDir}
	lines, err := linesProvider.ProvideLines()
	if err != nil {
		log.Fatal("Failed to read files: ", err)
	}
	fmt.Printf("Read in %v lines", len(lines))
}
