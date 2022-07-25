package the_primary_algorithms

// 以下是 the_primary_algorithms 关于dynamic programming的代码实现部分

// 1.climbStairs 爬楼梯
func climbStairs(n int) int {
	if n < 3 {
		return n
	}

	var p1, p2, next int = 1, 2, 0
	for i := 3; i < n+1; i++ {
		next = p1 + p2
		p1, p2 = p2, next
	}

	return p2
}

// 2.maxProfit1 寻找最大收益
func maxProfit1(prices []int) int {

	// 已知前[0, j-1]的最大收益以及包含边界的最大收益
	// [0, j]的最大收益需要在考虑当前收益边界
	// 如果边界收益累加 小于 当前边界收益 可能需要重新规划入手股票
	// 如果边界收益累加 大于 [0, j-1]的最大收益，当前的边界收益可能是最大的收益
	var preMax, preBoundaryMax int
	for i := 1; i < len(prices); i++ {

		// 对数组做差价处理，显示第二天开始后一天卖出股票的收益
		if prices[i]-prices[i-1] > preBoundaryMax+prices[i]-prices[i-1] {
			preBoundaryMax = 0
		}

		preBoundaryMax += prices[i] - prices[i-1]
		if preBoundaryMax > preMax {
			preMax = preBoundaryMax
		}

	}

	return preMax
}

// max ...
func maxProfit2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	// 初始化
	prices[1] = prices[1] - prices[0]
	if prices[1] < 0 {
		// 如果前两天收益为负，最小价格为第二天价格，还原即可
		// 并且置最大收益为0
		prices[0], prices[1] = prices[1], 0
	}

	// 记录[0, i-1] 的最小值，如果当前prices[i] - min > maxProfit, 则记录最大利润
	for i := 2; i < len(prices); i++ {
		if prices[i] < prices[0] {
			prices[0] = prices[i]
		} else if prices[i]-prices[0] > prices[1] {
			prices[1] = prices[i] - prices[0]
		}
	}

	return prices[1]
}

// 3.maxSubArray 类似 question2
// dp解法的优化版
// dp 已经dp[0, i-1]的最大收益，判断当前dp[0,i]的最大收益在于nums[i]是否超过dp[0, i-1]的最大收益
// if dp[0,i-1] + nums[i] > dp[0,i-1] then dp[0,i] = dp[0, i-1] + nums[i]
// else dp[0, i] = dp[0, i-1]
func maxSubArray(nums []int) int {
	preMax, preBoundaryMax := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {

		preBoundaryMax += nums[i]
		if nums[i] > preBoundaryMax {
			preBoundaryMax = nums[i]
		}

		if preBoundaryMax > preMax {
			preMax = preBoundaryMax
		}

	}

	return preMax
}

// 4.rob 如果nums代表偷窃房子的金额收益，让你在不能连续两间房子偷窃的前提下，获取最高金额
func rob(nums []int) int {
	// 初始化 dp
	dp := append([]int{0, 0}, nums...)

	for i := 2; i < len(dp); i++ {

		// 检查隔间偷窃的最高金额 与 不偷当前房子 的最大收益
		dp[i] = dp[i] + dp[i-2]
		if dp[i-1] > dp[i] {
			dp[i] = dp[i-1]
		}

		// 初始版本
		// if dp[i]+dp[i-2] > dp[i-1] {
		// 	dp[i] += dp[i-2]
		// } else {
		// 	dp[i] = dp[i-1]
		// }
	}

	return dp[len(dp)-1]
}
