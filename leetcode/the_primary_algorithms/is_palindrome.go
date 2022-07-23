package the_primary_algorithms

import (
	"fmt"
)

func isNumberOrLetter(char byte) bool {
	// ascii:
	//  数字： 48~57
	//  大写字母： 65~90
	//  小写字母： 97~122
	if char-48 < 10 || char-65 < 26 || char-97 < 26 {
		return true
	}

	return false
}

func isEqual(char1 byte, char2 byte) bool {
	if char1 == char2 || (char1 > 64 && char2 > 64) && (char1+32 == char2 || char2+32 == char1) {
		//case1: char1 = char2 , 可能是数字对数字，大写字母对大写字母，小写字母对小写字母
		//case2: 大写字母与小写字母的转换判断， ascii 平移32
		return true
	}

	return false
}

// IsPalindrome 检查字符串是否为回文字符串
func IsPalindrome(s string) bool {

	l, r := 0, len(s)-1
	for l < r {

		// 如果当前不为number or letter，就寻找下一个字符
		for l < r && !isNumberOrLetter(s[l]) {
			l++
		}

		for l < r && !isNumberOrLetter(s[r]) {
			r--
		}

		//case1: 由于左边界优先找到判断字符，如果左边界提前碰壁右边界，减少运算符操作，提前退出，平均优化4ms
		//case2: 发生右边界碰壁左边界同理
		//case3: 字符不相等，字符串的数字英文字符不是回文串
		if r < l || !isEqual(s[l], s[r]) {
			return false
		}

		l++
		r--
	}

	return true
}

// IsPalindromeLists 针对链表检查是否为回文链表
func IsPalindromeLists(head *ListNode) bool {
	var hash1, hash2 int
	x, multiple := 2, 1

	// 按照x进制，换算组合，从前往后的正序与反序，参考10进制
	// [1, 2, 2, 1] 10进制为 (正序)1221 (反序)1221
	//               4进制为 (正序)105  (反序)105
	for head != nil {
		hash1 = hash1*x + head.Val
		hash2 = hash2 + head.Val*multiple

		multiple *= x

		head = head.Next
	}

	fmt.Println(hash1, hash2)
	return hash1 == hash2
}
