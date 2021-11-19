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

func (words *Words) GetSentenceStart() (*Word, error) {
	if words.TotalWordCount == 0 {
		return nil, nil
	}

	probability := rand.Float64()
	var word *Word
	for word == nil {
		startableWords := make([]*Word, 10, words.TotalWordCount)
		for _, candidateWord := range words.Words {
			if candidateWord.HasStartProbability(probability) {
				startableWords = append(startableWords, candidateWord)
			}
		}

		if len(startableWords) == 0 {
			return nil, nil
		}

		word = startableWords[rand.Intn(len(startableWords))]
	}

	return word, nil
}
