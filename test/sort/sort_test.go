package sort_test

import (
	"algorithms/intro_algorithms/cp02"
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	fmt.Println("test insert sort func")

	arr := cp02.GetRandomArr(30, 100)
	cmp := cp02.Compare(false)

	fmt.Printf("arr: %v\n", arr)
	result := cp02.Insert_sort(arr, cmp)

	fmt.Printf("result: %v\n", result)
}

func TestValidSequence(t *testing.T) {
	a := make([]int, 20)
	for i := 0; i < 20; i++ {
		a[i] = i
	}

	fmt.Printf("result: %v\n", cp02.ValidSequence(a, cp02.Compare(true)))
}

func TestSelect(t *testing.T) {
	cp02.TestSort(true, 30, 100, cp02.Select_sort)
}

func TestExchange(t *testing.T) {
	arr := make([]int, 2)

	arr[0] = 1
	arr[1] = 2
	cp02.Exchange(arr, 0, 1)

	fmt.Print(arr)
}

func TestMerge(t *testing.T) {
	cp02.TestSort(true, 200, 100, cp02.MergeSort)
}
