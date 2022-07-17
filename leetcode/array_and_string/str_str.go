package array_and_string

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

// strStr ...
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
