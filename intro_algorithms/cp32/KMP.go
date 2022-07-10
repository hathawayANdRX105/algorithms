package cp32

import (
	"math"
)

// GetNext 主要是对pattern 对自身进行检索跳跃数组的生成
func GetNext(pattern string) []int {

	// 'next' 数组保存了匹配失败后能立即开始匹配字符的索引信息，节省了非必要的字符匹配
	next := make([]int, len(pattern))

	// j = next[j] 跳跃后作为这次匹配模式串字符的新的起始位置
	// 设置next[j] = -1为匹配模式串第一个字符与当前文本字符不相等的flag
	next[0] = -1

	for i, j := 1, 0; i < len(pattern)-1; {

		// case 1: 当前txt字符与模式串第一个字符不匹配时，j=-1，进入判断重新移位匹配
		// case 2: 连续匹配模式串字符
		if j < 0 || pattern[i] == pattern[j] {
			i++
			j++
			// 记录当前模式串字符能够跳跃的索引
			next[i] = j
		} else {
			// case 3: 当两个字符不匹配时，j索引跳跃
			j = next[j]
		}
	}

	return next
}

// time:  O(m+n)
// space: O(m)
// KMPMatcher txt字符串对pattern模式串的匹配
// 对应cp32.4 Knuth-Morris-Pratt 算法的实现改进
func KMPMatcher(txt, pattern string) string {
	next := GetNext(pattern)

	var t, p int
	maxP := math.MinInt

	for t < len(txt) {

		// case 1: p = -1 代表匹配失败后， 重新匹配模式串, t = t+1 & p = 0
		// case 2: txt[t] = patter[p] 时，进行匹配模式串
		if p < 0 || txt[t] == pattern[p] {
			p++
			t++
		} else {
			// 利用跳跃数据，减少不必要的匹配次数
			p = next[p]
		}

		// 如果 p 刚好为模式串长度，则说明匹配成功
		if p > maxP {
			maxP = p

			if p == len(pattern) {
				return pattern
			}
		}

	}

	return pattern[:maxP]
}
