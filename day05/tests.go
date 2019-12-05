package main

import "log"

func test01() {
	m := memory{1002, 4, 3, 4, 33}
	expected := memory{1002, 4, 3, 4, 99}
	outCh := make(chan int64)
	go computer{
		memory: m,
	}.runUntilHalt(1, outCh)
	if out, ok := <-outCh; ok {
		log.Panicf("expected test to not output anything but got %d", out)
	}
	for i := range m {
		if m[i] != expected[i] {
			log.Panicf("at index %d, expected %d but got %d", i, expected[i], m[i])

		}
	}
}

func test02() {
	m := memory{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	outCh := make(chan int64)
	go computer{
		memory: m,
	}.runUntilHalt(0, outCh)
	for out := range outCh {
		if out != 0 {
			panic("expected 0")
		}
	}
}

func test03() {
	m := memory{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	outCh := make(chan int64)
	go computer{
		memory: m,
	}.runUntilHalt(1, outCh)
	for out := range outCh {
		if out != 1 {
			panic("expected 1")
		}
	}
}

func test04() {
	m := memory{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}
	outCh := make(chan int64)
	go computer{
		memory: m,
	}.runUntilHalt(1, outCh)
	for out := range outCh {
		if out != 1 {
			panic("expected 1")
		}
	}
}

func test05() {
	m := memory{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}
	outCh := make(chan int64)
	go computer{
		memory: m,
	}.runUntilHalt(0, outCh)
	for out := range outCh {
		if out != 0 {
			panic("expected 0")
		}
	}
}

func test06() {
	m := memory{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	outCh := make(chan int64)
	go computer{
		memory: m,
	}.runUntilHalt(8, outCh)
	for out := range outCh {
		if out != 1 {
			panic("expected 1")
		}
	}
}

func test07() {
	m := memory{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	outCh := make(chan int64)
	go computer{
		memory: m,
	}.runUntilHalt(9, outCh)
	for out := range outCh {
		if out != 0 {
			panic("expected 0")
		}
	}
}

func test08() {
	m := memory{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}
	outCh := make(chan int64)
	go computer{
		memory: m,
	}.runUntilHalt(8, outCh)
	for out := range outCh {
		if out != 0 {
			panic("expected 0")
		}
	}
}

func test09() {
	m := memory{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}
	outCh := make(chan int64)
	go computer{
		memory: m,
	}.runUntilHalt(7, outCh)
	for out := range outCh {
		if out != 1 {
			panic("expected 1")
		}
	}
}

func test10() {
	m := memory{3, 3, 1108, -1, 8, 3, 4, 3, 99}
	outCh := make(chan int64)
	go computer{
		memory: m,
	}.runUntilHalt(8, outCh)
	for out := range outCh {
		if out != 1 {
			panic("expected 1")
		}
	}
}

func test11() {
	m := memory{3, 3, 1108, -1, 8, 3, 4, 3, 99}
	outCh := make(chan int64)
	go computer{
		memory: m,
	}.runUntilHalt(0, outCh)
	for out := range outCh {
		if out != 0 {
			panic("expected 0")
		}
	}
}

func test12() {
	m := memory{3, 3, 1107, -1, 8, 3, 4, 3, 99}
	outCh := make(chan int64)
	go computer{
		memory: m,
	}.runUntilHalt(0, outCh)
	for out := range outCh {
		if out != 1 {
			panic("expected 1")
		}
	}
}

func test13() {
	m := memory{3, 3, 1107, -1, 8, 3, 4, 3, 99}
	outCh := make(chan int64)
	go computer{
		memory: m,
	}.runUntilHalt(9, outCh)
	for out := range outCh {
		if out != 0 {
			panic("expected 0")
		}
	}
}

func test14() {
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
		go computer{
			memory: m.copy(),
		}.runUntilHalt(tc.in, outCh)
		for out := range outCh {
			if out != tc.expect {
				log.Panicf("%d: expected %d but got %d", i, tc.expect, out)
			}
		}
	}
}
