package array_and_string

func SearchInsert(nums []int, target int) int {
	startIndex := 0
	endIndex := len(nums) - 1
	var mid int

	for startIndex <= endIndex {
		mid = startIndex + (endIndex-startIndex)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			startIndex = mid + 1
		} else {
			endIndex = mid - 1
		}
	}

	return startIndex
}
