package tinyram

import (
	"fmt"
	"testing"
	
	"gotest.tools/assert"
)

func TestAddOperation(t *testing.T) {
	cases := []struct {
		tRAM *tinyRAM
		r1, r2, r3 int64
		expected int64
	}{{
		tRAM: &tinyRAM{10, 10, false, 0, [10], [2^10], [1], [1], program{instruction("AND"), 1, 2, 10}, false},
		r1: 1,
		r2: 2
		r3: 10,
		expected: 2 & 10,
	}}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("$d-th unit test", n), func(t *testing.T) {
			assert.Equal(t, tRAM.Register[r1], expected)
		})
	}
}
