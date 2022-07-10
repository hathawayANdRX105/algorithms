package array_and_string_test

import (
	"algorithms/leetcode/array_and_string"
	"testing"
)

func TestPivotIndex1(t *testing.T) {
	// nums := []int {1, 7, 3, 6, 5, 6}
	// nums := []int {1, 2, 3}
	nums := []int{2, 1, -1}

	result := array_and_string.PivotIndex1(nums)

	t.Logf("result: %v\n", result)
}

// dynamic program
func TestPivotIndex2(t *testing.T) {
	// nums := []int {1, 7, 3, 6, 5, 6}
	// nums := []int {1, 2, 3}
	// nums := []int {2, 1, -1}
	nums := []int{-1, -1, 0, 1, 1, 0}

	result := array_and_string.PivotIndex2(nums)

	t.Logf("result: %v\n", result)
}

// 2*ls + p = totalSum
func TestPivotIndex3(t *testing.T) {
	// nums := []int {1, 7, 3, 6, 5, 6}
	// nums := []int {1, 2, 3}
	// nums := []int {2, 1, -1}
	nums := []int{-1, -1, 0, 1, 1, 0}

	
	result := array_and_string.PivotIndex3(nums)

	t.Logf("result: %v\n", result)
}

func TestSearchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 2

	result := array_and_string.SearchInsert(nums, target)

	t.Logf("result: %v\n", result)

}

func TestMerge(t *testing.T) {
	// nums := [][]int{{1, 3}, {8, 10}, {15, 18},  {2, 6}}
	nums := [][]int{{1, 4}, {5, 6}}

	result := array_and_string.Merge(nums)

	t.Logf("result: %v\n", result)
}

func TestRotate(t *testing.T) {
	nums := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	array_and_string.Rotate(nums)
	result := nums

	t.Logf("result: %v\n", result)
}

func TestSetZeroes(t *testing.T) {
	// nums := [][]int{{1, 2, 3}, {4, 0, 6}, {7, 8, 9}}
	nums := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}

	array_and_string.SetZeroes(nums)
	result := nums

	t.Logf("result: %v\n", result)
}

func TestFindDiagonalOrder(t *testing.T) {
	nums := [][]int{{1, 2}, {3, 4}}
	// nums := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// nums := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}

	result := array_and_string.FindDiagonalOrder(nums)

	t.Logf("result: %v\n", result)
}

func TestLongestCommonPrefix(t *testing.T) {
	// strs := []string{"flower", "flow", "flight"}
	strs := []string{"dog", "racecar", "car"}

	result := array_and_string.LongestCommonPrefix(strs)

	t.Logf("result: %v\n", result)
}

func TestLongestPalindrodme(t *testing.T) {
	// string := "cbbd"
	// string := "babad"
	string := "aacabdkacaa"

	result := array_and_string.LongestPalindrome3(string)
	t.Logf("result: %v\n", result)
}

func TestLongestPalindromeWithManaCher(t *testing.T) {
	// string := "cbbd"
	// string := "babad"
	string := "aacabdkacaa"

	result := array_and_string.LongestPalindromeWithManacher(string)


	t.Logf("result: %v\n", result)
}

func TestReverseWords(t *testing.T) {
	// string := "the sky is blue"
	string := "  hello world  "

	result := array_and_string.ReverseWords(string)


	t.Logf("result: %v\n", result)
}

