package intermediate_algorithms

import (
	"fmt"
	"math"
	"strconv"
)

// 以下是 intermediate_algorithms 关于 mathmatics 的代码实现部分

// 1.isHappy 快乐数
func isHappy(n int) bool {
	set := map[int]struct{}{}

	var nextN, single int
	for n != 1 {

		// 取个位数平方组成下一个候选快乐数
		for n > 0 {
			single = n % 10
			nextN += single * single

			n /= 10
		}

		// 查询是否重复候选快乐数，如果是，则原先的 n 不是快乐数
		if _, ok := set[nextN]; ok {
			return false
		}

		// 寻找下一个候选快乐数，重置条件
		n = nextN
		nextN = 0
	}

	return true
}

// 2.trailingZeroes 阶乘后的零
// trailingZeroes1 保留取零的后几位，以5的倍数为跳跃数
func trailingZeroes1(n int) int {
	var count int

	for i, last := 5, 1; i <= n; i += 5 {
		last *= i * (i - 1)

		mask := 10
		for last > 0 && last%10 == 0 {
			count++
			mask++

			last /= 10
		}

		last %= mask
	}

	return count
}

// trailingZeroes2
// n! = 1 x 2 x 3 x 4 x 5 x 6 x 7 x 8 x 9 x 10 x ... x n
// n! = 1 x 2 x 3 x (2x2) x 5 x (2x3) x 7 x (2x2x2) x 9 x (2x5)
// 进位0， 2x5=10
// 由于 n阶乘中质因数2数量越多于5，所以只需要求质因数5的个数
func trailingZeroes2(n int) int {
	var count int

	for n > 0 {
		n /= 5
		count += n
	}

	return count
}

// 3.titleToNumber excel表序列号
// 26个字母做序号，相当于 26进制，从后往前计算，前进一个字符，就相当于26多一个次方
func titleToNumber(columnTitle string) int {
	var count int

	for i := 0; i < len(columnTitle); i++ {
		count = int(columnTitle[i]-64) + count*26

		// fmt.Println(multiple, columnTitle[i] - 64, columnTitle[i])
	}

	return count
}

// 4.myPow Pow(x, n)
// 按幂运算规则，进行二进制凑数
func myPow(x float64, n int) float64 {
	ans := 1.0
	symbol := n < 0

	// n 换算为二进制 有 n = 1 + 2 + 4 + ... + 2^k(偶数)
	// ans = x^1 * x^2 * x^4 * ... * x^k
	for n != 0 {
		if n%2 != 0 {
			ans *= x
		}

		x *= x
		n /= 2
	}

	// 如果 n为负数，求倒数的n次幂
	if symbol {
		ans = 1 / ans
	}

	return ans
}

// 5.mySqrt x的平方根
// 牛顿迭代法 f(x) = x^2 - n
// 在 f(x) 已知 点(x[i], f(x[i]))
// 切线方程为 y - f(x[i]) = f'(x[i])(x[i+1] - x[i])
// 求与x轴的点 (x[i+1], 0)
// 代入，f(x), f'(x)方程简化为 x[i+1] = x[i] - (x[i]^2 - n) / (2* x[i])
//
//	即 x[i+1] = x[i] - (x[i] - n / x[i]) / 2
func mySqrt(x int) int {

	x1, t, norm := 1.0, float64(x), 1.0

	for math.Abs(norm) > 1e-5 {
		norm = (x1 - t/x1) / 2
		x1 -= norm
	}

	return int(x1)
}

// 6.divide 不能使用乘除法，求余运算计算除数
// 详细原理跟注释在 bit_operation 中的 Divide 以及 bit_opperation docs 中
func divide(dividend int, divisor int) int {
	symbol1 := dividend < 0
	symbol2 := divisor < 0

	if symbol1 {
		dividend = ^dividend + 1
	}

	if symbol2 {
		divisor = ^divisor + 1
	}

	var ans int
	for i := 31; -1 < i; i-- {
		if dividend>>i >= divisor {
			ans += 1 << i
			dividend -= divisor << i
		}
	}

	if symbol1 != symbol2 && (symbol1 || symbol2) {
		ans = ^ans + 1
	}

	if ans > 1<<31-1 {
		return 1<<31 - 1
	}

	return ans
}

// 7.fractionToDecimal 分数到小数
func fractionToDecimal(numerator int, denominator int) string {
	// cond 1: 被除数是0
	if numerator == 0 {
		return "0"
	}

	// res为格式化 numerator/denominator 的小数表示
	// uniqueset 用于记录每次求除数的余数
	res := strconv.FormatFloat(float64(numerator)/float64(denominator), 'f', -1, 64)
	uniqueSet := map[uint64]int{}

	// 保证 求除数跟余数是正数
	if numerator < 0 {
		numerator = ^numerator + 1
	}

	if denominator < 0 {
		denominator = ^denominator + 1
	}

	// 防止 numerator/denominator是 -2^31 超大数进行 运算
	// rem 是 numerator 减去 k * denominator 的从小数点第一位开始的余数
	rem := uint64(numerator - (numerator/denominator)*denominator)
	denominator2 := uint64(denominator)

	// 寻找小数点，从小数点后开始
	var r int
	for r < len(res) && res[r] != '.' {
		r++
	}

	// case 1: 如果不存在小数点，说明整除，直接返回
	if r >= len(res) {
		// 整除
		return res
	}

	// case 2: 有损的计算下能找到循环小数
	// 防止最后一位四舍五入，res直接取除了最后一位的格式化字符串
	res = res[:len(res)-1]
	for r++; r < len(res); r++ {
		pos, ok := uniqueSet[rem]

		// 先找余数是否重复出现，后计算下一个除数的余数
		if !ok {
			uniqueSet[rem] = r

			rem *= 10
			rem -= uint64(res[r]-'0') * denominator2
		} else {
			// 如果余数重复出现，那么一定存在循环的计算，则将 res[pos:r] 用“（）”包夹返回
			return res[:pos] + "(" + res[pos:r] + ")"
		}
	}

	// case 3：没计算完的小数后面可能存在循环小数
	// 如果进行四舍五入的计算，格式化字符直到最后都没找到，并且最后没计算的一位余数不为0，说明存在可能的循环小数
	if rem != 0 {

		// 复制到2倍大的rune数组
		resRune := make([]rune, len(res), len(res)<<1)
		copy(resRune, []rune(res))

		// 寻找知道余数为0，且字符串长度小于 10^4 的表示
		for ; rem != 0 && r < 1e4; r++ {
			// 上一次 退出循环，包括后计算的余数需要先检查
			if pos, ok := uniqueSet[rem]; ok {
				return string(resRune[:pos]) + "(" + string(resRune[pos:]) + ")"
			}

			// 余数没有出现重复，则记录，计算添加当前的除数，计算下一个除数的余数
			uniqueSet[rem] = r

			lastNum := rem * 10 / denominator2
			resRune = append(resRune, rune(lastNum)+'0')

			rem = rem*10 - lastNum*denominator2
		}

		// case 4: 计算完有损的小数形式
		// 返回当前的计算小数的字符串
		return string(resRune)
	}

	// case 5: 第一次计算小数完全，且不存在循环小数
	return res
}
