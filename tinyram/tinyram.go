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
}

type tinyRAMInstance struct {
	ram  tinyRAM
	prog program
}

// execute current instruction pointed by the tinyRAM
func (r *tinyRAMInstance) ExecCurrentInstruction() {
	// TODO: implement
}

// execute whole program and return whether the calculation accepted of NOT.
func (r *tinyRAMInstance) Exec() bool {
	// TODO: implement
	return false
}
