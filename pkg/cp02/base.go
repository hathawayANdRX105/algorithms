package cp02

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func init() {
	seed := time.Now().UnixNano()
	rand.Seed(seed)
}

// Compare return a func which can decide order by pass a specific value
func Compare(isAsc bool) func(o1, o2 int) bool {
	if isAsc {
		return isLess
	}

	return isGreater
}

func isLess(o1, o2 int) bool {
	return o1 <= o2
}

func isGreater(o1, o2 int) bool {
	return o1 >= o2
}

// Exchange this fuc will exchange tow value
// arr must be a slice with pointer
// and e1 , e2 is the index of exchange
func Exchange(arr []int, e1, e2 int) {
	temp := arr[e1]
	arr[e1] = arr[e2]
	arr[e2] = temp
}

// ValidSequence pass a array to valid which is sort by order
func ValidSequence(arr []int, cmp func(o1, o2 int) bool) bool {
	if len(arr) < 2 {
		return true
	}

	for i := 1; i < len(arr); i++ {
		if !cmp(arr[i-1], arr[i]) {
			return false
		}
	}

	return true
}

// GetRandomArr generate a random array by pass indicate number
func GetRandomArr(arraySize, rangeSize int) []int {

	arr := make([]int, arraySize)
	for i := 0; i < arraySize; i++ {
		arr[i] = rand.Intn(rangeSize)
	}

	return arr
}

func TestSort(isAsc bool, arraySize, rangeSize int, sort func(arr []int, cmp func(o1, o2 int) bool) []int) {
	arr := GetRandomArr(arraySize, rangeSize)
	cmp := Compare(isAsc)
	targetArr := sort(arr, cmp)

	isSort := ValidSequence(targetArr, cmp)

	fmt.Printf("ramdom arr: %v\n", arr)
	fmt.Printf("sort arr:   %v\n", targetArr)
	fmt.Printf("isSort: %v\n", isSort)
}

func CountDown(t *testing.T, recordFunc func()) {
	startTime := time.Now()

	recordFunc()

	elapsedTime := time.Since(startTime)
	t.Logf("func elapse time:%s", elapsedTime)
	t.Log(elapsedTime)
}
