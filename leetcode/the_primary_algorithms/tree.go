package the_primary_algorithms

import "container/list"

// 以下代码是 the_primary_algorithms 关于树的题型实现部分

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 1.maxDepth 寻找二叉树的最大深度, 遍历全部
var depth int

func recursiveAcumulateDepth(node *TreeNode, level int) {
	if node == nil {
		if depth < level {
			depth = level
		}
		return
	}

	recursiveAcumulateDepth(node.Left, level+1)
	recursiveAcumulateDepth(node.Right, level+1)
}

// maxDepth1 针对叶子节点处理，会少很多操作数
func maxDepth1(root *TreeNode) int {
	depth = 0

	recursiveAcumulateDepth(root, 0)
	return depth
}

// maxDepth2 为通俗的遍历后操作
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth2(root.Left)
	right := maxDepth2(root.Right)

	if right > left {
		left = right
	}

	return left
}

// 2.验证当前树是否为一棵二叉搜索树
// 题目对：有效的二叉搜索树是 left.val < root.val && root.val < right.val
var lastestNode *TreeNode

func traveseIsValidBST(node *TreeNode) bool {
	if node == nil {
		return true
	}

	if !traveseIsValidBST(node.Left) {
		return false
	}

	if lastestNode != nil && lastestNode.Val >= node.Val {
		return false
	}
	lastestNode = node

	if !traveseIsValidBST(node.Right) {
		return false
	}

	return true
}

// isValidBST1 设置一个全局节点，记录中序遍历时最近访问的根节点，也就是当前节点的上一个节点（中序遍历视角）
func isValidBST1(root *TreeNode) bool {
	lastestNode = nil
	return traveseIsValidBST(root)
}

// isValidBST2 利用双链表实现栈，实现树的中序遍历，判断有效二叉树的合理，作为练习
func isValidBST2(root *TreeNode) bool {
	// pre 用来记录迭代当前节点的上一个节点, cur 用于遍历节点
	var pre, cur *TreeNode = nil, root
	// stack 用压栈实现树的中序遍历
	stack := list.New()

	for cur != nil || stack.Len() > 0 {
		// 将当前节点的左节点全部压栈
		for cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		}

		// 对当前节点进行中序遍历
		// 已经将目前所有的根节点压入栈，没有左节点，可能存在右节点
		// 对当前判断后，需要将右子树的左节点压入栈（参考中序遍历的顺序）
		if stack.Len() > 0 {

			cur = stack.Remove(stack.Back()).(*TreeNode)

			// 判断 有效二叉树节点val的有序性
			if pre != nil && !(pre.Val < cur.Val) {
				return false
			}

			// 迭代，记录上一个节点以及迭代右子树，进行中序遍历
			pre, cur = cur, cur.Right
		}
	}

	return true
}

func recursiveSymmetric(leftNode, rightNode *TreeNode) bool {
	if leftNode == nil && rightNode == nil {
		return true
	}

	// case1: leftNode = nil && rightNode != nil
	// case2: leftNode != nil && rightNode = nil
	if leftNode == nil || rightNode == nil {
		return false
	}

	// leftNode != nil && rightNode != nil
	// 先深度遍历到叶子结点判断，再判断左右根节点的值，这样会快些
	// 针对叶子节点的判断会比根节点的值相等来的快
	if recursiveSymmetric(leftNode.Left, rightNode.Right) && recursiveSymmetric(leftNode.Right, rightNode.Left) {
		if leftNode.Val == rightNode.Val {
			return true
		}
	}

	// case3: leftNode != nil && rightNode != nil && leftNode.Val != rightNode.Val
	// case4: recursiveSymmetric(leftNode.Left, rightNode.Right) = false
	// case5: recursiveSymmetric(leftNode.Right, rightNode.Left) = false
	return false
}

// 3.对称二叉树
func isSymmetric(root *TreeNode) bool {
	// 题目中有说明节点数 [1, 1000]，所以必存在根节点
	return recursiveSymmetric(root.Left, root.Right)
}

// recursiveLevelStorage ...
func recursiveLevelStorage(node *TreeNode, level int, order [][]int) [][]int {
	if node == nil {
		return order
	}

	if level+1 > len(order) {
		order = append(order, []int{})
	}
	order = recursiveLevelStorage(node.Left, level+1, order)
	order = recursiveLevelStorage(node.Right, level+1, order)

	order[level-1] = append(order[level-1], node.Val)

	return order
}

// 4.二叉树的层序遍历
// 递归拼接数组方式
func levelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	order := [][]int{}
	order = recursiveLevelStorage(root, 1, order)

	return order
}

// levelOrder2 先序遍历，处理每一层的节点，利用链表实现
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	ans := [][]int{}
	queue := list.New()
	queue.PushFront(root)

	// i代表每一层
	for i := 0; queue.Len() > 0; i++ {
		size := queue.Len()
		ans = append(ans, make([]int, size))

		// j代表每一层的个数
		for j := 0; j < size; j++ {
			cur := queue.Remove(queue.Front()).(*TreeNode)
			ans[i][j] = cur.Val

			if cur.Left != nil {
				queue.PushBack(cur.Left)
			}

			if cur.Right != nil {
				queue.PushBack(cur.Right)
			}
		}
	}

	return ans
}

// 5.sortedArrayToBST 给定一个升序数组，将其转换为高度平衡的二叉树
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	// 根据题目要求需要:
	//   当数组长度为奇数时，取中位数
	//   当数组长度为偶数时，取长度一半值后一位
	mid := (len(nums)-1)>>1 + (len(nums)+1)%2
	
	node := &TreeNode{Val: nums[mid]}
	node.Left = sortedArrayToBST(nums[:mid])
	node.Right = sortedArrayToBST(nums[mid+1:])

	return node
}
