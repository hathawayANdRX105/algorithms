package illustrate_for_offer

// 12. 矩阵中的路径
func exist(board [][]byte, word string) bool {
	r, c := len(board), len(board[0])

	// pre-check to speed up
	cMap := make(map[byte]struct{})

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			cMap[board[i][j]] = struct{}{}
		}
	}

	for i := range word {
		if _, ok := cMap[word[i]]; !ok {
			return false
		}
	}

	path := make([][]bool, 0, r)
	for i := 0; i < r; i++ {
		path = append(path, make([]bool, c))
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if board[i][j] == word[0] && findPath(board, word, i, j, 0) {
				return true
			}
		}
	}

	return false
}

func findPath(board [][]byte, word string, i, j, k int) bool {

	if board[i][j] != word[k] {
		return false
	}

	// board 只有大小写字母，用其他ascii符号替换作为不可通过符号
	tmp := board[i][j]
	board[i][j] = ' '
	k++

	defer func() {
		// 结束后回复
		board[i][j] = tmp
	}()

	if k == len(word) {
		return true
	}

	if i < len(board)-1 && findPath(board, word, i+1, j, k) {
		return true
	}

	if j < len(board[0])-1 && findPath(board, word, i, j+1, k) {
		return true
	}

	if i > 0 && findPath(board, word, i-1, j, k) {
		return true
	}

	if j > 0 && findPath(board, word, i, j-1, k) {
		return true
	}

	return false
}

// 13. 机器人的运动范围
func movingCount(m int, n int, k int) int {
	var count int
	path := make([][]bool, 0, m)
	for r := 0; r < m; r++ {
		path = append(path, make([]bool, n))
	}

	isRightStep := func(i, j int) bool {
		// 不在范围或者已经走过的格子直接返回false
		if i >= m || j >= n || path[i][j] {
			return false
		}

		// 当前格子不符合运动范围
		if k < decomposeSum(i)+decomposeSum(j) {
			return false
		}

		path[i][j] = true
		count++
		return true
	}
	dfsMoving(0, 0, isRightStep)

	return count
}

func decomposeSum(i int) int {
	var ans int
	for i > 0 {
		ans += i % 10
		i /= 10
	}
	return ans
}

func dfsMoving(i, j int, isRightStep func(i, j int) bool) {
	if !isRightStep(i, j) {
		return
	}

	dfsMoving(i, j+1, isRightStep)
	dfsMoving(i+1, j, isRightStep)
}

// TreeNode is defined for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 27. 二叉树的镜像
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// 交换左右子树
	root.Left, root.Right = mirrorTree(root.Right), mirrorTree(root.Left)
	return root
}

// 28. 对称的二叉树
// *号标记为基本子问题结构
//
/*	   1
	  / \
	 2   2   *
	/ \ / \  *
   3  4 4  3 *
*/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return dfsForSymmetric(root.Left, root.Right)
}

func dfsForSymmetric(l, r *TreeNode) bool {
	// 两个都为 nil
	if l == nil && r == nil {
		return true
	}

	// 只有一个为nil
	if l == nil || r == nil {
		return false
	}

	if l.Val != r.Val {
		return false
	}

	return dfsForSymmetric(l.Left, r.Right) && dfsForSymmetric(l.Right, r.Left)
}

// 54. 二叉搜索树的第 k 大节点
func kthLargest(root *TreeNode, k int) int {
	var ans int
	set := func(val int) bool {
		if k > 0 {
			ans = val
			k--
		}

		return k < 1
	}

	dfsForK(root, set)
	return ans
}

// 后序遍历，寻找第k大值
func dfsForK(node *TreeNode, set func(int) bool) {
	if node == nil {
		return
	}

	dfsForK(node.Right, set)
	if !set(node.Val) {
		dfsForK(node.Left, set)
	}
}

var depth int

// 55 - I. 二叉树的深度
func maxDepth(root *TreeNode) int {
	depth = 0
	dfsForDepth(root, 0)
	return depth
}

func dfsForDepth(node *TreeNode, k int) {
	if node == nil {
		if k-1 > depth {
			depth = k - 1
		}
		return
	}

	dfsForDepth(node.Left, k+1)
	dfsForDepth(node.Right, k+1)
}
