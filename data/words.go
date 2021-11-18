package data

import (
	"math/rand"
	"strings"
)

type Words struct {
	Words          map[string]*Word
	TotalWordCount int
}

func NewWords() *Words {
	words := new(Words)
	words.Words = make(map[string]*Word)
	return words
}

func (words *Words) AddWordOccurrence(occurrence string) *Word {
	key := strings.ToLower(occurrence)
	if existingWord, doesContain := words.Words[key]; doesContain {
		existingWord.IncrementOccurences(1)
		words.TotalWordCount++
		return existingWord
	} else {
		newWord := NewWord(occurrence)
		newWord.Occurrences = 1
		words.Words[key] = newWord
		words.TotalWordCount++
		return newWord
	}
}

func (words *Words) FindStartingWord(sentenceStartProbability float64) (*Word, error) {
	startableWords := make([]*Word, 10, words.TotalWordCount)
	for _, word := range words.Words {
		if word.HasStartProbability(sentenceStartProbability) {
			startableWords = append(startableWords, word)
		}
	}

	if len(startableWords) == 0 {
		return nil, nil
	}

	return startableWords[rand.Intn(len(startableWords))], nil
}
