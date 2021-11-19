package counting

import (
	"fmt"
	"nanogengo/data"
	"nanogengo/genio"
	"regexp"
	"strings"
)

type WordCounter interface {
	CountWords() (data.Words, error)
}

type LinesProviderWordCounter struct {
	LinesProvider genio.LinesProvider
}

func (wordCounter LinesProviderWordCounter) CountWords() (data.Words, error) {
	words := data.NewWords()
	lines, err := wordCounter.LinesProvider.ProvideLines()
	if err != nil {
		return *new(data.Words), err
	}

	for _, line := range lines {
		err := CountWords(line, words)
		if err != nil {
			return *new(data.Words), err
		}
	}

	return *words, nil
}

func CountWords(line string, words *data.Words) error {
	punctuationRegex, err := regexp.Compile("[.;\\!\\?]{1}$")
	if err != nil {
		return err
	}

	// Do the same pattern as punctuation regex, but without the $ terminator so that non-tail
	// non-alpha characters are replaced
	trimAllRegex, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return err
	}

	var previousWord *data.Word
	var previousPunctuation *data.Punctuation
	tokens := strings.Split(line, " ")
	for _, token := range tokens {
		trimmedToken := strings.TrimSpace(token)
		if len(trimmedToken) == 0 {
			continue
		}

		effectiveToken := trimAllRegex.ReplaceAllString(trimmedToken, "")
		currentWord := words.AddWordOccurrence(effectiveToken)

		if previousPunctuation != nil {
			if previousPunctuation.Terminator {
				currentWord.IncrementSentenceStart(1)
			}
			data.AddSuccessor(previousPunctuation, currentWord)
		} else if previousWord == nil {
			currentWord.IncrementSentenceStart(1)
		} else if previousWord != nil {
			data.AddSuccessor(previousWord, currentWord)
		}

		// If the current token doesn't terminate in a punctuation, then set it up so
		// that the current word is chained to the next token
		// Otherwise, set it up so that the next token will be chained to the punctuation following
		// the current word
		if !punctuationRegex.MatchString(trimmedToken) {
			previousPunctuation = nil
			previousWord = currentWord
		} else {
			punctuation := punctuationRegex.FindStringSubmatch(trimmedToken)[0]
			addedPunctuation, err := currentWord.AddPunctuation(punctuation)
			if err != nil {
				return fmt.Errorf("unable to add punctuation '%v' for word '%v': %v", punctuation, trimmedToken, err)
			}
			previousPunctuation = addedPunctuation
			previousWord = nil
		}
	}
	return nil
}
