package the_primary_algorithms

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	// 删除当前节点，转移成删除下一个节点
	// 当前节点保存下一个节点的值,删除下一个节点
	// 等待gc分析无法引用node
	node.Val = node.Next.Val
	node.Next = node.Next.Next

	// 删除游离节点，结束内存
	// node.Val = node.Next.Val
	// temp := node.Next
	// node.Next = node.Next.Next
	// temp.Next = nil
}
