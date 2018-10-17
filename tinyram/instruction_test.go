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
		expectedPc uint64
	}{
		{
			tRAM: &tinyRAM{
				Register:      []uint64{0, 0, 3},
				ConditionFlag: true,
				Pc: 0,
			},
			r1:       uint64(0),
			r2:       uint64(2),
			r3:       uint64(10),
			expected: uint64(3 & 10),
			expectedPc: uint64(1),
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			andOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)
			assert.Equal(t, tcc.expected, tcc.tRAM.Register[tcc.r1])
			assert.Equal(t, false, tcc.tRAM.ConditionFlag)
			assert.Equal(t, tcc.expectedPc, tcc.tRAM.Pc)
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
		expectedPc uint64
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
			expectedPc: uint64(1),
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			orOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)
			assert.Equal(t, tcc.expected, tcc.tRAM.Register[tcc.r1])
			assert.Equal(t, false, tcc.tRAM.ConditionFlag)
			assert.Equal(t, tcc.expectedPc, tcc.tRAM.Pc)
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
		expectedPc uint64
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
			expectedPc: uint64(1),
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			xorOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)
			assert.Equal(t, tcc.expected, tcc.tRAM.Register[tcc.r1])
			assert.Equal(t, false, tcc.tRAM.ConditionFlag)
			assert.Equal(t, tcc.expectedPc, tcc.tRAM.Pc)
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
		expectedPc uint64
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
			expectedPc: uint64(1),
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			notOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)
			assert.Equal(t, tcc.expected, tcc.tRAM.Register[tcc.r1])
			assert.Equal(t, false, tcc.tRAM.ConditionFlag)
			assert.Equal(t, tcc.expectedPc, tcc.tRAM.Pc)
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
		expectedPc uint64
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
			expectedPc: uint64(1),
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			addOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)
			assert.Equal(t, tcc.expected, tcc.tRAM.Register[tcc.r1])
			assert.Equal(t, false, tcc.tRAM.ConditionFlag)
			assert.Equal(t, tcc.expectedPc, tcc.tRAM.Pc)
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
		expectedPc uint64
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
			expectedPc: uint64(1),
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			subOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)
			assert.Equal(t, tcc.expected, tcc.tRAM.Register[tcc.r1])
			assert.Equal(t, false, tcc.tRAM.ConditionFlag)
			assert.Equal(t, tcc.expectedPc, tcc.tRAM.Pc)
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
		expectedPc uint64
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
			expectedPc: uint64(1),
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			mullOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)
			assert.Equal(t, tcc.expected, tcc.tRAM.Register[tcc.r1])
			assert.Equal(t, false, tcc.tRAM.ConditionFlag)
			assert.Equal(t, tcc.expectedPc, tcc.tRAM.Pc)
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
		expectedPc uint64
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
			expectedPc: uint64(1),
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			umulhOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)
			assert.Equal(t, tcc.expected, tcc.tRAM.Register[tcc.r1])
			assert.Equal(t, false, tcc.tRAM.ConditionFlag)
			assert.Equal(t, tcc.expectedPc, tcc.tRAM.Pc)
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
		expectedPc uint64
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
			expectedPc: uint64(1),
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			smulhOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)
			assert.Equal(t, tcc.expected, tcc.tRAM.Register[tcc.r1])
			assert.Equal(t, false, tcc.tRAM.ConditionFlag)
			assert.Equal(t, tcc.expectedPc, tcc.tRAM.Pc)
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
		expectedPc uint64
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
			expectedPc: uint64(1),
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			udivOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)
			assert.Equal(t, tcc.expected, tcc.tRAM.Register[tcc.r1])
			assert.Equal(t, false, tcc.tRAM.ConditionFlag)
			assert.Equal(t, tcc.expectedPc, tcc.tRAM.Pc)
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
		expectedPc uint64
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
			expectedPc: uint64(1),
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			umodOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)
			assert.Equal(t, tcc.expected, tcc.tRAM.Register[tcc.r1])
			assert.Equal(t, false, tcc.tRAM.ConditionFlag)
			assert.Equal(t, tcc.expectedPc, tcc.tRAM.Pc)
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
		expectedPc uint64
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
			expectedPc: uint64(1),
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
			expectedPc: uint64(1),
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			shlOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)
			assert.Equal(t, tcc.expectedValue, tcc.tRAM.Register[tcc.r1])
			assert.Equal(t, tcc.expectedFlag, tcc.tRAM.ConditionFlag)
			assert.Equal(t, tcc.expectedPc, tcc.tRAM.Pc)
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
		expectedPc uint64
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
			expectedPc: uint64(1),
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
			expectedPc: uint64(1),
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			shrOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)
			assert.Equal(t, tcc.expectedValue, tcc.tRAM.Register[tcc.r1])
			assert.Equal(t, tcc.expectedFlag, tcc.tRAM.ConditionFlag)
			assert.Equal(t, tcc.expectedPc, tcc.tRAM.Pc)
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
		expectedPc uint64
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
			expectedPc: uint64(1),
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
			expectedPc: uint64(1),
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			cmpeOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)
			assert.Equal(t, tcc.expectedValue, tcc.tRAM.Register[tcc.r1])
			assert.Equal(t, tcc.expectedFlag, tcc.tRAM.ConditionFlag)
			assert.Equal(t, tcc.expectedPc, tcc.tRAM.Pc)
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
		expectedPc uint64
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
			expectedPc: uint64(1),
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
			expectedPc: uint64(1),
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			cmpaOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)			
			assert.Equal(t, tcc.expectedFlag, tcc.tRAM.ConditionFlag)
			assert.Equal(t, tcc.expectedPc, tcc.tRAM.Pc)
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
		expectedPc uint64
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
			expectedPc: uint64(1),
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
			expectedPc: uint64(1),
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			cmpaeOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)			
			assert.Equal(t, tcc.expectedFlag, tcc.tRAM.ConditionFlag)
			assert.Equal(t, tcc.expectedPc, tcc.tRAM.Pc)
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
		expectedPc uint64
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
			expectedPc: uint64(1),
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
			expectedPc: uint64(1),
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			cmpgOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)			
			assert.Equal(t, tcc.expectedFlag, tcc.tRAM.ConditionFlag)
			assert.Equal(t, tcc.expectedPc, tcc.tRAM.Pc)
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
		expectedPc uint64
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
			expectedPc: uint64(1),
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
			expectedPc: uint64(1),
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			cmpgeOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)			
			assert.Equal(t, tcc.expectedFlag, tcc.tRAM.ConditionFlag)
			assert.Equal(t, tcc.expectedPc, tcc.tRAM.Pc)
		})
	}
}

func TestMovOperation(t *testing.T) {
	cases := []struct {
		tRAM     *tinyRAM
		r1       uint64
		r2       uint64
		r3       uint64
		expected uint64
		expectedPc uint64
	}{
		{
			tRAM: &tinyRAM{				
				Register:      []uint64{0, 0, 0},				
			},
			r1:       0,
			r2:       2,
			r3:       0,			
			expected: 2,
			expectedPc: uint64(1),
		},		
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			movOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)			
			assert.Equal(t, tcc.expected, tcc.tRAM.Register[0])
			assert.Equal(t, tcc.expectedPc, tcc.tRAM.Pc)
		})
	}
}

func TestCmovOperation(t *testing.T) {
	cases := []struct {
		tRAM     *tinyRAM
		r1       uint64
		r2       uint64
		r3       uint64
		expected uint64	
		expectedPc uint64			
	}{
		{
			tRAM: &tinyRAM{				
				Register:      []uint64{0, 0, 0},
				ConditionFlag: false,
			},
			r1:       0,
			r2:       2,
			r3:       0,			
			expected: 0,
			expectedPc: uint64(1),		
		},
		{
			tRAM: &tinyRAM{				
				Register:      []uint64{0, 0, 0},
				ConditionFlag: true,
			},
			r1:       0,
			r2:       2,
			r3:       0,			
			expected: 2,
			expectedPc: uint64(1),			
		},		
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			cmovOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)
			if tcc.tRAM.ConditionFlag {
				assert.Equal(t, tcc.expected, tcc.tRAM.Register[0])
			} else {
				assert.Equal(t, tcc.expected, tcc.tRAM.Register[0])
			}
			assert.Equal(t, tcc.expectedPc, tcc.tRAM.Pc)
		})
	}
}

func TestJmpOperation(t *testing.T) {
	cases := []struct {
		tRAM     *tinyRAM
		r1       uint64
		r2       uint64
		r3       uint64
		expectedPc uint64		
	}{
		{
			tRAM: &tinyRAM{				
				Register:      []uint64{0, 0, 0},
				Pc: 0,
			},
			r1:       5,
			r2:       0,
			r3:       0,			
			expectedPc: 5,			
		},	
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			jmpOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)
			assert.Equal(t, tcc.expectedPc, tcc.tRAM.Pc)
		})
	}
}

func TestCjmpOperation(t *testing.T) {
	cases := []struct {
		tRAM     *tinyRAM
		r1       uint64
		r2       uint64
		r3       uint64
		expectedPc uint64		
	}{
		{
			tRAM: &tinyRAM{				
				Register:      []uint64{0, 0, 0},
				ConditionFlag: false,
				Pc: 0,
			},
			r1:       5,
			r2:       0,
			r3:       0,			
			expectedPc: 1,			
		},
		{
			tRAM: &tinyRAM{				
				Register:      []uint64{0, 0, 0},
				ConditionFlag: true,
				Pc: 0,
			},
			r1:       5,
			r2:       0,
			r3:       0,			
			expectedPc: 5,			
		},	
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			cjmpOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)
			assert.Equal(t, tcc.expectedPc, tcc.tRAM.Pc)
		})
	}
}

func TestCnjmpOperation(t *testing.T) {
	cases := []struct {
		tRAM     *tinyRAM
		r1       uint64
		r2       uint64
		r3       uint64
		expectedPc uint64		
	}{
		{
			tRAM: &tinyRAM{				
				Register:      []uint64{0, 0, 0},
				ConditionFlag: false,
				Pc: 0,
			},
			r1:       5,
			r2:       0,
			r3:       0,			
			expectedPc: 5,			
		},
		{
			tRAM: &tinyRAM{				
				Register:      []uint64{0, 0, 0},
				ConditionFlag: true,
				Pc: 0,
			},
			r1:       5,
			r2:       0,
			r3:       0,			
			expectedPc: 1,			
		},	
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			cnjmpOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)
			assert.Equal(t, tcc.expectedPc, tcc.tRAM.Pc)
		})
	}
}

func TestStoreOperation(t *testing.T) {
	cases := []struct {
		tRAM *tinyRAM
		r1 uint64
		r2 uint64
		r3 uint64
		expectedValue uint64
		expectedPc uint64
	}{
		{
			tRAM: &tinyRAM{
				Register: []uint64{0, 0, 0},
				Memory: []uint64{0, 0, 0},
				Pc: 0,				
			},
			r1: 1,
			r2: 3,
			r3: 0,
			expectedValue: 3,
			expectedPc: 1,
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			storeOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)
			assert.Equal(t, tcc.expectedValue, tcc.tRAM.Memory[1])
			assert.Equal(t, tcc.expectedPc, tcc.tRAM.Pc)
		})
	}
}

func TestLoadOperation(t *testing.T) {
	cases := []struct {
		tRAM *tinyRAM
		r1 uint64
		r2 uint64
		r3 uint64
		expectedValue uint64
		expectedPc uint64
	}{
		{
			tRAM: &tinyRAM{
				Register: []uint64{0, 0, 0},
				Memory: []uint64{0, 0, 5},
				Pc: 1,				
			},
			r1: 1,
			r2: 2,
			r3: 0,
			expectedValue: 5,
			expectedPc: 2,
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			loadOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)
			assert.Equal(t, tcc.expectedValue, tcc.tRAM.Register[1])
			assert.Equal(t, tcc.expectedPc, tcc.tRAM.Pc)
		})
	}
}

func TestReadOperation(t *testing.T) {
	cases := []struct {
		tRAM *tinyRAM
		r1 uint64
		r2 uint64
		r3 uint64
		expectedValue uint64		
		expectedPc uint64
		expectedFlag bool
		expectedPrimaryInputCount uint64
		expectedAuxiliaryInputCount uint64
	}{
		{
			tRAM: &tinyRAM{
				Register: []uint64{0, 0, 0},
				Memory: []uint64{0, 0, 5},
				PrimaryInput: []uint64{3, 4, 5},
				Pc: 1,
				ConditionFlag: true,
				PrimaryInputCount: 1,
			},
			r1: 1,
			r2: 0,
			r3: 0,
			expectedValue: 4,
			expectedPc: 2,
			expectedFlag: false,
			expectedPrimaryInputCount: 2,
		},
		{
			tRAM: &tinyRAM{
				Register: []uint64{0, 0, 0},
				Memory: []uint64{0, 0, 5},
				PrimaryInput: []uint64{3, 4, 5},
				Pc: 1,
				ConditionFlag: false,
				PrimaryInputCount: 3,
			},
			r1: 1,
			r2: 0,
			r3: 0,
			expectedValue: 0,
			expectedPc: 2,
			expectedFlag: true,
			expectedPrimaryInputCount: 3,
		},
		{
			tRAM: &tinyRAM{
				Register: []uint64{0, 0, 0},
				Memory: []uint64{0, 0, 5},
				AuxiliaryInput: []uint64{3, 4, 5},
				Pc: 1,
				ConditionFlag: true,
				AuxiliaryInputCount: 2,
			},
			r1: 1,
			r2: 1,
			r3: 0,
			expectedValue: 5,
			expectedPc: 2,
			expectedFlag: false,
			expectedAuxiliaryInputCount: 3,
		},
		{
			tRAM: &tinyRAM{
				Register: []uint64{0, 0, 0},
				Memory: []uint64{0, 0, 5},
				AuxiliaryInput: []uint64{3, 4, 5},
				Pc: 1,
				ConditionFlag: false,
				AuxiliaryInputCount: 3,
			},
			r1: 1,
			r2: 1,
			r3: 0,
			expectedValue: 0,
			expectedPc: 2,
			expectedFlag: true,
			expectedAuxiliaryInputCount: 3,
		},
		{
			tRAM: &tinyRAM{
				Register: []uint64{0, 0, 0},
				Memory: []uint64{0, 0, 5},
				AuxiliaryInput: []uint64{3, 4, 5},
				Pc: 1,
				ConditionFlag: false,				
			},
			r1: 1,
			r2: 3,
			r3: 0,
			expectedValue: 0,
			expectedPc: 2,
			expectedFlag: true,			
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			readOperation(tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)
			assert.Equal(t, tcc.expectedValue, tcc.tRAM.Register[1])
			assert.Equal(t, tcc.expectedPc, tcc.tRAM.Pc)
			assert.Equal(t, tcc.expectedFlag, tcc.tRAM.ConditionFlag)
			assert.Equal(t, tcc.expectedPrimaryInputCount, tcc.tRAM.PrimaryInputCount)
			assert.Equal(t, tcc.expectedAuxiliaryInputCount, tcc.tRAM.AuxiliaryInputCount)
		})
	}
}

