package the_primary_algorithms

// TwoSum 解决无序数组的两数之和的寻找
func TwoSum(nums []int, target int) []int {
	if len(nums) < 3 {
		return []int{0, 1}
	}

	var indexes []int
	record := make(map[int]int)

	for index2, value := range nums {

		// 利用map存放差值-索引, 差值为target-value, 寻找迭代过程中可能的目标值
		// case1: 没找到，保存记录
		// case2: 找到，包装成数组返回
		if index1, ok := record[value]; !ok {
			// 存储后续迭代寻找的另一个值
			record[target-value] = index2
		} else {
			indexes = []int{index1, index2}
			break
		}
	}

	return indexes
}
