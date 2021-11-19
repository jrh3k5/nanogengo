package data

import (
	"errors"
	"math/rand"
	"strings"
)

type Successable interface {
	GetSuccessors() map[string]*WordSuccessor
	GetTotalSuccessorCount() int
	GetTotalSuccessorOccurrences() int
	IncrementSuccessorCount(delta int)
	IncrementSuccessorOccurrences(delta int)
}

type Word struct {
	Word                        string
	Occurrences                 int
	Successors                  map[string]*WordSuccessor
	TotalSuccessorCount         int
	TotalSuccessorOccurrences   int
	Punctuations                map[string]*Punctuation
	TotalPunctuationCount       int
	TotalPunctuationOccurrences int
	SentenceStartCount          int
}

func (word *Word) GetSuccessors() map[string]*WordSuccessor {
	return word.Successors
}

func (word *Word) GetTotalSuccessorCount() int {
	return word.TotalSuccessorCount
}

func (word *Word) GetTotalSuccessorOccurrences() int {
	return word.TotalPunctuationOccurrences
}

func (word *Word) IncrementSuccessorCount(delta int) {
	word.TotalSuccessorCount += delta
}

func (word *Word) IncrementSuccessorOccurrences(delta int) {
	word.TotalSuccessorOccurrences += delta
}

type WordSuccessor struct {
	Word        *Word
	Occurrences int
}

type Punctuation struct {
	Punctuation               string
	Occurrences               int
	Successors                map[string]*WordSuccessor
	TotalSuccessorCount       int
	TotalSuccessorOccurrences int
	Terminator                bool
}

func (punctuation *Punctuation) GetSuccessors() map[string]*WordSuccessor {
	return punctuation.Successors
}

func (punctuation *Punctuation) GetTotalSuccessorCount() int {
	return punctuation.TotalSuccessorCount
}

func (punctuation *Punctuation) GetTotalSuccessorOccurrences() int {
	return punctuation.TotalSuccessorOccurrences
}

func (puncutation *Punctuation) IncrementSuccessorCount(delta int) {
	puncutation.TotalSuccessorCount += delta
}

func (punctuation *Punctuation) IncrementSuccessorOccurrences(delta int) {
	punctuation.TotalSuccessorOccurrences += delta
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

func AddSuccessor(successable Successable, successor *Word) {
	if existingSuccessor, doesContain := successable.GetSuccessors()[successor.GetKey()]; doesContain {
		existingSuccessor.IncrementOccurrences(1)
	} else {
		newSuccessor := new(WordSuccessor)
		newSuccessor.Word = successor
		newSuccessor.Occurrences = 1
		successable.GetSuccessors()[successor.GetKey()] = newSuccessor
		successable.IncrementSuccessorCount(1)
	}
	successable.IncrementSuccessorOccurrences(1)
}

func (word *Word) AddPunctuation(punctuation string) (*Punctuation, error) {
	if len(strings.TrimSpace(punctuation)) == 0 {
		return nil, errors.New("invalid punctuation: " + punctuation)
	}

	if existingPunctuation, doesContain := word.Punctuations[punctuation]; doesContain {
		existingPunctuation.IncrementOccurrences(1)
		word.TotalPunctuationOccurrences++
		return existingPunctuation, nil
	} else {
		newPunctuation := new(Punctuation)
		newPunctuation.Punctuation = punctuation
		newPunctuation.Occurrences = 1
		newPunctuation.Terminator = punctuation == "?" || punctuation == "!" || punctuation == "."
		newPunctuation.Successors = make(map[string]*WordSuccessor)
		word.Punctuations[punctuation] = newPunctuation
		word.TotalPunctuationCount++
		word.TotalPunctuationOccurrences++
		return newPunctuation, nil
	}
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

// Gets, if probability indicates, puncutation to follow this word. Returns a nil pointer if there is no applicable punctuation at this time
func (word *Word) GetPunctuation() (*Punctuation, error) {
	if word.TotalPunctuationCount == 0 {
		return nil, nil
	}

	// Find the closest punctuation that does not fall below the probability
	var matchingPunctuation *Punctuation
	// Sometimes, the probability calculated can be too low, so keep looping until a punctuation is selected
	for matchingPunctuation == nil {
		puncProbability := rand.Float64()
		var matchingPunctuationProbability float64
		for _, candidate := range word.Punctuations {
			candidateProbability := float64(candidate.Occurrences) / float64(word.TotalPunctuationOccurrences)
			// If the current candidate is within the range of probability and is *more probable* than what was previously selected,
			// use that punctuation
			if candidateProbability >= puncProbability && candidateProbability >= matchingPunctuationProbability {
				matchingPunctuation = candidate
				matchingPunctuationProbability = candidateProbability
			}
		}
	}

	return matchingPunctuation, nil
}

func GetNextWord(successable Successable) (*Word, error) {
	if successable.GetTotalSuccessorCount() == 0 {
		return nil, nil
	}

	var matchingWord *WordSuccessor
	for matchingWord == nil {
		wordProbability := rand.Float64()
		var matchingWordProbability float64
		for _, candidate := range successable.GetSuccessors() {
			candidateProbability := float64(candidate.Occurrences) / float64(successable.GetTotalSuccessorOccurrences())
			// If the current candidate is within the range of probability and is *more probable* than what was previously selected,
			// use that word
			if candidateProbability >= wordProbability && candidateProbability >= matchingWordProbability {
				matchingWord = candidate
				matchingWordProbability = candidateProbability
			}
		}
	}

	return matchingWord.Word, nil
}
