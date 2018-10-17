package tinyram

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestParseRawAsm(t *testing.T) {
	cases := []struct {
		path     string
		expected program
	}{{
		path: "../example/testForParser.asm",
		expected: program{
			instructionToken{instruction("MOV"), 1, 2, 100},
			instructionToken{instruction("ADD"), 100, 100, 100},
			instructionToken{instruction("SUB"), 100, 2, 100},
			instructionToken{instruction("READ"), 100, 200, 3},
			instructionToken{instruction("STORE"), 10, 20, 30},
			instructionToken{instruction("UMULH"), 10, 300, 0},
			instructionToken{instruction("SHR"), 1, 0, 0},
		},
	}}

	for n, tc := range cases {
		tcc := tc
		t.Run(fmt.Sprintf("%d-th unit test", n), func(t *testing.T) {
			actual, err := parseRawAsm(tcc.path)
			if err != nil {
				t.Fatalf("parseRawAsm failed")
			}

			// check the length
			assert.Equal(t, len(tcc.expected), len(actual))

			for i := range tcc.expected {
				assert.Equal(t, tcc.expected[i], actual[i])
			}
		})
	}
}
