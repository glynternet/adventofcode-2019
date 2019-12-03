package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseInstruction(t *testing.T) {
	assert.Equal(t,
		instruction{distance: 987, direction: direction{x: 1}},
		parseInstruction("R987"))
}

func Test_directions_pointsCovered(t *testing.T) {
	assert.Len(t, directions{direction{}}.pointsCovered(), 1)
	covered := directions{direction{}, direction{x: 1}}.pointsCovered()
	assert.Len(t, covered, 2)
	assert.True(t, covered.contains(xy{}))
	assert.True(t, covered.contains(xy{x: 1}))

}

func Test_findIntersects(t *testing.T) {

	//is := findIntersects(`R2,
	//L2,U2,R2,D2`)
}
