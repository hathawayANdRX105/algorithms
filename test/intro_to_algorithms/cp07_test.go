package intro_to_algorithms_test

import (
	"algorithms/intro_algorithms/cp02"
	"algorithms/intro_algorithms/cp07"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := cp02.GetRandomArr(10, 100)

	cp07.QuickSort(arr, 0, len(arr)-1)

	t.Log(arr)
}

func TestRamdomizedQuickSort(t *testing.T) {
	arr := cp02.GetRandomArr(10, 100)

	cp07.RandomizedQuickSort(arr, 0, len(arr)-1)

	t.Log(arr)
}

func TestOptimizedQuickSort(t *testing.T) {
	arr := cp02.GetRandomArr(30, 100)

	cp07.OptimizedQuickSort(arr, 0, len(arr)-1, 5)
	t.Logf("optimized quick sort:%v\n", arr)
}

func TestHoarePartition(t *testing.T) {
	arr := cp02.GetRandomArr(10, 100)

	t.Logf("before partition:%v\n", arr)
	q := cp07.HoarePartition(arr, 0, len(arr)-1)
	t.Logf("after  partition:%v\n", arr)
	t.Logf("partition q:%v", q)
}
