package main

import (
	"fmt"
	"log"
	"nanogengo/counting"
	"nanogengo/genio"
	"os"
)

func main() {
	textFilesDir := os.Args[1]
	fmt.Printf("Read files from: %v\n", textFilesDir)

	linesProvider := genio.DirectoryLinesProvider{DirLocation: textFilesDir}
	wordCounter := counting.LinesProviderWordCounter{LinesProvider: linesProvider}
	words, err := wordCounter.CountWords()
	if err != nil {
		log.Fatalf("Unexpected error counting words: %v", err)
	}
	fmt.Printf("Counted %v words", len(words.Words))
}
