package tree

type AVLNode struct {
	Height int
}

func NewAVLTree[V any](cmp func(v1, v2 V) bool) *AVLTree[V] {

	return &AVLTree[V]{BasicTree[AVLNode, V]{cmp: cmp}}
}

type AVLTree[V any] struct {
	BasicTree[AVLNode, V]
}

func UpdateHeight[V any](node *TreeNode[AVLNode, V]) {
	if node == nil {
		return
	}

	var lh, rh int
	if node.Left != nil {
		lh = node.Left.Node.Height
	}

	if node.Right != nil {
		lh = node.Right.Node.Height
	}

	if rh > lh {
		lh = rh
	}

	node.Node.Height = lh + 1
}

func BalanceFactor[V any](node *TreeNode[AVLNode, V]) int {
	var lh, rh int
	if node.Left != nil {
		lh = node.Left.Node.Height
	}

	if node.Right != nil {
		rh = node.Right.Node.Height
	}

	return lh - rh
}

func LeftRotation[V any](node *TreeNode[AVLNode, V]) *TreeNode[AVLNode, V] {
	pivot := node.Right
	node.Right, pivot.Left = pivot.Left, node

	UpdateHeight(node)
	UpdateHeight(pivot)

	return pivot
}

func RightRotation[V any](node *TreeNode[AVLNode, V]) *TreeNode[AVLNode, V] {
	pivot := node.Left
	node.Left, pivot.Right = pivot.Right, node

	UpdateHeight(node)
	UpdateHeight(pivot)

	return pivot
}

// BalanceLeftTree  如果当前节点出现左子树不平衡, 则进行左子树平衡，判断插入节点是左子树的左边还是右边进行不同的处理
func BalanceLeftTree[V any](node *TreeNode[AVLNode, V], isLeft bool) *TreeNode[AVLNode, V] {
	if isLeft {
		node = RightRotation(node)
	} else {
		node.Left = LeftRotation(node.Left)
		node = RightRotation(node)
	}

	return node
}

// BalanceRightTree  如果当前节点出现右子树不平衡, 则进行右子树平衡，判断插入节点是右子树的左边还是右边进行不同的处理
func BalanceRightTree[V any](node *TreeNode[AVLNode, V], isRight bool) *TreeNode[AVLNode, V] {
	if isRight {
		node = LeftRotation(node)
	} else {
		node.Left = RightRotation(node.Left)
		node = LeftRotation(node)
	}

	return node
}

func (t *AVLTree[V]) insert(node *TreeNode[AVLNode, V], elem V, time int) *TreeNode[AVLNode, V] {
	if node == nil {
		return &TreeNode[AVLNode, V]{Val: elem, Time: time, Node: AVLNode{Height: 1}}
	}

	if t.cmp(elem, node.Val) {
		// 向左找
		node.Left = t.insert(node.Left, elem, time)

		if BalanceFactor(node) > 1 {
			node = BalanceLeftTree(node, t.cmp(elem, node.Left.Val))
		}
	} else if t.cmp(node.Val, elem) {
		// 向右找
		node.Right = t.insert(node.Right, elem, time)

		// 如果当前节点不平衡
		if BalanceFactor(node) < -1 {
			// 判断elem插入在右子树的方向
			node = BalanceRightTree(node, t.cmp(node.Right.Val, elem))
		}
	} else { // elem is equal to node.Val
		node.Time += time
		return node
	}

	UpdateHeight(node)
	return node
}

func (t *AVLTree[V]) InsertOnce(elem V) {
	t.Root = t.insert(t.Root, elem, 1)
}

// Insert ...
func (t *AVLTree[V]) Insert(elem V, time int) {
	t.Root = t.insert(t.Root, elem, time)
}

// remove ...
func (t *AVLTree[V]) remove(node *TreeNode[AVLNode, V], elem V, time int) *TreeNode[AVLNode, V] {
	if node == nil {
		return nil
	}

	if t.cmp(elem, node.Val) {
		node.Left = t.remove(node.Left, elem, time)
		UpdateHeight(node.Left)
	} else if t.cmp(node.Val, elem) {
		node.Right = t.remove(node.Right, elem, time)
		UpdateHeight(node.Right)
	} else {
		// sub-case 1:抵消次数,无需删除节点
		if node.Time-time > 0 {
			node.Time -= time
			return node
		}

		// sub-case 2: 删除节点是叶子节点
		if node.Left == nil && node.Right == nil {
			return nil
		}

		// sub-case 3: 删除节点是非叶子节点, 并且两个子树都存在,则用子树更高的前后继节点替换删除节点,并且删除替换节点
		if node.Left != nil && node.Right != nil {
			if node.Left.Node.Height > node.Right.Node.Height {
				pre := Maximum(node.Left)
				node.Val, node.Time = pre.Val, pre.Time

				node.Left = t.remove(node.Left, pre.Val, time)
				UpdateHeight(node.Left)

			} else {
				post := Minimum(node.Right)
				node.Val, node.Time = post.Val, post.Time

				node.Right = t.remove(node.Right, post.Val, time)
				UpdateHeight(node.Right)
			}
		} else {
			// sub-case 4: node.Left == nil or node.Right == nil
			// 将单子树节点保存到当前父节点, 因为不会出现不平衡结构,所以只有一个子树情况下，该子树只有一个节点
			child := node.Left
			if node.Right != nil {
				child = node.Right
			}

			node.Val, node.Time, node.Node.Height = child.Val, child.Time, 1
			node.Left, node.Right = nil, nil
		}

		// 删除节点之后，返回
		return node
	}

	// 更新高度，维护avltree
	factor := BalanceFactor(node)
	if factor > 1 {
		node = BalanceLeftTree(node, BalanceFactor(node.Left) >= 0)
	} else if factor < -1 {
		node = BalanceRightTree(node, BalanceFactor(node.Right) <= 0)
	}

	return node
}

func (t *AVLTree[V]) RemoveOnce(elem V) {
	t.Root = t.remove(t.Root, elem, 1)
}

func (t *AVLTree[V]) Remove(elem V, time int) {
	t.Root = t.remove(t.Root, elem, time)
}

func (t *AVLTree[V]) Search(elem V) *TreeNode[AVLNode, V] {
	return Search(t.Root, elem, t.cmp)
}

// IsAVLTree ...
func IsAVLTree[V any](t *AVLTree[V]) bool {
	var st []*TreeNode[AVLNode, V]
	st = append(st, t.Root)

	for len(st) > 0 {
		node := st[len(st)-1]
		st = st[:len(st)-1]

		factor := BalanceFactor(node)
		if factor > 1 || factor < -1 {
			return false
		}

		if node.Left != nil {
			st = append(st, node.Left)
		}

		if node.Right != nil {
			st = append(st, node.Right)
		}

	}

	return true
}
