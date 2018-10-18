package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/LayerXcom/gram/tinyram"
)

type rawParams struct {
	path,
	primaryInput,
	auxiliaryInput string
	numRegister uint64
	t           int
}

func main() {
	var ps rawParams
	flag.StringVar(&ps.path, "path", "", "path for .asm file: example.asm")
	flag.StringVar(&ps.primaryInput, "pi", "", "comma separated auxiliaryInput (type []uint64). Example: '1,100,200'")
	flag.StringVar(&ps.auxiliaryInput, "ai", "", "comma separated auxiliaryInput (type []uint64). Example: '1,100,200'")
	flag.Uint64Var(&ps.numRegister, "r", 0, "the number of register")
	flag.IntVar(&ps.t, "t", 0, "time bound on execution:")
	flag.Parse()

	if ext := strings.Split(ps.path, ".")[1]; ext != "asm" {
		fmt.Printf("invalid file extention: %s != .asm", ext)
		return
	}

	var pi, ai []uint64
	if ps.primaryInput != "" {
		for _, raw := range strings.Split(ps.primaryInput, ",") {
			input, err := strconv.ParseUint(raw, 10, 64)
			if err != nil {
				fmt.Printf("failed to parse %s in primaryInput to uint64: %v", raw, err)
				return
			}
			pi = append(pi, input)
		}
	}

	if ps.auxiliaryInput != "" {
		for _, raw := range strings.Split(ps.auxiliaryInput, ",") {
			input, err := strconv.ParseUint(raw, 10, 64)
			if err != nil {
				fmt.Printf("failed to parse %s in auxiliaryInput to uint64: %v", raw, err)
				return
			}
			ai = append(ai, input)
		}
	}

	ram, err := tinyram.GetTinyRAMInstance(ps.path, ps.numRegister, pi, ai)
	if err != nil {
		fmt.Println("tinyram.GetTinyRAMInstancef failed: ", err)
		return
	}

	fmt.Println("Execution Result: ", ram.Exec(ps.t))
}
