package tree

type Number interface {
	int | int8 | int16 | int32 | int64
}

func CmpNumber[T Number](v1, v2 T) bool {
	return v1 < v2
}

type Node interface {
	AVLNode | BSTNode
}

// TreeNode 使用组合,来统一不同树使用的树节点，Node用于特殊的树节点需要增加的功能
type TreeNode[N Node, V any] struct {
	Node  N
	Val   V
	Time  int
	Left  *TreeNode[N, V]
	Right *TreeNode[N, V]
}

// BasicTree 通过装饰来实现其他树的特殊处理
type BasicTree[N Node, V any] struct {
	Root *TreeNode[N, V]
	cmp  func(v1, v2 V) bool
}

func Preorder[N Node, V any](node *TreeNode[N, V], do func(node *TreeNode[N, V])) {
	if node == nil {
		return
	}

	do(node)
	Preorder(node.Left, do)
	Preorder(node.Right, do)
}

func Inorder[N Node, V any](node *TreeNode[N, V], do func(node *TreeNode[N, V])) {
	if node == nil {
		return
	}

	Inorder(node.Left, do)
	do(node)
	Inorder(node.Right, do)
}

func Postorder[N Node, V any](node *TreeNode[N, V], do func(node *TreeNode[N, V])) {
	if node == nil {
		return
	}

	Postorder(node.Left, do)
	Postorder(node.Right, do)
	do(node)
}

func Minimum[N Node, V any](node *TreeNode[N, V]) *TreeNode[N, V] {
	if node == nil {
		return nil
	}

	for node.Left != nil {
		node = node.Left
	}

	return node
}

func Maximum[N Node, V any](node *TreeNode[N, V]) *TreeNode[N, V] {
	if node == nil {
		return nil
	}

	for node.Right != nil {
		node = node.Right
	}

	return node
}

// Search[N Node, V any] if elem == node.Val then cmp(elem, node.Val) => false and cmp(node.Val, elem) => false
func Search[N Node, V any](node *TreeNode[N, V], elem V, cmp func(v1, v2 V) bool) *TreeNode[N, V] {
	if node == nil {
		return nil
	}

	if cmp(elem, node.Val) {
		return Search(node.Left, elem, cmp)
	} else if cmp(node.Val, elem) {
		return Search(node.Right, elem, cmp)
	}

	return node
}
