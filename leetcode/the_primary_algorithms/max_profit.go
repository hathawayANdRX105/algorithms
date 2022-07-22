package the_primary_algorithms

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
