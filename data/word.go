package data

import "strings"

type Word struct {
	Word                string
	Occurrences         int
	Successors          map[string]*WordSuccessor
	TotalSuccessorCount int
}

type WordSuccessor struct {
	Word        *Word
	Occurrences int
}

func NewWord(word string) *Word {
	newWord := new(Word)
	newWord.Word = word
	newWord.Successors = make(map[string]*WordSuccessor)
	return newWord
}

func (word *Word) Increment(delta int) {
	word.Occurrences = word.Occurrences + delta
}

func (wordSuccessor *WordSuccessor) Increment(delta int) {
	wordSuccessor.Occurrences = wordSuccessor.Occurrences + delta
}

func (word *Word) AddSuccessor(successor *Word) {
	if existingSuccessor, doesContain := word.Successors[successor.GetKey()]; doesContain {
		existingSuccessor.Increment(1)
	} else {
		newSuccessor := new(WordSuccessor)
		newSuccessor.Word = successor
		newSuccessor.Occurrences = 1
		word.Successors[successor.GetKey()] = newSuccessor
	}
	word.TotalSuccessorCount++
}

func (word *Word) GetKey() string {
	return strings.ToLower(word.Word)
}
