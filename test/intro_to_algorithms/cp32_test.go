package cp32_test

import (
	"algorithms/intro_algorithms/cp32"
	"testing"
)

func TestComputePrefixFunction(t *testing.T) {
	pattern := "ababaca"

	next := cp32.GetNext(pattern)

	t.Logf("%v", next)
}

func TestKMPMatcher(t *testing.T) {
	txt := "bacbababaabcbab"
	pattern := "ababaca"

	subStr := cp32.KMPMatcher(txt, pattern)

	t.Logf("%v", subStr)
}
