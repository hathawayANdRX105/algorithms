package chapter02

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {

	fmt.Println("test insert sort func")

	arr := GetRandomArr(30, 100)
	cmp := Compare(false)

	fmt.Printf("arr: %v\n", arr)
	result := Insert_sort(arr, cmp)

	fmt.Printf("result: %v\n", result)
}

func TestValidSquence(t *testing.T) {
	a := make([]int, 20)
	for i := 0; i < 20; i++ {
		a[i] = i
	}

	fmt.Printf("result: %v\n", ValidSquence(a, Compare(true)))
}

func TestSelect(t *testing.T) {
	TestSort(true, 30, 100, Select_sort)
}

func TestExchange(t *testing.T) {
	arr := make([]int, 2)

	arr[0] = 1
	arr[1] = 2
	Exchange(arr, 0, 1)

	fmt.Print(arr)

}

func TestMerge(t *testing.T) {
	TestSort(true, 200, 100, MergeSort)
}
