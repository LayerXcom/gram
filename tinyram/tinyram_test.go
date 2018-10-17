package tinyram

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

// execute the program noted in https://www.scipr-lab.org/doc/TinyRAM-spec-0.991.pdf
func TestExecution(t *testing.T) {
	cases := []struct {
		path string
		timeBound int
		expectedFlag bool
	}{
		{
			path: "../example/testForSuccess.asm",
			timeBound: 9,
			expectedFlag: true,
		},
		{
			path: "../example/testForFail.asm",
			timeBound: 8,	
			expectedFlag: false,		
		},
	}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {			
			tRAM, err := GetTinyRAMInstance(tcc.path, 64, []uint64{2}, []uint64{2})
			if err != nil {
				t.Fatalf("get the tinyRAM instnce failed")
			}

			tRAM.Exec(tcc.timeBound)
			assert.Equal(t, tcc.expectedFlag, tRAM.Accept)						
		})
	}
}
