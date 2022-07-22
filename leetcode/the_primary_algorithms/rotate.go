package the_primary_algorithms

// 约瑟夫环交换
// time:  O(n)
// space: O(1)
func Rotate1(nums []int, k int) {
	k %= len(nums)
	if k == 0 {
		// 不用翻转
		return
	}
	var move, startIndex, next int

	hold := nums[0]

	for {
		// find next jump index in cycle
		next = (next + k) % len(nums)

		// set and store
		hold, nums[next] = nums[next], hold
		move++

		// 当前是否遇到闭环
		if next == startIndex {
			// 移动步数够了，退出循环
			if move >= len(nums) {
				break
			}

			next++
			startIndex++
			hold = nums[next]
		}
	}
}

func Reverse(nums []int, l, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

func Rotate2(nums []int, k int) {
	k %= len(nums)
	if k == 0 {
		return
	}

	Reverse(nums, 0, len(nums)-1)
	Reverse(nums, 0, k-1)
	Reverse(nums, k, len(nums)-1)

}

// 矩阵顺时针旋转90°
func Rotate3(matrix [][]int) {

	// 反对角线翻转
	n := len(matrix)
	ui, bi, uj, bj := 0, n-1, 0, n-1
	for ui < bi {
		for uj < bj {
			// 保持ui，bj 不变
			matrix[ui][uj], matrix[bi][bj] = matrix[bi][bj], matrix[ui][uj]

			uj++
			bi--
		}

		ui, uj = ui+1, 0
		bi, bj = n-1, bj-1
	}

	for ui, bi = 0, n-1; ui < bi; ui, bi = ui+1, bi-1 {
		matrix[ui], matrix[bi] = matrix[bi], matrix[ui]
	}

}
