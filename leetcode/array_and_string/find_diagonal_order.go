package array_and_string

func FindDiagonalOrder(mat [][]int) []int {
	row, col := len(mat), len(mat[0])
	result := make([]int, row*col)
	
	var i, j, c int
	for c < row*col {

		if (i+j)%2 == 0 {
			// upward diagonal
			for 0 <= i && j < col {

				result[c] = mat[i][j]
				i--
				j++
				c++
			}

			if j < col {
				i += 1
			} else {
				i += 2
				j -= 1
			}

		} else {
			// downward diagonal
			for 0 <= j && i < row {
				result[c] = mat[i][j]
				i++
				j--
				c++
			}

			if i < row {
				j += 1
			} else {
				i -= 1
				j += 2
			}

		}
	}

	return result
}
