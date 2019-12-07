package main

import "log"

func day5test00() {
	m := memory{3, 0, 4, 0, 99}
	expected := memory{1, 0, 4, 0, 99}
	outCh := make(chan int64)
	inCh := make(chan int64)

	go computer{
		memory: m,
	}.runUntilHalt(inCh, outCh)
	inCh <- 1
	if out := <-outCh; out != 1 {
		log.Panicf("expected output of 1 but got %d", out)
	}
	for i := range m {
		if m[i] != expected[i] {
			log.Panicf("at index %d, expected %d but got %d", i, expected[i], m[i])

		}
	}
}

func day5test01() {
	m := memory{1002, 4, 3, 4, 33}
	expected := memory{1002, 4, 3, 4, 99}
	outCh := make(chan int64)
	inCh := make(chan int64, 1)
	go computer{
		memory: m,
	}.runUntilHalt(inCh, outCh)
	inCh <- 1
	if out, ok := <-outCh; ok {
		log.Panicf("expected test to not output anything but got %d", out)
	}
	for i := range m {
		if m[i] != expected[i] {
			log.Panicf("at index %d, expected %d but got %d", i, expected[i], m[i])

		}
	}
}

func day5test02() {
	m := memory{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	outCh := make(chan int64)
	inCh := make(chan int64)
	go computer{
		memory: m,
	}.runUntilHalt(inCh, outCh)
	inCh <- 0
	for out := range outCh {
		if out != 0 {
			panic("expected 0")
		}
	}
}

func day5test03() {
	m := memory{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	outCh := make(chan int64)
	inCh := make(chan int64)
	go computer{
		memory: m,
	}.runUntilHalt(inCh, outCh)
	inCh <- 1
	for out := range outCh {
		if out != 1 {
			panic("expected 1")
		}
	}
}

func day5test04() {
	m := memory{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}
	outCh := make(chan int64)
	inCh := make(chan int64)
	go computer{
		memory: m,
	}.runUntilHalt(inCh, outCh)
	inCh <- 1
	for out := range outCh {
		if out != 1 {
			panic("expected 1")
		}
	}
}

func day5test05() {
	m := memory{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}
	outCh := make(chan int64)
	inCh := make(chan int64)
	go computer{
		memory: m,
	}.runUntilHalt(inCh, outCh)
	inCh <- 0
	for out := range outCh {
		if out != 0 {
			panic("expected 0")
		}
	}
}

func day5test06() {
	m := memory{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	outCh := make(chan int64)
	inCh := make(chan int64)
	go computer{
		memory: m,
	}.runUntilHalt(inCh, outCh)
	inCh <- 8
	for out := range outCh {
		if out != 1 {
			panic("expected 1")
		}
	}
}

func day5test07() {
	m := memory{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	outCh := make(chan int64)
	inCh := make(chan int64)
	go computer{
		memory: m,
	}.runUntilHalt(inCh, outCh)
	inCh <- 9
	for out := range outCh {
		if out != 0 {
			panic("expected 0")
		}
	}
}

func day5test08() {
	m := memory{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}
	outCh := make(chan int64)
	inCh := make(chan int64)
	go computer{
		memory: m,
	}.runUntilHalt(inCh, outCh)
	inCh <- 8
	for out := range outCh {
		if out != 0 {
			panic("expected 0")
		}
	}
}

func day5test09() {
	m := memory{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}
	outCh := make(chan int64)
	inCh := make(chan int64)
	go computer{
		memory: m,
	}.runUntilHalt(inCh, outCh)
	inCh <- 7
	for out := range outCh {
		if out != 1 {
			panic("expected 1")
		}
	}
}

func day5test10() {
	m := memory{3, 3, 1108, -1, 8, 3, 4, 3, 99}
	outCh := make(chan int64)
	inCh := make(chan int64)
	go computer{
		memory: m,
	}.runUntilHalt(inCh, outCh)
	inCh <- 8
	for out := range outCh {
		if out != 1 {
			panic("expected 1")
		}
	}
}

func day5test11() {
	m := memory{3, 3, 1108, -1, 8, 3, 4, 3, 99}
	outCh := make(chan int64)
	inCh := make(chan int64)
	go computer{
		memory: m,
	}.runUntilHalt(inCh, outCh)
	inCh <- 0
	for out := range outCh {
		if out != 0 {
			panic("expected 0")
		}
	}
}

func day5test12() {
	m := memory{3, 3, 1107, -1, 8, 3, 4, 3, 99}
	outCh := make(chan int64)
	inCh := make(chan int64)
	go computer{
		memory: m,
	}.runUntilHalt(inCh, outCh)
	inCh <- 0
	for out := range outCh {
		if out != 1 {
			panic("expected 1")
		}
	}
}

func day5test13() {
	m := memory{3, 3, 1107, -1, 8, 3, 4, 3, 99}
	outCh := make(chan int64)
	inCh := make(chan int64)
	go computer{
		memory: m,
	}.runUntilHalt(inCh, outCh)
	inCh <- 9
	for out := range outCh {
		if out != 0 {
			panic("expected 0")
		}
	}
}

func day5test14() {
	m := memory{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	for i, tc := range []struct {
		in     int64
		expect int64
	}{
		{in: 0, expect: 999},
		{in: 8, expect: 1000},
		{in: 9, expect: 1001},
	} {
		outCh := make(chan int64)
		inCh := make(chan int64)
		go computer{
			memory: m.copy(),
		}.runUntilHalt(inCh, outCh)
		inCh <- tc.in
		for out := range outCh {
			if out != tc.expect {
				log.Panicf("%d: expected %d but got %d", i, tc.expect, out)
			}
		}
	}
}

func day5Tests() {
	input := []int64{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1102, 17, 65, 225, 102, 21, 95, 224, 1001, 224, -1869, 224, 4, 224, 1002, 223, 8, 223, 101, 7, 224, 224, 1, 224, 223, 223, 101, 43, 14, 224, 1001, 224, -108, 224, 4, 224, 102, 8, 223, 223, 101, 2, 224, 224, 1, 223, 224, 223, 1101, 57, 94, 225, 1101, 57, 67, 225, 1, 217, 66, 224, 101, -141, 224, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 1, 224, 1, 224, 223, 223, 1102, 64, 34, 225, 1101, 89, 59, 225, 1102, 58, 94, 225, 1002, 125, 27, 224, 101, -2106, 224, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 5, 224, 1, 224, 223, 223, 1102, 78, 65, 225, 1001, 91, 63, 224, 101, -127, 224, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 3, 224, 1, 223, 224, 223, 1102, 7, 19, 224, 1001, 224, -133, 224, 4, 224, 102, 8, 223, 223, 101, 6, 224, 224, 1, 224, 223, 223, 2, 61, 100, 224, 101, -5358, 224, 224, 4, 224, 102, 8, 223, 223, 101, 3, 224, 224, 1, 224, 223, 223, 1101, 19, 55, 224, 101, -74, 224, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 1, 224, 1, 224, 223, 223, 1101, 74, 68, 225, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 107, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 329, 1001, 223, 1, 223, 1008, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 344, 1001, 223, 1, 223, 7, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 359, 1001, 223, 1, 223, 8, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 374, 1001, 223, 1, 223, 1007, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 389, 101, 1, 223, 223, 8, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 404, 101, 1, 223, 223, 1108, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 419, 1001, 223, 1, 223, 1108, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 434, 101, 1, 223, 223, 1108, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 449, 101, 1, 223, 223, 1008, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 464, 101, 1, 223, 223, 7, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 479, 101, 1, 223, 223, 108, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 494, 101, 1, 223, 223, 107, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 509, 101, 1, 223, 223, 107, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 524, 1001, 223, 1, 223, 1107, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 539, 101, 1, 223, 223, 1008, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 554, 1001, 223, 1, 223, 8, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 569, 101, 1, 223, 223, 1007, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 584, 1001, 223, 1, 223, 1107, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 599, 101, 1, 223, 223, 7, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 614, 101, 1, 223, 223, 108, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 629, 1001, 223, 1, 223, 108, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 644, 101, 1, 223, 223, 1007, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 659, 101, 1, 223, 223, 1107, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 674, 1001, 223, 1, 223, 4, 223, 99, 226}
	log.Println("day5Part1")
	day5Part1(input)
	log.Println("day5Part2")
	day5Part2(input)

	day5test00()
	day5test01()
	day5test02()
	day5test03()
	day5test04()
	day5test05()
	day5test06()
	day5test07()
	day5test08()
	day5test09()
	day5test10()
	day5test11()
	day5test12()
	day5test13()
	day5test14()
}

func day5Part1(input []int64) {
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

func day5Part2(input []int64) {
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
