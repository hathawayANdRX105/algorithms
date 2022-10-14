package illustrate_for_offer

import (
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

// findNumberIn2DArray2 dfs 从右上角走到左下角 O(m + n)
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
