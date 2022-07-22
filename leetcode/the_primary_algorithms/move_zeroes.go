package the_primary_algorithms

// 在array_and_string 做过
func MoveZeroes(nums []int) {

	var s, f int
	for f < len(nums) {
		if nums[f] != 0 {
			nums[s], nums[f] = nums[f], nums[s]
			s++
		}

		f++
	}

}
