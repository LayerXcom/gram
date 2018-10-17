package tinyram

import (
	"fmt"	
	"testing"

	"gotest.tools/assert"
)

func TestAndOperation(t *testing.T) {
	cases := []struct {
		tRAM     *tinyRAM
		r1       uint64
		r2       uint64
		r3       uint64
		expected uint64
	}{
		{
			tRAM: &tinyRAM{
				Register:      []uint64{0, 0, 3},
				ConditionFlag: true,
			},
			r1:       uint64(0),
			r2:       uint64(2),
			r3:       uint64(10),
			expected: uint64(3 & 10),
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
		r1       uint64
		r2       uint64
		r3       uint64
		expected uint64
	}{
		{
			tRAM: &tinyRAM{
				Register:      []uint64{0, 0, 3},
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
		r1       uint64
		r2       uint64
		r3       uint64
		expected uint64
	}{
		{
			tRAM: &tinyRAM{
				Register:      []uint64{0, 0, 3},
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
		r1       uint64
		r2       uint64
		r3       uint64
		expected uint64
	}{
		{
			tRAM: &tinyRAM{
				Register:      []uint64{0, 0, 3},
				ConditionFlag: true,
			},
			r1:       0,
			r2:       2,
			r3:       0,
			// 2 ^ (2 ** 64 -1) 
			expected: 18446744073709551613,
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
		r1       uint64
		r2       uint64
		r3       uint64
		expected uint64
	}{
		{
			tRAM: &tinyRAM{				
				Register:      []uint64{0, 0, 3},
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
		r1       uint64
		r2       uint64
		r3       uint64
		expected uint64
	}{
		{
			tRAM: &tinyRAM{				
				Register:      []uint64{0, 0, 10},
				ConditionFlag: true,
			},
			r1:       0,
			r2:       2,
			r3:       3,			
			expected: 7,
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
		r1       uint64
		r2       uint64
		r3       uint64
		expected uint64
	}{
		{
			tRAM: &tinyRAM{				
				Register:      []uint64{0, 0, 3},
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
		r1       uint64
		r2       uint64
		r3       uint64
		expected uint64
	}{
		{
			tRAM: &tinyRAM{				
				Register:      []uint64{0, 0, 3},
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
		r1       uint64
		r2       uint64
		r3       uint64
		expected uint64
	}{
		{
			tRAM: &tinyRAM{				
				Register:      []uint64{0, 0, 3},
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
		r1       uint64
		r2       uint64
		r3       uint64
		expected uint64
	}{
		{
			tRAM: &tinyRAM{				
				Register:      []uint64{0, 0, 10},
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
		r1       uint64
		r2       uint64
		r3       uint64
		expected uint64
	}{
		{
			tRAM: &tinyRAM{				
				Register:      []uint64{0, 0, 10},
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

func TestShlOperation(t *testing.T) {
	cases := []struct {
		tRAM     *tinyRAM
		r1       uint64
		r2       uint64
		r3       uint64
		expectedValue uint64
		expectedFlag bool
	}{
		{
			tRAM: &tinyRAM{				
				Register:      []uint64{0, 0, 10},
				ConditionFlag: true,
			},
			r1:       0,
			r2:       2,
			r3:       3,			
			expectedValue: 80,
			expectedFlag: false,
		},
		{
			tRAM: &tinyRAM{				
				Register:      []uint64{0, 0, 1},
				ConditionFlag: false,
			},
			r1:       0,
			r2:       2,
			r3:       63,			
			// 1 << 63 =
			expectedValue: 9223372036854775808,
			expectedFlag: true,
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			shlOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)
			assert.Equal(t, tcc.expectedValue, tcc.tRAM.Register[tcc.r1])
			assert.Equal(t, tcc.expectedFlag, tcc.tRAM.ConditionFlag)
		})
	}
}

func TestShrOperation(t *testing.T) {
	cases := []struct {
		tRAM     *tinyRAM
		r1       uint64
		r2       uint64
		r3       uint64
		expectedValue uint64
		expectedFlag bool
	}{
		{
			tRAM: &tinyRAM{				
				Register:      []uint64{0, 0, 100},
				ConditionFlag: true,
			},
			r1:       0,
			r2:       2,
			r3:       3,			
			expectedValue: 12,
			expectedFlag: false,
		},
		{
			tRAM: &tinyRAM{				
				Register:      []uint64{0, 0, 4},
				ConditionFlag: false,
			},
			r1:       0,
			r2:       2,
			r3:       2,			
			// 4 >> 2 =
			expectedValue: 1,
			expectedFlag: true,
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			shrOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)
			assert.Equal(t, tcc.expectedValue, tcc.tRAM.Register[tcc.r1])
			assert.Equal(t, tcc.expectedFlag, tcc.tRAM.ConditionFlag)
		})
	}
}

func TestCmpeOperation(t *testing.T) {
	cases := []struct {
		tRAM     *tinyRAM
		r1       uint64
		r2       uint64
		r3       uint64
		expectedValue uint64
		expectedFlag bool
	}{
		{
			tRAM: &tinyRAM{				
				Register:      []uint64{0, 0, 0},
				ConditionFlag: true,
			},
			r1:       1,
			r2:       2,
			r3:       0,			
			expectedValue: 0,
			expectedFlag: false,
		},
		{
			tRAM: &tinyRAM{				
				Register:      []uint64{0, 0, 0},
				ConditionFlag: false,
			},
			r1:       2,
			r2:       2,
			r3:       0,			
			expectedValue: 0,
			expectedFlag: true,
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			cmpeOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)
			assert.Equal(t, tcc.expectedValue, tcc.tRAM.Register[tcc.r1])
			assert.Equal(t, tcc.expectedFlag, tcc.tRAM.ConditionFlag)
		})
	}
}

func TestCmpaOperation(t *testing.T) {
	cases := []struct {
		tRAM     *tinyRAM
		r1       uint64
		r2       uint64
		r3       uint64
		expectedValue uint64
		expectedFlag bool
	}{
		{
			tRAM: &tinyRAM{				
				Register:      []uint64{0, 0, 0},
				ConditionFlag: true,
			},
			r1:       1,
			r2:       2,
			r3:       0,			
			expectedValue: 0,
			expectedFlag: false,
		},
		{
			tRAM: &tinyRAM{				
				Register:      []uint64{0, 0, 0},
				ConditionFlag: false,
			},
			r1:       3,
			r2:       2,
			r3:       0,			
			expectedValue: 0,
			expectedFlag: true,
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			cmpaOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)			
			assert.Equal(t, tcc.expectedFlag, tcc.tRAM.ConditionFlag)
		})
	}
}

func TestCmpaeOperation(t *testing.T) {
	cases := []struct {
		tRAM     *tinyRAM
		r1       uint64
		r2       uint64
		r3       uint64
		expectedValue uint64
		expectedFlag bool
	}{
		{
			tRAM: &tinyRAM{				
				Register:      []uint64{0, 0, 0},
				ConditionFlag: true,
			},
			r1:       1,
			r2:       2,
			r3:       0,			
			expectedValue: 0,
			expectedFlag: false,
		},
		{
			tRAM: &tinyRAM{				
				Register:      []uint64{0, 0, 0},
				ConditionFlag: false,
			},
			r1:       2,
			r2:       2,
			r3:       0,			
			expectedValue: 0,
			expectedFlag: true,
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			cmpaeOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)			
			assert.Equal(t, tcc.expectedFlag, tcc.tRAM.ConditionFlag)
		})
	}
}

func TestCmpgOperation(t *testing.T) {
	cases := []struct {
		tRAM     *tinyRAM
		r1       uint64
		r2       uint64
		r3       uint64
		expectedValue uint64
		expectedFlag bool
	}{
		{
			tRAM: &tinyRAM{				
				Register:      []uint64{0, 0, 0},
				ConditionFlag: true,
			},
			r1:       1,
			r2:       2,
			r3:       0,			
			expectedValue: 0,
			expectedFlag: false,
		},
		{
			tRAM: &tinyRAM{				
				Register:      []uint64{0, 0, 0},
				ConditionFlag: false,
			},
			r1:       3,
			r2:       2,
			r3:       0,			
			expectedValue: 0,
			expectedFlag: true,
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			cmpgOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)			
			assert.Equal(t, tcc.expectedFlag, tcc.tRAM.ConditionFlag)
		})
	}
}

func TestCmpgeOperation(t *testing.T) {
	cases := []struct {
		tRAM     *tinyRAM
		r1       uint64
		r2       uint64
		r3       uint64
		expectedValue uint64
		expectedFlag bool
	}{
		{
			tRAM: &tinyRAM{				
				Register:      []uint64{0, 0, 0},
				ConditionFlag: true,
			},
			r1:       1,
			r2:       2,
			r3:       0,			
			expectedValue: 0,
			expectedFlag: false,
		},
		{
			tRAM: &tinyRAM{				
				Register:      []uint64{0, 0, 0},
				ConditionFlag: false,
			},
			r1:       2,
			r2:       2,
			r3:       0,			
			expectedValue: 0,
			expectedFlag: true,
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			cmpgeOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)			
			assert.Equal(t, tcc.expectedFlag, tcc.tRAM.ConditionFlag)
		})
	}
}
