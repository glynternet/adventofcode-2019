package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	bs, err := ioutil.ReadFile("/home/glyn/go/src/github.com/glynternet/adventofcode-2019/day03/input.txt")
	if err != nil {
		panic(err)
	}

	for i, tc := range []struct {
		in           string
		minIntersect int
	}{
		{
			in: `R2
R2`,
			minIntersect: 1,
		},
		{
			in: `R2
L2,U2,R4,D2`,
			minIntersect: 2,
		},
		{
			in: `R8,U5,L5,D3
U7,R6,D4,L4`,
			minIntersect: 6,
		},
		{
			in: `R75,D30,R83,U83,L12,D49,R71,U7,L72
U62,R66,U55,R34,D71,R55,D58,R83`,
			minIntersect: 159,
		},
		{
			in: `R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
U98,R91,D20,R16,D67,R40,U7,R15,U6,R7`,
			minIntersect: 135,
		},
	} {
		min := findMinDistanceIgnoringTime(tc.in)
		if min != tc.minIntersect {
			log.Panicf("%d: expected %d but got %d", i, tc.minIntersect, min)
		}
	}
	part01(string(bs))

	i := 0
	in := `R75,D30,R83,U83,L12,D49,R71,U7,L72
U62,R66,U55,R34,D71,R55,D58,R83`
	expected := 610
	time := findMinDistanceConsideringTime(in)
	if time != expected {
		log.Panicf("%d: expected %d but got %d", i, expected, time)
	}
	part02(string(bs))
}

func part01(bs string) {
	min := findMinDistanceIgnoringTime(bs)
	expected := 1337
	if min != expected {
		log.Panicf("%s: expected %d but got %d", "part01", expected, min)
	}
	log.Println(min)
}

func part02(bs string) {
	min := findMinDistanceConsideringTime(bs)
	expected := 65356
	if min != expected {
		log.Panicf("%s: expected %d but got %d", "part02", expected, min)
	}
	log.Println(min)
}

func findMinDistanceIgnoringTime(bs string) int {
	intersects := findIntersects(bs)
	min := 99999999
	for intersect := range intersects {
		if intersect.manhattanDistance() <= min {
			min = intersect.manhattanDistance()
		}
	}
	return min
}

func findMinDistanceConsideringTime(bs string) int {
	intersects := findIntersects(bs)
	min := 99999999
	for _, dist := range intersects {
		if dist <= min {
			min = dist
		}
	}
	return min
}

func findIntersects(bs string) map[xy]int {
	lines := strings.Split(bs, "\n")
	var wires []string
	for _, wire := range lines {
		if wire != "" {
			wires = append(wires, wire)
		}
	}
	if len(wires) != 2 {
		log.Panicf("expected %d wires got %d", 2, len(wires))
	}

	instructions := make([]directions, len(wires))
	for i, wire := range wires {
		instructionStrings := strings.Split(wire, ",")
		for _, instructionString := range instructionStrings {
			instructions[i] = append(instructions[i], parseInstruction(instructionString).split()...)
		}
	}

	points := func() []pointsCovered {
		var mps []pointsCovered
		for _, ins := range instructions {
			mps = append(mps, ins.pointsCovered())
		}
		return mps
	}()
	// this is fine when just 2
	intersects := make(map[xy]int)
	for p := range points[0] {
		if points[1].contains(p) {
			intersects[p] = points[0][p] + points[1][p]
		}
	}
	return intersects
}

type instruction struct {
	direction
	distance int
}

type direction struct {
	x, y int
}

type xy struct {
	x, y int
}

func (xy xy) manhattanDistance() int {
	x := xy.x
	if x < 0 {
		x *= -1
	}
	y := xy.y
	if y < 0 {
		y *= -1
	}
	return x + y
}

func parseInstruction(val string) instruction {
	letter := val[0]
	distance, err := strconv.Atoi(val[1:])
	if err != nil {
		panic(err)
	}
	var dir direction
	switch letter {
	default:
		log.Panicf("unexpected direction: %c", letter)
	case 'L':
		dir = direction{x: -1}
	case 'R':
		dir = direction{x: 1}
	case 'U':
		dir = direction{y: 1}
	case 'D':
		dir = direction{y: -1}
	}
	return instruction{
		direction: dir,
		distance:  distance,
	}
}

func (inst instruction) split() directions {
	var dirs directions
	for j := 0; j < inst.distance; j++ {
		dirs = append(dirs, inst.direction)
	}
	return dirs
}

type directions []direction

func (ds directions) pointsCovered() pointsCovered {
	var cursor xy
	points := make(pointsCovered)
	for i, d := range ds {
		cursor.x += d.x
		cursor.y += d.y
		if _, exists := points[cursor]; !exists {
			// after step with index 0, we have taken 1 step, so it needs a +1
			points[cursor] = i + 1
		}
	}
	return points
}

type pointsCovered map[xy]int

func (pc pointsCovered) contains(val xy) bool {
	_, ok := pc[val]
	return ok
}
