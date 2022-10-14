package bit_operation_test

import (
	"algorithms/pkg/bit_operation"
	"strconv"
	"testing"
)

// TestAdd ...
func TestAdd(t *testing.T) {
	n1, n2 := 0, -7
	t.Logf("n1 + n2 = %v\n", bit_operation.Add(n1, n2))
	t.Logf("n1 - n2 = %v\n", bit_operation.Sub(n1, n2))
}

// TestShiftForMinus ...
func TestShiftForMinus(t *testing.T) {
	// var i int32 = -1 << 32
	var i int32 = -1 << 31
	t.Log(strconv.FormatInt(int64(i), 2), i)

}

func TestMultiply(t *testing.T) {
	n1, n2 := 3, -7

	t.Logf("n1 * n2 = %v\n", bit_operation.Multiply(n1, n2))
}

func TestDivide(t *testing.T) {
	n1, n2 := 21, -3

	t.Logf("n1 / n2 = %v\n", bit_operation.Divide(n1, n2))
}
