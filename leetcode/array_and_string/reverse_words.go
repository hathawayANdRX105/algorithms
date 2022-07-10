package array_and_string

// ReverseWords ...
func ReverseWords(s string) string {
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
