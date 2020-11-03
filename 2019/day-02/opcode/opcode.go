// Package opcode contains functions for computing intcodes per
// https://adventofcode.com/2019/day/2.
package opcode

import "fmt"

// ComputeOpcode interprets intcode as an int slice based on opcodes
// contained in it, and returns the modified intcode as an int slice.
func ComputeOpcode(code []int) []int {
	for i := 0; code[i] != 99; i += 4 {
		switch code[i] {
		case 1:
			code[code[i+3]] = code[code[i+1]] + code[code[i+2]]
		case 2:
			code[code[i+3]] = code[code[i+1]] * code[code[i+2]]
		default:
			fmt.Errorf("Invalid opcode '%d'\n", code[i])
		}
	}
	return code
}
