package cp08

import (
	"algorithms/intro_algorithms/cp07"
	"math/rand"
	"time"
)

// linear_time_sort.go 实现 计数排序以及桶排序

func init() {
	rand.Seed(time.Now().UnixNano())
}

// CountingSort 计数排序 当 k = O(n) 时，时间复杂度为 O(n)
// time: O(n + k)
func CountingSort(A []int, k int) []int {
	n := len(A)
	B := make([]int, n)   // 跟A大小一致
	C := make([]int, k+1) // C[0~k]

	// 统计A数组格数
	for _, v := range A {
		C[v] += 1
	}

	// 补充小于或等于A中不存在元素的个数
	for i := 1; i < k+1; i++ {
		C[i] += C[i-1]
	}

	// 根据A中元素的个数填入B中, 为了填满B，同时减少个数
	for _, v := range A {
		B[C[v]-1] = v
		C[v] -= 1
	}

	return B
}

// partitionForRadix ...
func partitionForRadix(arr []int, pre, rear int, cmp func(o1, o2 int) bool) int {
	// 随机抽取比较值，期望O(lgn)
	i, j := cp07.GetRandomFromRange(pre, rear), pre
	arr[i], arr[rear] = arr[rear], arr[i]

	for i = j - 1; j < rear; j++ {
		if cmp(arr[j], arr[rear]) {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[rear] = arr[rear], arr[i+1]

	return i + 1
}

func quickSortForRadix(arr []int, pre, rear int, cmp func(o1, o2 int) bool) {
	// 尾递归(DFS): 通过缩圈 pre让 q 总是在另一个sort区间进行划分
	for pre < rear {
		q := partitionForRadix(arr, pre, rear, cmp)
		quickSortForRadix(arr, pre, q-1, cmp)
		pre = q + 1
	}

	// BFS
	// if rear <= pre {
	// 	return
	// }

	// q := partitionForRadix(arr, pre, rear, cmp)

	// quickSortForRadix(arr, pre, q-1, cmp)
	// quickSortForRadix(arr, q+1, rear, cmp)
}

// RadixSort 对十进制数进行排序，d代表了十进制下可以取的位数个数
func RadixSort(arr []int, d int) {
	for i, shift := 0, 1; i < d; i++ {
		quickSortForRadix(arr, 0, len(arr)-1, func(o1, o2 int) bool {
			// 获取对应的进制数
			return o1/shift%10 < o2/shift%10
		})
		shift *= 10
	}
}

type bucketNode struct {
	next *bucketNode
	val  float64
}

// TODO BucketSort
// insertSortForBucket ...
func insertSortForBucket(dummyHead *bucketNode) {
	p := dummyHead.next
	// 维护 dummyHead -> pre 的区间
	for p.next != nil {
		if p.val <= p.next.val {
			p = p.next
			continue
		}

		// insert
		k := dummyHead
		insertP := p.next
		for k.next.val < insertP.val {
			k = k.next
		}

		p.next = insertP.next
		insertP.next = k.next
		k.next = insertP
	}
}

// BucketSort ...
func BucketSort(arr []float64) {
	B := make([]*bucketNode, 10)

	var p *bucketNode

	// 将 arr 元素填入bucket中
	for _, v := range arr {

		i := int(v * 10)
		if B[i] == nil {
			B[i] = &bucketNode{}
		}

		p = B[i].next
		B[i].next = &bucketNode{next: p, val: v}

	}

	// 对 bucket insert sort
	for _, v := range B {
		if v == nil {
			continue
		}

		insertSortForBucket(v)
	}

	// 连接桶
	arr = arr[:0]
	for _, v := range B {
		if v == nil {
			continue
		}

		for v.next != nil {
			arr = append(arr, v.next.val)
			v = v.next
		}
	}
}
