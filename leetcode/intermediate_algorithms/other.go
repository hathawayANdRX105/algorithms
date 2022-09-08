package intermediate_algorithms

import (
	"strconv"
)

// 以下是 intermediate_algorithms 关于 other 的代码实现部分

// 1.getSum 两整数之和，不能直接使用+/-运算
func getSum(a int, b int) int {
	// a为保留进制数， b为进位数
	if b == 0 {
		return a
	}

	return getSum(a^b, (a&b)<<1)
}

// 2.evalRPN 逆波兰式运算
func evalRPN(tokens []string) int {

	stack := make([]int, 0, (len(tokens)+1)>>1)
	// var p int

	for _, v := range tokens {
		switch v {
		case "+":
			stack[len(stack)-2] += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "-":
			stack[len(stack)-2] -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "*":
			stack[len(stack)-2] *= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "/":
			stack[len(stack)-2] /= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		default:
			val, _ := strconv.Atoi(v)
			stack = append(stack, val)
		}

	}

	return stack[0]
}

// 3.majorityElement 多数元素
func majorityElement(nums []int) int {

	// 按次数抵消的方式，记录最多的众数
	ans, count := nums[0], 1

	for _, v := range nums {
		// case 1：记录数相同则，累加count，跳过，减少后面的判断
		if ans == v {
			count++
			continue
		}

		// case 2:当前数与记录数不同，减少count
		count--

		// case 3：如果count为零，需要记录新的可能众数
		if count == 0 {
			ans = v
			count++
		}
	}

	return ans
}

// 4.leastInterval 任务调度器
// [A, A, B], n = 2  --> A-B-x-A
func leastInterval(tasks []byte, n int) int {
	// n 代表了A-x-x-A 的间隔时间
	// cond 1: n 为0时，自由排列任务
	if n == 0 {
		return len(tasks)
	}

	// 输入n代表不同任务之间的间隔，为了最小化时间，需要将最大数量相同任务并排在一起
	// 目标：满足最多任务之间填充的间隔是最小时间
	// 模式：间隔 A-B-x-x-A-B-x-x + 尾部 A-B-C
	// 最小任务时间，有两个结果：1.原先不同顺序的并排，相当于数量总和时间 2.除了间隔填充任务外，需要尾部填充更多的任务，使得时间加长

	dict := make([]int, 26)
	// maxSize 代表某个任务的最多数量，可能是某几个任务数量都相同，则最多数量也相同
	// maxSizeOfCount 代表整个任务列表中最多数量的任务个数
	// ex:[A, A, B, B, C] --> maxSize=2 [A,A] / [B, B] // maxSizeOfCount=2 有两个任务[A,B]数量最多
	var maxSize, maxSizeOfCount int

	for _, c := range tasks {
		dict[c-'A']++
		if maxSize < dict[c-'A'] {
			maxSize = dict[c-'A']
		}
	}

	for _, size := range dict {
		if size == maxSize {
			maxSizeOfCount++
		}
	}

	n++       // n++ 代表 A-x-...-x 需要的间隔时间
	maxSize-- // maxSize-- 方便计算 总数量 A-x-...-x 间隔的时间
	ans := n*maxSize + maxSizeOfCount
	if ans > len(tasks) {
		return ans
	}

	return len(tasks)
}
