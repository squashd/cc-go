package diff

func LCSLines(X, Y []string) []string {
	m := len(X) + 1
	n := len(Y) + 1
	C := createEditTable(X, Y)
	return reconstructLines(C, X, Y, m-1, n-1)
}

func reconstructLines(C [][]int, X, Y []string, i, j int) []string {
	if i == 0 || j == 0 {
		return []string{}
	}
	nI := i - 1
	nJ := j - 1
	if X[nI] == Y[nJ] {
		return append(reconstructLines(C, X, Y, i-1, j-1), X[nI])
	}
	if C[i][j-1] > C[i-1][j] {
		return reconstructLines(C, X, Y, i, j-1)
	}
	return reconstructLines(C, X, Y, i-1, j)
}
