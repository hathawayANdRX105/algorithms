package array_and_string

// ReverseWords1 翻转字符串里的单词,保持单词顺序不变
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

// ReverseWords2 反转字符串中的单词 III, 保持单词顺序，翻转每个单词的字符
func ReverseWords2(s string) string {

	var l, r int
	strByteArr := []byte(s)

	// 单词翻转, word[i, j] 闭区间
	reverseSingle := func(b []byte, i, j int) {
		for i < j {
			b[i], b[j] = b[j], b[i]
			i++
			j--
		}
	}

	for ; r < len(s); r++ {
		if strByteArr[r] == ' ' {
			reverseSingle(strByteArr, l, r-1)

			l = r + 1
		}
	}

	// 题目说明字符串开头结尾不含空格，直接对最后一个单词翻转
	reverseSingle(strByteArr, l, r-1)

	// 考虑最后可能有空格结尾，则说明最后一个单词还没翻转
	// if strByteArr[r-1] != ' ' {
	// 	reverseSingle(strByteArr, l, r-1)
	// }

	return string(strByteArr)
}
