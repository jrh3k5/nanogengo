# nanogengo

[![Build Status](https://app.travis-ci.com/jrh3k5/nanogengo.svg?branch=main)](https://app.travis-ci.com/jrh3k5/nanogengo)

A NaNoGenMo implementation in Go. This was a convenient reason to teach myself more about the Go programming language.

## Usage

With the assembled executable, execute:

```
./nanogengo <directory containing text files>
```

Each text file is assumed to contain strings that are complete sentences (and each line may contain multiple sentences). This means that, if your text files contain sentences that wrap lines, that this will not necessarily generate a coherent novel. For example, this would violate the assumptions of this program:

```
this is a line
that wraps and completes the sentence on the next line.
```

## Building

### Preqrequisites

This project uses Go 1.17 to build.

### Compiling

Clone the repository and then run:

```
go build
```