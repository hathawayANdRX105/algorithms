package array_and_string

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
