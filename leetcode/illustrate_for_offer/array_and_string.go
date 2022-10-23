package illustrate_for_offer

import (
	"bytes"
	"math"
	"sort"
	"strconv"
	"strings"
)

// array_and_string.go 是 illustrate_for_offer 关于 数组与字符串 的代码实现部分

// 04.二维数组中的查找
// findNumberIn2DArray1 简单的二次二分法，先是列二分法对第一列帅选起始点，后对每行做二分查询
func findNumberIn2DArray1(matrix [][]int, target int) bool {
	if len(matrix) < 1 || len(matrix[0]) < 1 || target < matrix[0][0] {
		return false
	}

	l, r := 0, len(matrix)-1
	for l < r {
		mid := (l + r) >> 1
		if matrix[mid][0] == target {
			return true
		}

		if matrix[mid][0] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	if matrix[l][0] > target {
		l--
	}

	for l > -1 {
		cl, cr := 0, len(matrix[0])-1
		for cl <= cr {
			mid := (cl + cr) >> 1
			if matrix[l][mid] == target {
				return true
			}

			if matrix[l][mid] < target {
				cl = mid + 1
			} else {
				cr = mid - 1
			}
		}

		l--
	}

	return false
}

// findNumberIn2DArray2 bfsForOrder 从右上角走到左下角 O(m + n)
func findNumberIn2DArray2(matrix [][]int, target int) bool {
	if len(matrix) < 1 || len(matrix[0]) < 1 || matrix[0][0] > target {
		return false
	}

	n, m := len(matrix), len(matrix[0])
	r, c := 0, m-1
	for r < n && c > 0 {
		if matrix[r][c] == target {
			return true
		}

		if matrix[r][c] < target {
			r++
		} else {
			c--
		}
	}

	return false
}

// 05.替换空格
// replaceSpace
func replaceSpace(s string) string {
	if len(s) < 1 {
		return ""
	}

	ans := make([]byte, 0, len(s))

	for i := range s {
		ans = append(ans, s[i])

		if ans[len(ans)-1] == ' ' {
			ans[len(ans)-1] = '%'
			ans = append(ans, '2')
			ans = append(ans, '0')
		}
	}

	return string(ans)
}

// 11. 旋转数组的最小数字
func minArray(numbers []int) int {
	if len(numbers) < 2 || numbers[0] < numbers[len(numbers)-1] {
		return numbers[0]
	}

	// l 因为只指向大于 numbers[r]的数或最小值
	// r 为了保证 l 的正确性, 需要提前收敛
	l, r := 0, len(numbers)-1
	for l < r {
		m := (l + r) >> 1
		if numbers[m] < numbers[r] {
			r = m
		} else if numbers[m] > numbers[r] {
			// l -> 要么是大于numbers[r] 要么是最小值
			l = m + 1
		} else if numbers[m] == numbers[l] {
			// 前提 numbers[m] = number[r], 如果 numbers[m] = numbers[l] 则左右边界收缩
			// 避免 数组全是相同数 造成复杂度为 O(n)
			r--
			l++
		} else {
			r--
		}
	}

	return numbers[l]
}

// 17.打印从 1 到最大的 n 位数
func printNumbers(n int) []int {
	// 快速幂
	mul := 10
	max := 1
	for n > 0 {
		if n&1 == 1 {
			max *= mul
		}

		mul *= mul
		n >>= 1
	}

	// 1 ~ max - 1
	ans := make([]int, 0, max-1)
	for i := 1; i < max; i++ {
		ans = append(ans, i)
	}

	return ans
}

// 21.调整数组顺序使奇数位于偶数前面
func exchange(nums []int) []int {
	o, e := 0, len(nums)
	for o < e {
		if nums[o]&1 == 1 {
			o++
			continue
		}

		e--
		nums[o], nums[e] = nums[e], nums[o]
	}

	return nums
}

// 29.顺时针打印矩阵
func spiralOrder(matrix [][]int) []int {
	if len(matrix) < 1 || len(matrix[0]) < 1 {
		return nil
	}

	l, r, t, b := 0, len(matrix[0])-1, 0, len(matrix)-1
	ans := make([]int, 0, (r+1)*(b+1))
	// 四指针指定边界值，四种运动方式
	for {
		for i := l; i <= r; i++ {
			ans = append(ans, matrix[t][i])
		}
		t++
		if t > b {
			break
		}

		for i := t; i <= b; i++ {
			ans = append(ans, matrix[i][r])
		}
		r--
		if r < l {
			break
		}

		for i := r; i >= l; i-- {
			ans = append(ans, matrix[b][i])
		}
		b--
		if b < t {
			break
		}

		for i := b; i >= t; i-- {
			ans = append(ans, matrix[i][l])
		}
		l++
		if l > r {
			break
		}
	}

	return ans
}

// 39.数组中出现次数超过一半的数字
// majorityElement1 数组排序，中间值为超过一半数量的值
func majorityElement1(nums []int) int {
	sort.Ints(nums)
	m := (len(nums) - 1) / 2
	return nums[m]
}

// majorityElement2 摩尔投票 贪心
func majorityElement2(nums []int) int {
	cur, count := nums[0], 1

	for i := 1; i < len(nums); i++ {
		if cur == nums[i] {
			count++
			continue
		}

		count--
		if count == 0 {
			cur = nums[i]
			count++
		}
	}

	return cur
}

// 45. 把数组排成最小的数
// minNumber 直接拼接字符串性能也很快，先实现再说性能
func minNumber(nums []int) string {
	n := len(nums)
	strs := make([]string, n)
	for n > 0 {
		n--
		strs[n] = strconv.FormatInt(int64(nums[n]), 10) // less stack grow
	}

	// key: if a + b < b + a then a < b logically
	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] < strs[j]+strs[i]
	})

	return strings.Join(strs, "")
}

// 53 - I. 在排序数组中查找数字 I 并返回个数
func search(nums []int, target int) int {
	return getRightMargin(nums, target) - getRightMargin(nums, target-1)
}

func getRightMargin(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) >> 1

		if nums[m] <= target {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return l
}

// 53 - II. 0～n-1 中缺失的数字
// missingNumber1 数列和法
func missingNumber1(nums []int) int {
	var sum int

	for i := range nums {
		sum += nums[i]
	}
	n := len(nums)
	return (n+1)*n/2 - sum
}

// missingNumber2 异或运算
func missingNumber2(nums []int) int {
	// 0 1 2 3 正常情况
	// 0 1 3   缺失情况
	// index 直到 n - 2 的情况，因此从 n-1 开始，即len(nums)开始异或
	var ans int = len(nums)
	for i := range nums {
		ans ^= nums[i] ^ i
	}

	return ans
}

// missingNumber3 二分查询
func missingNumber3(nums []int) int {
	// 0 1 2 3 4 正常情况
	// 0 1 3 4 5 缺失情况
	// 如果 nums[mid] = index 则说明缺失之不在 mid之前
	l, r := 0, len(nums)-1
	for l < r {
		m := (l + r) >> 1

		if nums[m] == m {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return l
}

// 57. 和为 s 的两个数字, nums 是递增排序的数组
// twoSum1 为无序做法
func twoSum1(nums []int, target int) []int {
	m := make(map[int]struct{}, len(nums))

	for i := range nums {
		if _, ok := m[nums[i]]; ok {
			return []int{target - nums[i], nums[i]}
		}

		m[target-nums[i]] = struct{}{}
	}

	return nil
}

// twoSum2 是针对递增数组的优化
func twoSum2(nums []int, target int) []int {
	l, r := 0, len(nums)-1

	for nums[l]+nums[r] != target {
		if nums[l]+nums[r] < target {
			l++
		} else {
			r--
		}
	}

	nums[0], nums[1] = nums[l], nums[r]
	return nums[:2]
}

// the sum of arr[l, l + 1, ..., r]
func getLRSum(l, r int) int {
	return (r + l) * (r - l + 1) >> 1
}

// 57 - II. 和为 s 的连续正数序列[*双指针滑动区间]
func findContinuousSequence(target int) [][]int {
	var ans [][]int
	// 滑动区间, (target >> 1) + (target >> 1 ) + 1 < target
	// l, r 活动到 (target >> 1 ) + 1 停止
	l, r := 1, 2
	for l < r {
		if getLRSum(l, r) == target {
			// 当前子序列和 arr[l, l+1, ..., r] 为target
			subSlice := make([]int, 0, r-l+1)
			for i := l; i <= r; i++ {
				subSlice = append(subSlice, i)
			}
			ans = append(ans, subSlice)
			// 添加完当前子序列，前移
			l++
		} else if getLRSum(l, r) < target {
			r++
		} else {
			l++
		}
	}

	return ans
}

// 58 - I. 翻转单词顺序
func reverseWords(s string) string {
	s = strings.TrimRight(s, " ")
	ans := make([]byte, 0, len(s))
	n := len(s)
	l, r := n-1, n-1
	for l > -1 {
		if s[l] != ' ' {
			l--
			continue
		}

		for i := l - 1; i <= r; i++ {
			ans = append(ans, s[i])
		}
		ans = append(ans, ' ')

		// 翻转后跳过接下来的空格
		for l < n && s[l] == ' ' {
			l--
		}
		r = l
	}

	if r > -1 {
		for l <= r {
			l++
			ans = append(ans, s[l])
		}
	}

	return string(ans)
}

// 58 - II. 左旋转字符串
func reverseLeftWords(s string, n int) string {
	if n < 1 {
		return s
	}

	return s[n:] + s[:n]
}

// 61. 扑克牌中的顺子
func isStraight(nums []int) bool {
	sort.Ints(nums)

	var i, zeroCount int
	for nums[i] == 0 {
		i++
	}
	zeroCount = i // i刚好为0的个数

	for i+1 < len(nums) {
		// 出现同牌情况，不是顺子
		if nums[i] == nums[i+1] {
			return false
		}

		// 检查是否连续
		if nums[i]+1 != nums[i+1] {
			// 不连续，通过大小王抵消
			zeroCount -= nums[i+1] - nums[i] - 1
			if zeroCount < 0 {
				return false
			}
		}

		i++
	}

	return true
}

// 评论方法 最差O(n)，不需要排序
// https://leetcode.cn/leetbook/read/illustrate-lcof/e2lnqc/
func isStraight2(nums []int) bool {
	min, max := 14, 1
	for i := range nums {
		if nums[i] == 0 {
			continue
		}

		if nums[i] > max {
			max = nums[i]
		}

		if nums[i] < min {
			min = nums[i]
		}
	}

	// 顺子区间[除0外]最大值跟最小值差值必须小于4
	if max-min > 4 {
		return false
	}

	// 出现重复牌，则不是顺子
	visit := make([]bool, 14)
	for i := range nums {
		if nums[i] != 0 {
			if !visit[nums[i]] {
				visit[nums[i]] = true
			} else {
				return false
			}
		}
	}

	return true
}

// 66. 构建乘积数组
func constructArr(a []int) []int {
	n := len(a)
	ans := make([]int, n)

	// 前乘
	l := 1
	for i := 0; i < n; i++ {
		ans[i] = l
		l *= a[i]
	}

	// 后乘
	l = 1
	for n > 0 {
		n--
		ans[n] *= l
		l *= a[n]
	}

	return ans
}

// 67. 把字符串转换成整数
// the_primary_algorithms 中的 string.go 有相同实现，可以参考
func strToInt(str string) int {
	var ans, i int
	// 抛弃空格
	n := len(str)
	for i < n && str[i] == ' ' {
		i++
	}
	// 如果抛弃完空格后没有字符，则直接返回
	if i >= n {
		return 0
	}

	// 判断符号
	var positive bool
	if str[i] < '0' || str[i] > '9' {
		if str[i] == '+' {
			positive = true
		} else if str[i] != '-' {
			return 0
		}
		i++
	} else {
		positive = true
	}

	// 组成数字
	lowerMax := math.MaxInt32 / 10
	k := math.MaxInt32 % 10
	for i < n {
		unit := str[i] - '0'
		if unit < 0 || unit > 9 {
			break
		}

		if ans > lowerMax || (ans == lowerMax && unit > 7) {
			ans = math.MaxInt32
			k++
			break
		}

		ans = ans*10 + int(unit)
		i++

	}

	if !positive {
		ans = ^ans + 1
	}

	if k > math.MaxInt32%10 && !positive {
		ans--
	}

	return ans
}
