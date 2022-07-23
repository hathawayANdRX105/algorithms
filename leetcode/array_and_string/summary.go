package array_and_string

// 1.Generate 生成杨辉三角（全）
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


// 2.GetRow 生成杨辉三角(某行) 
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


// 3.ReverseWords2 反转字符串中的单词 III, 保持单词顺序，翻转每个单词的字符
func ReverseWords2(s string) string {

	var l, r int
	strByteArr := []byte(s)

	// 单词翻转, word[i, j] 闭区间
	reverseSingle := func(b []byte, i, j int) {
		for i < j {
			b[i], b[j] = b[j], b[i]
			i++
			j--
		}
	}

	for ; r < len(s); r++ {
		if strByteArr[r] == ' ' {
			reverseSingle(strByteArr, l, r-1)

			l = r + 1
		}
	}

	// 题目说明字符串开头结尾不含空格，直接对最后一个单词翻转
	reverseSingle(strByteArr, l, r-1)

	// 考虑最后可能有空格结尾，则说明最后一个单词还没翻转
	// if strByteArr[r-1] != ' ' {
	// 	reverseSingle(strByteArr, l, r-1)
	// }

	return string(strByteArr)
}

// 4.FindMin 寻找旋转数组的最小值
func FindMin(nums []int) int {

	l, r := 0, len(nums)-1

	for l < r {
		// 不建议(l+r) >> 1, l+r占用更多的内存
		middle := l + (r-l)>>1

		if nums[r] < nums[l] {
			// r < l , 存在翻转
			// 确定中位数在数组位置
			if nums[r] < nums[middle] {
				// l < r < m， 中位数在较大数子数组中
				l = middle + 1
			} else {
				// l < m < r, 中位数在较小子数组中
				r = middle
			}
		} else {
			// l < r
			return nums[l]
		}
	}

	return nums[l]
}


// 5.RemoveDuplicates 删除有序数组中的重复项
func RemoveDuplicates(nums []int) int {
	f, s := 1, 0
	for f < len(nums) {
		if nums[f] != nums[s] {
			s++
			nums[s] = nums[f]
		}

		f++
	}

	return s + 1
}


func RemoveDuplicatesBinarySearch(nums []int) int {
	if len(nums) < 2 {
		return 1
	}

	f, s := 1, 0
	for {
		if nums[f] != nums[s] {
			s++
			nums[s] = nums[f]
		}

		if f+1 == len(nums) {
			break
		}

		// 存在重复的数值，二分查询
		l, r := s, len(nums)-1

		for l < r {
			if nums[s] < nums[l+(r-l)>>1] {
				// 中位数大于重复的数值
				r = l + (r-l)>>1
			} else {
				// 中位数小于等于重复的数值
				l = l + 1 + (r-l)>>1
			}

		}

		f = l
	}

	return s + 1
}

// 6.MoveZeroes 移动零
func MoveZeroes(nums []int) {

	var s, f int
	for f < len(nums) {
		if nums[f] != 0 {
			nums[s], nums[f] = nums[f], nums[s]
			s++
		}

		f++
	}
}


