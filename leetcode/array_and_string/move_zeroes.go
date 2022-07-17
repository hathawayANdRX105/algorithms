package array_and_string

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
