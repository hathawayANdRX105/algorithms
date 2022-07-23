package the_primary_algorithms

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func recusiveRemove(node *ListNode, target int) int {
	if node == nil {
		return 0
	}

	nthFromEnd := recusiveRemove(node.Next, target)

	if nthFromEnd == target+1 {
		node.Val = node.Next.Val
		node.Next = node.Next.Next
	}

	return nthFromEnd + 1
}

// recusive remove method
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	point := &ListNode{Next: head}

	recusiveRemove(point, n)
	defer func() {
		point.Next = nil
	}()
	return point.Next
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}

	var l, r *ListNode
	l, r = dummyHead, head

	// 1 <= n <= size
	// 使用两个节点，先让右节点偏移n个单位，然后双节点一起遍历，知道右节点到达终点，左节点的Next就是要删除的节点
	for n > 0 {
		r = r.Next
		n--
	}

	for r != nil {
		l, r = l.Next, r.Next
	}

	l.Next = l.Next.Next

	return dummyHead.Next
}
