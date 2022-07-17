package cp04

import (
	"math"
)


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

// NonRecursiveFindMaximumSubArray is another implementation of finding maximum sub array
// just use the loops
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

// FindMaxSubArray 是寻找最大子数组和的线性方法的实现，主要根据三个情况进行分类处理
// time:  O(n)
// space: O(1)
// 对应书本 4.1-5 题
func FindMaxSubArray(a []int) (int, int, int) {

	// preMaxSum 是a[0, i-1] 之间最大的子数组和
	// preBoundaryMaxSum 是包括a[i]下的最大子数组和，可能是a[0, i-1] + a[i] 或者a[i]
	preMaxSum := math.MinInt
	preBoundaryMaxSum := math.MinInt

	var beginI, endI, tempI int
	for i := 0; i < len(a); i++ {

		// 针对a[i] 新加入影响最大子数组和的分类讨论
		if a[i] > preBoundaryMaxSum+a[i] {
			// 如果 a[i] 大于 上一个a[0, i-1] 的最大边界数组和
			// 可能是一个新的最大子数组和的起始，用tempI记录索引，更新最大边界数组和
			tempI = i
			preBoundaryMaxSum = a[i]
		} else {
			// 情况2：a[0, i] 最大边界和为 上一个边界和加上当前a[i]
			// 当前索引可能是最大边界的末尾，直接更新最大边界数组和
			preBoundaryMaxSum += a[i]
		}

		// 如果超过a[0, i-1] 的最大数组和，更新起始和末尾位置以及a[0~i]的最大和
		if preBoundaryMaxSum > preMaxSum {
			preMaxSum = preBoundaryMaxSum
			beginI = tempI
			endI = i
		}
	}

	return beginI, endI, preMaxSum
}
