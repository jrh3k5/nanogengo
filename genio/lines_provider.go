package genio

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type LinesProvider interface {
	// ProvideLines provides the lines to be used for NaNoGenGo
	// Returns a list of the strings
	// If there is an error, a nil list address and the error object is returned
	ProvideLines() ([]string, error)
}

// DirectoryLinesProvider provides lines from a given directory location
type DirectoryLinesProvider struct {
	DirLocation string
}

// ReaderLinesProvider provides lines from a given reader
type ReaderLinesProvider struct {
	Reader io.Reader
}

type ArrayLinesProvider struct {
	Lines []string
}

func (directory DirectoryLinesProvider) ProvideLines() ([]string, error) {
	files, err := ioutil.ReadDir(directory.DirLocation)

	if err != nil {
		return *new([]string), err
	}

	lines := make([]string, 0)
	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(strings.ToLower(file.Name()), ".txt") {
			continue
		}

		fullFileName := filepath.Join(directory.DirLocation, file.Name())
		fileRef, err := os.Open(fullFileName)
		if err != nil {
			return *new([]string), err
		}
		defer fileRef.Close()

		readerContainer := ReaderLinesProvider{Reader: fileRef}
		fileLines, err := readerContainer.ProvideLines()
		if err != nil {
			return *new([]string), err
		}
		lines = append(lines, fileLines...)
	}

	return lines, nil
}

func (reader ReaderLinesProvider) ProvideLines() ([]string, error) {
	lines := make([]string, 0)

	scanner := bufio.NewScanner(reader.Reader)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return *new([]string), err
	}

	return lines, nil
}

func (arrayProvider ArrayLinesProvider) ProvideLines() ([]string, error) {
	return arrayProvider.Lines, nil
}
