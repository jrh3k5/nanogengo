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
	assert.Equal(t, fishWord.TotalSuccessorCount, 2)
	// Verify fish successors
	assert.Assert(t, is.Contains(fishWord.Successors, "two"))
	assert.Equal(t, fishWord.Successors["two"].Occurrences, 1)
	assert.Assert(t, is.Contains(fishWord.Successors, "blue"))
	assert.Equal(t, fishWord.Successors["blue"].Occurrences, 1)
	assert.Equal(t, len(fishWord.Successors), 2, "Only 'blue' and 'two' should follow 'fish'")

	assert.Assert(t, is.Contains(words.Words, "one"))
	oneWord := words.Words["one"]
	assert.Equal(t, oneWord.Occurrences, 1)
	assert.Equal(t, oneWord.TotalSuccessorCount, 1)
	assert.Assert(t, is.Contains(oneWord.Successors, "fish"))
	assert.Equal(t, oneWord.Successors["fish"].Occurrences, 1)
	assert.Equal(t, len(oneWord.Successors), 1)

	assert.Assert(t, is.Contains(words.Words, "two"))
	twoWord := words.Words["two"]
	assert.Equal(t, twoWord.Occurrences, 1)
	assert.Equal(t, twoWord.TotalSuccessorCount, 1)
	assert.Assert(t, is.Contains(twoWord.Successors, "fish"))
	assert.Equal(t, twoWord.Successors["fish"].Occurrences, 1)
	assert.Equal(t, len(twoWord.Successors), 1)

	assert.Assert(t, is.Contains(words.Words, "red"))
	redWord := words.Words["red"]
	assert.Equal(t, redWord.Occurrences, 1)
	assert.Equal(t, redWord.TotalSuccessorCount, 1)
	assert.Assert(t, is.Contains(redWord.Successors, "fish"))
	assert.Equal(t, redWord.Successors["fish"].Occurrences, 1)
	assert.Equal(t, len(redWord.Successors), 1)

	assert.Assert(t, is.Contains(words.Words, "blue"))
	blueWord := words.Words["blue"]
	assert.Equal(t, blueWord.Occurrences, 1)
	assert.Equal(t, blueWord.TotalSuccessorCount, 1)
	assert.Assert(t, is.Contains(blueWord.Successors, "fish"))
	assert.Equal(t, blueWord.Successors["fish"].Occurrences, 1)
	assert.Equal(t, len(blueWord.Successors), 1)
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
	assert.Assert(t, is.Contains(foxWord.Successors, "hen"))
	assert.Equal(t, foxWord.Successors["hen"].Occurrences, 1)
	assert.Equal(t, len(foxWord.Successors), 1, "The punctuation should mean only 'Hen' follows fox, and it's once")

	assert.Assert(t, is.Contains(words.Words, "hen"))
	henWord := words.Words["hen"]
	assert.Equal(t, henWord.Occurrences, 1)
	assert.Equal(t, henWord.TotalSuccessorCount, 0)
	assert.Equal(t, len(henWord.Successors), 0)
}
