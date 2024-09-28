package main

import (
	"bufio"
	"fmt"
	"io"
)

func countFileCharsStream(reader io.Reader) (string, error) {
	charCount := 0
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		charCount++
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", charCount), nil
}

func countFileChars(reader io.Reader) (string, error) {
	bytes, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}

	str := string(bytes)

	charCount := 0
	for _ = range str {
		charCount++
	}

	return fmt.Sprintf("%d", charCount), nil
}
