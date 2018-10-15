package tinyram

type instruction string

// TODO: list all the instruction used by TinyRAM

type instructionToken struct {
	inst instruction
	r1   int64
	r2   int64
}

// check whether a given instructionToken is valid or NOT.
func validateToken(instr instruction) bool {
	// TODO: implement
	return true
}

// TODO: implement the behavior of all instructions
