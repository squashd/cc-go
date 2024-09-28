package diff

func createEditTable(X, Y []string) [][]int {
	m := len(X) + 1
	n := len(Y) + 1
	C := createMatrix(m, n)
	for i := range X {
		nI := i + 1
		for j := range Y {
			nJ := j + 1
			if X[i] == Y[j] {
				diagonal := C[nI-1][nJ-1]
				C[nI][nJ] = diagonal + 1
			} else {
				xPrev := C[nI-1][nJ]
				yPrev := C[nI][nJ-1]
				C[nI][nJ] = max(xPrev, yPrev)
			}
		}
	}
	return C
}

func createMatrix(m, n int) [][]int {
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
	}
	return matrix
}
