package illustrate_for_offer

import (
	"container/list"
	"sort"
)

/**
* Your CQueue object will be instantiated and called as such:
* obj := Constructor();
* obj.AppendTail(value);
* param_2 := obj.DeleteHead();
 */

// 09. 用两个栈实现队列
type CQueue struct {
	head, tail []int // 模仿栈
}

func Constructor1() CQueue {
	return CQueue{}
}

func (cq *CQueue) AppendTail(value int) {
	cq.tail = append(cq.tail, value)
}

func (cq *CQueue) DeleteHead() (top int) {
	if len(cq.head) < 1 {
		// 头栈与尾栈都没有元素
		if len(cq.tail) < 1 {
			top = -1
			return
		}

		// 转换尾栈
		h, t := cq.head, cq.tail
		for len(t) > 0 {
			end := len(t) - 1
			h = append(h, t[end])
			t = t[:end]
		}

		cq.head, cq.tail = h, t
	}

	top = cq.head[len(cq.head)-1]
	cq.head = cq.head[:len(cq.head)-1]
	return
}

/**
* Your MinStack object will be instantiated and called as such:
* obj := Constructor();
* obj.Push(x);
* obj.Pop();
* param_3 := obj.Top();
* param_4 := obj.Min();
 */

type minNode struct {
	val int
	min int
}

type MinStack struct {
	inStack []minNode
}

/** initialize your data structure here. */
func Constructor2() MinStack {
	return MinStack{}
}

func (ms *MinStack) Push(x int) {
	min := x
	if len(ms.inStack) > 0 && ms.inStack[len(ms.inStack)-1].min < min {
		min = ms.inStack[len(ms.inStack)-1].min
	}
	ms.inStack = append(ms.inStack, minNode{val: x, min: min})
}

func (ms *MinStack) Pop() {
	if len(ms.inStack) < 1 {
		return
	}

	ms.inStack = ms.inStack[:len(ms.inStack)-1]
}

func (ms *MinStack) Top() int {
	return ms.inStack[len(ms.inStack)-1].val
}

func (ms *MinStack) Min() int {
	return ms.inStack[len(ms.inStack)-1].min
}

// 31. 栈的压入、弹出序列 [*]
func validateStackSequences1(pushed []int, popped []int) bool {
	stack := make([]int, 0, len(pushed))
	var p int
	// push -> pop .... -> push
	// 模仿入栈跟出栈操作
	for i := range pushed {
		stack = append(stack, pushed[i])

		for len(stack) > 1 && stack[len(stack)-1] == popped[p] {
			p++
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) > 1
}

// validateStackSequences2 减少使用stack，利用指针进行模拟入栈出栈
func validateStackSequences2(pushed []int, popped []int) bool {
	if len(pushed) < 1 {
		return true
	}

	var t, p int
	for i := range pushed {
		t++
		pushed[t-1] = pushed[i]
		for t > 0 && pushed[t-1] == popped[p] {
			p++
			t--
		}
	}

	return t == 0
}

// 59 - I. 滑动窗口的最大值
// 后续影响前区间
func maxSlidingWindow(nums []int, k int) []int {
	for i := k - 1; i < len(nums); i++ {
		for j := i - 1; j > i-k && nums[i] > nums[j]; j-- {
			nums[j] = nums[i]
		}
	}

	return nums[:len(nums)-k+1]
}

type hp struct {
	nums []int
	sort.IntSlice
}

func (h *hp) Less(i, j int) bool {
	// heap 存索引，通过nums与索引比对
	return h.nums[h.IntSlice[i]] > h.nums[h.IntSlice[j]]
}

func (h *hp) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *hp) Pop() interface{} {
	// 不需要返回值使用，直接缩减
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return nil
}

// 59 - II. 队列的最大值[*]
/**
* Your MaxQueue object will be instantiated and called as such:
* obj := Constructor();
* param_1 := obj.Max_value();
* obj.Push_back(value);
* param_3 := obj.Pop_front();
 */

type MaxQueue struct {
	queue, sup *list.List
}

func Constructor() MaxQueue {
	return MaxQueue{queue: list.New(), sup: list.New()}
}

func (mq *MaxQueue) Max_value() int {
	if mq.sup.Len() < 1 {
		return -1
	}

	return mq.sup.Back().Value.(int)
}

func (mq *MaxQueue) Push_back(value int) {
	mq.queue.PushBack(value)

	for mq.sup.Len() > 0 && mq.sup.Back().Value.(int) < value {
		mq.sup.Remove(mq.sup.Back())
	}

	mq.sup.PushBack(value)
}

func (mq *MaxQueue) Pop_front() int {
	if mq.queue.Len() < 1 {
		return -1
	}

	pop := mq.queue.Remove(mq.queue.Front()).(int)
	for mq.sup.Len() > 0 && mq.sup.Front().Value.(int) == pop {
		mq.sup.Remove(mq.sup.Front())
	}

	return pop
}
