package tree

type BSTNode struct{}

// NewBinarySearchTree[T comparable] ...
func NewBinarySearchTree[V any](cmp func(v1, v2 V) bool) *BinarySearchTree[V] {
	return &BinarySearchTree[V]{BasicTree[BSTNode, V]{cmp: cmp}}
}

type BinarySearchTree[V any] struct {
	BasicTree[BSTNode, V]
}

// Search ...
func (t *BinarySearchTree[V]) Search(elem V) *TreeNode[BSTNode, V] {
	return Search(t.Root, elem, t.cmp)
}

// Insert ...
func (bst *BinarySearchTree[V]) Insert(insertVal V) {
	if bst.Root == nil {
		bst.Root = &TreeNode[BSTNode, V]{Val: insertVal, Time: 1}
		return
	}

	var p *TreeNode[BSTNode, V]
	r := bst.Root

	for r != nil {
		p = r

		if bst.cmp(insertVal, r.Val) {
			r = r.Left
		} else if bst.cmp(r.Val, insertVal) {
			r = r.Right
		} else {
			p.Time++
			return
		}

	}

	node := &TreeNode[BSTNode, V]{Val: insertVal, Time: 1}
	if bst.cmp(insertVal, p.Val) {
		p.Left = node
		return
	}

	p.Right = node
}

// Remove ...
func (bst *BinarySearchTree[V]) remove(node *TreeNode[BSTNode, V], elem V, time int) *TreeNode[BSTNode, V] {
	if node == nil { // case1: not found
		return node
	}

	// case 2 : 继续寻找
	if bst.cmp(elem, node.Val) {
		node.Left = bst.remove(node.Left, elem, time)
	} else if bst.cmp(node.Val, elem) {
		node.Right = bst.remove(node.Right, elem, time)
	} else { // case 3: found the val of node is equal to elem
		if node.Time-time > 0 {
			node.Time -= time
			return node
		}

		// sub case1: node 只有一个子树或无子树
		if node.Left == nil {
			return node.Right
		}

		if node.Right == nil {
			return node.Left
		}

		// sub case2: node 有两个子树，则用后继节点替换删除
		next := Minimum(node.Right)
		node.Val = next.Val
		node.Time = next.Time
		node.Right = bst.remove(node.Right, next.Val, time)

		return node
	}

	return node
}

// Remove ...
func (bst *BinarySearchTree[V]) RemoveOnce(elem V) {
	bst.Root = bst.remove(bst.Root, elem, 1)
}

func (bst *BinarySearchTree[V]) Remove(elem V, time int) {
	if time < 1 {
		return
	}
	bst.Root = bst.remove(bst.Root, elem, time)
}

// IsOrder ...
// func IsOrder[N Node, V any](bst *BinarySearchTree[V]) bool {
// 	var st []*TreeNode[N, V]
// 	st = append(st, bst.Root)

// 	for len(st) > 0 {
// 		p := st[len(st)-1]
// 		st = st[:len(st)-1]

// 		if p.Left != nil {
// 			if bst.cmp(p.Val, p.Left.Val) {
// 				return false
// 			}
// 			st = append(st, p.Left)
// 		}

// 		if p.Right != nil {
// 			if bst.cmp(p.Right.Val, p.Val) {
// 				return false
// 			}

// 			st = append(st, p.Right)
// 		}
// 	}

// 	return true
// }
