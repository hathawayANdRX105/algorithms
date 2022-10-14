package cp06_test

import (
	"algorithms/intro_algorithms/cp06"
	"algorithms/pkg/generic_sort"
	"testing"
)

func TestHeapSort(t *testing.T) {
	var comparator generic_sort.Comparator[byte]
	comparator = &generic_sort.NumberComparator[byte]{}
	arr := []byte{3, 4, 1, 2, 5, 7, 10}

	// var comparator generic_sort.Comparator[float32]
	// comparator = &generic_sort.NumberComparator[float32]{}
	// arr := []float32{3.3, 1.3, 4, 5.8, 10, 9, 2.2}

	sortedArr := cp06.HeapSort(arr, comparator)

	t.Log(sortedArr)
}

func TestPriorityQueue(t *testing.T) {
	var comparator generic_sort.Comparator[byte]
	comparator = &generic_sort.NumberComparator[byte]{}
	arr := []byte{3, 4, 1, 2, 5, 7, 10}

	// var comparator generic_sort.Comparator[float32]
	// comparator = &generic_sort.NumberComparator[float32]{}
	// arr := []float32{3.3, 1.3, 4, 5.8, 10, 9, 2.2}

	pq := cp06.BuildPriorityQueue(comparator, arr...)

	pq.Insert(6)
	pq.Print()

	pq.Replace(7, 8)
	pq.Print()
	t.Log(pq.Peek())

	pq.Insert(11)
	pq.Print()

	// for i := pq.Size(); 0 < i; i-- {
	// 	t.Logf("size:%v \t peek:%v\n", i, pq.GetPeek())
	// 	pq.Print()
	// }

}
