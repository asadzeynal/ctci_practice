package ch1

func rotateMatrix(matrix [][]int) [][]int {
	level := len(matrix) - 1
	numOfCycles := level
	for deepness := 0; deepness < level-1; deepness++ {
		i, j := deepness, deepness
		minCoord := deepness
		maxCoord := level - deepness
		for cycle := 0; cycle < numOfCycles; cycle++ {
			var el int
			el = matrix[i][j]
			for {
				if j >= i && j < maxCoord {
					j++
				} else if j == maxCoord && i < maxCoord {
					i++
				} else if i == maxCoord && j > minCoord {
					j--
				} else {
					i--
				}

				temp := matrix[i][j]
				matrix[i][j] = el
				el = temp

				if i == deepness && j == deepness {
					break
				}
			}
		}
		numOfCycles -= 2
	}

	return matrix
}

func rotateMatrixFaster(matrix [][]int) [][]int {
	n := len(matrix) - 1
	for depth := 0; depth < len(matrix)-2; depth++ {
		for offset := 0; offset < n-(depth*2); offset++ {
			el1 := matrix[depth][depth+offset]
			el2 := matrix[depth+offset][n-depth]
			el3 := matrix[n-depth][n-depth-offset]
			el4 := matrix[n-depth-offset][depth]

			matrix[depth][depth+offset] = el4
			matrix[depth+offset][n-depth] = el1
			matrix[n-depth][n-depth-offset] = el2
			matrix[n-depth-offset][depth] = el3
		}
	}

	return matrix
}
