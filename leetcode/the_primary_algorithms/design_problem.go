package the_primary_algorithms

import (
	"math/rand"
)

// 以下是 the_primary_algorithms 关于 design problem 的代码实现.

// 1.打乱数组
/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

type Solution struct {
	nums, res []int
}

func Constructor1(nums []int) Solution {
	res := make([]int, len(nums))
	copy(res, nums)

	return Solution{nums: nums, res: res}
}

func (this *Solution) Reset() []int {
	return this.nums
}

func (this *Solution) Shuffle() []int {
	for n := len(this.nums) - 1; 0 < n; n-- {
		random := rand.Intn(n + 1)
		this.res[n], this.res[random] = this.res[random], this.res[n]
		// fmt.Println(random, n, this.res)
	}

	return this.res
}

// 2.最小栈
/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

type MinNode struct {
	val  int
	min  int
	next *MinNode
}

type MinStack struct {
	head *MinNode
}

func Constructor2() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.head = &MinNode{val: val, min: val, next: this.head}
	if this.head.next != nil && this.head.min > this.head.next.min {
		this.head.min = this.head.next.min
	}
}

func (this *MinStack) Pop() {
	this.head = this.head.next
}

func (this *MinStack) Top() int {
	return this.head.val
}

func (this *MinStack) GetMin() int {
	return this.head.min
}
