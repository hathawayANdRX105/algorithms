package the_primary_algorithms

func IsValidSudoku(board [][]byte) bool {
	// 整数32位利用27位来分别保存当前数出现 行，列，数字方格位置

	occurence := make([]int, 9)

	for i, arr := range board {
		for j, value := range arr {
			if value == '.' {
				continue
			}

			rowShift := 1 << i
			colShift := 1 << (j + 9)
			gridShift := 1 << (i/3*3 + j/3 + 18)

			data := occurence[value-'1']
			if (data&rowShift > 0) || (data&colShift > 0) || (data&gridShift > 0) {
				// 出现相同行，列，方格任意情况两次，则为无效数独
				return false
			}

			// 更新
			occurence[value-'1'] += rowShift + colShift + gridShift
		}
	}

	return true
}
