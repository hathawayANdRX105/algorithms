package codewarsproblem

// Jaden Casing Strings
// https://www.codewars.com/kata/5390bac347d09b7da40006f6
// 每个英文单词转换位大写开头
func ToJadenCase(str string) string {

	var l int
	ans := []byte(str)
	for r := 0; r < len(str); r++ {
		if str[r] != ' ' {
			continue
		}

		//change to lower case and turn into upper case
		ans[l] = ans[l] | 0x20 ^ 0x20
		l = r + 1
	}

	ans[l] = ans[l] | 0x20 ^ 0x20

	return string(ans)
}

// Count the divisors of a number
// https://www.codewars.com/kata/542c0f198e077084c0000c2e
// time: O(lgn)
// Find combinations, and limit the push out boundaries, to be able to exit early
// e.g. Divisorts(64)
// ([1, 64], [2, 32], [4, 16], [8, 8])
// [l, r] exit cond is l >= r, just loop for 8 times, not 64
func Divisors(n int) int {
	// e.g. n = 1, ans = 1, (1)
	// e.g. n = 2, ans = 2, (1, 2)
	if n < 3 {
		return n
	}

	// escape (1, n) pair, start from ans = 2
	ans := 2
	l, r := 2, n/2
	// [1, n], [2, n/2] or [2, n/2 + 1]
	for ; l < r; l++ {
		// the result moded by most of number isn't target
		if n%l != 0 {
			r = n/l + 1
			continue
		}

		ans += 2
		r = n / l
	}

	l-- // cancel out for the last add oper
	// e.g. 25 => [5, 5] should add 1， but ans had added 2, then substract 1 for correction
	if l == r && n%l == 0 {
		ans--
	}

	return ans
}

// Beginner Series #3 Sum of Numbers
// https://www.codewars.com/kata/55f2b110f61eb01779000053
// time: O(lgn)
func GetSum1(a, b int) int {
	if a > b {
		a, b = b, a
	}

	var ans int
	for a < b {
		ans += a + b
		a++
		b--
	}

	if a == b {
		ans += a
	}

	return ans
}

// 计算，思维再打开些
// Sn = a1 + ... + an =  (a1 + an) * n / 2
func GetSum2(a, b int) int {
	if a > b {
		a, b = b, a
	}

	// err => (b - a + 1) / 2 * (a + b)
	// (b - a + 1) / 2 may loss precision
	// (7) / 2 => 3
	return (a + b) * (b - a + 1) / 2
}
