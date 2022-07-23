package array_and_string

import (
	"algorithms/pkg/sort"
	"math"
)

// 1.ReverseString 利用双指针思想进行遍历一个byte数组
// 更节省内存的实现可以尝试单指针遍历半个数组，通过中心对称计算另一边的位置
func ReverseString(s []byte) {
	if len(s) == 0 {
		return
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

}


// 2.ArrayPairSum 使用了 快速3切分排序算法
// 数组拆分I
func ArrayPairSum(nums []int) int {
	sort.QuickSortBy3Way(nums)

	var sum int
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}

	return sum
}


// 3.TwoSum 寻找两数之和的索引
//   注意下面两个算法都是针对numbers为有序情况
//time:  O(n)
//space: O(1)
//前提:   numbers是有序数组
//TwoSum
func TwoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	if len(numbers) > 2 {

		for l < r {
			if numbers[l]+numbers[r] < target {
				// l + r < target, l右移 使和变大
				l++
			} else if numbers[l]+numbers[r] > target {
				// l + r > target , 使和变小
				r--
			} else {
				// l + r = target
				break
			}
		}
	}

	numbers[0] = l + 1
	numbers[1] = r + 1
	return numbers[:2]
}

//time:  O(lg(n))
//space: O(1)
//前提:   numbers是有序数组
//BinarySearchTwoSum 在TowSum 原基础上，通过二分查询来快速节省不必要的匹配
func BinarySearchTwoSum(numbers []int, target int) []int {

	l, r := 0, len(numbers)-1
	if len(numbers) > 2 {

		for l < r {
			m := l + (r-l)>>1
			if numbers[l]+numbers[m] > target {
				// case 1: 当 l + m > target时, r取中位数, 使和变小
				r = m
			} else if numbers[m]+numbers[r] < target {
				// case 2: 当 m + r < target时, l取中位数, 使和变大
				l = m
			} else if numbers[l]+numbers[r] == target {
				// case 3: l + r = target , 结束匹配

				break
			} else {
				if numbers[l]+numbers[r] < target {
					// case 4: l + r < target, l自增, 使和变大
					l++
				} else {
					// case 5: l + r > target, r自减， 使和变小
					r--
				}
			}
		}
	}

	numbers[0] = l + 1
	numbers[1] = r + 1
	return numbers[:2]
}

// 4.RemoveElement 原地移除 目标值为val的数
// 快慢指针
func RemoveElement(nums []int, val int) int {

	// 'fast' 代表快指针，遍历一遍数组，检查当前索引位置是否为val
	// 'slow' 代表慢指针，代表填入下一个不为val的索引位置, 同时也等于不含val的数组大小
	var fast, slow int
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	return slow
}


/*
 5.FindMaxConsecutiveOnes 寻找最长数值为1的连续子数组
 param:
   f: 遍历数组，判断当前数是否为断层（值为0），如果为断层，则f代表当前连续的1的末尾
   s: 指向当前连续的1的开头
   maxLen: 记录连续1的最长长度
*/
func FindMaxConsecutiveOnes(nums []int) int {
	var maxLen, f, s int
	
	for f < len(nums) {
		if nums[f] != 1 {
			if f-s > maxLen {
				maxLen = f - s
			}

			// 重新更新断层位置
			s = f + 1
		}

		f++
	}

	// 检查最后连续长度
	if f-s > maxLen {
		maxLen = f - s
	}

	return maxLen
}


// 6.MinSubArrayLen 寻找最短的子数组，且子数组之和 >= target
func MinSubArrayLen(target int, nums []int) int {
	minLen := len(nums) + 1
	var sum, f, s int

	// 如果未累加到末尾 或 已经到了末尾时，还有当前子数组和排除首部之和小于target，将退出
	for f < len(nums) {
		// 极端情况：如果存在一个数大于target，说明最短的子数组为1
		if nums[f] >= target {
			return 1
		}

		sum += nums[f]
		f++

		// 缩短子数组和，使之满足sum-nums[s]>=target的同时，长度最短
		for sum >= target {
			if f-s < minLen {
				minLen = f - s
			}

			sum -= nums[s]
			s++
		}
	}

	if minLen == len(nums)+1 {
		minLen = 0
	}

	return minLen
}

// deprecate 副产品算法
// MinSubArrayLenWithTarget 寻找目标和为target的最小连续子数组长度
func MinSubArrayLenWithTarget(target int, nums []int) int {
	var sum, f, s int
	minLen := math.MaxInt

	for ; f < len(nums); f++ {
		sum += nums[f]

		// case 1:如果 sum > target，则减去当前首部数值，并且s前移
		if sum > target {
			sum -= nums[s]
			s++
		}

		// case 2:如果刚好加上nums[f], sum = target，进行判断目标和为target的最小的长度
		// case 3:如果刚好减去nums[s], sum = target，进行判断目标和为target的最小的长度
		// 因为sum时当前 nums[s, f] 闭合区间之和，因此长度为 f - s + 1
		if sum == target && f-s+1 < minLen {
			minLen = f - s + 1
		}

	}

	// fmt.Println("-------------------------------")
	// case 4: 如果累加到末尾时，sum 仍大于 target，可能存留子数组和为target的情况，进行排除
	for sum > target && s < len(nums) {
		sum -= nums[s]
		s++

		// 由于 f = len(nums), 并且sum为当前nums[s, f-1]闭合区间之和，因此长度计算为 f - s
		if sum == target && f-s < minLen {
			minLen = f - s
		}
	}

	// 如果minLen 没有变化，说明minLen 为0，nums[0, last]不存在和为target的子数组
	if minLen == math.MaxInt {
		minLen = 0
	}

	return minLen
}

