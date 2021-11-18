package data

type Word struct {
	Word        string
	Occurrences int
}

func NewWord(word string) *Word {
	newWord := new(Word)
	newWord.Word = word
	return newWord
}

func (word *Word) Increment(delta int) {
	word.Occurrences = word.Occurrences + delta
}
