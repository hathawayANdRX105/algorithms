package array_and_string

// GetRow ...
func GetRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	} else if rowIndex == 1 {
		return []int{1, 1}
	}

	rowArr := make([]int, rowIndex+1)
	rowArr[0] = 1
	rowArr[1] = 1

	for row := 2; row < rowIndex+1; row++ {

		pre := rowArr[0]
		for i := 1; i <= row; i++ {
			// 记录上一层后一个两两相加数
			post := rowArr[i]

			// 当前层的中间值为上一层的pre+post, 同时post也是rowArr[i]
			rowArr[i] += pre

			// 更新上一层的前一个两两相加数
			pre = post
		}

	}

	return rowArr
}
