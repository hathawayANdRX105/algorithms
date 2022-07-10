package main

import "algorithms/intro_algorithms/cp02"


func main() {
	arraySize, rangeSize := 100, 100

	cp02.TestSort(true, arraySize, rangeSize, cp02.MergeSort)

}

