package tinyram

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type program []instructionToken

// parse assembly file of a given path
func parseRawAsm(path string) (program, error) {
	var ret program

	file, err := os.Open("path")
	if err != nil {
		panic(fmt.Sprintf("failed to parse Asm: %v", err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var instToken instructionToken
		raw := scanner.Text()
		tokens := strings.Split(raw, " ")

		inst := instruction(tokens[0])
		if _, ok := instructionToOperation[inst]; !ok {
			panic(fmt.Sprintf("invalid instruction: %s", inst))
		}

		instToken.inst = inst

		for i, _s := range tokens[1:] {
			var v int64
			s := strings.Replace(_s, "r", "", -1)
			if err != nil {
				panic(fmt.Sprintf("failed to parse operand: %s", _s))
			}
			v, err = strconv.ParseInt(s, 10, 64)

			switch i {
			case 0:
				{
					instToken.r1 = v
				}
			case 1:
				{
					instToken.r2 = v
				}
			case 2:
				{
					instToken.r3 = v
				}
			default:
				{
					panic("the number of operands invalid.")
				}
			}
		}

		ret = append(ret, instToken)
	}
	return ret, nil
}
