package array_and_string

// 1.PivotIndex1 寻找数组的中心索引
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

// 2.SearchInsert 二分法
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

// partition according the start value in range
func partitionForMerge(arr [][]int, lo, hi int) int {
	cmpValue := arr[lo][0]
	i, j := lo, hi+1

	for {
		for i++; arr[i][0] < cmpValue; i++ {
			if i == hi {
				break
			}
		}

		for j--; arr[j][0] > cmpValue; j-- {
			if j == lo {
				break
			}
		}

		if i >= j {
			break
		}

		//exchange value of index i, j
		arr[i], arr[j] = arr[j], arr[i]
	}

	//exchange cmpValue with value of index j
	arr[lo], arr[j] = arr[j], arr[lo]

	return j
}

func quickSortForMerge(intervals [][]int, lo, hi int) {
	if hi <= lo {
		return
	}

	pivot := partitionForMerge(intervals, lo, hi)
	quickSortForMerge(intervals, lo, pivot-1)
	quickSortForMerge(intervals, pivot+1, hi)
}

// 3.Merge 合并区间
func Merge(intervals [][]int) [][]int {
	size := len(intervals)
	if size < 2 {
		return intervals
	}

	quickSortForMerge(intervals, 0, size-1)

	var uniqueIndex int
	uniqueIntervals := make([][]int, size)
	uniqueIntervals[0] = intervals[0]
	
	for i := 1; i < size; i++ {

		if intervals[i-1][1] >= intervals[i][0] {
			// range merge
			// choose the smaller start range num
			intervals[i][0] = intervals[i-1][0]

			// choose the bigger end range num
			if intervals[i-1][1] > intervals[i][1]{
				intervals[i][1] = intervals[i-1][1]
			}
			
			// update merge range array
			uniqueIntervals[uniqueIndex] = intervals[i]

			// look for the next range whether need to be merge
		} else {
			// stop merge for uniqueIndex and shift next one to keep looking.
			uniqueIndex++
			uniqueIntervals[uniqueIndex] = intervals[i]
		}
	}

	return uniqueIntervals[:uniqueIndex+1]
}


