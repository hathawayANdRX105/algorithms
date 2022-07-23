package the_primary_algorithms

// 以下代码是关于 the_primary_algorithms 链表的实现部分
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

// 1.deleteNode 删除当前节点
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

// 2.removeNthFromEnd1 删除倒数nth的节点
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

// 3.reverseList 翻转链表
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

// 4.MergeTwoLists 合并两个有序链表
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	// 重构一条新的有序链表
	head := &ListNode{}
	p := head
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			p.Next = list1
			list1 = list1.Next
		} else {
			p.Next = list2
			list2 = list2.Next
		}

		p = p.Next
	}

	if list1 == nil {
		p.Next = list2
	} else {
		p.Next = list1
	}

	return head.Next
}

// 5.IsPalindromeLists 针对链表检查是否为回文链表
func IsPalindromeLists(head *ListNode) bool {
	var hash1, hash2 int
	x, multiple := 2, 1

	// 按照x进制，换算组合，从前往后的正序与反序，参考10进制
	// [1, 2, 2, 1] 10进制为 (正序)1221 (反序)1221
	//               4进制为 (正序)105  (反序)105
	for head != nil {
		hash1 = hash1*x + head.Val
		hash2 = hash2 + head.Val*multiple

		multiple *= x

		head = head.Next
	}

	return hash1 == hash2
}

// 6.hasCycle检查链表是否存在闭环
func hasCycle(head *ListNode) bool {
	// 减少平均使用内存
	if head == nil {
		return false
	}

	// s2 = v2*t
	// s1 = v1*t
	// s2 = step * s1, 存在某一个时刻发生相撞，如果存在环
	// 保证fast的跳跃，如果遇上nil说明，为链表,否则有环等待相遇判断
	jump := head
	for jump != nil && jump.Next != nil {

		head = head.Next
		jump = jump.Next.Next

		if head == jump {
			return true
		}
	}

	return false
}
