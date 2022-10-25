package intermediate_algorithms

import (
	"math"
)

// 1.canJump 跳跃游戏
// 从第一个索引能否连续跳到最后一个索引，步距之间的中转索引都可以跳
func canJump(nums []int) bool {
	if len(nums) < 2 {
		return true
	}

	// 经过上一个语句，长度不为1，但是第一个跳跃值为0
	if nums[0] == 0 {
		return false
	}

	// 如果当前索引+最远跳跃 到达目标节点，则记录转移
	// 目标节点可能是 终点 或 到达终点的中转节点
	lastRelayIndex := len(nums) - 1
	for i := lastRelayIndex - 1; -1 < i; i-- {
		if nums[i]+i >= lastRelayIndex {
			lastRelayIndex = i
		}
	}

	// 如果最后一个中转节点是 第一个索引，说明存在至少一条跳跃路径，从起点到达终点
	return lastRelayIndex == 0
}

// 2.uniquePaths 不同路径
func uniquePaths(m int, n int) int {
	// cond1: 当 m = 1 或 n = 1 时，只有一条路
	if m == 1 || n == 1 {
		return 1
	}

	// cond2: 当 m = 2 或 n = 2 时， 最终结果都是 max(m, n)
	if m == 2 || n == 2 {
		if m < n {
			m = n
		}

		return m
	}

	cell := make([]int, n+1)

	// 初始化 2xn 的不同路径，[0，1，2,...,n]
	for i := 1; i < n+1; i++ {
		cell[i] = i
	}

	for i := 2; i < m; i++ {
		for j := 1; j < n+1; j++ {
			cell[j] += cell[j-1]
		}

	}

	return cell[n]
}

// 3.coinChange 零钱兑换
func coinChange(coins []int, amount int) int {

	if amount == 0 {
		return 0
	}

	amount += 1
	dp := make([]int, 1, amount)

	mask := 1 << 31
	// 1.先遍历背包，再遍历物品
	// for i := 1; i < amount; i++ {

	// 	dp = append(dp, mask)
	// // dp[i] 记录组成金额 i的最少硬币数量
	// 	for j := 0; j < len(coins); j++ {
	// 		if  coins[j] <= i  && dp[i-coins[j]]+1 < dp[i] {
	// 			dp[i] = dp[i-coins[j]] + 1
	// 		}
	// 	}
	// }

	// 2.遍历物品，再遍历背包
	// 初始化
	for i := 1; i < amount; i++ {
		dp = append(dp, mask)
	}

	for j := 0; j < len(coins); j++ {
		for i := coins[j]; i < amount; i++ {
			if dp[i-coins[j]] != mask && dp[i-coins[j]]+1 < dp[i] {
				dp[i] = dp[i-coins[j]] + 1
			}
		}
	}

	// 特殊情况，没有任何组合能组成目标金额
	if dp[amount-1] == mask {
		dp[amount-1] = -1
	}

	return dp[amount-1]
}

// 4.lengthOfLIS1 最长递增子序列
// time: O(nlgn)
func lengthOfLIS1(nums []int) int {

	// dp 用来记录不同长度的 最小序列值
	var maxR int // 记录 dp记录的最右边界
	dp := make([]int, 0, len(nums))

	for i := 1; i < len(nums); i++ {

		// O(n) 一起添加新值
		dp = append(dp, math.MaxInt)

		l, r := 0, len(dp)
		// 二分查询 取下底最接近 nums[i] 的长度位置
		for l < r {
			m := l + (r-l+1)>>1

			// 左靠右保 与中值无关（？）
			if dp[m] < nums[i] {
				l = m + 1
			} else {
				r = m
			}
		}

		// 记录当前长度 r 下 最小的末尾值
		if dp[l] < nums[i] {
			dp[l] = nums[i]
		}

		// 记录长度
		if maxR < l {
			maxR = l
		}
	}

	return maxR + 1
}

// lengthOfLIS2 O(n^2)
func lengthOfLIS2(nums []int) int {

	// dp 记录 nums同下标 开头的最长子序列长度
	maxLen := 1
	dp := make([]int, len(nums))

	for i := len(nums) - 1; -1 < i; i-- {

		// 每个子序列长度初始化为1
		dp[i] = 1

		// 遍历当前 序列值 能否与 已知的最长子序列长度 进行组合
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1

				// 记录最长子序列长度
				if maxLen < dp[i] {
					maxLen = dp[i]
				}
			}
		}

	}

	return maxLen
}
