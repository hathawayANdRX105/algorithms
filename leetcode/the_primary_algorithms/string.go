package the_primary_algorithms

import "bytes"

// 以下代码是关于 the_primary_algorithms 字符串的实现部分

// 1.reverseString 反转字符串 已经在 array_and_string 中实现
func reverseString(s []byte) {
	l, r := 0, len(s)-1

	for l < r {
		s[l], s[r] = s[r], s[l]

		l++
		r--
	}
}


// 2.ReverseInt 整数反转
func ReverseInt(x int) int {

	var p int
	stack := make([]int, 10)
	maxInt := 1<<31 - 1
	minInt := -1 << 31

	// 取个位数压栈
	for x != 0 {
		stack[p] = x % 10
		x /= 10
		p++
	}

	for multiple := 1; 0 < p; p-- {
		if stack[p-1]*multiple > maxInt-x || stack[p-1]*multiple < minInt-x {
			// 防止越界，保持增量范围合理
			// case1: x为正数， 正增量应小于 maxInt - x
			// case2: x为负数， 负增量应大于 minInt - x
			return 0
		}

		// 没有越界
		x += stack[p-1] * multiple
		multiple *= 10
	}

	return x
}

// 3.FirstUniqChar 寻找字符串中第一个唯一的字符
func FirstUniqChar(s string) int {

	// 字母数组， 前5位数记录索引，后5位数记录出现次数
	dict := make([]int, 26)

	for i, v := range s {
		// 题目限制s长度[1, 1e5]
		// 1e5 后记录出现次数，1e5前记录最近的出现索引
		dict[v-'a'] = (dict[v-'a']/1e5+1)*1e5 + i
	}

	first := len(s)
	for i := 0; i < len(dict); i++ {
		// 检查出现次数为1的字符，并且索引越小就更新索引记录
		if dict[i]/1e5 == 1 && dict[i]%1e5 < first {
			first = dict[i] % 1e5
		}
	}

	// 初始化处理
	if first == len(s) {
		return -1
	}

	return first
}

// 4.IsAnagram 针对t对s异位字符是否有效 **
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make([]rune, 26)

	for _, v := range s {
		count[v-'a']++
	}

	for _, v := range t {
		count[v-'a']--

		if count[v-'a'] < 0 {
			return false
		}
	}

	return true
}

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

// 5.IsPalindrome 检查字符串是否为回文字符串
// s含有其他字符，不单止英文（大小写），数字
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

// 6.MyAtoi 将字符串转换为整数 ，如果前面为英文则默认为0
func MyAtoi(s string) int {

	var l, ans int
	// 清除空格
	for l < len(s) && s[l] == ' ' {
		l++
	}

	// 如果越界，则转换失败
	if l >= len(s) {
		return 0
	}

	// 判断正负，如果字符开头有英文则转换失败
	var isMinus bool
	if s[l] == '-' {
		isMinus = true
		l++
	} else if s[l] == '+' {
		l++
	} else if s[l]-'0' < 0 || s[l]-'0' > 9 {
		// 清除空格后，如果开头不是’-‘， ’+‘， ’1-9‘,则转换失败
		return 0
	}

	for r := l; r < len(s); r++ {
		temp := s[r] - '0'
		if temp < 0 || temp > 9 {
			// 如果不是数字，将跳出
			break
		}

		// 假设s[r]是数字,判断越界，能进行以下处理判断，说明s[r]是即将十进制进位的最后一位
		// case1： 如果当前ans已经大于214748364，说明超过 2^31 不满足正负数最大值，跳出处理
		// case2: 如果当前ans为214748364 对即将加入的最后一位进行大小判断，
		//   sub case1: 如果ans为正数，temp <= 7 为正常情况，否则以最大值跳出循环，根据后续符号减一处理
		//   sub case2: 如果ans为负数，temp <= 8 为正常情况，但是temp=8时，可以立即退出，再作符号处理返回结果即可
		if ans > 214748364 || (ans == 214748364 && temp > 7) {
			ans = 1 << 31
			break
		}

		ans = ans*10 + int(temp)
	}

	// 根据正负来做细节处理
	if isMinus {
		ans *= -1
	} else if ans == 1<<31 {
		ans -= 1
	}

	return ans
}

// 7.strStr kmp算法 在 array_and_string 的 string 以及 intro_algorithms 第三十二章的kmp算法 已经实现

// 8.CountAndSay 外观数列的实现
func CountAndSay(n int) string {
	var buffer bytes.Buffer
	component := "1"
	for i := 1; i < n; i++ {

		// 快慢指针遍历上一个外观数列
		var l, r int
		for r < len(component) {
			if component[r] != component[l] {

				// 写入统计数(r-l)，以及统计字符component[l]
				buffer.WriteByte(byte(r-l) + '0')
				buffer.WriteByte(component[l])

				l = r
			}
			r++
		}
		// 写入最后的字符统计以及字符
		buffer.WriteByte(byte(r-l) + '0')
		buffer.WriteByte(component[l])

		// 更新外观数列，以及复用buffer 字节数组
		component = buffer.String()
		buffer.Reset()
	}

	return component
}

// 9.longestCommonPrefix 已经在 array_and_string 中的 string 实现
// 寻找字符串数组中的最长的公共前缀
