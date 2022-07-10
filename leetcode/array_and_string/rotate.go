package array_and_string

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
