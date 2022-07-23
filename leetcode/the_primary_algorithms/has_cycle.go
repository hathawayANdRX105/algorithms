package the_primary_algorithms

func hasCycle(head *ListNode) bool {

	// s2 = v2*t
	// s1 = v1*t
	// s2 = step * s1, 存在某一个时刻发生相撞，如果存在环
	// 保证fast的跳跃，如果遇上nil说明，为链表,否则有环等待相遇判断
	slow, fast := head, head
	for fast != nil && fast.Next != nil {

		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}

	}

	return false
}
