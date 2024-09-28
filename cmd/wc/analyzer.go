package main

import (
	"fmt"
	"io"
)

type (
	AnalyserFunc func(io.Reader) (string, error)

	analyser struct {
		fileFuncs map[string]AnalyserFunc
	}
)

func newFileAnalyser() analyser {
	return analyser{
		fileFuncs: make(map[string]AnalyserFunc),
	}
}

func (a analyser) registerFileFunc(flag string, fn AnalyserFunc) {
	a.fileFuncs[flag] = fn
}

func (a analyser) run(flag string, reader io.Reader) (string, error) {
	fn, ok := a.fileFuncs[flag]
	if !ok {
		return "", fmt.Errorf("Invalid flag:", flag)
	}

	numString, err := fn(reader)
	if err != nil {
		return "", fmt.Errorf("Error running function:", err)
	}
	return numString, nil
}
