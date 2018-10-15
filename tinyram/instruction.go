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

var instructionToOperation = map[instruction]func(tRam *tinyRAM, r1, r2, r3, int64) {
	AND: addOperation,
	OR: orOperation,     
	XOR: xorOperation,    
	NOT: notOperation,    
	ADD: addOperatiopn,    
	SUB: subOperation,    
	MULL: mullOperation,   
	UMULH: umulhOperation,  
	SMULH: smulhOperation,  
	UDIV: udivOperation,   
	UMOD: umodOperation,   
	SHL: shlOperation,    
	SHR: shrOperation,    
	CMPE: cmpeOperation,   
	CMPA: cmpaOperation,   
	CMPAE: cmpaeOperation,  
	CMPG: cmpgOperation,   
	CMPGE: cmpgeOperation,  
	MOV: movOperation,    
	CMOV: cmovOperation,   
	JMP: jmpOperation,    
	CJMP: cjmpOperation,   
	CNJMP: cnjmpOperation,  
	STORE: storeOperation,  
	LOAD: loadOperation,   
	READ: readOperation,   
	ANSWER: answerOperation,
}

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