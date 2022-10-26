package illustrate_for_offer

// hammingWeight 15. 二进制中 1 的个数
func hammingWeight(num uint32) int {
	var count uint32
	for num > 0 {
		count += num & 1
		num >>= 1
	}

	return int(count)
}

// myPow 16. 数值的整数次方
func myPow(x float64, n int) float64 {
	// 对x^(-n)进行倒数处理
	if n < 0 {
		x = 1 / x
		n = ^n + 1
	}

	// 快速幂 n^m
	// 利用 n, n^2, n^4 ...,通过检查m来快速求出n^m
	ans := 1.0
	for n > 0 {
		if n&1 == 1 {
			ans *= x
		}
		x *= x
		n >>= 1
	}

	return ans
}

// singleNumbers 56 - I. 数组中数字出现的次数
// 数组中只有两个数字是单次出现，请找出来 [ O(2n) ]
func singleNumbers(nums []int) []int {
	// 如果是寻找单个数字，可以直接全部异或找到
	var t int
	for _, v := range nums {
		t ^= v
	}

	// 但是出现两个单次数字，全部异或之后， t = a ^ b,因为异或运行，t二进制上出现为1必定是a，b不同情况
	// 可能是该位上a为1或0，b为0或1，因此借此找到最低位不同处进行划分，其他相同元素重复出现也是可以异或抵消
	div := 1
	for t&div == 0 {
		div++
	}

	var a, b int
	for _, v := range nums {
		if v&div != 0 {
			a ^= v
		} else {
			b ^= v
		}
	}

	return []int{a, b}
}

// add 65. 不用加减乘除做加法
func add(a int, b int) int {
	// a是进位数， b是上一个排除进位数的保留值
	if a&b == 0 {
		return a | b
	}

	return add((a&b)<<1, (a|b)^(a&b))
}

func optimizedAdd(a, b int) int {
	for a&b != 0 {
		// 进位数，左移之后是真正的进位数
		carry := a & b
		// b需要先保留a与b，所有位信息，然后排除掉需要进位位置的1
		b = (a | b) ^ carry
		// a用来保存下一个需要进位检查的值
		a = carry << 1
	}

	return a | b
}
