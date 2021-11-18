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
		wordCounts, err := CountWords(line)
		if err != nil {
			return *new(data.Words), err
		}
		words.MergeWords(wordCounts)
	}

	return *words, nil
}

func CountWords(line string) (map[string]*data.Word, error) {
	alphaOnlyRegex, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return *new(map[string]*data.Word), err
	}

	wordMap := make(map[string]*data.Word)
	tokens := strings.Split(line, " ")
	for _, token := range tokens {
		trimmedToken := strings.TrimSpace(token)
		if len(trimmedToken) == 0 {
			continue
		}

		effectiveToken := alphaOnlyRegex.ReplaceAllString(trimmedToken, "")
		tokenKey := strings.ToLower(effectiveToken)
		if existingWord, doesContain := wordMap[tokenKey]; doesContain {
			existingWord.Increment(1)
		} else {
			newWord := data.NewWord(effectiveToken)
			newWord.Occurrences = 1
			wordMap[tokenKey] = newWord
		}
	}
	return wordMap, nil
}
