package genio

import (
	"bufio"
	"container/list"
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
	ProvideLines() (list.List, error)
}

// DirectoryLinesProvider provides lines from a given directory location
type DirectoryLinesProvider struct {
	DirLocation string
}

// ReaderLinesProvider provides lines from a given reader
type ReaderLinesProvider struct {
	Reader io.Reader
}

func (directory DirectoryLinesProvider) ProvideLines() (list.List, error) {
	files, err := ioutil.ReadDir(directory.DirLocation)

	if err != nil {
		return *new(list.List), err
	}

	lines := list.New()
	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(strings.ToLower(file.Name()), ".txt") {
			continue
		}

		fullFileName := filepath.Join(directory.DirLocation, file.Name())
		fileRef, err := os.Open(fullFileName)
		if err != nil {
			return *new(list.List), err
		}
		defer fileRef.Close()

		readerContainer := ReaderLinesProvider{Reader: fileRef}
		fileLines, err := readerContainer.ProvideLines()
		if err != nil {
			return *new(list.List), err
		}
		lines.PushBackList(&fileLines)
	}

	return *lines, nil
}

func (reader ReaderLinesProvider) ProvideLines() (list.List, error) {
	lines := list.New()

	scanner := bufio.NewScanner(reader.Reader)
	for scanner.Scan() {
		lines.PushBack(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return *new(list.List), err
	}

	return *lines, nil
}
