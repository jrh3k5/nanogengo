package data

type Words struct {
	Words          map[string]*Word
	TotalWordCount int
}

func NewWords() *Words {
	words := new(Words)
	words.Words = make(map[string]*Word)
	return words
}

func (words *Words) MergeWords(toMerge map[string]*Word) *Words {
	for key, value := range toMerge {
		if existingWord, doesContain := words.Words[key]; doesContain {
			existingWord.Increment(value.Occurrences)
		} else {
			words.Words[key] = value
		}
		words.TotalWordCount += value.Occurrences
	}
	return words
}
