package tinyram

type tinyRAM struct {
	WordSize      int64
	NumRegister   int64
	ConditionFlag bool

	// pc ... Program Counter
	Pc int64

	// Register ... the length should be $NumRegister$
	// and  each item be $WordSize$-bits
	Register []int64

	// Memory ... the length should be $2^WordSize$
	// and each word should be $WordSize$-bits
	Memory []int64

	// PrimaryInput ... represents primary input of the computation
	// each word should be $WordSize$-bits
	PrimaryInput []int64

	// AuxiliaryInput ... represents auxiliary input of the computation
	// each word should be $WordSize$-bits
	AuxiliaryInput []int64

	// Prog ... READ-ONLY program
	Prog  program

	// represents the correctness of the computation
	// and is set to be false by default.
	Accept  bool
}

// execute current instruction pointed by the tinyRAM
func (r *tinyRAM) ExecCurrentInstruction() {
	// TODO: implement
}

// execute whole program and return whether the calculation accepted of NOT.
func (r *tinyRAM) Exec() bool {
	// TODO: implement
	return false
}

// get the pointer of tinyRAMInstance with a given ASM program.
// TODO: pass tinyRAM parameters
func GetTinyRAMInstance(asmPath string) (*tinyRAM, error) {
	// TODO: implement
	return nil, nil
}
