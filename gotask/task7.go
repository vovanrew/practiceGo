package gotask

func DiagonalReverse(matrix [][]int) [][]int {
	var temp int
	for i := 0; i <= len(matrix) - 1; i++ {
		for j := i + 1; j < len(matrix[i]); j++ {
			temp = matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
		}
	}

	return matrix
}