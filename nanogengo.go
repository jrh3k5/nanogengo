package main

import (
	"fmt"
	"log"
	"nanogengo/counting"
	"nanogengo/genio"
	"nanogengo/probability"
	"os"
)

func main() {
	textFilesDir := os.Args[1]
	fmt.Printf("Read files from: %v\n", textFilesDir)

	linesProvider := genio.DirectoryLinesProvider{DirLocation: textFilesDir}
	wordCounter := counting.LinesProviderWordCounter{LinesProvider: linesProvider}
	words, err := wordCounter.CountWords()
	if err != nil {
		log.Fatalf("Unexpected error counting words: %v\n", err)
	}
	fmt.Printf("Counted %v words\n", len(words.Words))

	wordsProbability := probability.WordsBackedWordsProbability{Words: words}
	firstWord, err := wordsProbability.GetSentenceStart()
	if err != nil {
		log.Fatalf("Failed to get the first word: %v\n", err)
	}
	if firstWord == nil {
		log.Fatal("Unable to select a matching starting first word\n")
	}
	fmt.Printf("First word is: %v", firstWord.Word)
}
