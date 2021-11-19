package main

import (
	"fmt"
	"log"
	"math/rand"
	"nanogengo/counting"
	"nanogengo/data"
	"nanogengo/genio"
	"os"
	"time"
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

	firstWord, err := words.GetSentenceStart()
	if err != nil {
		log.Fatalf("Failed to get the first word: %v\n", err)
	}
	if firstWord == nil {
		log.Fatal("Unable to select a matching starting first word\n")
	}
	fmt.Printf("First word is: %v\n", firstWord.Word)

	punctuation, err := firstWord.GetPunctuation()
	if err != nil {
		log.Fatalf("Failed to get a punctuation: %v\n", err)
	}

	if punctuation == nil {
		fmt.Printf("No punctuation available.\n")
	} else {
		fmt.Printf("Word's next punctuation is: '%v'\n", punctuation.Punctuation)
	}

	nextWord, err := data.GetNextWord(firstWord)
	if err != nil {
		log.Fatalf("Failed to get the next word: %v\n", err)
	}

	if nextWord == nil {
		fmt.Printf("No next word available.\n")
	} else {
		fmt.Printf("Word's next word is: '%v'\n", nextWord.Word)
	}

	// TODO: use probability of words following puncutation
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}
