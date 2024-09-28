package main

import (
	"fmt"
	"io"
	"strings"
)

func countFileLines(reader io.Reader) (string, error) {
	content, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}

	str := string(content)
	split := strings.Split(str, "\n")

	return fmt.Sprintf("%d", len(split)-1), nil
}
