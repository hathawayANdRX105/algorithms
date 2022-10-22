package illustrate_for_offer

import "encoding/pem"

// 18. 删除链表的节点
// ListNode is used to defined for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 用dummyHead
func deleteNode(head *ListNode, val int) *ListNode {
	dummyH := &ListNode{Val: val + 1, Next: head}
	head = dummyH
	for head.Next != nil && head.Next.Val != val {
		head = head.Next
	}
	head.Next = head.Next.Next

	return dummyH.Next
}

// 检查头部再遍历
func deleteNode2(head *ListNode, val int) *ListNode {
	if head != nil && head.Val == val {
		return head.Next
	}

	h := head
	for head.Next != nil && head.Next.Val != val {
		head = head.Next
	}
	head.Next = head.Next.Next
	return h
}

// 22. 链表中倒数第 k 个节点
func getKthFromEnd(head *ListNode, k int) *ListNode {
	p := head
	for p != nil && k > 0 {
		p = p.Next
		k--
	}

	// 防止k过大，链表不足，直接返回
	if p == nil {
		return head
	}

	for p.Next != nil {
		head, p = head.Next, p.Next
	}

	return head.Next
}

// 24. 反转链表
func reverseList(head *ListNode) *ListNode {
	var tail *ListNode
	for head != nil {
		t := head
		head = head.Next

		t.Next = tail
		tail = t
	}

	return tail
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	p := l3
	for l1 != nil && l2 != nil {
		var t *ListNode

		if l2.Val < l1.Val {
			t = l2
			l2 = l2.Next
		} else {
			t = l1
			l1 = l1.Next
		}
		p.Next = t
		p = t
	}

	if l1 != nil {
		p.Next = l1
	}

	if l2 != nil {
		p.Next = l2
	}

	return l3.Next
}

// 52. 两个链表的第一个公共节点
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var ka, kb int
	var p *ListNode
	p = headA
	for p != nil {
		p = p.Next
		ka++
	}

	p = headB
	for p != nil {
		p = p.Next
		kb++
	}

	if kb > ka {
		headA, headB = headB, headA
		ka, kb = kb, ka
	}

	// 利用差距同步相同链表
	k := ka - kb
	for k > 0 {
		k--
		headA = headA.Next
	}

	for headA != headB {
		headA, headB = headA.Next, headB.Next
	}

	return headA
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	pa, pb := headA, headB
	for pa != pb {
		if pa != nil {
			pa = pa.Next
		} else {
			pa = headB
		}

		if pb != nil {
			pb = pb.Next
		} else {
			pb = headA
		}
	}

	return pa
}
