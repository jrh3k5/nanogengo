package data

import (
	"errors"
	"math/rand"
	"strings"
)

type Successable interface {
	Successors() map[string]*WordSuccessor
	TotalSuccessorCount() int
	TotalSuccessorOccurrences() int
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
		word.TotalSuccessorCount++
	}
	word.TotalSuccessorOccurrences++
}

func (punctuation *Punctuation) AddSuccessor(successor *Word) {
	if existingSuccessor, doesContain := punctuation.Successors[successor.GetKey()]; doesContain {
		existingSuccessor.IncrementOccurrences(1)
	} else {
		newSuccessor := new(WordSuccessor)
		newSuccessor.Word = successor
		newSuccessor.Occurrences = 1
		punctuation.Successors[successor.GetKey()] = newSuccessor
		punctuation.TotalSuccessorCount++
	}
	punctuation.TotalSuccessorOccurrences++
}

func (word *Word) AddPunctuation(punctuation string) error {
	if len(strings.TrimSpace(punctuation)) == 0 {
		return errors.New("invalid punctuation: " + punctuation)
	}

	if existingPunctuation, doesContain := word.Punctuations[punctuation]; doesContain {
		existingPunctuation.IncrementOccurrences(1)
	} else {
		newPunctuation := new(Punctuation)
		newPunctuation.Punctuation = punctuation
		newPunctuation.Occurrences = 1
		word.Punctuations[punctuation] = newPunctuation
		word.TotalPunctuationCount++
	}
	word.TotalPunctuationOccurrences++
	return nil
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

// Gets the next word for the given word. Returns nil if there is no word to follow.
func (word *Word) GetNextWord() (*Word, error) {
	if word.TotalSuccessorCount == 0 {
		return nil, nil
	}

	var matchingWord *WordSuccessor
	for matchingWord == nil {
		wordProbability := rand.Float64()
		var matchingWordProbability float64
		for _, candidate := range word.Successors {
			candidateProbability := float64(candidate.Occurrences) / float64(word.TotalPunctuationOccurrences)
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
