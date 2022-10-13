package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// parseInput ...
func parseInput(buf *bufio.Reader) ([]int, []int) {
	size, _ := strconv.Atoi(ReadLine(buf))
	// fmt.Printf("size:%v", size)
	xSlice := make([]int, size)
	ySlice := make([]int, size)

	xSlice = fillSliceBySplitStr(strings.Split(ReadLine(buf), " "), xSlice)
	ySlice = fillSliceBySplitStr(strings.Split(ReadLine(buf), " "), ySlice)

	return xSlice, ySlice
}

func FindMaximumGroup(x []int, y []int) {
	for i := 0; i < len(x); i++ {
		y[i] -= x[i]
	}

	sort.Ints(y)

	l, r := 0, len(y)-1
	var count int
	for l < r && y[r] > -1 {
		if y[r]+y[l] < 0 {
			l++
			continue
		}

		count++
		r--
		l++
	}

	fmt.Println(count)
}

func RunForFindMaximumGroup() {
	buf := bufio.NewReader(os.Stdin)
	defer os.Stdin.Close()

	queSize, _ := strconv.Atoi(ReadLine(buf))

	for queSize > 0 {
		x, y := parseInput(buf)
		FindMaximumGroup(x, y)
		queSize--
	}
}
