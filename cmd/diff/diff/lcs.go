package diff

func LargestCommonSequence(X, Y string) string {
	m := len(X)
	n := len(Y)

	C := createMatrix(m+1, n+1)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if X[i-1] == Y[j-1] {
				diagonal := C[i-1][j-1] + 1
				C[i][j] = diagonal
			} else {
				xPrev := C[i-1][j]
				yPrev := C[i][j-1]
				C[i][j] = max(xPrev, yPrev)
			}
		}
	}

	return reconstructLCS(C, X, Y, m, n)
}

func reconstructLCS(C [][]int, X, Y string, i, j int) string {
	if i == 0 || j == 0 {
		return ""
	}
	if X[i-1] == Y[j-1] {
		return reconstructLCS(C, X, Y, i-1, j-1) + string(X[i-1])
	}
	if C[i-1][j] > C[i][j-1] {
		return reconstructLCS(C, X, Y, i-1, j)
	}
	return reconstructLCS(C, X, Y, i, j-1)
}
