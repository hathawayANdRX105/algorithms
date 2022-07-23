package the_primary_algorithms

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var newHead, nextHead *ListNode

	for head != nil {
		// 保存旧链表的下一个
		nextHead = head.Next

		// 修改当前的旧链表的头节点指向新链表的头
		head.Next = newHead

		// 更新新旧链表的头节点
		head, newHead = nextHead, head
	}

	return newHead
}
