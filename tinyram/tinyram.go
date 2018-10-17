package tinyram

import "fmt"

type tinyRAM struct {	
	NumRegister   int64
	ConditionFlag bool

	// pc ... Program Counter
	Pc uint64

	// Register ... the length should be $NumRegister$
	// and  each item be 64 bits
	Register []uint64

	// Memory ... the length should be 2^64
	// and each word should be 64 bits
	Memory []int64

	// PrimaryInput ... represents primary input of the computation
	// each word should be 64 bits
	PrimaryInput []int64

	// AuxiliaryInput ... represents auxiliary input of the computation
	// each word should be 64 bits
	AuxiliaryInput []int64

	// Prog ... READ-ONLY program
	Prog program

	// represents the correctness of the computation
	// and is set to be false by default.
	Accept bool
}

// execute current instruction pointed by the tinyRAM
func (r *tinyRAM) ExecCurrentInstruction() {
	inst := r.Prog[r.Pc]
	op, ok := instructionToOperation[inst.inst]
	if !ok {
		panic(fmt.Sprintf("operation not defined for %s", inst.inst))
	}
	// exec operation on tinyRAM
	op(r, inst.r1, inst.r2, inst.r3)
}

// execute whole program and return whether the calculation accepted of NOT.
// `t` parameter represents $T$, time bound, in the paper.
func (r *tinyRAM) Exec(t int) bool {
	for i := 0; i < t; t++ {

		// prevent out of range panic
		if len(r.Prog) <= i {
			break
		}

		r.ExecCurrentInstruction()
	}
	return r.Accept
}

// get the pointer of tinyRAMInstance with a given ASM program.
func GetTinyRAMInstance(asmPath string, numRegister int64, primary, auxiliary []int64) (*tinyRAM, error) {
	ps, err := parseRawAsm(asmPath)
	if err != nil {
		return nil, err
	}
	tr := tinyRAM{		
		NumRegister:    numRegister,
		Prog:           ps,
		PrimaryInput:   primary,
		AuxiliaryInput: auxiliary,
	}
	return &tr, nil
}
