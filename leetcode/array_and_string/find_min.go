package array_and_string

// FindMin ...
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
