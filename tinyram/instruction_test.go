package tinyram

import (
	"fmt"	
	"testing"

	"gotest.tools/assert"
)

func TestAndOperation(t *testing.T) {
	cases := []struct {
		tRAM     *tinyRAM
		r1       int64
		r2       int64
		r3       int64
		expected int64
	}{
		{
			tRAM: &tinyRAM{
				Register:      []int64{0, 0, 3},
				ConditionFlag: true,
			},
			r1:       0,
			r2:       2,
			r3:       10,
			expected: 3 & 10,
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			andOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)
			assert.Equal(t, tcc.expected, tcc.tRAM.Register[tcc.r1])
			assert.Equal(t, false, tcc.tRAM.ConditionFlag)
		})
	}
}

func TestOrOperation(t *testing.T) {
	cases := []struct {
		tRAM     *tinyRAM
		r1       int64
		r2       int64
		r3       int64
		expected int64
	}{
		{
			tRAM: &tinyRAM{
				Register:      []int64{0, 0, 3},
				ConditionFlag: true,
			},
			r1:       0,
			r2:       2,
			r3:       10,
			expected: 3 | 10,
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			orOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)
			assert.Equal(t, tcc.expected, tcc.tRAM.Register[tcc.r1])
			assert.Equal(t, false, tcc.tRAM.ConditionFlag)
		})
	}
}

func TestXorOperation(t *testing.T) {
	cases := []struct {
		tRAM     *tinyRAM
		r1       int64
		r2       int64
		r3       int64
		expected int64
	}{
		{
			tRAM: &tinyRAM{
				Register:      []int64{0, 0, 3},
				ConditionFlag: true,
			},
			r1:       0,
			r2:       2,
			r3:       10,
			expected: 3 ^ 10,
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			xorOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)
			assert.Equal(t, tcc.expected, tcc.tRAM.Register[tcc.r1])
			assert.Equal(t, false, tcc.tRAM.ConditionFlag)
		})
	}
}

func TestNotOperation(t *testing.T) {
	cases := []struct {
		tRAM     *tinyRAM
		r1       int64
		r2       int64
		r3       int64
		expected int64
	}{
		{
			tRAM: &tinyRAM{
				Register:      []int64{0, 0, 3},
				ConditionFlag: true,
			},
			r1:       0,
			r2:       2,
			r3:       0,
			expected: 4611686018427387901,
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			notOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)
			assert.Equal(t, tcc.expected, tcc.tRAM.Register[tcc.r1])
			assert.Equal(t, false, tcc.tRAM.ConditionFlag)
		})
	}
}

func TestAddOperation(t *testing.T) {
	cases := []struct {
		tRAM     *tinyRAM
		r1       int64
		r2       int64
		r3       int64
		expected int64
	}{
		{
			tRAM: &tinyRAM{				
				Register:      []int64{0, 0, 3},
				ConditionFlag: true,
			},
			r1:       0,
			r2:       2,
			r3:       10,
			expected: 3 + 10,
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			addOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)
			assert.Equal(t, tcc.expected, tcc.tRAM.Register[tcc.r1])
			assert.Equal(t, false, tcc.tRAM.ConditionFlag)
		})
	}
}

func TestSubOperation(t *testing.T) {
	cases := []struct {
		tRAM     *tinyRAM
		r1       int64
		r2       int64
		r3       int64
		expected int64
	}{
		{
			tRAM: &tinyRAM{				
				Register:      []int64{0, 0, 10},
				ConditionFlag: true,
			},
			r1:       0,
			r2:       2,
			r3:       3,
			// 4611686018427387903 + 3 - 10 =
			expected: 4611686018427387910,
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			subOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)
			assert.Equal(t, tcc.expected, tcc.tRAM.Register[tcc.r1])
			assert.Equal(t, false, tcc.tRAM.ConditionFlag)
		})
	}
}

func TestMullOperation(t *testing.T) {
	cases := []struct {
		tRAM     *tinyRAM
		r1       int64
		r2       int64
		r3       int64
		expected int64
	}{
		{
			tRAM: &tinyRAM{				
				Register:      []int64{0, 0, 3},
				ConditionFlag: true,
			},
			r1:       0,
			r2:       2,
			r3:       10,
			expected: 3 * 10,
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			mullOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)
			assert.Equal(t, tcc.expected, tcc.tRAM.Register[tcc.r1])
			assert.Equal(t, false, tcc.tRAM.ConditionFlag)
		})
	}
}

func TestUmlhOperation(t *testing.T) {
	cases := []struct {
		tRAM     *tinyRAM
		r1       int64
		r2       int64
		r3       int64
		expected int64
	}{
		{
			tRAM: &tinyRAM{				
				Register:      []int64{0, 0, 3},
				ConditionFlag: true,
			},
			r1:       0,
			r2:       2,
			r3:       10,
			expected: 3 * 10,
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			umulhOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)
			assert.Equal(t, tcc.expected, tcc.tRAM.Register[tcc.r1])
			assert.Equal(t, false, tcc.tRAM.ConditionFlag)
		})
	}
}

func TestSmulhOperation(t *testing.T) {
	cases := []struct {
		tRAM     *tinyRAM
		r1       int64
		r2       int64
		r3       int64
		expected int64
	}{
		{
			tRAM: &tinyRAM{				
				Register:      []int64{0, 0, 3},
				ConditionFlag: true,
			},
			r1:       0,
			r2:       2,
			r3:       10,
			expected: 3 * 10,
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			smulhOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)
			assert.Equal(t, tcc.expected, tcc.tRAM.Register[tcc.r1])
			assert.Equal(t, false, tcc.tRAM.ConditionFlag)
		})
	}
}

func TestUdivOperation(t *testing.T) {
	cases := []struct {
		tRAM     *tinyRAM
		r1       int64
		r2       int64
		r3       int64
		expected int64
	}{
		{
			tRAM: &tinyRAM{				
				Register:      []int64{0, 0, 10},
				ConditionFlag: true,
			},
			r1:       0,
			r2:       2,
			r3:       3,
			// 10 / 3 = 3 with a remainder 1
			expected: 3,
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			udivOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)
			assert.Equal(t, tcc.expected, tcc.tRAM.Register[tcc.r1])
			assert.Equal(t, false, tcc.tRAM.ConditionFlag)
		})
	}
}

func TestUmodOperation(t *testing.T) {
	cases := []struct {
		tRAM     *tinyRAM
		r1       int64
		r2       int64
		r3       int64
		expected int64
	}{
		{
			tRAM: &tinyRAM{				
				Register:      []int64{0, 0, 10},
				ConditionFlag: true,
			},
			r1:       0,
			r2:       2,
			r3:       3,
			// 10 / 3 = 3 with a remainder 1
			expected: 1,
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			umodOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)
			assert.Equal(t, tcc.expected, tcc.tRAM.Register[tcc.r1])
			assert.Equal(t, false, tcc.tRAM.ConditionFlag)
		})
	}
}
