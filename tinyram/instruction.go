package tinyram

type instruction string

// TODO: list all the instruction used by TinyRAM

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

var instructionToOperation = map[instruction]func(tRam *tinyRAM, r1, r2, r3 int64){
	AND:    andOperation,
	OR:     orOperation,
	XOR:    xorOperation,
	NOT:    notOperation,
	ADD:    addOperatiopn,
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
	r1   int64
	r2   int64
	r3   int64
}

// TODO: implement the behavior of all instructions
func andOperation(tRAM *tinyRAM, r1, r2, r3 int64) {
	tRAM.Register[r1] = tRAM.Register[r2] & r3
	if tRAM.Register[r1] == 0 {
		tRAM.ConditionFlag = true
	} else if tRAM.Register[r1] != 0 {
		tRAM.ConditionFlag = false
	}
}

func orOperation(tRAM *tinyRAM, r1, r2, r3 int64) {
	tRAM.Register[r1] = tRAM.Register[r2] | r3
	if tRAM.Register[r1] == 0 {
		tRAM.ConditionFlag = true
	} else if tRAM.Register[r1] != 0 {
		tRAM.ConditionFlag = false
	}
}

func xorOperation(tRAM *tinyRAM, r1, r2, r3 int64) {
	tRAM.Register[r1] = tRAM.Register[r2] ^ r3
	if tRAM.Register[r1] == 0 {
		tRAM.ConditionFlag = true
	} else if tRAM.Register[r1] != 0 {
		tRAM.ConditionFlag = false
	}
}

func notOperation(tRAM *tinyRAM, r1, r2, r3 int64) {
	result := tRAM.Register[r1] &^ r2
	if result == 0 {
		tRAM.ConditionFlag = true
	} else if result != 0 {
		tRAM.ConditionFlag = false
	}
}

func addOperatiopn(tRAM *tinyRAM, r1, r2, r3 int64) {

}

func subOperation(tRAM *tinyRAM, r1, r2, r3 int64) {

}

func mullOperation(tRAM *tinyRAM, r1, r2, r3 int64) {

}

func umulhOperation(tRAM *tinyRAM, r1, r2, r3 int64) {

}

func smulhOperation(tRAM *tinyRAM, r1, r2, r3 int64) {

}

func udivOperation(tRAM *tinyRAM, r1, r2, r3 int64) {

}

func umodOperation(tRAM *tinyRAM, r1, r2, r3 int64) {

}

func shlOperation(tRAM *tinyRAM, r1, r2, r3 int64) {

}

func shrOperation(tRAM *tinyRAM, r1, r2, r3 int64) {

}

func cmpeOperation(tRAM *tinyRAM, r1, r2, r3 int64) {

}

func cmpaOperation(tRAM *tinyRAM, r1, r2, r3 int64) {

}

func cmpaeOperation(tRAM *tinyRAM, r1, r2, r3 int64) {

}

func cmpgOperation(tRAM *tinyRAM, r1, r2, r3 int64) {

}

func cmpgeOperation(tRAM *tinyRAM, r1, r2, r3 int64) {

}

func movOperation(tRAM *tinyRAM, r1, r2, r3 int64) {

}

func cmovOperation(tRAM *tinyRAM, r1, r2, r3 int64) {

}

func jmpOperation(tRAM *tinyRAM, r1, r2, r3 int64) {

}

func cjmpOperation(tRAM *tinyRAM, r1, r2, r3 int64) {

}

func cnjmpOperation(tRAM *tinyRAM, r1, r2, r3 int64) {

}

func storeOperation(tRAM *tinyRAM, r1, r2, r3 int64) {

}

func loadOperation(tRAM *tinyRAM, r1, r2, r3 int64) {

}

func readOperation(tRAM *tinyRAM, r1, r2, r3 int64) {

}

func answerOperation(tRAM *tinyRAM, r1, r2, r3 int64) {

}
