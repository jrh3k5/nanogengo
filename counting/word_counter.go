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
	punctuationRegex, err := regexp.Compile("[^a-zA-Z0-9]+$")
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

		effectiveToken := punctuationRegex.ReplaceAllString(trimmedToken, "")
		currentWord := words.AddWordOccurrence(effectiveToken)

		if previousWord == nil {
			currentWord.IncrementSentenceStart(1)
		}

		// If there's any non-alphanumeric characters at the end, then consider this the end of the
		// sentence, so 'reset' the tracking so that the end of this sentence isn't mistakenly
		// tracked as preceding the beginning of the next sentence
		if !punctuationRegex.MatchString(trimmedToken) {
			if previousWord != nil {
				previousWord.AddSuccessor(currentWord)
			}
			previousWord = currentWord
		} else {
			punctuation := punctuationRegex.FindStringSubmatch(trimmedToken)[0]
			err = currentWord.AddPunctuation(punctuation)
			if err != nil {
				return fmt.Errorf("unable to add punctuation '%v' for word '%v': %v", punctuation, trimmedToken, err)
			}
			previousWord = nil
		}
	}
	return nil
}
