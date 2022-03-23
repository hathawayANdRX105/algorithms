package main

import (
	cp2 "algorithms/pkg/chapter02"
)

func main() {
	isAsc := true
	arraySize, rangeSize := 100, 100

	cp2.TestSort(isAsc, arraySize, rangeSize, cp2.MergeSort)
}
