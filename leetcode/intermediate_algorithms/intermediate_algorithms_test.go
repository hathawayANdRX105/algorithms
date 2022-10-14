package intermediate_algorithms_test

import (
	"strconv"
	"testing"
)

func TestFractionToDecimal(t *testing.T) {
	t.Log(strconv.FormatFloat(4.0/333.0, 'f', 64, 64))
}

func TestSlice(t *testing.T) {
	ex1 := []int{2, 1}

	t.Log(ex1[:1], ex1[2:])

	ex1 = ex1[1:]
	t.Log(ex1)

	ex1 = ex1[1:]
	t.Log(ex1)
}

func TestStrNum(t *testing.T) {
	n1 := "2"

	t.Log(n1[0] - '0')

}
