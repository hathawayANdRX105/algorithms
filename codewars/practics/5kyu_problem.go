package codewarsproblem

// getAnagramsMask is the function that calculate for mask which remove pairs of letter and total ascii char
// so use mask and total can find the unique pattern for a word
// a ^ a = 0 , a ^ 0 = a then a ^ b ^ a = b
func getAnagramsMask(word string) (rune, rune) {
	var mask, total rune
	for _, s := range word {
		mask ^= s
		total += s
	}

	return mask, total
}

// Where my anagrams at?
// https://www.codewars.com/kata/523a86aa4230ebb5420001e1
func Anagrams(word string, words []string) []string {
	mask, total := getAnagramsMask(word)

	// save mem by use original words array
	var l, r int // words[0:l] maintain the anagrams string
	for r < len(words) {
		mask2, total2 := getAnagramsMask(words[r])
		if mask == mask2 && total == total2 {
			words[l] = words[r]
			l++
		}

		r++
	}

	return words[:l]
}

type PosPeaks struct {
	Pos   []int
	Peaks []int
}
// Pick peaks
// https://www.codewars.com/kata/5279f6fe5ab7f447890006a7
// 寻找不包含两侧边界的顶峰, 由一个数组组成
func PickPeaks(array []int) PosPeaks {

	var pos []int
	var peaks []int

	var p int
	for i := 0; i+1 < len(array); i++ {
		if array[i] < array[i+1] { // only when the element incresing then update p
			p = i + 1
		} else if p > 0 && array[i] > array[i+1] {
			// only when p was set and next element decreses
			pos = append(pos, p)
			peaks = append(peaks, array[p])
			p = 0 // reset for next candidate
		}
	}

	return PosPeaks{pos, peaks}
}

// Numbe of trailing zeros of N!
// https://www.codewars.com/kata/52f787eb172a8b4ae1000a34
// 找公因数 2 x 5的组合， 由于 5远比2个数少，需要找多少个5
func Zeros(n int) int {
	var count int
	for n > 0 {
		n /= 5
		count += n
	}

	return count
}

// ISBN-10 Validation
// https://www.codewars.com/kata/51fc12de24a9d8cb0e000001
// ISBN-10 前九位是0-9， 最后一位为 0-9，或者x代表10
func ValidISBN10(isbn string) bool {
	var cal, num int
	end := len(isbn) - 1
	for i := 0; i < end; i++ {
		num = int(isbn[i] - '0')
		if num > 9 { // the first nine number is not 0-9
			return false
		}

		cal += (i + 1) * num
	}

	// deal with the last number
	if isbn[end] > '9' {
		// ascii(x) = 120
		// turn last char into lower case and subtract to 110
		num = int(isbn[end] | 0x20 - 110)
	} else {
		num = int(isbn[end] - '0')
	}
	cal += 10 * num

	return cal%11 == 0 // implement proper solution here
}

// Valid Parentheses 匹配括号
// https://www.codewars.com/kata/52774a314c2333f0a7000688
// 旧方法： 栈匹配左括号
func ValidParentheses(parens string) bool {
	if len(parens)%2 != 0 {
		return false
	}

	st := make([]byte, 0, len(parens))
	for i := 0; i < len(parens); i++ {
		if parens[i] == ')' {
			if len(st) < 1 { // if stack's empty then matching left parenthese is failure
				return false
			}

			st = st[:len(st)-1]
			continue
		}

		st = append(st, '(')
	}

	return len(st) == 0 // check if stack is empty
}

func ValidParentheses2(parens string) bool {
	// the number of left parensthese should be equal to right parensthese
	if len(parens)%2 != 0 {
		return false
	}

	var match int
	for i := 0; i < len(parens); i++ {
		// when it encounter with '('
		if parens[i] == '(' {
			match++
			continue
		}

		// when it encounter with ')'
		match--
		if match < 0 {
			return false
		}
	}

	return match == 0
}

// printMap ...
func printMap(m [][]int) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			print(m[i][j], " ")
		}

		println()
	}
}

