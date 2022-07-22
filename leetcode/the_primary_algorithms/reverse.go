package the_primary_algorithms

func ReverseInt(x int) int {

	var p int
	stack := make([]int, 10)
	maxInt := 1<<31 - 1
	minInt := -1 << 31

	// 取个位数压栈
	for x != 0 {
		stack[p] = x % 10
		x /= 10
		p++
	}

	for multiple := 1; 0 < p; p-- {
		if stack[p-1]*multiple > maxInt-x || stack[p-1]*multiple < minInt-x {
			// 防止越界，保持增量范围合理
			// case1: x为正数， 正增量应小于 maxInt - x
			// case2: x为负数， 负增量应大于 minInt - x
			return 0
		}

		// 没有越界
		x += stack[p-1] * multiple
		multiple *= 10
	}

	return x
}
