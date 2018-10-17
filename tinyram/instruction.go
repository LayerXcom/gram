package tinyram

import (
	"math"
)

type instruction string

const (
	AND    instruction = "AND"
	OR     instruction = "OR"
	XOR    instruction = "XOR"
	NOT    instruction = "NOT"
	ADD    instruction = "ADD"
	SUB    instruction = "SUB"
	MULL   instruction = "MULL"
	UMULH  instruction = "UMULH"
	SMULH  instruction = "SMULH"
	UDIV   instruction = "UDIV"
	UMOD   instruction = "UMOD"
	SHL    instruction = "SHL"
	SHR    instruction = "SHR"
	CMPE   instruction = "SMPE"
	CMPA   instruction = "CMPA"
	CMPAE  instruction = "CMPAE"
	CMPG   instruction = "CMPG"
	CMPGE  instruction = "CMPGE"
	MOV    instruction = "MOV"
	CMOV   instruction = "CMOV"
	JMP    instruction = "JMP"
	CJMP   instruction = "CJMP"
	CNJMP  instruction = "CNJMP"
	STORE  instruction = "STORE"
	LOAD   instruction = "LOAD"
	READ   instruction = "READ"
	ANSWER instruction = "ANSWER"
)

var instructionToOperation = map[instruction]func(tRam *tinyRAM, r1, r2, r3 uint64){
	AND:    andOperation,
	OR:     orOperation,
	XOR:    xorOperation,
	NOT:    notOperation,
	ADD:    addOperation,
	SUB:    subOperation,
	MULL:   mullOperation,
	UMULH:  umulhOperation,
	SMULH:  smulhOperation,
	UDIV:   udivOperation,
	UMOD:   umodOperation,
	SHL:    shlOperation,
	SHR:    shrOperation,
	CMPE:   cmpeOperation,
	CMPA:   cmpaOperation,
	CMPAE:  cmpaeOperation,
	CMPG:   cmpgOperation,
	CMPGE:  cmpgeOperation,
	MOV:    movOperation,
	CMOV:   cmovOperation,
	JMP:    jmpOperation,
	CJMP:   cjmpOperation,
	CNJMP:  cnjmpOperation,
	STORE:  storeOperation,
	LOAD:   loadOperation,
	READ:   readOperation,
	ANSWER: answerOperation,
}

type instructionToken struct {
	inst instruction
	r1   uint64
	r2   uint64
	r3   uint64
}

//
// Bit operations
//

func andOperation(tRAM *tinyRAM, r1, r2, r3 uint64) {
	tRAM.Register[r1] = tRAM.Register[r2] & r3	
	if tRAM.Register[r1] == 0 {
		tRAM.ConditionFlag = true
	} else {
		tRAM.ConditionFlag = false
	}
	tRAM.Pc++
}

func orOperation(tRAM *tinyRAM, r1, r2, r3 uint64) {
	tRAM.Register[r1] = tRAM.Register[r2] | r3
	if tRAM.Register[r1] == 0 {
		tRAM.ConditionFlag = true
	} else {
		tRAM.ConditionFlag = false
	}
	tRAM.Pc++
}

func xorOperation(tRAM *tinyRAM, r1, r2, r3 uint64) {
	tRAM.Register[r1] = tRAM.Register[r2] ^ r3
	if tRAM.Register[r1] == 0 {
		tRAM.ConditionFlag = true
	} else {
		tRAM.ConditionFlag = false
	}
	tRAM.Pc++
}

func notOperation(tRAM *tinyRAM, r1, r2, r3 uint64) {
	tRAM.Register[r1] = r2 ^ math.MaxUint64
	if tRAM.Register[r1] == 0 {
		tRAM.ConditionFlag = true
	} else {
		tRAM.ConditionFlag = false
	}
	tRAM.Pc++
}

//
// Integer operations
//

func addOperation(tRAM *tinyRAM, r1, r2, r3 uint64) {
	tRAM.Register[r1] = tRAM.Register[r2] + r3
	if (tRAM.Register[r1] >> 63) & 1 == 1 {
		tRAM.ConditionFlag = true
	} else {
		tRAM.ConditionFlag = false
	}
	tRAM.Pc++
}

func subOperation(tRAM *tinyRAM, r1, r2, r3 uint64) {	
	tRAM.Register[r1] = uint64(math.Abs(float64(tRAM.Register[r2] - r3)))
	if tRAM.Register[r2] <= r3 {		
		tRAM.ConditionFlag = true
	} else {		
		tRAM.ConditionFlag = false
	}
	tRAM.Pc++
}

func mullOperation(tRAM *tinyRAM, r1, r2, r3 uint64) {
	tRAM.Register[r1] = tRAM.Register[r2] * r3
	if math.IsInf(float64(tRAM.Register[r1]), 1) {
		tRAM.ConditionFlag = true
	} else {
		tRAM.ConditionFlag = false
	}
	tRAM.Pc++
}

func umulhOperation(tRAM *tinyRAM, r1, r2, r3 uint64) {
	tRAM.Register[r1] = tRAM.Register[r2] * r3
	if math.IsInf(float64(tRAM.Register[r1]), 1) {
		tRAM.ConditionFlag = true
	} else {
		tRAM.ConditionFlag = false
	}
	tRAM.Pc++
}

func smulhOperation(tRAM *tinyRAM, r1, r2, r3 uint64) {
	tRAM.Register[r1] = tRAM.Register[r2] * r3
	if math.IsInf(float64(tRAM.Register[r1]), 1) {
		tRAM.ConditionFlag = true
	} else {
		tRAM.ConditionFlag = false
	}
	tRAM.Pc++
}

func udivOperation(tRAM *tinyRAM, r1, r2, r3 uint64) {
	if r3 == 0 {
		tRAM.ConditionFlag = true
	} else {
		tRAM.Register[r1] = tRAM.Register[r2] / r3
		tRAM.ConditionFlag = false
	}
	tRAM.Pc++
}

func umodOperation(tRAM *tinyRAM, r1, r2, r3 uint64) {
	if r3 == 0 {
		tRAM.ConditionFlag = true
	} else {
		tRAM.Register[r1] = tRAM.Register[r2] % r3
		tRAM.ConditionFlag = false
	}
	tRAM.Pc++
}

//
// Shift operations
//

func shlOperation(tRAM *tinyRAM, r1, r2, r3 uint64) {
	tRAM.Register[r1] = tRAM.Register[r2] << r3
	if (tRAM.Register[r1] >> 63) & 1 == 1 {
		tRAM.ConditionFlag = true
	} else {
		tRAM.ConditionFlag = false
	}
	tRAM.Pc++
}

func shrOperation(tRAM *tinyRAM, r1, r2, r3 uint64) {
	tRAM.Register[r1] = tRAM.Register[r2] >> r3
	if tRAM.Register[r1] & 1 == 1 {
		tRAM.ConditionFlag = true
	} else {
		tRAM.ConditionFlag = false
	}
	tRAM.Pc++
}

//
// Compare operations
//

func cmpeOperation(tRAM *tinyRAM, r1, r2, r3 uint64) {
	if r1 == r2 {
		tRAM.ConditionFlag = true
	} else {
		tRAM.ConditionFlag = false
	}
	tRAM.Pc++
}

func cmpaOperation(tRAM *tinyRAM, r1, r2, r3 uint64) {
	if r1 > r2 {
		tRAM.ConditionFlag = true
	} else {
		tRAM.ConditionFlag = false
	}
	tRAM.Pc++
}

func cmpaeOperation(tRAM *tinyRAM, r1, r2, r3 uint64) {
	if r1 >= r2 {
		tRAM.ConditionFlag = true
	} else {
		tRAM.ConditionFlag = false
	}
	tRAM.Pc++
}

func cmpgOperation(tRAM *tinyRAM, r1, r2, r3 uint64) {
	if r1 > r2 {
		tRAM.ConditionFlag = true
	} else {
		tRAM.ConditionFlag = false
	}
	tRAM.Pc++
}

func cmpgeOperation(tRAM *tinyRAM, r1, r2, r3 uint64) {
	if r1 >= r2 {
		tRAM.ConditionFlag = true
	} else {
		tRAM.ConditionFlag = false
	}
	tRAM.Pc++
}

//
// Move operations
//

func movOperation(tRAM *tinyRAM, r1, r2, r3 uint64) {
	tRAM.Register[r1] = r2
	tRAM.Pc++
}

func cmovOperation(tRAM *tinyRAM, r1, r2, r3 uint64) {
	if tRAM.ConditionFlag {
		tRAM.Register[r1] = r2
	}
	tRAM.Pc++
}

//
// Jump operations
//

func jmpOperation(tRAM *tinyRAM, r1, r2, r3 uint64) {
	tRAM.Pc = r1
}

func cjmpOperation(tRAM *tinyRAM, r1, r2, r3 uint64) {
	if tRAM.ConditionFlag {
		tRAM.Pc = r1
	} else {
		tRAM.Pc++
	}

}

func cnjmpOperation(tRAM *tinyRAM, r1, r2, r3 uint64) {
	if !tRAM.ConditionFlag {
		tRAM.Pc = r1
	} else {
		tRAM.Pc++
	}
}

//
// Memory operations
//

func storeOperation(tRAM *tinyRAM, r1, r2, r3 uint64) {
	tRAM.Memory[r1] = r2
	tRAM.Pc++
}

func loadOperation(tRAM *tinyRAM, r1, r2, r3 uint64) {
	tRAM.Register[r1] = tRAM.Memory[r2]
	tRAM.Pc++
}

//
// Input operations
//

func readOperation(tRAM *tinyRAM, r1, r2, r3 uint64) {
	if r2 == 0 {
		if tRAM.PrimaryInputCount <= uint64(len(tRAM.PrimaryInput)) - 1 {
			tRAM.Register[r1] = tRAM.PrimaryInput[tRAM.PrimaryInputCount]
			tRAM.PrimaryInputCount++
			tRAM.ConditionFlag = false
		} else {
			tRAM.Register[r1] = 0
			tRAM.ConditionFlag = true
		}				
	} else if r2 == 1 {
		if tRAM.AuxiliaryInputCount <= uint64(len(tRAM.AuxiliaryInput)) - 1 {
			tRAM.Register[r1] = tRAM.AuxiliaryInput[tRAM.AuxiliaryInputCount]
			tRAM.AuxiliaryInputCount++
			tRAM.ConditionFlag = false
		} else {
			tRAM.Register[r1] = 0
			tRAM.ConditionFlag = true
		}
	} else {
		tRAM.Register[r1] = 0
		tRAM.ConditionFlag = true
	}
	tRAM.Pc++
}

//
// Answer operations
//

// The program accepted if the return value is 0
func answerOperation(tRAM *tinyRAM, r1, r2, r3 uint64) {
	// return r1
}
