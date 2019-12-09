package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/glynternet/adventofcode-2019/assert"
)

func main() {
	//day5Tests()
	//part01testCases()
	//part01()
	//
	part02testCases()
	part02()
}

func part01() {
	input := memory{3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 34, 51, 64, 81, 102, 183, 264, 345, 426, 99999, 3, 9, 102, 2, 9, 9, 1001, 9, 4, 9, 4, 9, 99, 3, 9, 101, 4, 9, 9, 102, 5, 9, 9, 1001, 9, 2, 9, 4, 9, 99, 3, 9, 101, 3, 9, 9, 1002, 9, 5, 9, 4, 9, 99, 3, 9, 102, 3, 9, 9, 101, 3, 9, 9, 1002, 9, 4, 9, 4, 9, 99, 3, 9, 1002, 9, 3, 9, 1001, 9, 5, 9, 1002, 9, 5, 9, 101, 3, 9, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 99}
	max := int64(-999999)
	for _, seq := range generateSequences(0) {
		ampCount := 5
		out := runChained(ampCount, input.copy(), seq)
		if out > max {
			max = out
		}
	}
	assert.True(max == int64(46248), "")
	fmt.Println(max)
}

func part02() {
	input := memory{3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 34, 51, 64, 81, 102, 183, 264, 345, 426, 99999, 3, 9, 102, 2, 9, 9, 1001, 9, 4, 9, 4, 9, 99, 3, 9, 101, 4, 9, 9, 102, 5, 9, 9, 1001, 9, 2, 9, 4, 9, 99, 3, 9, 101, 3, 9, 9, 1002, 9, 5, 9, 4, 9, 99, 3, 9, 102, 3, 9, 9, 101, 3, 9, 9, 1002, 9, 4, 9, 4, 9, 99, 3, 9, 1002, 9, 3, 9, 1001, 9, 5, 9, 1002, 9, 5, 9, 101, 3, 9, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 99}
	max := int64(-999999)
	for _, seq := range generateSequences(5) {
		ampCount := 5
		out := runLooped(ampCount, input.copy(), seq)
		if out > max {
			max = out
		}
	}
	//assert.True(max == int64(46248), "")
	fmt.Println(max)
}

// need to start at 0 for part 1 and 5 for part 2
func generateSequences(start int64) [][]int64 {
	var seqs [][]int64
	end := 5 + start
	for a := start; a < end; a++ {
		for b := start; b < end; b++ {
			if b == a {
				continue
			}
			for c := start; c < end; c++ {
				if c == b || c == a {
					continue
				}
				for d := start; d < end; d++ {
					if d == c || d == b || d == a {
						continue
					}
					for e := start; e < end; e++ {
						if e == d || e == c || e == b || e == a {
							continue
						}
						seqs = append(seqs, []int64{a, b, c, d, e})
					}
				}
			}
		}
	}
	return seqs
}

func part01testCases() {
	for i, tc := range []struct {
		thrusterSignal []int64
		expectOut      int64
		memory         memory
	}{
		{
			thrusterSignal: []int64{4, 3, 2, 1, 0},
			memory:         memory{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0},
			expectOut:      43210,
		},
		{
			thrusterSignal: []int64{0, 1, 2, 3, 4},
			memory: memory{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23,
				101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0},
			expectOut: 54321,
		},
		{
			thrusterSignal: []int64{1, 0, 4, 3, 2},
			memory: memory{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33,
				1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0},
			expectOut: 65210,
		},
	} {
		ampCount := 5
		out := runChained(ampCount, tc.memory, tc.thrusterSignal)
		assert.Truef(tc.expectOut == out, "[%d] expected %d but got %d", i, tc.expectOut, out)
		log.Println(out)
	}
}

func part02testCases() {
	for i, tc := range []struct {
		thrusterSignal []int64
		expectOut      int64
		memory         memory
	}{
		{
			thrusterSignal: []int64{9, 8, 7, 6, 5},
			memory: memory{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26,
				27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5},
			expectOut: 139629729,
		},
		{
			thrusterSignal: []int64{9, 7, 8, 5, 6},
			memory: memory{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54,
				-5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4,
				53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10},
			expectOut: 18216,
		},
	} {
		ampCount := 5
		out := runLooped(ampCount, tc.memory, tc.thrusterSignal)
		assert.Truef(tc.expectOut == out, "[%d] expected %d but got %d", i, tc.expectOut, out)
		log.Println(out)
	}
}

func runLooped(ampCount int, input memory, thrusterSignal []int64) int64 {
	inChans := make([]chan int64, ampCount)
	outChans := make([]chan int64, ampCount)
	ampControllers := make([]computer, ampCount)
	for i := 0; i < ampCount; i++ {
		ampControllers[i] = computer{
			id:      i,
			memory:  input.copy(),
			pointer: 0,
		}
		inChans[i] = make(chan int64, 1)
		outChans[i] = make(chan int64, 1)
		// could be a data thing here
		inChans[i] <- thrusterSignal[i]
		log.Printf("[%d] initial input:%d sent", i, thrusterSignal[i])
	}

	for i := 0; i < ampCount-1; i++ {
		log.Printf("[%d] chaining", i)
		go func(index int) {
			log.Printf("[%d] channels chained", index)
			for {
				val, ok := <-outChans[index]
				nextInChan := inChans[index+1]
				if !ok {
					log.Printf("closing input channel: %d", index+1)
					close(nextInChan)
					return
				}
				nextInChan <- val
			}
		}(i)
	}

	wg := sync.WaitGroup{}
	var last int64
	wg.Add(1)
	go func() {
		log.Printf("channels looping")
		for {
			val, ok := <-outChans[ampCount-1]
			if !ok {
				log.Printf("closing input channel: %d", 0)
				close(inChans[0])
				wg.Done()
				return
			}
			log.Println(val)
			last = val
			inChans[0] <- val
		}
	}()

	for i := 0; i < ampCount; i++ {
		go func(index int) {
			ampControllers[index].runUntilHalt(inChans[index], outChans[index])
		}(i)
	}

	inChans[0] <- 0
	//var last int64
	//for val := range outChans[ampCount-1] {
	//	last = val
	//}
	wg.Wait()
	return last
}

func runChained(ampCount int, input memory, thrusterSignal []int64) int64 {
	inChans := make([]chan int64, ampCount)
	outChans := make([]chan int64, ampCount)
	ampControllers := make([]computer, ampCount)
	for i := 0; i < ampCount; i++ {
		ampControllers[i] = computer{
			id:      i,
			memory:  input.copy(),
			pointer: 0,
		}
		inChans[i] = make(chan int64, 1)
		outChans[i] = make(chan int64, 1)
		// could be a data thing here
		inChans[i] <- thrusterSignal[i]
		log.Printf("[%d] initial input sent", i)
	}

	for i := 0; i < ampCount-1; i++ {
		log.Printf("[%d] chaining", i)
		go func(index int) {
			val, ok := <-outChans[index]
			if !ok {
				return
			}
			inChans[index+1] <- val
			log.Printf("[%d] channels chained", index)
		}(i)
	}

	for i := 0; i < ampCount; i++ {
		go func(index int) {
			ampControllers[index].runUntilHalt(inChans[index], outChans[index])
		}(i)
	}

	inChans[0] <- 0
	var last int64
	for val := range outChans[ampCount-1] {
		last = val
	}
	return last
}

type computer struct {
	id int
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
	ops := initialOpCodes(c.id, input, output)
	for {
		//log.Printf("%+v", c)
		instruction := parseParameterInstruction(c.memory.getValue(c.pointer))
		code := instruction.opcode()
		if code == 99 {
			log.Printf("[%d] closing output channel", c.id)
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

func initialOpCodes(id int, input <-chan int64, outCh chan<- int64) map[opcode]func(parameterInstruction) operation {
	return map[opcode]func(parameterInstruction) operation{
		1: getAdditionOp,
		2: getMultiplicationOp,
		3: getStoreOp(id, input),
		4: getOutputOp(id, outCh),
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

func getStoreOp(id int, input <-chan int64) func(parameterInstruction) operation {
	return func(pi parameterInstruction) operation {
		storePositionParamMode := pi.parameterMode(0)
		return func(c *computer) {
			storePos := storePositionParamMode.getParamValue(*c, 0)
			log.Printf("[%d] waiting for input", id)
			i, ok := <-input
			if !ok {
				log.Printf("[%d] CHANNEL IS CLOSED", id)
				return
			}
			log.Printf("[%d] input received:%d", id, i)
			c.memory[storePos] = i
			c.pointer += 1 + 1
		}
	}
}

func getOutputOp(id int, output chan<- int64) func(parameterInstruction) operation {
	return func(pi parameterInstruction) operation {
		outputPositionParamMode := pi.parameterMode(0)
		return func(c *computer) {
			outputPos := outputPositionParamMode.getParamValue(*c, 0)
			value := c.memory.getValue(int(outputPos))
			log.Printf("[%d] sending value:%d", id, value)
			output <- value
			log.Printf("[%d] sent value:%d", id, value)
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
