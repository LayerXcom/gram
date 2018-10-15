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
// `t` parameter represents $T$, time bound, in the paper.
func (r *tinyRAMInstance) Exec(t int) bool {
	for i := 0; i < t; t++ {
		if len(r.prog) <= t {
			break
		}
		inst := r.prog[r.ram.Pc]

		if inst.inst == instruction("RETURN") {
			// TODO: check the target value of RETURN operation
			return true
		} else {
			r.ExecCurrentInstruction()
		}
	}
	return false
}

// get the pointer of tinyRAMInstance with a given ASM program.
func GetTinyRAMInstance(asmPath string, wordSize, numRegister int64) (*tinyRAMInstance, error) {
	ps, err := parseRawAsm(asmPath)
	if err != nil {
		return nil, err
	}
	tr := tinyRAM{
		WordSize:wordSize,
		NumRegister: numRegister,
	}
	return &tinyRAMInstance{
		ram: tr,
		prog: ps,
	}, nil
}
