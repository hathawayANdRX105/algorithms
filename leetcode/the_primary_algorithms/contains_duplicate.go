package the_primary_algorithms

func ContainsDuplicate1(nums []int) bool {

	size := len(nums)
	table := make([]int, size)

	for _, v := range nums {
		mod := v % size

		if mod < 0 {
			mod += size
		}

		if v == 0 {
			// 针对值为0，做特殊情况处理
			v = 10e9 + 1
		}

		for mod < size && table[mod] != 0 {
			if table[mod] == v {
				return true
			}

			mod = (mod + 1) % size
		}

		table[mod] = v
	}

	return false
}


// 使用map
func ContainsDuplicate2(nums []int) bool {

	uniqueMap := make(map[int]struct{}, len(nums))

	for _, v := range nums {

		if _, ok := uniqueMap[v]; ok {
			return true
		}

		uniqueMap[v] = struct{}{}

	}

	return false
}
