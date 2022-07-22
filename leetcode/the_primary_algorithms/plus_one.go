package the_primary_algorithms

func PlusOne(digits []int) []int {

	for i := len(digits) - 1; 0 <= i; i-- {
		if (digits[i]+1)/10 == 1 {
			//能进位
			digits[i] = 0

			// 进位后缺少空间，申请新的数组
			if i-1 < 0 {
				newDigits := make([]int, len(digits)+1)
				newDigits[0] = 1
				return newDigits
			}
		} else {
			digits[i] += 1
			break
		}
	}

	return digits
}
