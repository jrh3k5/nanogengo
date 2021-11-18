package counting

import (
	"nanogengo/genio"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCounting(t *testing.T) {
	lines := []string{"one fish two fish", "red fish blue fish"}
	linesProvider := genio.ArrayLinesProvider{Lines: lines}
	wordCounter := LinesProviderWordCounter{LinesProvider: linesProvider}
	words, err := wordCounter.CountWords()
	assert.NilError(t, err, "Error occurred while counting words")

	assert.Equal(t, words.TotalWordCount, 8)

	assert.Assert(t, is.Contains(words.Words, "fish"))
	fishWord := words.Words["fish"]
	assert.Equal(t, fishWord.Occurrences, 4)

	assert.Assert(t, is.Contains(words.Words, "one"))
	oneWord := words.Words["one"]
	assert.Equal(t, oneWord.Occurrences, 1)

	assert.Assert(t, is.Contains(words.Words, "two"))
	twoWord := words.Words["two"]
	assert.Equal(t, twoWord.Occurrences, 1)

	assert.Assert(t, is.Contains(words.Words, "red"))
	redWord := words.Words["red"]
	assert.Equal(t, redWord.Occurrences, 1)

	assert.Assert(t, is.Contains(words.Words, "blue"))
	blueWord := words.Words["blue"]
	assert.Equal(t, blueWord.Occurrences, 1)
}

// TestCountsStripPunctuation verifies that punctuation is not counted when counting words
func TestCountsStripPunctuation(t *testing.T) {
	lines := []string{"Fox. Fox? Fox  Hen"}
	linesProvider := genio.ArrayLinesProvider{Lines: lines}
	wordCounter := LinesProviderWordCounter{LinesProvider: linesProvider}
	words, err := wordCounter.CountWords()
	assert.NilError(t, err, "Error occurred while counting words")

	assert.Equal(t, words.TotalWordCount, 4)

	assert.Assert(t, is.Contains(words.Words, "fox"))
	foxWord := words.Words["fox"]
	assert.Equal(t, foxWord.Occurrences, 3)

	assert.Assert(t, is.Contains(words.Words, "hen"))
	henWord := words.Words["hen"]
	assert.Equal(t, henWord.Occurrences, 1)
}
