package cp07

import (
	"algorithms/intro_algorithms/cp02"
	"math/rand"
	"time"
)

// 以下QuickSort 都针对int类型，作为练习

func init() {
	rand.Seed(time.Now().UnixNano())
}

// partirion ...
func partition(arr []int, pre, rear int) int {
	// arr[pre, i] 维护 小于等于 arr[rear]的值
	// arr[i+1, j] 维护 大于 arr[rear] 的值
	i, j := pre-1, pre
	for j < rear {
		if arr[j] <= arr[rear] {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}

		j++
	}

	// 交换arr[i+1]与arr[rear]
	arr[i+1], arr[rear] = arr[rear], arr[i+1]
	
	return i + 1
}

// QuickSort ...
func QuickSort(arr []int, pre, rear int) {
	if pre >= rear {
		return
	}

	// 切分子数组arr[pre, rear]
	q := partition(arr, pre, rear)

	// 递归排序 arr[pre, q-1] 跟 arr[q+1, rear]
	QuickSort(arr, pre, q-1)
	QuickSort(arr, q+1, rear)
}

// GetRandomFromRange 针对int类型
func GetRandomFromRange(lo, hi int) int {
	return lo + rand.Intn(hi-lo)
}

// randomizedPartition ...
func randomizedPartition(arr []int, pre, rear int) int {

	// 从 arr[pre, rear] 范围中随机选出 某一个值进行划分
	q := GetRandomFromRange(pre, rear)
	arr[q], arr[rear] = arr[rear], arr[q]

	return partition(arr, pre, rear)
}

// RandomizedQuickSort 的平均期望算法复杂度为O(nlgn)
func RandomizedQuickSort(arr []int, pre, rear int) {
	if pre >= rear {
		return
	}

	// 随机划分
	q := randomizedPartition(arr, pre, rear)

	// 递归排序
	RandomizedQuickSort(arr, pre, q-1)
	RandomizedQuickSort(arr, q+1, rear)
}

// OptimizedQuickSort 针对数组划分排序的数量 小于 k 时，使用插入排序，加速
func OptimizedQuickSort(arr []int, pre, rear, k int) {

	if rear-pre < k {
		cp02.InsertSort(arr[pre:rear+1], func(o1, o2 int) bool { return o1 <= o2 })
		return
	}

	q := partition(arr, pre, rear)

	OptimizedQuickSort(arr, pre, q-1, k)
	OptimizedQuickSort(arr, q+1, rear, k)
}

// HoarePartition C.R.hoare 设计最初的划分版本, 对应思考题7.1
func HoarePartition(arr []int, pre, rear int) int {
	x := pre

	pre, rear = pre-1, rear+1
	for {
		for pre < rear && arr[pre+1] <= arr[x] {
			pre++
		}

		for pre < rear && arr[rear-1] > arr[x] {
			rear--
		}

		if pre >= rear {
			arr[pre], arr[x] = arr[x], arr[pre]
			return pre
		}

		arr[pre+1], arr[rear-1] = arr[rear-1], arr[pre+1]
	}
}
