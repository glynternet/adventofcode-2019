package main

import (
	"log"
)

func main() {
	test00()
	test01()
	test02()
	test03()
	test04()
	test05()
	test06()
	test07()
	test08()
	test09()
	test10()
	test11()
	test12()
	test13()
	test14()
	input := []int64{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1102, 17, 65, 225, 102, 21, 95, 224, 1001, 224, -1869, 224, 4, 224, 1002, 223, 8, 223, 101, 7, 224, 224, 1, 224, 223, 223, 101, 43, 14, 224, 1001, 224, -108, 224, 4, 224, 102, 8, 223, 223, 101, 2, 224, 224, 1, 223, 224, 223, 1101, 57, 94, 225, 1101, 57, 67, 225, 1, 217, 66, 224, 101, -141, 224, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 1, 224, 1, 224, 223, 223, 1102, 64, 34, 225, 1101, 89, 59, 225, 1102, 58, 94, 225, 1002, 125, 27, 224, 101, -2106, 224, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 5, 224, 1, 224, 223, 223, 1102, 78, 65, 225, 1001, 91, 63, 224, 101, -127, 224, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 3, 224, 1, 223, 224, 223, 1102, 7, 19, 224, 1001, 224, -133, 224, 4, 224, 102, 8, 223, 223, 101, 6, 224, 224, 1, 224, 223, 223, 2, 61, 100, 224, 101, -5358, 224, 224, 4, 224, 102, 8, 223, 223, 101, 3, 224, 224, 1, 224, 223, 223, 1101, 19, 55, 224, 101, -74, 224, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 1, 224, 1, 224, 223, 223, 1101, 74, 68, 225, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 107, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 329, 1001, 223, 1, 223, 1008, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 344, 1001, 223, 1, 223, 7, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 359, 1001, 223, 1, 223, 8, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 374, 1001, 223, 1, 223, 1007, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 389, 101, 1, 223, 223, 8, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 404, 101, 1, 223, 223, 1108, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 419, 1001, 223, 1, 223, 1108, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 434, 101, 1, 223, 223, 1108, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 449, 101, 1, 223, 223, 1008, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 464, 101, 1, 223, 223, 7, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 479, 101, 1, 223, 223, 108, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 494, 101, 1, 223, 223, 107, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 509, 101, 1, 223, 223, 107, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 524, 1001, 223, 1, 223, 1107, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 539, 101, 1, 223, 223, 1008, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 554, 1001, 223, 1, 223, 8, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 569, 101, 1, 223, 223, 1007, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 584, 1001, 223, 1, 223, 1107, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 599, 101, 1, 223, 223, 7, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 614, 101, 1, 223, 223, 108, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 629, 1001, 223, 1, 223, 108, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 644, 101, 1, 223, 223, 1007, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 659, 101, 1, 223, 223, 1107, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 674, 1001, 223, 1, 223, 4, 223, 99, 226}
	log.Println("part01")
	part01(input)
	log.Println("part02")
	part02(input)
}

func part01(input []int64) {
	inCh := make(chan int64)
	outCh := make(chan int64)
	go computer{memory: memory(input).copy()}.runUntilHalt(inCh, outCh)
	inCh <- 1
	var count int
	expected := []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 15259545}
	for val := range outCh {
		if expected[count] != val {
			log.Panicf("expected %d but got %d", expected[count], val)
		}
		//log.Println(val)
		count++
	}
}

func part02(input []int64) {
	inCh := make(chan int64)
	outCh := make(chan int64)
	go computer{memory: memory(input).copy()}.runUntilHalt(inCh, outCh)
	inCh <- 5
	var count int
	expected := []int64{7616021}
	for val := range outCh {
		if expected[count] != val {
			log.Panicf("expected %d but got %d", expected[count], val)
		}
		//log.Println(val)
		count++
	}
}

type computer struct {
	memory
	pointer int
}

// input to task is the memory of a computer
type memory []int64

func (m memory) copy() memory {
	out := make(memory, len(m))
	copy(out, m)
	return out
}

func (c computer) runUntilHalt(input <-chan int64, output chan<- int64) {
	ops := initialOpCodes(<-input, output)
	for {
		//log.Printf("%+v", c)
		instruction := parseParameterInstruction(c.memory.getValue(c.pointer))
		code := instruction.opcode()
		if code == 99 {
			close(output)
			return
		}

		getOp, ok := ops[code]
		if !ok {
			log.Panicf("invalid opcode:%d at pointer:%d", code, c.pointer)
		}
		getOp(instruction)(&c)
	}
}

type opcode int
type operation func(*computer)

func initialOpCodes(input int64, outCh chan<- int64) map[opcode]func(parameterInstruction) operation {
	return map[opcode]func(parameterInstruction) operation{
		1: getAdditionOp,
		2: getMultiplicationOp,
		3: getStoreOp(input),
		4: getOutputOp(outCh),
		5: getJumpIfTrueOp,
		6: getJumpIfFalseOp,
		7: getLessThanOp,
		8: getEqualOp,
	}
}

func getAdditionOp(pi parameterInstruction) operation {
	lhsParamMode := pi.parameterMode(0)
	rhsParamMode := pi.parameterMode(1)
	storeParamMode := pi.parameterMode(2)
	return func(c *computer) {
		lhs := lhsParamMode.getParamValue(*c, 0)
		rhs := rhsParamMode.getParamValue(*c, 1)
		storePos := storeParamMode.getParamValue(*c, 2)
		c.memory[storePos] = c.memory[lhs] + c.memory[rhs]
		c.pointer += 3 + 1
	}
}

func getMultiplicationOp(pi parameterInstruction) operation {
	lhsParamMode := pi.parameterMode(0)
	rhsParamMode := pi.parameterMode(1)
	storeParamMode := pi.parameterMode(2)
	return func(c *computer) {
		lhs := lhsParamMode.getParamValue(*c, 0)
		rhs := rhsParamMode.getParamValue(*c, 1)
		storePos := storeParamMode.getParamValue(*c, 2)
		c.memory[storePos] = c.memory[lhs] * c.memory[rhs]
		c.pointer += 3 + 1
	}
}

func getStoreOp(input int64) func(parameterInstruction) operation {
	return func(pi parameterInstruction) operation {
		storePositionParamMode := pi.parameterMode(0)
		return func(c *computer) {
			storePos := storePositionParamMode.getParamValue(*c, 0)
			c.memory[storePos] = input
			c.pointer += 1 + 1
		}
	}
}

func getOutputOp(output chan<- int64) func(parameterInstruction) operation {
	return func(pi parameterInstruction) operation {
		outputPositionParamMode := pi.parameterMode(0)
		return func(c *computer) {
			outputPos := outputPositionParamMode.getParamValue(*c, 0)
			output <- c.memory.getValue(int(outputPos))
			c.pointer += 1 + 1
		}
	}
}

func getJumpIfTrueOp(pi parameterInstruction) operation {
	firstParamMode := pi.parameterMode(0)
	secondParamMode := pi.parameterMode(1)
	return func(c *computer) {
		firstParam := firstParamMode.getValueOtherFunc(*c, 0)
		secondParam := secondParamMode.getValueOtherFunc(*c, 1)
		if firstParam != 0 {
			c.pointer = int(secondParam)
			return
		}
		c.pointer += 2 + 1
		return
	}
}

func getJumpIfFalseOp(pi parameterInstruction) operation {
	firstParamMode := pi.parameterMode(0)
	secondParamMode := pi.parameterMode(1)
	return func(c *computer) {
		firstParam := firstParamMode.getValueOtherFunc(*c, 0)
		secondParam := secondParamMode.getValueOtherFunc(*c, 1)
		if firstParam == 0 {
			c.pointer = int(secondParam)
			return
		}
		c.pointer += 2 + 1
	}
}

func getLessThanOp(pi parameterInstruction) operation {
	firstParamMode := pi.parameterMode(0)
	secondParamMode := pi.parameterMode(1)
	thirdParamMode := pi.parameterMode(2)
	return func(c *computer) {
		// getting param value
		firstParam := firstParamMode.getValueOtherFunc(*c, 0)
		// getting param value
		secondParam := secondParamMode.getValueOtherFunc(*c, 1)
		// getting param value
		thirdParam := thirdParamMode.getParamValue(*c, 2)
		var storeValue int64
		if firstParam < secondParam {
			storeValue = 1
		}
		c.memory[int(thirdParam)] = storeValue
		c.pointer += 3 + 1
	}
}

func getEqualOp(pi parameterInstruction) operation {
	firstParamMode := pi.parameterMode(0)
	secondParamMode := pi.parameterMode(1)
	thirdParamMode := pi.parameterMode(2)
	return func(c *computer) {
		firstParam := firstParamMode.getValueOtherFunc(*c, 0)
		secondParam := secondParamMode.getValueOtherFunc(*c, 1)
		thirdParam := thirdParamMode.getParamValue(*c, 2)
		var storeValue int64
		if firstParam == secondParam {
			storeValue = 1
		}
		c.memory[int(thirdParam)] = storeValue
		c.pointer += 3 + 1
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

func (pm parameterMode) getParamValue(c computer, i int) int64 {
	paramStartIndex := c.pointer + 1
	switch pm {
	case 0:
		// position mode
		return c.memory.getValue(paramStartIndex + i)
	case 1:
		// immediate mode
		return int64(paramStartIndex + i)
	}
	log.Panicf("unsupported parameter mode: %d", pm)
	return 987654321
}

func (pm parameterMode) getValueOtherFunc(c computer, i int) int64 {
	paramStartIndex := c.pointer + 1
	valueAtParam := c.memory.getValue(paramStartIndex + i)
	switch pm {
	case 0:
		// position mode
		return c.memory.getValue(int(valueAtParam))
	case 1:
		// immediate mode
		return valueAtParam
	}
	log.Panicf("unsupported parameter mode: %d", pm)
	return 987654321
}
