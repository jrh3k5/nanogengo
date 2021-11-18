package data

import "strings"

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
		existingWord.Increment(1)
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
