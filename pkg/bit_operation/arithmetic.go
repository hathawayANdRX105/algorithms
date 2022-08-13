package bit_operation

// Add 利用位运算进行加法运算
// a^b 代表了按位异或运算，保留了无进位和
// a&b 代表了按位与运算下，需要进位的位置
// 那么(a&b)<<1 表示 进位和
func Add(a, b int) int {
	var ans int

	// 让每次的进位和都需要比对无进位和是否仍需要进位，如果没有，则说明完成加法
	for b != 0 {
		ans = a ^ b
		b = (a & b) << 1

		// fmt.Printf("a:%v, b:%v\n", strconv.FormatInt(int64(a), 2), strconv.FormatInt(int64(b), 2))
		// fmt.Printf("a^b:%v, a&b<<1:%v\n", strconv.FormatInt(int64(a^b), 2), strconv.FormatInt(int64((a&b)<<1), 2))
		// fmt.Println()

		a = ans
	}

	return ans
}

// recursiveAdd Add 的递归版本
func recursiveAdd(sum, carry int) int {
	// 如果不存在进位，则返回结果
	if carry == 0 {
		return sum
	}

	return recursiveAdd(sum^carry, (sum&carry)<<1)
}

// negative 正负数取反码都需要进行补位
func negative(a int) int {
	// 正数表示为负数，正数取反码，需要对应负数的补码，需要最后一位加1
	// 负数同理
	return Add(^a, 1)
}

func Sub(a, b int) int {
	// 交换律： a - b => a + (-b)
	return Add(a, negative(b))
}

// Multiply ...
// 原理： b个a相加
/* 快速的倍数相加:
*  ex1: 假设 a=3, b=7， 求 3*7的结果， 7的二进制为 0111， 3的二进制为 0011
* 	    可以分解为 0001 个 0011， 0010个 0011， 0100个 0011
*    	即 1个3， 2个3， 4个3， 相当于 1个3， 1个6， 1个12
*
*  ex2: 假设 a=3, b=5， 求 3*5的结果， 5的二进制为 0101， 3的二进制为 0011
* 	    可以分解为 0001 个 0011， 0100个 0011
* 	    即 1个3， 4个3， 相当于 1个3， 1个12
 */
func Multiply(a, b int) int {
	if b < 0 {
		// 由于位移运算， 当b>0时，无论a是否为正数都不收影响
		// 但当 b<0 时，终止条件为右移到0，符文位
		a = negative(a)
		b = negative(b)
	}

	var ans int

	for b != 0 {
		// 检查当前倍数下的a是否组成最终的结果数，如果组成则相加，否则跳过
		if b&1 != 0 {
			ans = Add(ans, a)
		}

		// 倍数的转移
		a <<= 1
		b >>= 1
	}

	return ans
}

// isNegative ...
func isNegative(a int) bool {
	return a&(1<<31) != 0
}

// Divide ...
func Divide(a, b int) int {
	symbolA := isNegative(a)
	symbolB := isNegative(b)

	if symbolA {
		a = negative(a)
	}

	if symbolB {
		b = negative(b)
	}

	var ans int

	// 从全局进行shift，当 a << i到某i时刚好大于时，说明 能够包含 b << i
	// 即 b * (1<<i) 或 b * 2^i ，ans只需要加上 2^i的倍数，最终凑成最大的组成数即可
	for i := 31; -1 < i; i-- {
		if (a >> i) >= b {
			ans += 1 << i
			a -= b << i
		}
	}

	if symbolA != symbolB && (symbolA || symbolB) {
		ans = negative(ans)
	}

	return ans
}
