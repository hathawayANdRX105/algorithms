package array_and_string

// LongestCommonPrefix ...
func LongestCommonPrefix(strs []string) string {

	result := strs[0]
	var commonIndex int
	for commonIndex < len(strs[0]) {

		// check the single letter whether are queal for each word
		var i int
		for ; i < len(strs); i++ {

			// if current word contains non-letter and the letter of commonIndex isn't equal to previous word letter of commonIndex then stop.
			if commonIndex >= len(strs[i]) || result[commonIndex] != strs[i][commonIndex] {
				break
			}
		}

		// only if iterate all word then commonIndex plus 1.
		if i == len(strs) {
			commonIndex++
		} else {
			break
		}

	}

	return result[:commonIndex]
}
