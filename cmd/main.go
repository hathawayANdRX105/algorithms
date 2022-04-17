package main

import (
	cp2 "algorithms/pkg/cp02"
)

func main() {
	arraySize, rangeSize := 100, 100

	cp2.TestSort(true, arraySize, rangeSize, cp2.MergeSort)
}
