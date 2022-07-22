package the_primary_algorithms

// 利用map做长度短的数组的唯一集，利用快慢指针对长度长的数组做并集判断以及存放数
func Intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	uniqueSet := make(map[int]int, len(nums1))

	// 对出现的数字做统计
	for _, v := range nums1 {
		uniqueSet[v] += 1
	}

	// 节约内存，对nums2快慢指针来保存并集结果
	var l, r int
	for r < len(nums2) {
		if count, ok := uniqueSet[nums2[r]]; ok && count > 0 {
			uniqueSet[nums2[r]] -= 1
			nums2[l] = nums2[r]
			l++
		}
		r++
	}

	return nums2[:l]
}
