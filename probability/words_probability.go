package probability

import (
	"math/rand"
	"nanogengo/data"
	"time"
)

type WordsProbability interface {
	GetSentenceStart() (data.Word, error)

	GetPunctuation(word data.Word) (data.Punctuation, error)

	GetNextWord(word data.Word) (data.Word, error)
}

type WordsBackedWordsProbability struct {
	Words data.Words
}

func (prob WordsBackedWordsProbability) GetSentenceStart() (*data.Word, error) {
	if prob.Words.TotalWordCount == 0 {
		return nil, nil
	}

	probability := rand.Float64()
	var word *data.Word
	var err error
	for word == nil {
		word, err = prob.Words.FindStartingWord(probability)
		if err != nil {
			return nil, err
		}
	}

	return word, nil
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}
