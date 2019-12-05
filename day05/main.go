package main

import (
	"log"
)

func main() {
	test00()
	test01()

	input := []int64{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1102, 17, 65, 225, 102, 21, 95, 224, 1001, 224, -1869, 224, 4, 224, 1002, 223, 8, 223, 101, 7, 224, 224, 1, 224, 223, 223, 101, 43, 14, 224, 1001, 224, -108, 224, 4, 224, 102, 8, 223, 223, 101, 2, 224, 224, 1, 223, 224, 223, 1101, 57, 94, 225, 1101, 57, 67, 225, 1, 217, 66, 224, 101, -141, 224, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 1, 224, 1, 224, 223, 223, 1102, 64, 34, 225, 1101, 89, 59, 225, 1102, 58, 94, 225, 1002, 125, 27, 224, 101, -2106, 224, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 5, 224, 1, 224, 223, 223, 1102, 78, 65, 225, 1001, 91, 63, 224, 101, -127, 224, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 3, 224, 1, 223, 224, 223, 1102, 7, 19, 224, 1001, 224, -133, 224, 4, 224, 102, 8, 223, 223, 101, 6, 224, 224, 1, 224, 223, 223, 2, 61, 100, 224, 101, -5358, 224, 224, 4, 224, 102, 8, 223, 223, 101, 3, 224, 224, 1, 224, 223, 223, 1101, 19, 55, 224, 101, -74, 224, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 1, 224, 1, 224, 223, 223, 1101, 74, 68, 225, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 107, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 329, 1001, 223, 1, 223, 1008, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 344, 1001, 223, 1, 223, 7, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 359, 1001, 223, 1, 223, 8, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 374, 1001, 223, 1, 223, 1007, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 389, 101, 1, 223, 223, 8, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 404, 101, 1, 223, 223, 1108, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 419, 1001, 223, 1, 223, 1108, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 434, 101, 1, 223, 223, 1108, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 449, 101, 1, 223, 223, 1008, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 464, 101, 1, 223, 223, 7, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 479, 101, 1, 223, 223, 108, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 494, 101, 1, 223, 223, 107, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 509, 101, 1, 223, 223, 107, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 524, 1001, 223, 1, 223, 1107, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 539, 101, 1, 223, 223, 1008, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 554, 1001, 223, 1, 223, 8, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 569, 101, 1, 223, 223, 1007, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 584, 1001, 223, 1, 223, 1107, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 599, 101, 1, 223, 223, 7, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 614, 101, 1, 223, 223, 108, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 629, 1001, 223, 1, 223, 108, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 644, 101, 1, 223, 223, 1007, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 659, 101, 1, 223, 223, 1107, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 674, 1001, 223, 1, 223, 4, 223, 99, 226}
	part01(input)
}

func part01(input []int64) {
	m := memory(input)
	outCh := make(chan int64)
	go m.runUntilHalt(1, outCh)
	var count int
	expected := []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 15259545}
	for val := range outCh {
		if expected[count] != val {
			log.Panicf("expected %d but got %d", expected[count], val)
		}
		log.Println(val)
		count++
	}
}

func test00() {
	m := memory{3, 0, 4, 0, 99}
	expected := memory{1, 0, 4, 0, 99}
	outCh := make(chan int64)
	go m.runUntilHalt(1, outCh)
	if out := <-outCh; out != 1 {
		log.Panicf("expected output of 1 but got %d", out)
	}
	for i := range m {
		if m[i] != expected[i] {
			log.Panicf("at index %d, expected %d but got %d", i, expected[i], m[i])

		}
	}
}

func test01() {
	m := memory{1002, 4, 3, 4, 33}
	expected := memory{1002, 4, 3, 4, 99}
	outCh := make(chan int64)
	go m.runUntilHalt(1, outCh)
	if out, ok := <-outCh; ok {
		log.Panicf("expected test to not output anything but got %d", out)
	}
	for i := range m {
		if m[i] != expected[i] {
			log.Panicf("at index %d, expected %d but got %d", i, expected[i], m[i])

		}
	}
}

// input to task is the memory of a computer
type memory []int64

func (m memory) copyAndRestore(noun, verb int64) memory {
	out := make(memory, len(m))
	copy(out, m)
	out[1] = noun
	out[2] = verb
	return out
}

func (m memory) runUntilHalt(input int64, output chan<- int64) memory {
	var pointer int
	ops := initialOpCodes(input, output)
	for {
		//log.Printf("%+v", m)
		instruction := parseParameterInstruction(m.getValue(pointer))
		code := instruction.opcode()
		if code == 99 {
			close(output)
			return m
		}

		getOp, ok := ops[code]
		if !ok {
			log.Panicf("invalid value at pointer %d %d", pointer, code)
		}
		operation, paramCount := getOp(m, instruction, pointer)
		operation(m)
		pointer += 1 + paramCount
	}
}

type opcode int
type operation func(memory)

func initialOpCodes(input int64, outCh chan<- int64) map[opcode]func(memory, parameterInstruction, int) (operation, int) {
	return map[opcode]func(memory, parameterInstruction, int) (operation, int){
		1: memory.getAdditionOp,
		2: memory.getMultiplicationOp,
		3: getStoreOp(input),
		4: getOutputOp(outCh),
	}
}

func (m memory) getAdditionOp(pi parameterInstruction, pointer int) (operation, int) {
	lhs := pi.parameterMode(0).getGetter(m)(pointer + 1)
	rhs := pi.parameterMode(1).getGetter(m)(pointer + 2)
	storePos := pi.parameterMode(2).getGetter(m)(pointer + 3)
	return func(in memory) {
		in[storePos] = in[lhs] + in[rhs]
	}, 3
}

func (m memory) getMultiplicationOp(pi parameterInstruction, pointer int) (operation, int) {
	lhs := pi.parameterMode(0).getGetter(m)(pointer + 1)
	rhs := pi.parameterMode(1).getGetter(m)(pointer + 2)
	storePos := pi.parameterMode(2).getGetter(m)(pointer + 3)
	return func(in memory) {
		in[storePos] = in[lhs] * in[rhs]
	}, 3
}

func getStoreOp(input int64) func(memory, parameterInstruction, int) (operation, int) {
	return func(m memory, pi parameterInstruction, pointer int) (operation, int) {
		storePos := pi.parameterMode(0).getGetter(m)(pointer + 1)
		return func(in memory) {
			in[storePos] = input
		}, 1
	}
}

func getOutputOp(output chan<- int64) func(memory, parameterInstruction, int) (operation, int) {
	return func(m memory, pi parameterInstruction, pointer int) (operation, int) {
		outputPos := pi.parameterMode(0).getGetter(m)(pointer + 1)
		return func(_ memory) {
			output <- m.getValue(int(outputPos))
		}, 1
	}
}

func (m memory) getValue(cursor int) int64 {
	return m[cursor]
}

func (m memory) getOutputVariable() int64 {
	return m[0]
}

type parameterInstruction []int64

func parseParameterInstruction(s int64) parameterInstruction {
	return splitDigits(s)
}

func (pi parameterInstruction) opcode() opcode {
	if len(pi) == 0 {
		return 0
	}
	if len(pi) == 1 {
		return opcode(pi[0])
	}
	return opcode(pi[len(pi)-2]*10 + pi[len(pi)-1])
}

func (pi parameterInstruction) parameterMode(index int) parameterMode {
	if len(pi) < 2+1+index {
		return 0
	}
	firstParamIndex := len(pi) - 3
	return parameterMode(pi[firstParamIndex-index])
}

func splitDigits(s int64) []int64 {
	if s < 10 {
		return []int64{s}
	}
	return append(splitDigits(s/10), s%10)
}

type parameterMode int64

func (pm parameterMode) getGetter(m memory) func(i int) int64 {
	switch pm {
	case 0:
		return func(i int) int64 {
			return m.getValue(i)
		}
	case 1:
		return func(i int) int64 {
			return int64(i)
		}
	default:
		log.Panicf("unsupported parameter mode: %d", pm)
	}
	return nil
}
