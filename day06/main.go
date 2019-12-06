package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type node struct {
	name   string
	parent *node
}

func (n node) String() string {
	if n.parent == nil {
		return n.name
	}
	return fmt.Sprintf("%s:%s", n.name, n.parent)
}

func main() {
	test01()
	part01()
	part02()
}

func test01() {
	input := `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`

	total := calculateOrbits(input)
	assertTrue(total == 42, "")
	fmt.Println(total)
}

func part01() {
	bs, err := ioutil.ReadFile("/home/glyn/go/src/github.com/glynternet/adventofcode-2019/day06/input.txt")
	assertTrue(err == nil, "err should not happen")
	orbits := calculateOrbits(string(bs))
	fmt.Println(orbits)
}

func part02() {
	bs, err := ioutil.ReadFile("/home/glyn/go/src/github.com/glynternet/adventofcode-2019/day06/input.txt")
	assertTrue(err == nil, "err should not happen")
	fmt.Println(diffBetweenYouAndSanta(generateGraph(string(bs))))
}

type graph map[string]*node

func (g graph) forEach(fn func(*node)) {
	for _, n := range g {
		fn(n)
	}
}

func calculateOrbits(input string) int {
	nodes := generateGraph(input)
	var total int
	for _, n := range nodes {
		total += n.orbits()
	}
	return total
}

func diffBetweenYouAndSanta(g graph) int {
	yourAncestors := make(map[string]struct{})
	y := g["YOU"]
	log.Printf("YOU: %d", y.orbits())
	y.walk(func(n *node) bool {
		yourAncestors[n.name] = struct{}{}
		return true
	})

	s := g["SAN"]
	log.Printf("SAN: %d", s.orbits())
	var commonAncestor string
	s.walk(func(n *node) bool {
		if _, ok := yourAncestors[n.name]; ok {
			commonAncestor = n.name
			return false
		}
		return true
	})

	c := g[commonAncestor].orbits()
	log.Printf("COMMON: %d", c)
	return (y.orbits() - 1) + (s.orbits() - 1) - 2*(c)
}

func generateGraph(input string) graph {
	nodes := make(graph)
	for _, line := range strings.Split(input, "\n") {
		names, ok := parseLine(line)
		if !ok {
			continue
		}
		parent := names[0]
		child := names[1]

		_, pExists := nodes[parent]
		_, cExists := nodes[child]

		if cExists && pExists {
			nodes[child].parent = nodes[parent]
			continue
		}

		// if parent exists add child to graph pointing to parent
		if pExists {
			nodes[child] = &node{
				name:   child,
				parent: nodes[parent],
			}
			continue
		}

		// if child already exists, create parent and point child to it
		if cExists {
			nodes[parent] = &node{
				name: parent,
			}
			nodes[child].parent = nodes[parent]
			continue
		}

		// neither exist
		nodes[parent] = &node{
			name: parent,
		}
		nodes[child] = &node{
			name:   child,
			parent: nodes[parent],
		}
		continue
	}
	return nodes
}

func (n *node) walk(walkFn func(*node) (doContinue bool)) {
	if !walkFn(n) {
		return
	}
	if n.parent != nil {
		n.parent.walk(walkFn)
	}
}

func (n *node) orbits() int {
	var nodeCount int
	n.walk(func(n *node) bool {
		nodeCount++
		return true
	})
	return nodeCount - 1
}

func parseLine(line string) ([]string, bool) {
	ns := strings.Split(line, ")")
	if len(ns) == 1 {
		log.Printf("encountered strange line: %s", line)
		if ns[0] == "" {
			return nil, false
		}
		log.Panicf("received line with 1 node that wasn't empty")
	}
	assertLength(2, ns)
	assertTrue(ns[0] != "" && ns[1] != "", "should both be not empty")
	return ns, true
}

func assertLength(length int, in []string) {
	if len(in) != length {
		log.Panicf("expected length %d but got %d: %v", length, len(in), in)
	}
}

func assertTrue(b bool, s string) {
	if b != true {
		log.Panicf("should be true: %s", s)
	}
}

func assertTruef(b bool, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	if b != true {
		log.Panicf("should be true: %s", msg)
	}
}

func assertFalse(b bool) {
	if b != false {
		log.Panicf("should be false")
	}
}
