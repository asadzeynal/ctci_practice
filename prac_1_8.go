package main

func zeroMatrix(matrix [][]int) [][]int {
	n := len(matrix)
	zeroRows := make(map[int]struct{}, 0)
	zeroCols := make(map[int]struct{}, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				zeroRows[i] = struct{}{}
				zeroCols[j] = struct{}{}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if _, ok := zeroRows[i]; ok {
				matrix[i][j] = 0
				continue
			}
			if _, ok := zeroCols[j]; ok {
				matrix[i][j] = 0
			}
		}
	}

	return matrix
}
