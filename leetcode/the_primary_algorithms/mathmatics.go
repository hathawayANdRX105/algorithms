package the_primary_algorithms

import (
	"fmt"
	"strconv"
)

// 1.fizzBuzz 根据求余打印出字符串数组
func fizzBuzz(n int) []string {
	ans := make([]string, n)

	for i := 1; i < n+1; i++ {

		if i%3 == 0 {
			ans[i-1] = "Fizz"
		}

		if i%5 == 0 {
			ans[i-1] += "Buzz"
		}

		if len(ans[i-1]) == 0 {
			ans[i-1] = strconv.Itoa(i)
		}

	}

	return ans
}

// 2.CountPrimes 欧拉线性筛
// 初始化枚举
var count []int

func init1() {
	// 题设 寻找小于 n 的素数数量，因此创建 n 大小数组，多出索引为0
	// prime[0] 代表素数数量
	maxN := 5000000
	prime := make([]int, maxN)
	count = make([]int, maxN+1)

	// 欧拉筛法 O(n)
	for i := 2; i < maxN; i++ {
		if prime[i] == 0 {
			count[i] = count[i-1] + 1
			prime[count[i]] = i
		} else {
			count[i] = count[i-1]
		}

		// 筛除关于当前素数的倍数（非素数）
		for j := 1; j <= count[i] && i*prime[j] < maxN; j++ {

			prime[i*prime[j]] = 1

			// 如果当前数已经被某项前素数筛选过，减少重复退出
			if i%prime[j] == 0 {
				break
			}
		}
	}

	// 埃氏筛法O(nlglgn)
	// var primeCount int // 非枚举法
	visit := make([]bool, maxN)
	for i := 2; i < maxN; i++ {
		// 枚举记录 count
		count[i] = count[i-1]

		if !visit[i] {
			// primeCount++
			count[i] += 1

			for j := i; j < maxN; j += i {
				visit[j] = true
			}
		}

	}
}

// 搭配提前枚举
func CountPrimes(n int) int {
	if n == 0 {
		return 0
	}

	return count[n-1]
}

// 3.IsPowerOfThree 判断当前数是否为3的k次幂
func IsPowerOfThree(n int) bool {
	// 1.算数法
	// n = 3^x
	// lgn = xlg3; x = lgn/lg3 如果x为整数，则n为3的x次幂
	// lgn 中 n > 0
	// if n <= 0 {
	// 	return false
	// }

	// x := math.Log10(float64(n)) / math.Log10(3.)
	// x -= float64(int32(x))
	// return x <= 1e-10 // float64精度转换问题

	// 2.利用2^31 - 1的最大3次幂1162261467 来求n的余数
	// 3^k = 1162261467     其他任何3的k-h次幂都可以用最大幂一部分表示
	return n > 0 && 1162261467%n == 0
}

// 4.romanToInt 将罗马字符串转化为数字
var romanDict = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func RomanToInt(s string) int {
	var ans, pre int

	for i := len(s) - 1; 0 <= i; i-- {
		fmt.Println(i, string(s[i]), romanDict[s[i]], romanDict)
		if romanDict[s[i]] < pre {
			ans -= romanDict[s[i]]
		} else {
			ans += romanDict[s[i]]
		}

		pre = romanDict[s[i]]
	}

	return ans
}
