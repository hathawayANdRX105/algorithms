package illustrate_for_offer

import "strconv"

// findNthDigit

func findNthDigit(n int) int {
	// [1-9] 9个元素，9个个位数
	// [10-99] 90个元素， 90 * 2 个个位数
	// [100-999] 900 个元素， 900*3 个个位数
	digit := 1    // 用于记录几位数
	start := 1    // 用于记录开始的数值， 1-10-100-1000...
	curCount := 9 // 记录每个区间的总个位数
	for n > curCount {
		n -= curCount

		digit++
		start *= 10
		curCount = start * digit * 9 // 1 * 1 * 9 // 2 * 10 * 9
	}

	n--
	// 此时n代表在某个区间开始，n个个位数，迁移
	// 因此 (n-1)/digit 代表需要加上值
	// i.e. 100,101 [1,2,3,4,5,6]
	// 如果n<4,则不需要加，直接100开始，如果每超过三则需要加1
	// 并且相对于每三个小值区间选择最后的个位数，因此需要取余(%digit)
	// 偏移之后开始判断拼接最近数字
	v := start + n/digit
	return int(strconv.FormatInt(int64(v), 10)[n%digit] - '0') // 利用库函数转为字符串方便取值
}
