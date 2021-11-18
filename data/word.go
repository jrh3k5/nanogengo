package data

import "strings"

type Word struct {
	Word                string
	Occurrences         int
	Successors          map[string]*Word
	TotalSuccessorCount int
}

func NewWord(word string) *Word {
	newWord := new(Word)
	newWord.Word = word
	newWord.Successors = make(map[string]*Word)
	return newWord
}

func CopyWord(word *Word) *Word {
	newWord := NewWord(word.Word)
	newWord.Occurrences = word.Occurrences
	for key, value := range word.Successors {
		newWord.Successors[key] = CopyWord(value)
	}
	return newWord
}

func (word *Word) Increment(delta int) {
	word.Occurrences = word.Occurrences + delta
}

func (word *Word) AddSuccessor(successor *Word) {
	if existingSuccessor, doesContain := word.Successors[successor.GetKey()]; doesContain {
		existingSuccessor.Increment(1)
	} else {
		newSuccessor := CopyWord(successor)
		word.Successors[newSuccessor.GetKey()] = newSuccessor
	}
}

func (word *Word) GetKey() string {
	return strings.ToLower(word.Word)
}
