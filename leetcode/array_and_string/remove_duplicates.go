package array_and_string

func RemoveDuplicates(nums []int) int {
	f, s := 1, 0
	for f < len(nums) {
		if nums[f] != nums[s] {
			s++
			nums[s] = nums[f]
		}

		f++
	}

	return s + 1
}


func RemoveDuplicatesBinarySearch(nums []int) int {
	if len(nums) < 2 {
		return 1
	}

	f, s := 1, 0
	for {
		if nums[f] != nums[s] {
			s++
			nums[s] = nums[f]

		}

		if f+1 == len(nums) {
			break
		}

		// 存在重复的数值，二分查询
		l, r := s, len(nums)-1

		for l < r {
			if nums[s] < nums[l+(r-l)>>1] {
				// 中位数大于重复的数值
				r = l + (r-l)>>1
			} else {
				// 中位数小于等于重复的数值
				l = l + 1 + (r-l)>>1
			}

		}

		f = l
	}

	return s + 1
}
