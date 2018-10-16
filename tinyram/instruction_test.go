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
			expected: 9223372036854775805,
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