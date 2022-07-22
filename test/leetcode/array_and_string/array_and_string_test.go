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

func TestReverseWords1(t *testing.T) {
	// string := "the sky is blue"
	string := "  hello world  "

	result := array_and_string.ReverseWords1(string)

	t.Logf("result: %v\n", result)
}

func TestReverseWords2(t *testing.T) {
	// string := "the sky is blue"
	string := " Let's take LeetCode contest "

	result := array_and_string.ReverseWords2(string)

	t.Logf("result: %v\n", result)
}

func TestStrStr(t *testing.T) {
	txt := "hello"
	pattern := "ll"

	result := array_and_string.StrStr(txt, pattern)

	t.Logf("txt => %s, result => %v\n", txt, result)
}

func TestReverseString(t *testing.T) {
	str := []byte("hello")
	array_and_string.ReverseString(str)

	t.Logf("reverse :%v\n", string(str))
}

func TestArrayPairSum(t *testing.T) {
	nums := []int{6, 2, 6, 5, 1, 2}

	sum := array_and_string.ArrayPairSum(nums)

	t.Logf("sum :%v\n", sum)
}

func TestTwoSum(t *testing.T) {
	// nums := []int{2, 7, 11, 15}
	// target := 9

	// nums := []int{2, 3, 4}
	// target := 6

	// nums := []int{-1, 0}
	// target := -1

	nums := []int{-1, -1, -1, -1, 1, 1}
	target := 2

	// nums := []int{1, 2, 3, 4, 4, 9, 56, 90}
	// target := 8

	result1 := array_and_string.TwoSum(nums, target)
	result2 := array_and_string.BinarySearchTwoSum(nums, target)

	t.Logf("indexes :%v\n", result1)
	t.Logf("indexes :%v\n", result2)
}

func TestRemoveElement(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2

	result := array_and_string.RemoveElement(nums, val)

	t.Logf("lenght:%v\n nums:%v\n", result, nums)
}

func TestFindMaxConsecutiveOnes(t *testing.T) {
	// nums := []int{1, 0, 1, 1, 0, 1}
	nums := []int{1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 1, 1, 1, 1}

	consecutiveLen := array_and_string.FindMaxConsecutiveOnes(nums)

	t.Logf("lenght:%v\n", consecutiveLen)
}

func TestMinSubArrayLen(t *testing.T) {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7

	// nums := []int{1, 4, 4}
	// target := 4

	// nums := []int{1, 1, 1, 1, 1, 1, 1, 1}
	// target := 11

	minLen := array_and_string.MinSubArrayLen(target, nums)

	t.Logf("min len:%v\n", minLen)
}

func TestGenerate(t *testing.T) {

	numRows := 10

	yangHuiTriangle := array_and_string.Generate(numRows)

	for _, v := range yangHuiTriangle {
		t.Logf("%v\n", v)
	}

}

func TestGetRow(t *testing.T) {

	rowIndex := 10

	result := array_and_string.GetRow(rowIndex)

	t.Logf("%v\n", result)
}

func TestFindMin(t *testing.T) {

	// nums := []int{4, 5, 6, 7, 0, 1, 2}
	nums := []int{10, 11, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := array_and_string.FindMin(nums)

	t.Logf("min :%v\n", result)
}

func TestRemoveDuplicates(t *testing.T) {

	// nums := []int{4, 5, 6, 7, 0, 1, 2}
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 5, 5, 5, 5, 5, 5, 5, 6, 7, 8, 9, 9}

	result := array_and_string.RemoveDuplicates(nums)
	// result := array_and_string.RemoveDuplicatesBinarySearch(nums)

	t.Logf("nums: %v\n length :%v\n", nums[:result], result)
}

func TestMoveZeroes(t *testing.T) {

	nums := []int{0, 1, 0, 3, 12}

	array_and_string.MoveZeroes(nums)

	t.Logf("nums: %v\n", nums)
}
