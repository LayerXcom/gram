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

var instructionToOperation = map[instruction]func(tRam *tinyRAM, r1, r2, r3 operand){
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
	r1   operand
	r2   operand
	r3   operand
}

type operand struct {
	isIndex bool
	value   int64
}

// check whether a given instructionToken is valid or NOT.
func validateToken(instr instruction) bool {
	// TODO: implement
	return true
}

// TODO: implement the behavior of all instructions
func andOperation(tRAM *tinyRAM, r1, r2, r3 operand) {
	tRAM.Register[r1.value] = tRAM.Register[r2.value] & tRAM.Register[r3.value]
	if tRAM.WordSize == 0 {
		tRAM.ConditionFlag = true
	} else if tRAM.WordSize != 0 {
		tRAM.ConditionFlag = false
	}
}

func orOperation(tRAM *tinyRAM, r1, r2, r3 operand) {

}

func xorOperation(tRAM *tinyRAM, r1, r2, r3 operand) {

}

func notOperation(tRAM *tinyRAM, r1, r2, r3 operand) {

}

func addOperation(tRAM *tinyRAM, r1, r2, r3 operand) {

}

func subOperation(tRAM *tinyRAM, r1, r2, r3 operand) {

}

func mullOperation(tRAM *tinyRAM, r1, r2, r3 operand) {

}

func umulhOperation(tRAM *tinyRAM, r1, r2, r3 operand) {

}

func smulhOperation(tRAM *tinyRAM, r1, r2, r3 operand) {

}

func udivOperation(tRAM *tinyRAM, r1, r2, r3 operand) {

}

func umodOperation(tRAM *tinyRAM, r1, r2, r3 operand) {

}

func shlOperation(tRAM *tinyRAM, r1, r2, r3 operand) {

}

func shrOperation(tRAM *tinyRAM, r1, r2, r3 operand) {

}

func cmpeOperation(tRAM *tinyRAM, r1, r2, r3 operand) {

}

func cmpaOperation(tRAM *tinyRAM, r1, r2, r3 operand) {

}

func cmpaeOperation(tRAM *tinyRAM, r1, r2, r3 operand) {

}

func cmpgOperation(tRAM *tinyRAM, r1, r2, r3 operand) {

}

func cmpgeOperation(tRAM *tinyRAM, r1, r2, r3 operand) {

}

func movOperation(tRAM *tinyRAM, r1, r2, r3 operand) {

}

func cmovOperation(tRAM *tinyRAM, r1, r2, r3 operand) {

}

func jmpOperation(tRAM *tinyRAM, r1, r2, r3 operand) {

}

func cjmpOperation(tRAM *tinyRAM, r1, r2, r3 operand) {

}

func cnjmpOperation(tRAM *tinyRAM, r1, r2, r3 operand) {

}

func storeOperation(tRAM *tinyRAM, r1, r2, r3 operand) {

}

func loadOperation(tRAM *tinyRAM, r1, r2, r3 operand) {

}

func readOperation(tRAM *tinyRAM, r1, r2, r3 operand) {

}

func answerOperation(tRAM *tinyRAM, r1, r2, r3 operand) {

}
