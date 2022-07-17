package array_and_string

import (
	"math"
)

func MinSubArrayLen(target int, nums []int) int {

	minLen := len(nums) + 1
	var sum, f, s int

	// 如果未累加到末尾 或 已经到了末尾时，还有当前子数组和排除首部之和小于target，将退出
	for f < len(nums) {
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

// MinSubArrayLenWithTarget 寻找目标和为target的最小连续子数组长度
// deprecate
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
