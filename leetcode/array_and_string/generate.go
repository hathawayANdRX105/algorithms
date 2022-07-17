package array_and_string

func Generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	} else if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}

	yangHuiTriangle := make([][]int, numRows)
	yangHuiTriangle[0] = []int{1}
	yangHuiTriangle[1] = []int{1, 1}

	for row := 2; row < numRows; row++ {
		rowArr := make([]int, row+1)

		// 第row层，除了前后两个数为1外，中间数是由row-1层的数两两相加获得,除去最后两个数的组合
		for i := 0; i < row-1; i++ {
			rowArr[i+1] = yangHuiTriangle[row-1][i] + yangHuiTriangle[row-1][i+1]
		}

		rowArr[0] = 1
		rowArr[row] = 1
		yangHuiTriangle[row] = rowArr
	}

	return yangHuiTriangle
}
