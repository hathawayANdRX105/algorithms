package array_and_string

import "algorithms/pkg/sort"

func ArrayPairSum(nums []int) int {
	sort.QuickSortBy3Way(nums)

	var sum int
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}

	return sum
}
