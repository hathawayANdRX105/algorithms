package array_and_string

// 1.Rotate 旋转矩阵
func Rotate(matrix [][]int) {

	// rotate first dimension arrays
	size := len(matrix) - 1
	lo, hi := 0, size

	for lo < hi {
		matrix[lo], matrix[hi] = matrix[hi], matrix[lo]

		lo++
		hi--
	}

	// rotate diagonal elements
	var i, j int
	for i <= size {

		j = i + 1
		for j <= size {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]

			j++
		}

		i++
	}
}

// 2.SetZeroes 零矩阵
func SetZeroes(matrix [][]int) {

	row, col := len(matrix), len(matrix[0])
	// i ~ row, row + j ~ row + col
	record := make([]bool, row+col)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] != 0 {
				continue
			} else {
				record[i] = true
				record[row+j] = true
			}
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if record[i] || record[row+j] {
				matrix[i][j] = 0
			}
		}
	}

}


// 3.FindDiagonalOrder 矩阵对角线遍历
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

