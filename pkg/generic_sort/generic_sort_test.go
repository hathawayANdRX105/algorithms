package generic_sort_test

import (
	"algorithms/intro_algorithms/cp02"
	"algorithms/pkg/generic_sort"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	fmt.Println("test insert sort func")

	arr := cp02.GetRandomArr(30, 100)
	cmp := cp02.Compare(false)

	fmt.Printf("arr: %v\n", arr)
	result := cp02.InsertSort(arr, cmp)

	fmt.Printf("result: %v\n", result)
}

func TestValidSequence(t *testing.T) {
	a := make([]int, 20)
	for i := 0; i < 20; i++ {
		a[i] = i
	}

	t.Logf("result: %v\n", cp02.ValidSequence(a, cp02.Compare(true)))
}

func TestSelect(t *testing.T) {
	cp02.TestSort(true, 30, 100, cp02.SelectSort)
}

func TestExchange(t *testing.T) {
	arr := make([]int, 2)

	arr[0] = 1
	arr[1] = 2
	cp02.Exchange(arr, 0, 1)

	fmt.Print(arr)
}

func TestMerge(t *testing.T) {
	cp02.TestSort(true, 20, 10, cp02.MergeSort)
}

// TestPointerExchange 尝试寻找 利用指针交换两个值，如果交换值还存在细粒度的操作，需要提供操作空间
func TestPointerExchange(t *testing.T) {

	exchange := func(i, j *any) (any, any) {
		// inside here can do some atomic operations.
		return *j, *i
	}

	tempArr := []any{1, 2, 3, 4, 5, 6}

	t.Logf("%T, %v", tempArr, tempArr)
	tempArr[0], tempArr[len(tempArr)-1] = exchange(&tempArr[0], &tempArr[len(tempArr)-1])
	t.Logf("%T, %v", tempArr, tempArr)
}

func TestComparator(t *testing.T) {

	// t.Log("abc" < "abb")
	c := &generic_sort.NumberComparator[byte]{}

	tempArr := []byte{120, 100}
	t.Log(c)
	t.Log(c.IsLess(&tempArr[0], &tempArr[1]))
	t.Log(c.IsEqual(tempArr[0], tempArr[1]))
	t.Log(c.Swap(&tempArr[0], &tempArr[1]))
}

// TestShuffle ...
func TestShuffle(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	tempArr := make([]int, 10)
	for i := 0; i < len(tempArr); i++ {
		tempArr[i] = rand.Int()
	}

	fmt.Printf("before shuffle %v \n", tempArr)

	rand.Shuffle(len(tempArr), func(i, j int) {
		tempArr[i], tempArr[j] = tempArr[j], tempArr[i]
	})

	fmt.Printf("after shuffle %v \n", tempArr)

}

// TestQuickSort ...
func TestQuickSortByPartition(t *testing.T) {
	tempArr := cp02.GetRandomArr(20, 100)

	fmt.Printf("before sort: %v\n", tempArr)

	generic_sort.QuickSortByPartition(tempArr)
	fmt.Printf("after sort: %v\n", tempArr)

}

// TestQuickSort ...
func TestQuickSortBy3Way(t *testing.T) {
	tempArr := cp02.GetRandomArr(20, 100)

	fmt.Printf("before sort: %v\n", tempArr)

	generic_sort.QuickSortBy3Way(tempArr)
	fmt.Printf("after sort: %v\n", tempArr)

}

func TestIndex(t *testing.T) {
	tempArr := []int{1, 2, 3}

	// tempArr[-1] is valid.
	t.Log(tempArr[1])

	t.Log()
}
