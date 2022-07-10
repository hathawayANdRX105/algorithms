package array_and_string

import (
	"fmt"
)

// dp solved
// time:  O(n^2)
// space: O(n^2)
func LongestPalindrome1(s string) string {
	length := len(s)
	if length < 2 {
		return s
	}

	dp := make([][]bool, length)

	// initialize diagonal with true meant that single letters are palindrome.
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
		dp[i][i] = true
	}

	var beginIndex int
	maxLen := 1

	for j := 1; j < length; j++ {
		for i := 0; i < j; i++ {

			if s[i] != s[j] {
				dp[i][j] = false
			} else if j-i < 3 {
				// check s[i, j] is the base odd or even palindrome sub string
				// the base odd or even palindrome sub string length is 2 or 3
				// the sub string length: j - i + 1
				// boundary: j - i + 1 < 4
				// ex: aa, aba

				dp[i][j] = true
			} else {
				// check whether s[i+1, j-1] is palindrome string
				dp[i][j] = dp[i+1][j-1]
			}

			if dp[i][j] && j-i+1 > maxLen {
				// record longest palindrome index
				// s[j, i] is palindrome and j-i+1 is greater than longJ-longI+1
				beginIndex = i
				maxLen = j - i + 1
			}
		}
	}
	

	return s[beginIndex : beginIndex+maxLen]
}



// dp
// time:  O(n^2)
// space: O(n)
func LongestPalindrome2(s string) string {
	length := len(s)
	if length < 2 {
		return s
	}

	dp := make([]bool, length)

	var beginIndex int
	maxLen := 1

	for j := 1; j < length; j++ {

		for i := 0; i < j; i++ {

			if s[i] != s[j] {
				dp[i] = false
			} else if j-i < 3 {
				// check s[i, j] is the base odd or even palindrome sub string
				// the base odd or even palindrome sub string length is 2 or 3
				// the sub string length: j - i + 1
				// boundary: j - i + 1 < 4
				// ex: aa, aba

				dp[i] = true
			} else {
				// Given previous state, check whether s[i+1, j-1] is palindrome string
				dp[i] = dp[i+1]
			}

			if dp[i] && j-i+1 > maxLen {
				// record longest palindrome index
				// s[j, i] is palindrome and j-i+1 is greater than longJ-longI+1
				beginIndex = i
				maxLen = j - i + 1
			}
		}

		// diagonal should be true that meant single letter also is palindrome sub string.
		dp[j] = true

	}

	return s[beginIndex : beginIndex+maxLen]
}

// center expande method
// time:  O(n^2)
// space: O(1)
func LongestPalindrome3(s string) string {
	length := len(s)
	if length < 2 {
		return s
	}

	expandAroundCenter := func(s string, left, right int) int {
		for 0 <= left && right < len(s) && s[left] == s[right] {
			left--
			right++
		}

		// length = right - left + 1 - 2 (because it post movement)
		return right - left - 1
	}

	var startIndex, maxLen int
	for i := 0; i < length; i++ {
		oddLen := expandAroundCenter(s, i, i)
		evenLen := expandAroundCenter(s, i, i+1)

		fmt.Println(i, oddLen, evenLen)
		if oddLen > evenLen {
			evenLen = oddLen
		}

		if evenLen > maxLen {
			maxLen = evenLen
			startIndex = i - (maxLen-1)/2
		}

	}

	return s[startIndex : startIndex+maxLen]
}

// preProcess 对每个字符间隔插入#，使之成为奇回文形式，并且为了方便遍历在首位各插入不相同#的两个符号
func preProcess(s string) string {
	size := len(s)
	processString := make([]byte, 2*(size+1)+1)

	for i := 0; i < size; i++ {
		processString[2*i+1] = '#'
		processString[2*(i+1)] = s[i]
	}

	processString[0] = '^'
	processString[2*size+1] = '#'
	processString[2*(size+1)] = '$'

	return string(processString)
}

// LongestPalindromeWithManacher
// 马拉夫算法：根据已有的回文，来快速推算当前字符为中心的回文长度,减少一定次数的遍历
// time:  O(n)
// space: O(n)
func LongestPalindromeWithManacher(s string) string {
	if len(s) < 2 {
		return s
	}

	T := preProcess(s)       // 字符间隔插入#
	p := make([]int, len(T)) //记录字符回文的长度

	var C, R int
	var startCenter, maxLen int
	for i := 1; i < len(p)-1; i++ {

		if R > i {
			i_mirror := 2*C - i

			// 取最小
			if R-i > p[i_mirror] {
				p[i] = p[i_mirror]
			} else {
				p[i] = R - i
			}

		} else {
			// i = R or i > R
			p[i] = 0
		}

		// 中心拓展
		for T[i-p[i]-1] == T[i+p[i]+1] {
			p[i] += 1
		}

		// 记录最大回文长度
		if p[i] > maxLen {
			startCenter = i
			maxLen = p[i]
		}

		// 更新超过当前R长度的 中心点
		if i+p[i] > R {
			C = i
			R = i + p[i]
		}

	}

	startIndex := (startCenter - maxLen) / 2
	return s[startIndex : startIndex+maxLen]
}
