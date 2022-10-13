package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func countDominoes(n, m int) int {
	// 单行单列的多米诺骨牌
	if n < 2 {
		return m - 1
	}

	if m < 2 {
		return n - 1
	}

	return n * (m - 1)
}

// https://codeforces.com/problemset/problem/1725/A
func RunforDominoes() {
	// in:3 4
	buf := bufio.NewReader(os.Stdin)
	line, _, err := buf.ReadLine()
	if err != nil {
		fmt.Println(err)
		return
	}

	p := bytes.IndexByte(line, ' ')
	n, _ := strconv.Atoi(string(line)[:p])
	m, _ := strconv.Atoi(string(line)[p+1:])

	fmt.Println(countDominoes(n, m))
}
