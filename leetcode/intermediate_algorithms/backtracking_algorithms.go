package intermediate_algorithms

// 以下是 intermediate_algorithms 关于 backtracking_algorithms 的代码实现

var digitAlphabet [][]byte = [][]byte{
	{'a', 'b', 'c'},
	{'d', 'e', 'f'},
	{'g', 'h', 'i'},
	{'j', 'k', 'l'},
	{'m', 'n', 'o'},
	{'p', 'q', 'r', 's'},
	{'t', 'u', 'v'},
	{'w', 'x', 'y', 'z'},
}

// buildStringCombination 通过路径生成不同的组合，在末尾处返回结果
func buildStringCombination(digits string, subString string, ans []string) []string {
	if len(digits) < 1 {
		return append(ans, subString)
	}

	// fmt.Println(digits, subString)

	for _, alpha := range digitAlphabet[digits[0]-'2'] {
		ans = buildStringCombination(digits[1:], subString+string(alpha), ans)
	}

	return ans
}

// 1.letterCombinations 电话号码的字母组合
func letterCombinations(digits string) []string {

	var ans []string
	if len(digits) < 1 {
		return ans
	} else if len(digits) < 6 {
		ans = make([]string, 0, 3*len(digits))
	} else {
		ans = make([]string, 0, 15+(len(digits)-5)*4)
	}

	ans = buildStringCombination(digits, "", ans)

	return ans
}

// buildParenthesisCombination
// l, r  : 分别代表了 左右括号剩余使用的数量
// subStr: 是到当前函数递归路线的组成字符串结果
// add   : 传参函数，ans数组在内存堆中，进行添加最终结果
func buildParenthesisCombination(l, r int, subStr string, addAns func(endStr string)) {
	if l == 0 && r == 0 {
		addAns(subStr)
	}

	// 如果左括号还能继续使用
	if l != 0 {
		buildParenthesisCombination(l-1, r, subStr+"(", addAns)
	}

	// 保证 右括号总匹配左括号
	if l < r {
		buildParenthesisCombination(l, r-1, subStr+")", addAns)
	}
}

// 2.generateParenthesis 括号生成
func generateParenthesis(n int) []string {

	// ans 由于不知道准确大小，构建在堆中
	ans := []string{}

	// 通过传参，利用指针为ans添加元素, 减少函数栈的大小内存使用
	addAns := func(endStr string) {
		ans = append(ans, endStr)
	}

	// 提前进入第一个括号的构建，比 buildParenthesisCombination(n, n, "", add) 快一点
	buildParenthesisCombination(n-1, n, "(", addAns)

	return ans
}

// buildOrder ...
// 原理跟 buildParenthesisCombination 类似，让不确定大小的二维数组在内存堆中构建，然后添加
func buildOrder(subNums []int, subAns []int, addAns func(subAns []int)) {
	if len(subNums) < 1 {
		addAns(subAns)
	}

	for i := 0; i < len(subNums); i++ {
		// 将选择的元素交换到 第一索引的位置，保证传递的subNums 不包含选择的元素
		subNums[0], subNums[i] = subNums[i], subNums[0]

		buildOrder(subNums[1:], append(subAns, subNums[0]), addAns)

		// 保证其他需要选择元素的顺序，需要交换回来
		subNums[0], subNums[i] = subNums[i], subNums[0]
	}

}

// 3.permute 全排列
func permute(nums []int) [][]int {
	// ans 由于不确定大小，构建在内存堆中
	ans := [][]int{}
	addAns := func(subAns []int) {
		ans = append(ans, subAns)
	}

	buildOrder(nums, []int{}, addAns)

	return ans
}

// 4.subsets 求子集
// 利用好回溯算法的每个递归函数
func subsets(nums []int) [][]int {
	ans := [][]int{{}}

	for i := 0; i < len(nums); i++ {
		// 由于需要添加从零遍历，并且添加 ans， 不能使用len(ans)作为停止条件
		stopSize := len(ans)

		for j := 0; j < stopSize; j++ {
			// 分配内存
			subSet := make([]int, 0, len(ans[j])+1)

			// 复制 每前 n 个一维数组的元素
			subSet = append(subSet, ans[j]...)

			// 组合当前 选择排列的元素
			subSet = append(subSet, nums[i])

			ans = append(ans, subSet)
		}

	}

	return ans
}

// matchLetterPath ...
func matchLetterPath(retainWord string, board [][]byte, i, j int) bool {
	if len(retainWord) < 1 {
		return true
	}

	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != retainWord[0] {
		// cond 1 : 0 <= i < row && 0 <= j < col
		// cond 2 : 当前点遍历过
		// cond 3 : 当前字符不相等
		return false
	}

	// 替换字符，以作遍历过的marker
	passedMarker := board[i][j]
	board[i][j] = 0

	defer func() {
		// 恢复原来 board字符
		board[i][j] = passedMarker
	}()

	// 寻找下一个字符
	return matchLetterPath(retainWord[1:], board, i, j+1) || matchLetterPath(retainWord[1:], board, i+1, j) || matchLetterPath(retainWord[1:], board, i, j-1) || matchLetterPath(retainWord[1:], board, i-1, j)
}

// 5.exist 单词搜索
func exist(board [][]byte, word string) bool {

	// 补充：可以提前遍历，利用map做 unique set 判断 work单词都存在

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {

			if board[i][j] == word[0] {
				passedMarker := board[i][j]
				board[i][j] = 0
				if matchLetterPath(word[1:], board, i, j+1) || matchLetterPath(word[1:], board, i+1, j) || matchLetterPath(word[1:], board, i, j-1) || matchLetterPath(word[1:], board, i-1, j) {
					return true
				}

				board[i][j] = passedMarker
			}
		}

	}

	return false
}
