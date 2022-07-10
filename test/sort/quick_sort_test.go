package sort_test

import (
	"algorithms/intro_algorithms/cp02"
	"algorithms/pkg/sort"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

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

	sort.QuickSortByPartition(tempArr)
	fmt.Printf("after sort: %v\n", tempArr)

}

// TestQuickSort ...
func TestQuickSortBy3Way(t *testing.T) {
	tempArr := cp02.GetRandomArr(20, 100)

	fmt.Printf("before sort: %v\n", tempArr)

	sort.QuickSortBy3Way(tempArr)
	fmt.Printf("after sort: %v\n", tempArr)

}
