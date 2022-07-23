package the_primary_algorithms

// 以下代码是关于 the_primary_algorithms 的 array 的实现部分

// 1.removeDuplicates 删除排序数组中重复项 已经在 array_and_string 的 summary 中实现


// 2.MaxProfit
// prices 数组一定有两个价格，否则无法产生利益，不需要提前判断长度
func MaxProfit(prices []int) int {
	var max int
	for i := 1; i < len(prices); i++ {
		if prices[i-1] < prices[i] {
			max += prices[i] - prices[i-1]
		}
	}

	return max
}

// 3.旋转数组
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

func ReverseArray(nums []int, l, r int) {
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

	ReverseArray(nums, 0, len(nums)-1)
	ReverseArray(nums, 0, k-1)
	ReverseArray(nums, k, len(nums)-1)

}

func ContainsDuplicate1(nums []int) bool {

	size := len(nums)
	table := make([]int, size)

	for _, v := range nums {
		mod := v % size

		if mod < 0 {
			mod += size
		}

		if v == 0 {
			// 针对值为0，做特殊情况处理
			v = 10e9 + 1
		}

		for mod < size && table[mod] != 0 {
			if table[mod] == v {
				return true
			}

			mod = (mod + 1) % size
		}

		table[mod] = v
	}

	return false
}


// 4.ContainsDuplicate2 存在重复元素
// 使用map
func ContainsDuplicate2(nums []int) bool {
	uniqueMap := make(map[int]struct{}, len(nums))
	
	for _, v := range nums {
		if _, ok := uniqueMap[v]; ok {
			return true
		}

		uniqueMap[v] = struct{}{}
	}

	return false
}

// 5.SingleNumber 寻找数组中唯一出现的数
// 位运算：异或运算 a^b^a = b , 0^a = a, a^a = 0, 满足交换律
func SingleNumber(nums []int) int {
	var mask int
	for _, v := range nums {
		mask ^= v
	}

	return mask
}

// 6.Intersect 做两个数组的唯一集
// 利用map做长度短的数组的唯一集，利用快慢指针对长度长的数组做并集判断以及存放数
func Intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	uniqueSet := make(map[int]int, len(nums1))

	// 对出现的数字做统计
	for _, v := range nums1 {
		uniqueSet[v] += 1
	}

	// 节约内存，对nums2快慢指针来保存并集结果
	var l, r int
	for r < len(nums2) {
		if count, ok := uniqueSet[nums2[r]]; ok && count > 0 {
			uniqueSet[nums2[r]] -= 1
			nums2[l] = nums2[r]
			l++
		}
		r++
	}

	return nums2[:l]
}

// 7.PlusOne 对数组组成表示的数字，如何处理进位问题
func PlusOne(digits []int) []int {

	for i := len(digits) - 1; 0 <= i; i-- {
		if (digits[i]+1)/10 == 1 {
			//能进位
			digits[i] = 0

			// 进位后缺少空间，申请新的数组
			if i-1 < 0 {
				newDigits := make([]int, len(digits)+1)
				newDigits[0] = 1
				return newDigits
			}
		} else {
			digits[i] += 1
			break
		}
	}

	return digits
}

// 8.moveZeroes 已经在 array_and_string 的 summary 中实现


// 9.TwoSum 解决无序数组的两数之和的寻找
func TwoSum(nums []int, target int) []int {
	if len(nums) < 3 {
		return []int{0, 1}
	}

	var indexes []int
	record := make(map[int]int)

	for index2, value := range nums {

		// 利用map存放差值-索引, 差值为target-value, 寻找迭代过程中可能的目标值
		// case1: 没找到，保存记录
		// case2: 找到，包装成数组返回
		if index1, ok := record[value]; !ok {
			// 存储后续迭代寻找的另一个值
			record[target-value] = index2
		} else {
			indexes = []int{index1, index2}
			break
		}
	}

	return indexes
}

// 10.IsValidSudoku 检查目前已填入的9x9数独是否为一个合理的数独
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


// 11.Rotate3 旋转图像 矩阵顺时针旋转90°
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
