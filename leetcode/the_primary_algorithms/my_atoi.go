package the_primary_algorithms

// MyAtoi ...
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

		// 假设s[r]是数字,判断越界
		if ans > 214748364 || (ans == 214748364 && temp > 7) {
			ans = 1 << 31
			break
		}

		ans = ans*10 + int(temp)
	}

	if isMinus {
		ans *= -1
	} else if ans == 1<<31 {
		ans -= 1
	}

	return ans
}
