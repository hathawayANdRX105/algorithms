package the_primary_algorithms

import (
	"container/list"
	"math"
)

// 1.HammingWeight1 利用 x = lgn / lg2 求出底数，减去1<<x
func HammingWeight(num uint32) int {
	var count, shift int
	var i uint32 = 1
	for 0 < num {
		// 使用math.log 求底数
		shift = int(math.Log(float64(num)) / math.Log(2.))

		num = num - (i << shift)
		count++
	}

	return count
}

func HammingWeight2(num uint32) int {
	var count int
	for 0 < num {
		// 检查二进制最后一位，是否为1
		if num&1 > 0 {
			count++
		}

		num = num >> 1
	}

	return count
}

// 2.hammingDistance 寻找 x 与 y 在二进制上 1出现位置不同的个数
func hammingDistance(x int, y int) int {
	var distance int
	x ^= y
	for x > 0 {
		if x&1 > 0 {
			distance++
		}

		x = x >> 1
	}

	return distance
}

// 3.reverseBits 翻转整个uint32二进制位置
func reverseBits(num uint32) uint32 {
	var ans uint32

	// 注意到32位中第32位是不用偏移，因此实际上只有31次偏移
	for i := 1; i < 32; i++ {
		ans += num & 1
		num = num >> 1
		ans = ans << 1
	}

	ans += num & 1
	return ans
}

// 4.generate 杨辉三角的生成 已经在 array_and_string 的 summary.go 中实现

// 5.isValid 检查字符串中括号是否对称
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := list.New()

	for i := 0; i < len(s); i++ {
		if 0 < stack.Len() && s[i]-stack.Back().Value.(byte) < 3 && s[i] != stack.Back().Value {
			stack.Remove(stack.Back())
		} else {
			stack.PushBack(s[i])
		}
	}

	return stack.Len() == 0
}

// isValid2 数组栈 节约一点内存
func isValid2(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	var p int
	stack := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		if 0 < p && s[i]-stack[p-1] < 3 && s[i] != stack[p-1] {
			p--
		} else if p < len(stack) {
			stack[p] = s[i]
			p++
		} else {
			stack = append(stack, s[i])
			p++
		}
	}

	return p == 0
}

// 6.missingNumber 寻找[0, n] 中没有出现的数字
// 思路：由于全部出现时 [0, n] => 0 + 1 + ... + n = (n+1)*n/2
// total = (n+1)n/2  减去所有数后便是没有出现的数，包括0
func missingNumber(nums []int) int {
	total := (len(nums) + 1) * len(nums) / 2

	for _, v := range nums {
		total -= v
	}

	return total
}

// 利用异或运算 a^a = 0
// [0, n] 只会出现 n-1个数，索引迭代只会到n-1
// 因此从mask为n开始，能遍历[0,n] 2n-1 个数
func missingNumber2(nums []int) int {
	mask := len(nums)

	for i, v := range nums {
		mask ^= i ^ v
	}

	return mask
}
