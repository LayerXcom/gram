package tinyram

type tinyRAM struct {
	// TODO: implement
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

// get the pointer of tinyRAMInstance with a given ASM program.
// TODO: pass tinyRAM parameters
func GetTinyRAMInstance(asmPath string) (*tinyRAMInstance, error) {
	// TODO: implement
	return nil, nil
}
