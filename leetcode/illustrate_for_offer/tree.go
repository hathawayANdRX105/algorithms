package illustrate_for_offer

// lowestCommonAncestor 68 - I. 二叉搜索树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var path1, path2 []*TreeNode

	preorder(root, p, func(node *TreeNode) {
		path1 = append(path1, node)
	})

	preorder(root, q, func(node *TreeNode) {
		path2 = append(path2, node)
	})

	var i int
	n := len(path1)
	if len(path2) < n {
		n = len(path2)
	}

	for i < n && path1[i] == path2[i] {
		i++
	}

	return path1[i-1]
}

func preorder(node, target *TreeNode, do func(*TreeNode)) {
	if node == nil {
		return
	}

	do(node)
	if target.Val == node.Val {
		return
	}

	if node.Val < target.Val {
		preorder(node.Right, target, do)
	} else {
		preorder(node.Left, target, do)
	}
}

// optimizedLowestCommonAncestor 遇到p, q节点需要分开探索的节点就是最近的公共节点
func optimizedLowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// 递归
	v := root.Val
	if v > p.Val && v > q.Val {
		return optimizedLowestCommonAncestor1(root.Left, p, q)
	}

	if v < p.Val && v < q.Val {
		return optimizedLowestCommonAncestor1(root.Right, p, q)
	}

	return root
}

func optimizedLowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	// 迭代
	pp, qq := p.Val, q.Val
	for root != nil {
		v := root.Val
		if v > pp && v > qq {
			root = root.Left
			continue
		}

		if v < pp && v < qq {
			root = root.Right
			continue
		}

		return root
	}

	return nil
}

// lowestCommonAncestor2 68 - II. 二叉树的最近公共祖先 ，非二叉搜索树
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	var path1, path2 []*TreeNode
	preorder2(root, p.Val, func(node *TreeNode) {
		path1 = append(path1, node)
	})
	preorder2(root, q.Val, func(node *TreeNode) {
		path2 = append(path2, node)
	})

	var closest *TreeNode
	for l1, l2 := len(path1)-1, len(path2)-1; l1 > -1 && l2 > -1 && path1[l1].Val == path2[l2].Val; l1, l2 = l1-1, l2-1 {
		closest = path1[l1]
		path1 = path1[:l1]
		path2 = path2[:l2]
	}

	return closest
}

func preorder2(node *TreeNode, target int, do func(*TreeNode)) bool {
	if node == nil {
		return false
	}

	if node.Val == target {
		do(node)
		return true
	}

	if preorder2(node.Left, target, do) {
		do(node)
		return true
	}

	if preorder2(node.Right, target, do) {
		do(node)
		return true
	}

	return false
}
