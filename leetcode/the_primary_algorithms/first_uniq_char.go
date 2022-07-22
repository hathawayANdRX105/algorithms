package the_primary_algorithms

func FirstUniqChar(s string) int {

	// 字母数组， 前5位数记录索引，后5位数记录出现次数
	dict := make([]int, 26)

	for i, v := range s {
		dict[v-'a'] = (dict[v-'a']/1e5+1)*1e5 + i
	}

	first := len(s)
	for i := 0; i < len(dict); i++ {
		if dict[i]/1e5 == 1 && dict[i]%1e5 < first {
			first = dict[i] % 1e5
		}
	}

	if first == len(s) {
		return -1
	}

	return first
}
