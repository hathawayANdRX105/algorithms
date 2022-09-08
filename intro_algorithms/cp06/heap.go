package cp06

import (
	"algorithms/pkg/generic_sort"
	"fmt"
)

// swim 对当前下标i，i/2进行判断，上浮，维护堆性质
func swim[T generic_sort.Number](i int, array []T, comparator generic_sort.Comparator[T]) {

	// 如果当前索引i节点逻辑上比索引i/2节点大，即当前节点比根节点大，进行上浮
	for 0 < i && comparator.IsLess(&array[(i-1)/2], &array[i]) {
		array[(i-1)/2], array[i] = comparator.Swap(&array[(i-1)/2], &array[i])
		i = (i - 1) / 2
	}
}

// sink 对下标 i的值进行下沉
func sink[T generic_sort.Number](i, size int, array []T, comparator generic_sort.Comparator[T]) {
	var j int
	for i*2+1 < size {
		j = i*2 + 1

		// 如果逻辑上 array[j] 小于 array[j+1]
		// 当前根节点的两个子节点 array[j+1] 是逻辑最大值, j偏移一位
		if j+1 < size && comparator.IsLess(&array[j], &array[j+1]) {
			j++
		}

		// 同上同理， 如果 根节点比两个子节点逻辑上都大，已经满足逻辑上的最大堆，跳出循环
		if comparator.IsLess(&array[j], &array[i]) {
			break
		}

		// h.array[j] > h.array[i] 维护当前节点最大堆性质
		// 交换i，j位置 ，继续维护 原先以索引j节点的逻辑最大堆结构
		array[i], array[j] = comparator.Swap(&array[i], &array[j])
		i = j
	}
}

// HeapSort
func HeapSort[T generic_sort.Number](element []T, comparator generic_sort.Comparator[T]) []T {
	// heap 利用slice作为堆的维护对象，如果使用原来下标数组element
	// 则 当 i 为根节点索引时，子树节点为 i*2 + 1, (i+1)*2
	size := len(element)
	array := element

	// build heap 从下往上，从非叶子节点构建最大堆
	for i := size / 2; 0 <= i; i-- {
		sink(i, size, array, comparator)
	}

	// heap sort 堆排序
	for 0 < size {
		// 交换 root 与最后的叶子节点
		// 因为root 储存这最大或最小值， 排到最后的叶子节点，交换之后让原先叶子节点的值进行下沉
		array[0], array[size-1] = comparator.Swap(&array[0], &array[size-1])
		size-- // 同时减少堆大小，避免越界影响位置

		// 维护最大/小堆
		sink(0, size, array, comparator)
	}

	return array
}

/*
 PriorityQueue:利用slice作为堆的维护对象
 param:
   array 如果使用原来下标数组element，当 i 为根节点索引时，子树节点为 i*2 + 1, (i+1)*2
         因此将堆的开始下标从1开始，当 i 为根节点索引时，子树节点为 i*2, i*2 + 1
   size 为维护堆的大小，当排序Sort方法调用后，size将为1
   comparator 用来约束不同类型的比较，交换.
*/

type PriorityQueue[T generic_sort.Number] struct {
	array      []T
	size       int
	comparator generic_sort.Comparator[T]
}

func (pq *PriorityQueue[T]) Size() int {
	return pq.size
}

func (pq *PriorityQueue[T]) IsEmpty() bool {
	return pq.size == 0
}

// expandArrSize 扩增pq的切片大小
func (pq *PriorityQueue[T]) expandArrSize() {
	oldArr := pq.array
	pq.array = make([]T, pq.size, (pq.size+1)<<1)
	copy(pq.array, oldArr)
}

// Insert 插入优先队列中
func (pq *PriorityQueue[T]) Insert(element T) bool {
	// 如果当前容量不足，则扩增
	if pq.size+1 >= cap(pq.array) {
		pq.expandArrSize()
	}

	pq.array = append(pq.array, element)
	// 添加到末尾，然后进行逻辑上浮
	swim(pq.size, pq.array, pq.comparator)
	pq.size++

	return true
}

// Peek 类似 maximum / minumum
func (pq *PriorityQueue[T]) Peek() T {
	if pq.size <= 0 {
		panic("priority queue is empty!")
	}

	return pq.array[0]
}

func (pq *PriorityQueue[T]) GetPeek() T {
	if pq.size <= 0 {
		panic("priority queue is empty!")
	}

	// 将逻辑最优先值与最后一值交换，并且暂时保留
	pq.size--
	pq.array[0], pq.array[pq.size] = pq.array[pq.size], pq.array[0]
	peek := pq.array[pq.size]

	// 下沉交换值
	sink(0, pq.size, pq.array, pq.comparator)

	return peek
}

func (pq *PriorityQueue[T]) Replace(i int, key T) bool {
	// 判断当前队列是否存有一定数量元素
	if pq.size <= 0 {
		panic("priority queue is empty!")
	}

	// 判断索引i是否合理
	if i < 0 || i > pq.size {
		panic(fmt.Errorf("the index i value %v is valid", i))
	}

	// 根据情况
	// 如果交换索引i的值逻辑小于key，则需要进行逻辑下沉，否则逻辑上浮
	if pq.comparator.IsLess(&pq.array[i], &key) {
		pq.array[i] = key
		swim(i, pq.array, pq.comparator)
	} else {
		pq.array[i] = key
		sink(i, pq.size, pq.array, pq.comparator)
	}

	return true
}

func (pq *PriorityQueue[T]) Print() {
	fmt.Printf("%v\n", *pq)
}

func BuildPriorityQueue[T generic_sort.Number](comparator generic_sort.Comparator[T], elements ...T) *PriorityQueue[T] {
	pq := &PriorityQueue[T]{comparator: comparator, size: len(elements), array: elements}

	// 构建逻辑堆
	for i := pq.size/2 - 1; 0 <= i; i-- {
		sink(i, pq.size, pq.array, pq.comparator)
	}

	return pq
}
