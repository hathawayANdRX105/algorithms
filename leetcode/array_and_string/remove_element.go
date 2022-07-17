package array_and_string

func RemoveElement(nums []int, val int) int {

	// 'fast' 代表快指针，遍历一遍数组，检查当前索引位置是否为val
	// 'slow' 代表慢指针，代表填入下一个不为val的索引位置, 同时也等于不含val的数组大小
	var fast, slow int
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	return slow
}
