package main

import (
	"fmt"
	"io"
)

func countFileBytes(reader io.Reader) (string, error) {
	bytes, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", len(bytes)), nil
}
