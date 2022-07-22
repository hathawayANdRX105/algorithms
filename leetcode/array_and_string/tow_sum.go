package array_and_string

//time:  O(n)
//space: O(1)
//前提:   numbers是有序数组
//TwoSum
func TwoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	if len(numbers) > 2 {

		for l < r {
			if numbers[l]+numbers[r] < target {
				// l + r < target, l右移 使和变大
				l++
			} else if numbers[l]+numbers[r] > target {
				// l + r > target , 使和变小
				r--
			} else {
				// l + r = target
				break
			}
		}
	}

	numbers[0] = l + 1
	numbers[1] = r + 1
	return numbers[:2]
}

//time:  O(lg(n))
//space: O(1)
//前提:   numbers是有序数组
//BinarySearchTwoSum 在TowSum 原基础上，通过二分查询来快速节省不必要的匹配
func BinarySearchTwoSum(numbers []int, target int) []int {

	l, r := 0, len(numbers)-1
	if len(numbers) > 2 {

		for l < r {
			m := l + (r-l)>>1
			if numbers[l]+numbers[m] > target {
				// case 1: 当 l + m > target时, r取中位数, 使和变小
				r = m
			} else if numbers[m]+numbers[r] < target {
				// case 2: 当 m + r < target时, l取中位数, 使和变大
				l = m
			} else if numbers[l]+numbers[r] == target {
				// case 3: l + r = target , 结束匹配

				break
			} else {
				if numbers[l]+numbers[r] < target {
					// case 4: l + r < target, l自增, 使和变大
					l++
				} else {
					// case 5: l + r > target, r自减， 使和变小
					r--
				}
			}
		}
	}

	numbers[0] = l + 1
	numbers[1] = r + 1
	return numbers[:2]
}
