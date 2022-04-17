package test

import (
	"fmt"
	"testing"

	cp2 "algorithms/pkg/cp02"
)

func TestInsert(t *testing.T) {
	fmt.Println("test insert sort func")

	arr := cp2.GetRandomArr(30, 100)
	cmp := cp2.Compare(false)

	fmt.Printf("arr: %v\n", arr)
	result := cp2.Insert_sort(arr, cmp)

	fmt.Printf("result: %v\n", result)
}

func TestValidSequence(t *testing.T) {
	a := make([]int, 20)
	for i := 0; i < 20; i++ {
		a[i] = i
	}

	fmt.Printf("result: %v\n", cp2.ValidSequence(a, cp2.Compare(true)))
}

func TestSelect(t *testing.T) {
	cp2.TestSort(true, 30, 100, cp2.Select_sort)
}

func TestExchange(t *testing.T) {
	arr := make([]int, 2)

	arr[0] = 1
	arr[1] = 2
	cp2.Exchange(arr, 0, 1)

	fmt.Print(arr)
}

func TestMerge(t *testing.T) {
	cp2.TestSort(true, 200, 100, cp2.MergeSort)
}
