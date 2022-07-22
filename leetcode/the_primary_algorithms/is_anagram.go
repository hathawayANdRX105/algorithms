package the_primary_algorithms

// 111
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
