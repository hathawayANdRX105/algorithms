package array_and_string

/*
 FindMaxConsecutiveOnes
 param:
   f: 遍历数组，判断当前数是否为断层（值为0），如果为断层，则f代表当前连续的1的末尾
   s: 指向当前连续的1的开头
   maxLen: 记录连续1的最长长度
*/
func FindMaxConsecutiveOnes(nums []int) int {

	var maxLen, f, s int
	for f < len(nums) {
		if nums[f] != 1 {
			if f-s > maxLen {
				maxLen = f - s
			}

			// 重新更新断层位置
			s = f + 1
		}

		f++
	}

	// 检查最后连续长度
	if f-s > maxLen {
		maxLen = f - s
	}

	return maxLen
}
