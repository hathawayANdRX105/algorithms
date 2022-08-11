package intermediate_algorithms

// 以下是 intermediate_algorithms 关于 linked_list 的代码实现部分

// 1.addTowNumbers 两数相加 进位问题
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

// 原先版本，不加dummyNode
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	head := l1     // 记录头节点
	var futher int // 记录进位

	// 处理双链表下一个节点都不为空，合并，并且处理进位问题
	for l1.Next != nil && l2.Next != nil {

		l1.Val += l2.Val + futher

		futher = l1.Val / 10
		l1.Val %= 10

		l1 = l1.Next
		l2 = l2.Next
	}

	// 防止 l1 , l2 刚好同长度情况未添加最后一个节点情况
	if l1 != nil && l2 != nil {
		l1.Val += l2.Val + futher

		futher = l1.Val / 10
		l1.Val %= 10
	}

	// l1 为 nil 时， l2 可能还有节点
	if l1.Next == nil {
		l1.Next = l2.Next
		l2.Next = nil // 游离，交给gc处理
	}

	// 处理后续进位
	for 0 < futher && l1.Next != nil {
		l1.Next.Val += futher

		futher = l1.Next.Val / 10
		l1.Next.Val %= 10

		l1 = l1.Next
	}

	// 对于最后一个进位问题，需要创建新的节点保存
	if futher > 0 {
		l1.Next = &ListNode{Val: futher}
	}

	return head
}

// 简化版： 双链表前面添加dummyNode 方便最后一个进位处理，但是平均速度下降
func simpleAddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var futher int
	var head *ListNode

	// 添加dummyNode
	head = &ListNode{Next: l2}
	l2 = head

	head = &ListNode{Next: l1}
	l1 = head

	for l1.Next != nil && l2.Next != nil {

		l1.Next.Val += l2.Next.Val + futher

		futher = l1.Next.Val / 10
		l1.Next.Val %= 10

		l1 = l1.Next
		l2 = l2.Next
	}

	// l1 为 nil 时， l2 可能还有节点
	// 如果 l1 不为 nil，则说明 l2 到尾
	if l1.Next == nil {
		// 接着 l2 后续节点, 为后面处理进位做准备
		l1.Next = l2.Next
		l2.Next = nil // 游离，交给gc处理
	}

	// 处理后续进位
	for 0 < futher {
		// 最后进位处理
		if l1.Next == nil {
			l1.Next = &ListNode{Val: 1}
			break
		}

		l1.Next.Val += futher

		futher = l1.Next.Val / 10
		l1.Next.Val %= 10

		l1 = l1.Next
	}

	return head.Next
}

// 2.oddEvenList 奇偶链表
func oddEvenList(head *ListNode) *ListNode {
	// quick case: 当链表长度小于2时，链表不变
	if head == nil || head.Next == nil {

		return head
	}

	// head 指向奇链表开头， head2 指向偶链表开头
	// p1 指向奇链表末尾， p2 指向偶链表末尾
	p1, p2 := head, head.Next
	head2 := p2

	// 当没有 候选奇数节点，或 已经更新最后一个奇数节点，没有偶数节点时退出
	for p2 != nil && p2.Next != nil {
		// 跨一个节点连接
		p1.Next = p1.Next.Next
		p2.Next = p2.Next.Next

		// 更新末尾节点
		p1, p2 = p1.Next, p2.Next
	}

	// 连接两个链表
	p1.Next = head2

	return head
}

// 3.getIntersectionNode 相交链表

/*
 * 遇到关于双链表的题可以尝试步距计算角度切入

 * 双指针遍历， p1 从链表A开始， p2 从链表B开始
 * 到各自链表尾时，p1 重新从链表B开始， p2 重新从链表A开始
 * 假设链表A，链表B长度各自为 LenA + sameLen, LenB + sameLen
 * case 1:当两指针相遇时
 *        则 totalS1 =  LenA + sameLen + LenB, totalS2 = LenB + sameLen + LenA
 *
 * case 2:两指针最终不相遇，相当于遍历 m + n
 *
 */
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB

	for p1 != p2 {

		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next

		}

		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next

		}

	}

	return p1
}

// time:  O(m + n)
// space: O(m)
func getIntersectionNode2(headA, headB *ListNode) *ListNode {

	// set 集合
	unique := make(map[*ListNode]struct{})

	for headA != nil {
		unique[headA] = struct{}{}

		headA = headA.Next
	}

	for headB != nil {

		if _, ok := unique[headB]; ok {
			return headB
		}

		headB = headB.Next
	}

	return nil
}
