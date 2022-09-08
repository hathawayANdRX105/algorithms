package intro_to_algorithms_test

import (
	"algorithms/intro_algorithms/cp08"
	"testing"
)

func TestCountingSort(t *testing.T) {
	// A := []int{2, 5, 3, 0, 2, 3, 0, 3}
	// k := 5

	A := []int{2, 5, 3, 0, 1, 7, 2, 3, 0, 3, 7}
	k := 7

	B := cp08.CountingSort(A, k)
	t.Log(B)
}

func TestRadixSort(t *testing.T) {
	arr := []int{329, 457, 657, 839, 436, 720, 355}
	d := 3

	cp08.RadixSort(arr, d)

	t.Log(arr)
}

func TestBucketSort(t *testing.T) {
	arr := []float64{0.78, 0.17, 0.39, 0.26, 0.72, 0.94, 0.21, 0.12, 0.23, 0.68}
	cp08.BucketSort(arr)

	t.Logf("after bucket sort:%v\n", arr)
}
