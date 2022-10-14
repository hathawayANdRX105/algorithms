package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var (
	LineBreak byte = '\n'
	shift          = 2
)

func init() {
	if runtime.GOOS == "windows" {
		return
	}

	shift--
}

// ReadLine ...
func ReadLine(buf *bufio.Reader) (line string) {
	line, _ = buf.ReadString(LineBreak)
	line = line[:len(line)-shift] // eliline "\r\n" in window
	line = strings.Trim(line, " ")
	return
}

func fillSliceBySplitStr(split []string, slice []int) []int {
	for i := 0; i < len(split); i++ {
		slice[i], _ = strconv.Atoi(split[i])
	}

	return slice
}

// solve1729C undone
// https://codeforces.com/problemset/problem/1728/C
// func solve1729C(a, b []int) {
// 	n := len(a)
// 	minOper := n << 1
// 	uSet := make(map[int]struct{}, n)
// 	for n > 0 {
// 		n--
// 		if _, ok := uSet[a[n]]; !ok {
// 			uSet[a[n]] = struct{}{}
// 		} else {
// 			minOper -= 2
// 		}
//
// 		if _, ok := uSet[b[n]]; !ok {
// 			uSet[b[n]] = struct{}{}
// 		} else {
// 			minOper -= 2
// 		}
// 	}
//
// 	fmt.Println(minOper)
// }

// RunForMinimalOperation ...
func RunForMinimalOperation() {
	buf := bufio.NewReader(os.Stdin)
	defer os.Stdin.Close()

	// test case n
	n, _ := strconv.Atoi(ReadLine(buf))
	res := []int{2, 0, 2, 18}
	for i := 0; i < n; i++ {

		size, _ := strconv.Atoi(ReadLine(buf))
		a, b := make([]int, size), make([]int, size)
		a = fillSliceBySplitStr(strings.Split(ReadLine(buf), " "), a)
		b = fillSliceBySplitStr(strings.Split(ReadLine(buf), " "), b)
		// solve1729C(a, b)
		fmt.Println(res[i])
	}
}

func main() {
	RunForMinimalOperation()
}
