package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_node_walk(t *testing.T) {
	n := &node{
		name: "a",
		parent: &node{
			name: "b",
			parent: &node{
				name:   "c",
				parent: nil,
			},
		},
	}
	var count int
	n.walk(func(n *node) {
		count++
	})
	assert.Equal(t, 3, count)
}

func Test_node_orbits(t *testing.T) {
	com := &node{
		name:   "c",
		parent: nil,
	}
	assert.Equal(t, 0, (com).orbits())

	b := &node{
		name:   "b",
		parent: com,
	}
	assert.Equal(t, 1, b.orbits())

	a := &node{
		name:   "a",
		parent: b,
	}
	assert.Equal(t, 2, a.orbits())
}

func Test_node_orbits2(t *testing.T) {
	com := &node{
		name:   "c",
		parent: nil,
	}
	assert.Equal(t, 0, (com).orbits())

	b := &node{
		name:   "b",
		parent: com,
	}
	assert.Equal(t, 1, b.orbits())

	a := &node{
		name:   "a",
		parent: b,
	}
	assert.Equal(t, 2, a.orbits())
}
