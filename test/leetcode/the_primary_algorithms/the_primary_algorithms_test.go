package the_primary_algorithms_test

import (
	"algorithms/leetcode/the_primary_algorithms"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}

	maxPrices := the_primary_algorithms.MaxProfit(prices)

	t.Logf("max prices :%v\n", maxPrices)

}

// test rotate1,2
func TestRotate(t *testing.T) {
	// nums := []int{1, 2, 3, 4, 5, 6}
	// k := 4

	// nums := []int{1, 2, 3, 4, 5, 6}
	// k := 7

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27}
	k := 38

	// 约瑟夫环
	// the_primary_algorithms.Rotate1(nums, k)

	// 镜像翻转
	the_primary_algorithms.Rotate2(nums, k)

	t.Logf("rotate array :%v\n", nums)

}

// test rotate3
func TestRotate3(t *testing.T) {
	// matrix := [][]int{
	// 	{1, 2, 3, 4},
	// 	{5, 6, 7, 8},
	// 	{9, 10, 11, 12},
	// 	{13, 14, 15, 16},
	// }

	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	the_primary_algorithms.Rotate3(matrix)

	// print matrix
	for i := 0; i < len(matrix); i++ {
		t.Log(matrix[i])
	}
}

func TestContainsDuplicate(t *testing.T) {
	nums := []int{1, 2, 3, 0, 1}

	// result := the_primary_algorithms.ContainsDuplicate1(nums)
	result := the_primary_algorithms.ContainsDuplicate2(nums)

	t.Logf("rotate array :%v\n", result)

}

func TestSingleNumber(t *testing.T) {
	nums := []int{1, 2, 1, 3, 2}

	result := the_primary_algorithms.SingleNumber(nums)

	t.Logf("result:%v\n", result)
}

func TestIntersect(t *testing.T) {
	// nums1 := []int{1, 2, 1, 3, 2}
	// nums2 := []int{1, 2, 1}

	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}

	result := the_primary_algorithms.Intersect(nums1, nums2)

	t.Logf("result:%v\n", result)
}

func TestPlusOne(t *testing.T) {
	digits := []int{1, 9, 9}
	result := the_primary_algorithms.PlusOne(digits)

	t.Logf("result:%v\n", result)
}

func TestTowSum(t *testing.T) {
	nums := []int{2, 7, 1, 0, 8, 11, 15}
	target := 18

	indexes := the_primary_algorithms.TwoSum(nums, target)

	t.Logf("nums:%v\n target:%v\n indexes:%v", nums, target, indexes)
}

func TestIsValidSudoku(t *testing.T) {
	// ture
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

	// false board[0][0] => 8 == board[3][0] => 8
	// board := [][]byte{
	// 	{'8','3','.','.','7','.','.','.','.'},
	// 	{'6','.','.','1','9','5','.','.','.'} ,
	// 	{'.','9','8','.','.','.','.','6','.'} ,
	// 	{'8','.','.','.','6','.','.','.','3'} ,
	// 	{'4','.','.','8','.','3','.','.','1'} ,
	// 	{'7','.','.','.','2','.','.','.','6'} ,
	// 	{'.','6','.','.','.','.','2','8','.'} ,
	// 	{'.','.','.','4','1','9','.','.','5'} ,
	// 	{'.','.','.','.','8','.','.','7','9'}}

	result := the_primary_algorithms.IsValidSudoku(board)

	// result := '8'

	// result -= '1'
	t.Logf("result:%v", result)
}

func TestReverseInt(t *testing.T) {

	x := 2147483647

	result := the_primary_algorithms.ReverseInt(x)

	t.Logf("reverse int result:%v\n", result)
}

func TestFirstUniqChar(t *testing.T) {

	s := "leetcode"
	// s := "eetcoode"

	index := the_primary_algorithms.FirstUniqChar(s)

	t.Logf("index:%v --- string:%s\n", index, s)
}

func TestIsAnagram(t *testing.T) {
	// s1 := "xaaddy"
	// s2 := "xbbccy"

	s1 := "aacc"
	s2 := "ccac"

	result := the_primary_algorithms.IsAnagram(s1, s2)

	t.Logf("resutl:%v", result)
}

// TestIsPalindrome ...
func TestIsPalindrome(t *testing.T) {

	// s := "A man, a plan, a canal: Panama"
	// s:= "OP"
	s := "0P"

	result := the_primary_algorithms.IsPalindrome(s)

	t.Logf("result:%v", result)
	// t.Log(the_primary_algorithms.IsNumberOrLetter('A'))
	// t.Log(the_primary_algorithms.IsNumberOrLetter('a'))
	// t.Log(the_primary_algorithms.IsNumberOrLetter(9))
	// t.Log(the_primary_algorithms.IsEqual('A', 'a'))
	// t.Log(the_primary_algorithms.IsEqual('A', 'a'))
	// t.Log(the_primary_algorithms.IsEqual('a', 'B'))
}

func TestMyAtoi(t *testing.T) {

	// s := "-2147483648"
	// s := "   -2147483649"
	// s := "   -2147483659"

	// s := "meet with 124"

	// s := ""
	// s := "   "

	// s := "   1"
	// s := "   -1"

	// s := "2147483637"
	// s := "2147483657"

	s := " -1010023630o4"
	ans := the_primary_algorithms.MyAtoi(s)

	t.Logf("result:%v", ans)

}

func TestCountAndSay(t *testing.T) {
	n := 10
	result := the_primary_algorithms.CountAndSay(n)

	t.Logf("result string:%s", result)
}

func TestMergeTwoLists(t *testing.T) {
	l1 := &the_primary_algorithms.ListNode{Val: 1}
	l1.Next = &the_primary_algorithms.ListNode{Val: 1}
	l1.Next.Next = &the_primary_algorithms.ListNode{Val: 2}

	l2 := &the_primary_algorithms.ListNode{Val: 1}
	l2.Next = &the_primary_algorithms.ListNode{Val: 2}
	l2.Next.Next = &the_primary_algorithms.ListNode{Val: 3}

	mergeList := the_primary_algorithms.MergeTwoLists(l1, l2)

	for mergeList != nil {
		t.Logf("%v --- ", mergeList.Val)
	}
}

func TestIsPalindromeLists(t *testing.T) {
	// container := []int{1, 2, 2, 1}
	// container := []int{1, 2, 3, 2, 1}
	// container := []int{1, 1, 0, 1}
	// container := []int{1, 0, 1, 1}
	container := []int{7, 0, 8, 5, 6, 5, 7, 9, 2, 4, 1, 2, 8, 3, 9, 6, 6, 0, 8, 6, 9, 5, 7, 4, 1, 0}
	
	head := &the_primary_algorithms.ListNode{}
	p := head

	for _, v := range container {
		p.Next = &the_primary_algorithms.ListNode{Val: v}
		p = p.Next
	}

	result := the_primary_algorithms.IsPalindromeLists(head.Next)

	t.Logf("result:%v\n", result)
}
