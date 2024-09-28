package main

import (
	"bytes"
	"fmt"
	"io"
	"sync"
)

func countFileAll(reader io.Reader) (string, error) {
	var wg sync.WaitGroup
	var chars, lines, words string
	var charErr, lineErr, wordErr error
	var errs []error

	content, _ := io.ReadAll(reader)
	wg.Add(3)
	go func() {
		defer wg.Done()
		chars, charErr = countFileChars(bytes.NewReader(content))
		if charErr != nil {
			errs = append(errs, charErr)
		}
	}()
	go func() {
		defer wg.Done()
		lines, lineErr = countFileLines(bytes.NewReader(content))
		if lineErr != nil {
			errs = append(errs, lineErr)
		}
	}()
	go func() {
		defer wg.Done()
		words, wordErr = countFileWords(bytes.NewReader(content))
		if wordErr != nil {
			errs = append(errs, wordErr)
		}
	}()
	wg.Wait()

	if len(errs) > 0 {
		return "", fmt.Errorf("Error running functions:", errs)
	}

	formatted := fmt.Sprintf("%s %s %s", lines, words, chars)
	return formatted, nil
}
