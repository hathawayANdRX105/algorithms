package cp04

import (
	"math"
)

func FindMaxCrossingSubArray(arr []int, low, mid, high int) (int, int, int) {
	// find maximum sub arr from mid to low
	sum := 0
	leftSum, leftIndex := math.MinInt, mid

	for i := mid; low <= i; i-- {
		sum += arr[i]

		if sum > leftSum {
			leftSum = sum
			leftIndex = i
		}
	}

	// find maximum sub arr from mid to high
	sum = 0
	rightSum, rightIndex := math.MinInt, mid

	for i := mid + 1; i <= high; i++ {
		sum += arr[i]

		if sum > rightSum {
			rightSum = sum
			rightIndex = i
		}
	}

	// return maximum cross sub arr
	return leftIndex, rightIndex, leftSum + rightSum
}

// FindMaximumSubArray recursive
func FindMaximumSubArray(arr []int, low, high int) (int, int, int) {
	if low == high {
		return low, high, arr[low]
	}

	mid := (high-low)/2 + low

	// divide tow-sub arr to find maximum sub arr between low-mid and mid+1-high
	leftLow, leftHigh, leftSum := FindMaximumSubArray(arr, low, mid)
	rightLow, rightHigh, rightSum := FindMaximumSubArray(arr, mid+1, high)

	crossLow, crossHigh, crossSum := FindMaxCrossingSubArray(arr, low, mid, high)

	// compare three sub arr and return the maximum sub arr
	if rightSum > leftSum {
		leftLow, leftHigh, leftSum = rightLow, rightHigh, rightSum
	}

	if crossSum > leftSum {
		leftLow, leftHigh, leftSum = crossLow, crossHigh, crossSum
	}

	return leftLow, leftHigh, leftSum
}

// NonRecursiveFindMaximumSubArray is another implementation of find maximum sub array
// just use the loop
func NonRecursiveFindMaximumSubArray(arr []int, low, high int) (int, int, int) {
	subArrLow, subArrHigh, subArrSum := low, low, arr[low]

	for i := low + 1; i <= high; i++ {
		var curSum int

		for j := i; j >= 0; j-- {
			curSum += arr[j]

			if curSum > subArrSum {
				subArrLow = j
				subArrHigh = i
				subArrSum = curSum
			}
		}
	}

	return subArrLow, subArrHigh, subArrSum
}

//ForceFindMaximumSubArray is written for homework
func ForceFindMaximumSubArray(arr []int, low, high int) (int, int, int) {
	if low == high {
		return low, high, arr[low]
	}

	left, right, sum := low, low, arr[low]

	for i := low; i <= high; i++ {
		tempSum := arr[i]

		for j := i + 1; j <= high; j++ {
			tempSum += arr[j]

			if tempSum > sum {
				left = i
				right = j
				sum = tempSum
			}
		}
	}
	return left, right, sum
}
