package illustrate_for_offer

import "math"

// 10 - I. 斐波那契数列
func fib(n int) int {
	if n < 2 {
		return n
	}

	dp := make([]int, 2, n+1)
	dp[1] = 1
	for i := 2; i < n+1; i++ {
		dp = append(dp, (dp[i-2]+dp[i-1])%1000000007)
	}

	return dp[n]
}

// 节省内存 使用三个变量交替，因为子问题紧凑 dp[i] = dp[i-1] + dp[i-2]
func optimizedFib(n int) int {
	if n < 2 {
		return n
	}

	l, r := 0, 1
	for i := 2; i < n+1; i++ {
		tmp := (l + r) % 1000000007
		l = r
		r = tmp
	}

	return r
}

// 10 - II. 青蛙跳台阶问题
// numWays 前4结果： 1, 1, 2, 3, 5
func numWays(n int) int {
	dp := make([]int, 2, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < n+1; i++ {
		dp = append(dp, (dp[i-1]+dp[i-2])%1000000007)
	}

	return dp[n]
}

func optimizedNumWays(n int) int {
	l, r := 1, 1
	for i := 2; i < n+1; i++ {
		tmp := (l + r) % 1000000007
		l, r = r, tmp
	}

	return r
}

// maxSubArray 子问题为 到n为止(包括下标n）的子数组最大和是 前一个最大子数组和 还是 当前元素大
// 详解 => https://leetcode.cn/leetbook/read/illustrate-lcof/xsus9h/
func maxSubArray(nums []int) int {
	n := len(nums)
	max := nums[0]
	dp := make([]int, 0, n)
	dp = append(dp, max)

	for i := 1; i < n; i++ {
		dp = append(dp, nums[i])

		// 前子数组和为正数，可能存在相加为最大值
		if dp[i-1] > 0 {
			dp[i] += dp[i-1]
		}

		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}

func optimizedMaxSubArray(nums []int) int {
	max := -101
	var boundarySum int
	for i := range nums {
		boundarySum += nums[i]

		if nums[i] > boundarySum {
			boundarySum = nums[i]
		}

		if boundarySum > max {
			max = boundarySum
		}
	}

	return max
}

// 46. 把数字翻译成字符串
func translateNum(num int) int {
	if num < 1 {
		return 1
	}

	var nums []int
	for num > 0 {
		nums = append(nums, num%10)
		num /= 10
	}

	n := len(nums)
	single := make([]int, n+1)
	double := make([]int, n+1)
	// 前提条件:每个数字都能代表1个字母
	single[0], single[1] = 1, 1
	for i := 2; i < n+1; i++ {
		single[i] = single[i-1] + double[i-1]

		// 倒序分析，本质上跟正序无差异
		if nums[i-1] > 0 && nums[i-1]*10+nums[i-2] < 26 {
			double[i] = double[i-2] + single[i-2]
		}
	}

	return single[n] + double[n]
}

// optimized 2
func optimizedTranslateNum(num int) int {
	if num < 1 {
		return 1
	}

	var nums []int
	for num > 0 {
		nums = append(nums, num%10)
		num /= 10
	}
	n := len(nums)

	// 法1： dp 数组解决
	//	dp := make([]int, n+1)
	//	dp[0], dp[1] = 1, 1
	//	for i := 2; i < n+1; i++ {
	//
	//		dp[i] = dp[i-1]
	//		// 倒序分析，本质上跟正序无差异
	//		if nums[i-1] > 0 && nums[i-1]*10+nums[i-2] < 26 {
	//			dp[i] += dp[i-2]
	//		}
	//	}
	//
	//	return dp[n]

	// 法2：状态变量转移
	l, r := 1, 1
	for i := 2; i < n+1; i++ {
		cur := r
		// 倒序分析，本质上跟正序无差异
		if nums[i-1] > 0 && nums[i-1]*10+nums[i-2] < 26 {
			cur += l
		}

		l, r = r, cur
	}

	return r
}

// 47. 礼物的最大价值
// maxValue 利用二维数组 dp 计算从上从左来的加权最大值
func maxValue(grid [][]int) int {
	r, c := len(grid), len(grid[0])
	dp := make([][]int, 0, r)
	level := make([]int, 0, c)
	level = append(level, grid[0][0])
	for ci := 1; ci < c; ci++ {
		level = append(level, level[ci-1]+grid[0][ci])
	}
	dp = append(dp, level)

	for ri := 1; ri < r; ri++ {
		level := make([]int, 0, c)
		level = append(level, dp[ri-1][0]+grid[ri][0])

		for ci := 1; ci < c; ci++ {
			max := level[ci-1]
			if max < dp[ri-1][ci] {
				max = dp[ri-1][ci]
			}
			level = append(level, max+grid[ri][ci])
		}
		dp = append(dp, level)
	}

	// for i:=0; i < r; i++{
	//     fmt.Println(dp[i])
	// }
	return dp[r-1][c-1]
}

// 优化，将多层持久化为同一层，减少不必要的空间浪费
func optimizedMaxValue(grid [][]int) int {
	r, c := len(grid), len(grid[0])
	level := make([]int, c)

	for ri := 0; ri < r; ri++ {
		level[0] += grid[ri][0]
		for ci, l2r := 1, level[0]; ci < c; ci++ {
			if l2r > level[ci] {
				l2r += grid[ri][ci]
				level[ci] = l2r
			} else {
				level[ci] += grid[ri][ci]
				// l2r 需要继承当前层从左到右上一个dp状态的最优解
				l2r = level[ci]
			}
		}
	}

	return level[c-1]
}

// 49. 丑数 [*]
// nthUglyNumber 详解：https://leetcode.cn/leetbook/read/illustrate-lcof/5065hi/
func nthUglyNumber(n int) int {
	// 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 前十个丑数
	// 7之前都为顺序
	if n < 7 {
		return n
	}
	dp := make([]int, 0, n)
	dp = append(dp, 1)

	var a, b, c int
	for n > 1 {
		n--
		ma, mb, mc := dp[a]*2, dp[b]*3, dp[c]*5
		minVal := min(min(ma, mb), mc)
		dp = append(dp, minVal)

		// 如果 ma, mb, mc 出现相同值时，a, b, c 指针都向前移动一位
		// i.e. 2x3 = 3x2, 可能选取重复值，避免此情况，需要将a,b 同时自增
		// 所以下面判断不需要联合
		if minVal == ma {
			a++
		}
		if minVal == mb {
			b++
		}
		if minVal == mc {
			c++
		}
	}

	return dp[len(dp)-1]
}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}

// lastRemaining 不用优化版解释，不好懂
//  62. 圆圈中最后剩下的数字
//     解析文章： https://blog.csdn.net/u011500062/article/details/72855826
//     leetcode解析：https://leetcode.cn/circle/article/BOoxAL/
func lastRemaining(n int, m int) int {
	dp := make([]int, n)

	// 经典约瑟夫环解， 假设下面都用m=3
	// 当n=1时， [0]的解为0， 初始解无过程
	// 当n=2时， [0, 1]的解为 1
	//		过程：->0, 0->1, 1->0(x0)
	// 当n=3时， [0, 1, 2]的解为 1
	//		过程：->0, 0->1, 1->2(x2), 2->0, 0->1, 1->0(x0)
	// 当n=4时， [0, 1, 2, 3]的解为 0
	//		过程：->0, 0->1, 1->2(x2), 2->3, 3->0, 0->1(x1), 1->3, 3->0, 0->3(x3)

	// 对比n=2与n=3过程可以发现，[[1->2(x2), 2->0, 0->1]]这过程是独属于n=3
	// 对比n=3与n=4过程可以发现，[[2->3, 3->0, 0->1(x1)]]这过程是独属于n=4
	for i := 1; i < n; i++ {
		dp[i] = (dp[i-1] + m) % (i + 1)
	}
	return dp[n-1]
}

func optimizedLastRemaining(n int, m int) int {
	var last int

	// i 代表 多少个数，与使用dp的方法含义不一样[dp[i]代表下标, i+1 代表个数]
	for i := 2; i < n; i++ {
		last = (last + m) % i
	}
	return last
}

// maxProfit  63. 股票的最大利润 [转换为求最大子数组和]
func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}

	// prices => [1, 3, 2, 4]
	// 转换为dp为[2,-1,2]
	// 最大利润为4-1=3
	// 但可以转换为，后一天对前一天的利润净赚，(3-1) + (2-3) + (4-2) => 4-1,画图能更好的理解
	// 此处dp只用作于mental model

	var max, boundarySum int
	n := len(prices) - 1
	dp := make([]int, 0, n)
	for i := 0; i < n; i++ {
		dp = append(dp, prices[i+1]-prices[i])
		boundarySum += dp[i]

		if dp[i] > boundarySum {
			boundarySum = dp[i]
		}

		if boundarySum > max {
			max = boundarySum
		}
	}

	//fmt.Println(dp)

	return max
}

func optimizedMaxProfit1(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	n := len(prices) - 1
	var max, boundarySum int
	for i := 0; i < n; i++ {
		netValue := prices[i+1] + prices[i]
		boundarySum += netValue

		if netValue > boundarySum {
			boundarySum = netValue
		}

		if boundarySum > max {
			max = boundarySum
		}
	}

	return max
}

func optimizedMaxProfit2(prices []int) int {
	if len(prices) < 1 {
		return 0
	}

	n := len(prices)
	var max int
	min := math.MaxInt
	for i := 0; i < n; i++ {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i]-min > max {
			max = prices[i] - min
		}
	}

	return max
}
