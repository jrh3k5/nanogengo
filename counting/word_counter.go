package counting

import (
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
	alphaOnlyRegex, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return err
	}

	var previousWord *data.Word
	tokens := strings.Split(line, " ")
	for _, token := range tokens {
		trimmedToken := strings.TrimSpace(token)
		if len(trimmedToken) == 0 {
			continue
		}

		effectiveToken := alphaOnlyRegex.ReplaceAllString(trimmedToken, "")
		currentWord := words.AddWordOccurrence(effectiveToken)
		if previousWord != nil {
			previousWord.AddSuccessor(currentWord)
		}
		previousWord = currentWord
	}
	return nil
}
