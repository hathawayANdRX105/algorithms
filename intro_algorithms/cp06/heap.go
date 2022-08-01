package cp06

import (
	"algorithms/pkg/generic_sort"
	"fmt"
)

/*
 Heap:利用slice作为堆的维护对象
 param:
   array 如果使用原来下标数组element，当 i 为根节点索引时，子树节点为 i*2 + 1, (i+1)*2
         因此将堆的开始下标从1开始，当 i 为根节点索引时，子树节点为 i*2, i*2 + 1
   size 为维护堆的大小，当排序Sort方法调用后，size将为1
   Comparator 用来约束不同类型的比较，交换.
*/

// swim 对当前下标i，i/2进行判断，上浮，维护堆性质
func swim[T generic_sort.Number](i int, array []T, comparator generic_sort.Comparator[T]) {

	// 如果当前索引i节点逻辑上比索引i/2节点大，即当前节点比根节点大，进行上浮
	for 1 < i && comparator.IsLess(&array[i/2], &array[i]) {
		array[i/2], array[i] = comparator.Swap(&array[i/2], &array[i])
		i /= 2
	}
}

// sink 对下标 i的值进行下沉
func sink[T generic_sort.Number](i, size int, array []T, comparator generic_sort.Comparator[T]) {
	var j int
	for i*2 <= size {
		j = i * 2

		// 如果逻辑上 array[j] 小于 array[j+1]
		// 当前根节点的两个子节点 array[j+1] 是逻辑最大值, j偏移一位
		if j+1 <= size && comparator.IsLess(&array[j], &array[j+1]) {
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
	// 因此将堆的开始下标从1开始，当 i 为根节点索引时，子树节点为 i*2, i*2 + 1
	size := len(element)
	array := append(make([]T, 1, len(element)+1), element...)

	// build heap 从下往上，从非叶子节点构建最大堆
	for i := size / 2; 0 < i; i-- {
		sink(i, size, array, comparator)
	}

	// heap sort 堆排序
	for 1 < size {
		// 交换 root 与最后的叶子节点
		// 因为root 储存这最大或最小值， 排到最后的叶子节点，交换之后让原先叶子节点的值进行下沉
		array[1], array[size] = comparator.Swap(&array[1], &array[size])
		size-- // 同时减少堆大小，避免越界影响位置
		
		// 维护最大/小堆
		sink(1, size, array, comparator)
	}

	return array[1:]
}
