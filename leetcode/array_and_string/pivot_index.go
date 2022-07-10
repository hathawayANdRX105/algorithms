package array_and_string

// O(n^2) iteration for nums and compare the sum between left and right
func PivotIndex1(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	for i := 0; i < len(nums); i++ {
		leftSum := sumLeftArray(nums, i)
		rightSum := sumRightArray(nums, i)

		if leftSum == rightSum {
			return i
		}
	}

	return -1
}

func sumLeftArray(nums []int, index int) int {
	if index <= 0 {
		return 0
	}

	var tempSum int
	for i := 0; i < index; i++ {
		tempSum += nums[i]
	}

	return tempSum
}

func sumRightArray(nums []int, index int) int {
	lastIndex := len(nums) - 1
	if index >= lastIndex {
		return 0
	}

	var tempSum int

	for i := lastIndex; i > index; i-- {
		tempSum += nums[i]
	}

	return tempSum
}

// dynamic program
func PivotIndex2(nums []int) int {
	size := len(nums)
	if size < 2 {
		return 0
	}

	leftSum := make([]int, size)
	rightSum := make([]int, size)

	// initilaze the first value of rightSum array
	// sum all the value of nums which index range from 1 to size - 1.
	for i := 1; i < size; i++ {
		rightSum[0] += nums[i]
	}

	if leftSum[0] == rightSum[0] {
		return 0
	}

	for pivotIndex := 1; pivotIndex < size; pivotIndex++ {
		leftSum[pivotIndex] = leftSum[pivotIndex-1] + nums[pivotIndex-1]
		rightSum[pivotIndex] = rightSum[pivotIndex-1] - nums[pivotIndex]

		if leftSum[pivotIndex] == rightSum[pivotIndex] {
			return pivotIndex
		}

	}

	return -1
}

// if ls is equal to rs then 2*ls + p = totalSum
func PivotIndex3(nums []int) int {

	var totalSum int
	var leftSum int
	size := len(nums)

	for i := 0; i < size; i++ {
		totalSum += nums[i]
	}

	for pivot := 0; pivot < size; pivot++ {
		if leftSum*2+nums[pivot] == totalSum {
			return pivot
		} else {
			leftSum += nums[pivot]
		}
	}

	return -1

}
