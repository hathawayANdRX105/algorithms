package illustrate_for_offer

// 03. 数组中重复的数字
func findRepeatNumber(nums []int) int {
	var uMap map[int]struct{}
	for _, v := range nums {
		if _, ok := uMap[v]; ok {
			return v
		}

		uMap[v] = struct{}{}
	}

	return 0
}

// 利用 nums 元素值落在[0, n-1]范围做hash数列，如果找到相同值则返回，并且len(nums) > n-1
func findRepeatNumber2(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}

			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}

	return -1
}

// 48. 最长不含重复字符的子字符串
func lengthOfLongestSubstring(s string) int {
	var l, r, maxLen int
	n := len(s)
	char := make([]bool, (1<<7)-1)
	for ; r < n; r++ {
		if char[s[r]] {
			if r-l > maxLen {
				maxLen = r - l
			}

			for l <= r && s[l] != s[r] {
				char[s[l]] = false
				l++
			}
			l++
			continue
		}

		char[s[r]] = true
	}

	if r-l > maxLen {
		maxLen = r - l
	}

	return maxLen
}

// 50. 第一个只出现一次的字符, s只包含小写字母
func firstUniqChar(s string) byte {
	if len(s) < 1 {
		return ' '
	}

	cc := make([]int, 26)
	for _, v := range s {
		cc[v-'a']++
	}

	for _, v := range s {
		if cc[v-'a'] == 1 {
			return byte(v)
		}
	}

	return ' '
}
