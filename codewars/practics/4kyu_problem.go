package codewarsproblem

import (
	"fmt"
	"math"

	"golang.org/x/tools/go/analysis/passes/ifaceassert"
)

// Sudoku Solution Validator 数独验证
// https://www.codewars.com/kata/529bf0e9bdf7657179000008
func ValidateSolution(m [][]int) bool {
	row, col := len(m)-1, len(m[0])-1

	var rowSum int                  // row sum should be 45
	colSumArr := make([]int, col+1) // col sum should be 45
	gridSum := make([]int, 3)       // gridSum[i] is the ith sum of 3x3 grid every three row

	for row > -1 {
		for c := col; c > -1; c-- {
			v := m[row][c]
			if v == 0 { // invalid number
				return false
			}

			rowSum += v
			colSumArr[c] += v
			gridSum[c/3] += v
		}

		if rowSum != 45 {
			return false
		}

		if row%3 == 0 {
			for i := 0; i < len(gridSum); i++ {
				if gridSum[i] != 45 {
					return false
				}

				gridSum[i] = 0
			}
		}

		rowSum = 0
		row--
	}

	for _, v := range colSumArr {
		if v != 45 {
			return false
		}
	}

	return true
}

func ValidateSolution2(m [][]int) bool {
	for r := 0; r < 9; r++ {
		row, col, box := 0, 0, 0
		for c := 0; c < 9; c++ {
			i := (r/3)*3 + c/3
			j := (r%3)*3 + c%3

			row ^= 1 << uint(m[r][c])
			col ^= 1 << uint(m[c][r])
			box ^= 1 << uint(m[i][j])
		}
		if row != 1022 || col != 1022 || box != 1022 {
			return false
		}
	}
	return true
}

func limit(a1, a2 int64) int64 {
	if a1 < 0 {
		return a2
	}

	return a1
}

// recursiveDecompose ...
func recursiveDecompose(retain, n int64, add func(nextN int64)) bool {
	if retain == 0 {
		if n != 0 { // ignore n = 0, add(0) also is correct
			add(n)
		}
		return true
	}

	for n > 0 {
		// nSqureSum = 1^2 + .. + n^2 or overflow 64bit then use max int64(1<<63 - 1)
		nSqureSum := limit(n*(n+1)*(2*n+1)/6, math.MaxInt64)

		// if nSqureSum is less than retain, there're no possible decomposes made of retain
		if nSqureSum < retain {
			return false
		}

		retainSubN := retain - n*n
		nextN := int64(math.Sqrt(float64(retainSubN))) // calculate the most closely nextN of sqrt of retainSubN

		if nextN >= n { // keep asc order
			nextN = n - 1
		}

		if recursiveDecompose(retainSubN, nextN, add) {
			add(n)
			return true
		}

		n--
	}

	return false
}

// Square into Squares. Protect trees!
// https://www.codewars.com/kata/54eb33e5bc1a25440d000891
func Decompose(n int64) []int64 {
	var ans []int64
	add := func(nextN int64) { // save memory by pointer
		ans = append(ans, nextN)
	}

	recursiveDecompose(n*n, n-1, add)

	return ans
}
