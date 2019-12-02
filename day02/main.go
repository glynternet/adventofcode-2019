package main

import (
	"fmt"
	"log"
)

// input to task is the memory of a computer
type memory []int64

func main() {
	for i, tc := range []struct {
		in  memory
		out memory
	}{
		{
			in:  memory{1, 0, 0, 0, 99},
			out: memory{2, 0, 0, 0, 99},
		},
		{
			in:  memory{2, 3, 0, 3, 99},
			out: memory{2, 3, 0, 6, 99},
		},
		{
			in:  memory{2, 4, 4, 5, 99, 0},
			out: memory{2, 4, 4, 5, 99, 9801},
		},
		{
			in:  memory{1, 1, 1, 4, 99, 5, 6, 0, 99},
			out: memory{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	} {
		out := tc.in.runUntilHalt()
		for address := range out {
			expected := tc.out.getValue(address)
			actual := out.getValue(address)
			if actual != expected {
				log.Printf("case:%d unexpected values at address %d: expected:%d, actual:%d, output:%+v",
					//log.Panicf("case:%d unexpected values at address %d: expected:%d, actual:%d, output:%+v",
					i, address, expected, actual, out)
			}
		}
	}
	in := memory{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 9, 1, 19, 1, 5, 19, 23, 2, 9, 23, 27, 1, 27, 5, 31, 2, 31, 13, 35, 1, 35, 9, 39, 1, 39, 10, 43, 2, 43, 9, 47, 1, 47, 5, 51, 2, 13, 51, 55, 1, 9, 55, 59, 1, 5, 59, 63, 2, 6, 63, 67, 1, 5, 67, 71, 1, 6, 71, 75, 2, 9, 75, 79, 1, 79, 13, 83, 1, 83, 13, 87, 1, 87, 5, 91, 1, 6, 91, 95, 2, 95, 13, 99, 2, 13, 99, 103, 1, 5, 103, 107, 1, 107, 10, 111, 1, 111, 13, 115, 1, 10, 115, 119, 1, 9, 119, 123, 2, 6, 123, 127, 1, 5, 127, 131, 2, 6, 131, 135, 1, 135, 2, 139, 1, 139, 9, 0, 99, 2, 14, 0, 0}
	part01(in)
	part02(in)
}

func part01(in memory) {
	restored := in.copyAndRestore(12, 2)
	done := restored.runUntilHalt()
	outVar := done.getOutputVariable()
	expected := int64(5110675)
	if outVar != expected {
		log.Panicf("Expected part01 to equal %d but got %d", expected, outVar)
	}
	fmt.Println(outVar)
}

func part02(in memory) {
	var found bool
	for noun := int64(0); noun <= 99 && !found; noun++ {
		for verb := int64(0); verb <= 99 && !found; verb++ {
			restored := in.copyAndRestore(noun, verb)
			done := restored.runUntilHalt()
			if done.getOutputVariable() == 19690720 {
				log.Printf("Noun: %d, Verb: %d", noun, verb)
				answer := 100*noun + verb
				log.Printf("100 * %d + %d = %d", noun, verb, answer)
				expected := int64(4847)
				if expected != answer {
					log.Panicf("Expected part02 to equal %d but got %d", expected, answer)
				}
				found = true
			}
		}
	}

}

func (in memory) copyAndRestore(noun, verb int64) memory {
	out := make(memory, len(in))
	copy(out, in)
	out[1] = noun
	out[2] = verb
	return out
}

func (in memory) runUntilHalt() memory {
	var cursor int
	for {
		code := opcode(in.getValue(cursor))
		if code == 99 {
			return in
		}

		getOp, ok := ops[code]
		if !ok {
			log.Panicf("invalid value at cursor %d %d", cursor, code)
		}
		operation, paramCount := getOp(&in, cursor)
		operation(&in)
		cursor += 1 + paramCount
	}
}

type opcode int
type operation func(*memory)

var ops = map[opcode]func(*memory, int) (operation, int){
	1: (*memory).getAdditionOp,
	2: (*memory).getMultiplicationOp,
}

func (in memory) getAdditionOp(cursor int) (operation, int) {
	lhs := in.getValue(cursor + 1)
	rhs := in.getValue(cursor + 2)
	outPos := in.getValue(cursor + 3)
	return func(m *memory) {
		in[outPos] = in[lhs] + in[rhs]
	}, 3
}

func (in memory) getMultiplicationOp(cursor int) (operation, int) {
	lhs := in.getValue(cursor + 1)
	rhs := in.getValue(cursor + 2)
	outPos := in.getValue(cursor + 3)
	return func(m *memory) {
		in[outPos] = in[lhs] * in[rhs]
	}, 3
}

func (in memory) getValue(cursor int) int64 {
	val := in[cursor]
	return val
}

func (in memory) getOutputVariable() int64 {
	return in[0]
}
