package codewarsproblem

import (
	"strings"
)

// String Tops
// https://www.codewars.com/kata/59b7571bbf10a48c75000070
func Tops(msg string) string {
	n := len(msg)
	if n < 1 {
		return ""
	}

	i, step := 1, 0
	var ans []byte
	for i < n {
		ans = append(ans, msg[i])

		step += 2
		i += 1 + (step << 1)
	}

	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}

	return string(ans)
}

// Counting Duplicates 寻找字符中重复的字母
// https://www.codewars.com/kata/54bf1c2cd5b56cc47f0007a1
func duplicate_count(s1 string) int {
	//your code goes here
	unique := make(map[rune]int)

	for _, v := range s1 {
		_, ok := unique[v|0x20]
		if !ok {
			unique[v|0x20] = 1
			continue
		}

		unique[v|0x20]++
	}

	var ans int
	for _, v := range unique {
		if v > 1 {
			ans++
		}
	}

	return ans //Instead of returning '1', you have to return integer 'i' as answer of solution.
}

// CamelCase Method 骆驼峰转换
// https://www.codewars.com/kata/587731fda577b3d1b0001196
// time: O(n)
func CamelCase(s string) string {
	if len(s) < 1 {
		return s
	}

	bb := []byte(strings.TrimRight(s, " "))

	// tow pointer fix it
	var l, r int
	for ; r < len(bb); l, r = l+1, r+1 {
		if bb[r] != ' ' {
			bb[l] = bb[r]
			continue
		}

		// change ascii letter into upper case only when bb[r] encounter ' ' space
		r++
		bb[l] = bb[r] | 0x20 ^ 0x20 // to upper case
	}

	bb[0] = bb[0] | 0x20 ^ 0x20
	return string(bb[:l])
}

// flipWord
// time: O(lgm), which m is the length of word
func flipWord(word []byte, l, r int) {
	// center shrinkage
	for l < r {
		word[l], word[r] = word[r], word[l]
		l++
		r--
	}
}

// Stop gninnipS My sdroW! 翻转句子中长度超过5的单词
// https://www.codewars.com/kata/5264d2b162488dc400000001
// time: O(n + lgn)
func SpinWords(str string) string {
	n := len(str)
	sentence := []byte(str)

	// l points to the first letter of every word
	// r encounters ' ' will check the length of current word is greater than 5, if so then flip it
	var l, r int
	for ; r < n; r++ {
		if sentence[r] != ' ' {
			continue
		}

		if r-l > 4 { // if the length of word is greater than 5, then flip it
			flipWord(sentence, l, r-1)
		}
		l = r + 1 // l jumps to the first letter of next word
	}

	// in case the last word  if it need to be flipped
	if r-l > 4 {
		flipWord(sentence, l, r-1)
	}

	return string(sentence)
} // SpinWords

// Find The Parity Outlier  寻找数组中的奇偶异常值
// https://www.codewars.com/kata/5526fc09a1bbd946250002dc
func FindOutlier(integers []int) int {
	// ex   find  lastBitSum
	// 000 -> 1   0
	// 001 -> 1   1
	// 110 -> 0   2
	// 111 -> 0   3
	lastBit := 1
	if (integers[0]&1)+(integers[1]&1)+(integers[2]&1) > 1 {
		lastBit = 0
	}

	for i := 0; i < len(integers); i++ {
		if integers[i]&1 == lastBit {
			lastBit = i
			break
		}
	}

	return integers[lastBit]
}

// Tortoise racing 龟兔赛跑
// https://www.codewars.com/kata/55e2adece53b4cdcb900006c
// v1, v2为速度， g为相距距离，都大于0
func Race(v1, v2, g int) [3]int {
	if v2 < v1 { // B will never catch A
		return [3]int{-1, -1, -1}
	}

	time := float64(g / (v2 - v1))

	// 转换为 hour， minutes ， second
	return [3]int{int(time), int(time*60) % 60, int(time*3600) % 60}
}

// Build Tower 构建字符串金字塔
// https://www.codewars.com/kata/576757b1df89ecf5bd00073b
func TowerBuilder(nFloors int) []string {
	floorLen := 1 + (nFloors-1)<<1 // the length of every floor
	var ans []string
	for nFloors > 0 {
		nFloors-- // nth Floor was made of n-1 space and floorLen - (n-1)*2 star and n-1 space
		floorByte := make([]byte, 0, floorLen)

		// append ' '
		for i := 0; i < nFloors; i++ {
			floorByte = append(floorByte, ' ')
		}

		// append '*'
		for s := floorLen - nFloors<<1; s > 0; s-- {
			floorByte = append(floorByte, '*')
		}

		// append the same space floorByte[0:nFloors)
		floorByte = append(floorByte, floorByte[:nFloors]...)

		ans = append(ans, string(floorByte))
	}

	return ans
}
