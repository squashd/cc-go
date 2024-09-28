package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func countFileWordsStream(reader io.Reader) (string, error) {
	wordCount := 0
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		wordCount++
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", wordCount), nil
}

func countFileWords(reader io.Reader) (string, error) {
	bytes, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}

	str := string(bytes)
	lines := strings.Split(str, "\n")

	wordCount := 0
	for _, line := range lines {
		words := strings.Fields(line)
		wordCount += len(words)
	}

	return fmt.Sprintf("%d", wordCount), nil
}
