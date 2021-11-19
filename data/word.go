package data

import "strings"

type Word struct {
	Word                  string
	Occurrences           int
	Successors            map[string]*WordSuccessor
	TotalSuccessorCount   int
	Punctuations          map[string]*Punctuation
	TotalPunctuationCount int
	SentenceStartCount    int
}

type WordSuccessor struct {
	Word        *Word
	Occurrences int
}

type Punctuation struct {
	Punctuation string
	Occurrences int
}

func NewWord(word string) *Word {
	newWord := new(Word)
	newWord.Word = word
	newWord.Successors = make(map[string]*WordSuccessor)
	newWord.Punctuations = make(map[string]*Punctuation)
	return newWord
}

func (word *Word) IncrementOccurences(delta int) {
	word.Occurrences = word.Occurrences + delta
}

func (word *Word) IncrementSentenceStart(delta int) {
	word.SentenceStartCount = word.SentenceStartCount + delta
}

func (wordSuccessor *WordSuccessor) IncrementOccurrences(delta int) {
	wordSuccessor.Occurrences = wordSuccessor.Occurrences + delta
}

func (punctuation *Punctuation) IncrementOccurrences(delta int) {
	punctuation.Occurrences = punctuation.Occurrences + delta
}

func (word *Word) AddSuccessor(successor *Word) {
	if existingSuccessor, doesContain := word.Successors[successor.GetKey()]; doesContain {
		existingSuccessor.IncrementOccurrences(1)
	} else {
		newSuccessor := new(WordSuccessor)
		newSuccessor.Word = successor
		newSuccessor.Occurrences = 1
		word.Successors[successor.GetKey()] = newSuccessor
	}
	word.TotalSuccessorCount++
}

func (word *Word) AddPunctuation(punctuation string) {
	if existingPunctuation, doesContain := word.Punctuations[punctuation]; doesContain {
		existingPunctuation.IncrementOccurrences(1)
	} else {
		newPunctuation := new(Punctuation)
		newPunctuation.Occurrences = 1
		word.Punctuations[punctuation] = newPunctuation
	}
	word.TotalPunctuationCount++
}

func (word *Word) GetKey() string {
	return strings.ToLower(word.Word)
}

// HasStartProbability measures whether or not the probability of this word
// is as equal or more likely to occur than the given probability
func (word *Word) HasStartProbability(probability float64) bool {
	startProbability := float64(word.SentenceStartCount) / float64(word.Occurrences)
	return startProbability-probability >= 0
}

// // Gets, if probability indicates, puncutation to follow this word. Returns a nil pointer if there is no applicable punctuation at this time
// func (word *Word) GetPunctuation() (Punctuation, error) {
// 	if word.TotalPunctuationCount == 0 {
// 		return *new(Punctuation), nil
// 	}

// }

// // Gets the next word for the given word. Returns nil if there is no word to follow.
// func (word *Word) GetNextWord() (Word, error) {

// }
