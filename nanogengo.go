package main

import (
	"fmt"
	"log"
	"math/rand"
	"nanogengo/counting"
	"nanogengo/data"
	"nanogengo/genio"
	"os"
	"strings"
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

	currentWord, err := words.GetSentenceStart()
	if err != nil {
		log.Fatalf("Failed to get the first word: %v\n", err)
	}
	if currentWord == nil {
		log.Fatal("Unable to select a matching starting first word\n")
	}

	currentSentenceLength := 0
	var previousPunctuation *data.Punctuation
	for wordCount := 0; wordCount < 500; wordCount++ {
		currentSentenceLength++
		toPrint := strings.ToLower(currentWord.Word)
		// Only capitalize if this is the very first word, or if this is following punctuation and the previous punctuation was a terminator
		if currentSentenceLength == 1 && previousPunctuation == nil || (previousPunctuation != nil && previousPunctuation.Terminator) {
			toPrint = strings.Title(toPrint)
		}
		fmt.Print(toPrint)

		// Prevent the next token from accidentally thinking it follows punctuation
		previousPunctuation = nil

		var punctuation *data.Punctuation
		// Avoid awkwardness of too-short sentences by only getting a punctuation if there's enough
		// sentence to punctuate
		if currentSentenceLength > 2 {
			punctuation, err = currentWord.GetPunctuation()
			if err != nil {
				log.Fatalf("Unable to get punctuation for word '%v': %v", currentWord.Word, err)
			}
		}

		if punctuation != nil {
			fmt.Print(punctuation.Punctuation)
			previousPunctuation = punctuation
			currentWord, err = data.GetNextWord(punctuation)
			if err != nil {
				log.Fatalf("Unable to get next word after punctuation '%v': %v", punctuation.Punctuation, err)
			}

			// If there is no word found to follow the punctuation, then select a new first word
			if currentWord == nil {
				currentWord, err = words.GetSentenceStart()
				if err != nil {
					log.Fatalf("Failed to get a first word after punctuation '%v' yielded no successor: %v\n", punctuation.Punctuation, err)
				}
				if currentWord == nil {
					log.Fatalf("Unable to select a matching starting first word after punctuation '%v' yielded no successor\n", punctuation.Punctuation)
				}
			}

			currentSentenceLength = 0
		} else {
			nextWord, err := data.GetNextWord(currentWord)
			if err != nil {
				log.Fatalf("Unable to get next word after punctuation '%v': %v", currentWord.Word, err)
			}
			if nextWord == nil {
				log.Fatalf("Unexpected termination of no word and no punctuation available after word '%v'", currentWord.Word)
			}
			currentWord = nextWord
		}
		fmt.Print(" ")
	}

	// fmt.Printf("First word is: %v\n", firstWord.Word)

	// punctuation, err := firstWord.GetPunctuation()
	// if err != nil {
	// 	log.Fatalf("Failed to get a punctuation: %v\n", err)
	// }

	// if punctuation == nil {
	// 	fmt.Printf("No punctuation available.\n")
	// } else {
	// 	fmt.Printf("Word's next punctuation is: '%v'\n", punctuation.Punctuation)
	// }

	// nextWord, err := data.GetNextWord(firstWord)
	// if err != nil {
	// 	log.Fatalf("Failed to get the next word: %v\n", err)
	// }

	// if nextWord == nil {
	// 	fmt.Printf("No next word available.\n")
	// } else {
	// 	fmt.Printf("Word's next word is: '%v'\n", nextWord.Word)
	// }
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}
