package diff

import (
	"fmt"
	"strings"
)

func Diff(X []string, Y []string) {
	lcs := LCSLines(X, Y)
	var i, j int
	var sb strings.Builder

	for _, line := range lcs {
		// Deletions in X
		for i < len(X) && X[i] != line {
			sb.WriteString(fmt.Sprintf("\033[91m- %s\033[0m\n", X[i]))
			i++
		}
		// Insertions in Y
		for j < len(Y) && Y[j] != line {
			sb.WriteString(fmt.Sprintf("\033[92m+ %s\033[0m\n", Y[j]))
			j++
		}
		// Matching lines
		if i < len(X) && j < len(Y) && X[i] == Y[j] {
			sb.WriteString(fmt.Sprintf("  %s\n", line))
			i++
			j++
		}
	}

	// Remaining deletions from X
	for i < len(X) {
		sb.WriteString(fmt.Sprintf("\033[91m- %s\033[0m\n", X[i]))
		i++
	}

	// Remaining insertions from Y
	for j < len(Y) {
		sb.WriteString(fmt.Sprintf("\033[92m+ %s\033[0m\n", Y[j]))
		j++
	}

	fmt.Print(sb.String())
}
