package the_primary_algorithms

// SingleNumber
// 位运算：异或运算 a^b^a = b , 0^a = a, a^a = 0, 满足交换律
func SingleNumber(nums []int) int {
	var mask int
	for _, v := range nums {
		mask ^= v
	}

	return mask
}
