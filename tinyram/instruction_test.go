package tinyram

import (
	"fmt"
	"testing"
	
	"gotest.tools/assert"
)

func TestAndOperation(t *testing.T) {
	cases := []struct {
		tRAM *tinyRAM
		r1  int64
		r2  int64
		r3  int64
		expected int64
	}
	{
		{
			tRAM: &tinyRAM{					
				Register: []int64{0,0,3},
				ConditionFlag: true,		
			},
			r1: 0,
			r2: 2,
			r3: 10,
			expected: 3 & 10,
		}
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("$d-th unit test", n), func(t *testing.T) {
			andOperation(*tcc.tRAM, tcc.r1, tcc.r2, tcc.r3)
			assert.Equal(t, tcc.expected, *tcc.tRAM.Register[r1])
			assert.Equal(t, false, tcc.tRAM.ConditionFlag)
		})
	}
}
