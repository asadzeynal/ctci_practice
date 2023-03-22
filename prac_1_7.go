package main

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
