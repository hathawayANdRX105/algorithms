package intro_to_algorithms_test


import (
	"algorithms/intro_algorithms/cp04"
	"testing"
)

func TestFindMaximumSubArray(t *testing.T) {

	arr := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	// arr := []int{1, -3, 20, 47, -80, 13, -3, -25, 20, -13, 33, -30, -16, -23, 18, 20, -7, 12, 39, 10, -20, 1, -5, -22, 15, -4, 7, 39, 12, -22, -1, 14, -23, -2, 33, 19, -18, 10, -33, -21, -8, 22}
	// arr := []int{ -20, -33, -16, -5, -1, -10, -48 }
	length := len(arr) - 1

	// low, high, sum := cp04.ForceFindMaximumSubArray(arr, 0, length)
	low, high, sum := cp04.FindMaximumSubArray(arr, 0, length)
	low2, high2, sum2 := cp04.NonRecursiveFindMaximumSubArray(arr, 0, length)

	t.Logf("Recursive    low:%d, high:%d, sum:%d", low, high, sum)
	t.Logf("NonRecursive low:%d, high:%d, sum:%d", low2, high2, sum2)

}

func TestFindMaxSubArray(t *testing.T) {

	// arr := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	// arr := []int{1, -3, 20, 47, -80, 13, -3, -25, 20, -13, 33, -30, -16, -23, 18, 20, -7, 12, 39, 10, -20, 1, -5, -22, 15, -4, 7, 39, 12, -22, -1, 14, -23, -2, 33, 19, -18, 10, -33, -21, -8, 22}
	arr := []int{ -20, -33, -16, -5, -1, -10, -48 }
	length := len(arr) - 1
	begin, end, sum := cp04.FindMaxSubArray(arr)
	low, high, sum2 := cp04.FindMaximumSubArray(arr, 0, length)

	t.Logf("Recursive    low:%d, high:%d, sum:%d", low, high, sum)
	t.Logf("Linear       low:%d, high:%d, sum:%d", begin, end, sum2)

}


