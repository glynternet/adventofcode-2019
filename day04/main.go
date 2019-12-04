package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	part01()
	part02()
}

func part01() {
	var count int
	for i := 264360; i <= 746325; i++ {
		if splitChars(i).all(num.containsDouble, num.neverDecreases) {
			count++
		}
	}
	if count != 945 {
		log.Panicf("expected 945 for part 1")
	}
	fmt.Println(count)
}

func part02() {
	var count int
	for i := 264360; i <= 746325; i++ {
		if splitChars(i).all(num.containsNonOverlappingDigitPair, num.neverDecreases) {
			count++
		}
	}
	if count != 617 {
		log.Panicf("expected 617 for part 2")
	}
	fmt.Println(count)
}

type num [6]int

func (n num) all(fns ...func(num) bool) bool {
	for _, fn := range fns {
		if !fn(n) {
			return false
		}
	}
	return true
}

func (n num) containsDouble() bool {
	return len(n.getDoublesDigitCounts()) > 0
}

func (n num) containsNonOverlappingDigitPair() bool {
	for _, count := range n.getDoublesDigitCounts() {
		if count == 1 {
			return true
		}
	}
	return false
}

func (n num) getDoublesDigitCounts() map[int]int {
	is := make(map[int]int)
	for i := 0; i < 5; i++ {
		if n[i] == n[i+1] {
			is[n[i]] += 1
		}
	}
	return is
}

func (n num) neverDecreases() bool {
	for i := 0; i < 5; i++ {
		if n[i] > n[i+1] {
			return false
		}
	}
	return true
}

func splitChars(in int) num {
	str := strconv.Itoa(in)
	if len(str) != 6 {
		log.Panicf("expected 6 chars")
	}
	var out [6]int
	for i := 0; i < 6; i++ {
		atoi, err := strconv.Atoi(string(str[i]))
		if err != nil {
			panic(err)
		}
		out[i] = atoi
	}
	return out
}
