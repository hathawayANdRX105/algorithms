package array_and_string

import "fmt"

// 1.LongestCommonPrefix 寻找字符串数组中每个字符串的最长公共前缀
func LongestCommonPrefix(strs []string) string {

	result := strs[0]
	var commonIndex int
	for commonIndex < len(strs[0]) {

		// check the single letter whether are queal for each word
		var i int
		for ; i < len(strs); i++ {

			// if current word contains non-letter and the letter of commonIndex isn't equal to previous word letter of commonIndex then stop.
			// case1: 保证第一个字符串的遍历指针在其他字符串的遍历范围内
			// case2: 统计其他字符串在相同位置是否为同一个字符
			if commonIndex >= len(strs[i]) || result[commonIndex] != strs[i][commonIndex] {
				break
			}
		}

		// only if iterate all word then commonIndex plus 1.
		// 判断统计
		if i == len(strs) {
			commonIndex++
		} else {
			break
		}

	}

	return result[:commonIndex]
}


// 2.LongestPalindrome 最长回文字串
//  包含动态规划，动态规划内存优化，马拉夫算法实现
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

		// 中心拓展, 比对
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

// 3.ReverseWords1 翻转字符串里的单词,保持单词顺序不变
func ReverseWords1(s string) string {
	stack := make([]string, len(s), len(s))

	var size int
	var i, j int

	for {
		// find next word where strat
		// pass by white space ' ' and find the index of word start
		for i = j; i < len(s) && s[i] == ' '; {
			i++
		}

		// find out the index of word end
		for j = i; j < len(s) && s[j] != ' '; {
			j++
		}

		if i < j {
			// push stack
			stack[size] = s[i:j]
			size++

		}

		if j >= len(s) {
			break
		}

	}

	// // add last word
	// stack[size] = s[i:j]
	// size++

	result := ""
	// pop word and concat
	for i := size - 1; 0 <= i; i-- {
		result += stack[i] + " "
	}

	return result[:len(result)-1]
}



// buildNext ...
func buildNext(pattern string) []int {
	next := make([]int, len(pattern))
	next[0] = -1

	for i, j := 1, 0; i < len(pattern)-1; {
		if j < 0 || pattern[i] == pattern[j] {
			i++
			j++

			next[i] = j
		} else {
			j = next[j]
		}
	}

	return next
}

// 4.strStr 参考 intro_algorithms 第三十二章的kmp实现, 内有详细代码注释
func StrStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	next := buildNext(needle)

	var t, p int
	for t < len(haystack) {
		if p < 0 || haystack[t] == needle[p] {
			p++
			t++
		} else {
			p = next[p]
		}

		if p == len(needle) {
			return t - p
		}
	}

	return -1
}
