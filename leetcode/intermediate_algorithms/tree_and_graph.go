package intermediate_algorithms

import (
	"container/list"
)

// 以下是 intermediate_alorithms 关于 tree_and_graph 的算法代码实现部分

// Definition for a binary tree node.
/*
 * type TreeNode struct {
 *	Val   int
 *	Left  *TreeNode
 *	Right *TreeNode
 * }
 *
 */

// inorder ...
func inorder(node *TreeNode, appendElement func(val int)) {
	if node == nil {
		return
	}

	inorder(node.Left, appendElement)
	appendElement(node.Val)
	inorder(node.Right, appendElement)
}

// 1.inorderTraversal 二叉树的中序遍历
func inorderTraversal(root *TreeNode) []int {

	ans := []int{}
	appendElement := func(val int) {
		ans = append(ans, val)
	}

	inorder(root, appendElement)

	return ans
}

// 2.zigzagLevelOrder 二叉树的锯齿形层次遍历
// 下面用双向链表，可以利用切片进行优化，内存
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	ans := [][]int{}
	orderList, reverseOrderList := list.New(), list.New()

	orderList.PushFront(root)

	for {
		if orderList.Len() == 0 {
			break
		}

		// 当前层正序遍历
		levelArr := make([]int, orderList.Len())
		for i := 0; i < len(levelArr); i++ {
			// 正序删除
			node := orderList.Remove(orderList.Front()).(*TreeNode)

			levelArr[i] = node.Val

			// 反序插入
			if node.Left != nil {
				reverseOrderList.PushBack(node.Left)
			}

			if node.Right != nil {
				reverseOrderList.PushBack(node.Right)
			}
		}

		ans = append(ans, levelArr)

		if reverseOrderList.Len() == 0 {
			break
		}

		// 下一层进行反序遍历
		levelArr = make([]int, reverseOrderList.Len())
		for i := 0; i < len(levelArr); i++ {
			// 反序删除
			node := reverseOrderList.Remove(reverseOrderList.Back()).(*TreeNode)

			levelArr[i] = node.Val

			// 正序插入
			if node.Right != nil {
				orderList.PushFront(node.Right)
			}

			if node.Left != nil {
				orderList.PushFront(node.Left)
			}
		}

		ans = append(ans, levelArr)
	}

	return ans
}

// 优化：可以先O(n) 遍历一遍中序切片，用map进行存放，构建树时能快速计算
func buildTreeTraverse(preorder []int, inorder []int) *TreeNode {
	if len(inorder) < 1 {
		return nil
	}

	node := TreeNode{Val: preorder[0]}

	// O(n) 查找根节点
	var i int
	for i < len(inorder) && inorder[i] != preorder[0] {
		i++
	}

	node.Left = buildTreeTraverse(preorder[1:i+1], inorder[:i])
	node.Right = buildTreeTraverse(preorder[i+1:], inorder[i+1:])

	return &node
}

// 3.buildTree 从前序与中序遍历序列构建二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	return buildTreeTraverse(preorder, inorder)
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 4.connect 填充每个节点的下一个右侧节点指针
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	// 利用双指针进行串联每层节点，遍历该层的左右子树，并串联
	var p, preHead *Node
	// dummyP 为记录链表的游离头点
	dummyP := &Node{Next: root}

	for dummyP.Next != nil {
		preHead = dummyP.Next
		p = dummyP

		// 重置dummyP 为下一个串联节点的开头， 即preHead的左子节点
		dummyP.Next = preHead.Left
		for preHead != nil {
			if preHead.Left != nil {
				p.Next = preHead.Left
				p = p.Next
			}

			if preHead.Right != nil {
				p.Next = preHead.Right
				p = p.Next
			}

			preHead = preHead.Next
		}
	}

	return root
}

func fillTraverse(node *TreeNode, k int, addElem func(val int) bool) {
	if node == nil {
		return
	}

	fillTraverse(node.Left, k, addElem)

	// 中序遍历二叉搜索树
	if addElem(node.Val) {
		return
	}

	fillTraverse(node.Right, k, addElem)
}

// 5.kthSmallest 寻找二叉搜索树中的第k小的元素
func kthSmallest(root *TreeNode, k int) int {
	// nth 记录第k小(顺序)， ans为第k小值
	var nth, ans int

	addElem := func(val int) bool {
		if nth < k {
			ans = val
			nth++
		}

		return nth == k
	}

	fillTraverse(root, k, addElem)

	return ans
}

// fillLandsWithSea 假设标记岛屿的i，j都是合法的，减少末梢检查判断次数
func fillLandsWithSea(grid [][]byte, i, j int) {
	if grid[i][j] < '1' {
		return
	}

	grid[i][j] = '0'

	if i+1 < len(grid) {
		fillLandsWithSea(grid, i+1, j)
	}

	if j+1 < len(grid[0]) {
		fillLandsWithSea(grid, i, j+1)
	}

	if i > 0 {
		fillLandsWithSea(grid, i-1, j)
	}

	if j > 0 {
		fillLandsWithSea(grid, i, j-1)
	}
}

// 6.numIslands 岛屿数量
func numIslands(grid [][]byte) int {
	var count int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// 如果是‘海’则跳过
			if grid[i][j] != '1' {
				continue
			}

			// DFS:遇到岛屿，则填成‘海’, 并计数
			count++
			fillLandsWithSea(grid, i, j)
		}
	}

	return count
}
