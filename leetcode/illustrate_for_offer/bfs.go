package illustrate_for_offer

import "container/list"

// 32 - I. 从上到下打印二叉树
// levelOrder1 从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印
func levelOrder1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ans []int
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		size := queue.Len()
		for size > 0 {
			size--
			pop := queue.Remove(queue.Front()).(*TreeNode)
			ans = append(ans, pop.Val)

			if pop.Left != nil {
				queue.PushBack(pop.Left)
			}

			if pop.Right != nil {
				queue.PushBack(pop.Right)
			}
		}
	}

	return ans
}

// 32 - II. 从上到下打印二叉树 II 打印二叉树的每层子树值
// levelOrder2 使用container/list实现队列
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var ans [][]int
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		size := queue.Len()
		level := make([]int, 0, size)
		for size > 0 {
			size--
			pop := queue.Remove(queue.Front()).(*TreeNode)
			level = append(level, pop.Val)

			if pop.Left != nil {
				queue.PushBack(pop.Left)
			}
			if pop.Right != nil {
				queue.PushBack(pop.Right)
			}
		}

		ans = append(ans, level)
	}

	return ans
}

// levelOrder3 与 levelOrder2 同一解，但是使用数组队列，优化速度
func levelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var ans [][]int
	var q1, q2 []*TreeNode
	q1 = append(q1, root)

	for len(q1) > 0 {
		level := make([]int, 0, len(q1))
		for i, size := 0, len(q1); i < size; i++ {
			pop := q1[i]
			level = append(level, pop.Val)

			if pop.Left != nil {
				q2 = append(q2, pop.Left)
			}
			if pop.Right != nil {
				q2 = append(q2, pop.Right)
			}
		}

		ans = append(ans, level)
		q1, q2 = q2, q1
		q2 = q2[:0]
	}

	return ans
}

// 32 - III. 从上到下打印二叉树 III
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var ans [][]int
	var s1, s2 []*TreeNode
	s1 = append(s1, root)

	for len(s1) > 0 {
		s2 = s2[:0]
		size := len(s1)
		level := make([]int, 0, size)
		for size > 0 {
			// 倒序遍历
			size--
			pop := s1[size]
			level = append(level, pop.Val)

			// 顺序放入
			if pop.Left != nil {
				s2 = append(s2, pop.Left)
			}
			if pop.Right != nil {
				s2 = append(s2, pop.Right)
			}
		}
		// 添加到ans中
		ans = append(ans, level)

		size = len(s2)
		if size < 1 {
			break
		}
		// q1遍历过，q2可能还有下一层需要遍历，但方向相反
		s1 = s1[:0]
		level = make([]int, 0, size)
		for size > 0 {
			size--
			// 顺序遍历
			pop := s2[size]
			level = append(level, pop.Val)

			// 倒叙放入
			if pop.Right != nil {
				s1 = append(s1, pop.Right)
			}
			if pop.Left != nil {
				s1 = append(s1, pop.Left)
			}
		}
		ans = append(ans, level)
	}

	return ans
}
