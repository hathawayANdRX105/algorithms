package illustrate_for_offer

// buildTree 07. 重建二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) < 1 {
		return nil
	}

	node := TreeNode{Val: preorder[0]}
	var i int
	for preorder[0] != inorder[i] {
		i++
	}

	// 确认好范围，slice开闭区间 slice[s:e)
	// preorder[node, L, R]
	// inorder[L, node, R]
	// postorder[L, R, node]
	node.Left = buildTree(preorder[1:i+1], inorder[:i])
	node.Right = buildTree(preorder[i+1:], inorder[i+1:])

	return &node
}

// isSubStructure 26. 树的子结构
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	// 开局 B为空则不属于A的子结构
	if B == nil {
		return false
	}

	return dfs(A, B)
}

func dfs(node, target *TreeNode) bool {
	// dfs到底没找到
	if node == nil {
		return false
	}

	// 找到根节点相同，判断是否属于子结构，如果不是则继续遍历左右子树寻找
	if node.Val == target.Val && isSameNode(node.Left, target.Left) && isSameNode(node.Right, target.Right) {
		return true
	}

	return dfs(node.Left, target) || dfs(node.Right, target)
}

func isSameNode(a, b *TreeNode) bool {
	// b的节点数量少于或等于a
	if b == nil {
		return true
	}

	// 前提：b != nil
	// 如果 a没有当前节点，或者 a,b节点值不相等则说明b不是a的子结构
	if a == nil || a.Val != b.Val {
		return false
	}

	return isSameNode(a.Left, b.Left) && isSameNode(a.Right, b.Right)
}


// sumNums 64. 求1+2+…+n 求和公式
func sumNums(n int) int {
	return (n + 1) * n >> 1
}