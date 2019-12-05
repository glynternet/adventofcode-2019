package main

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseParamaters(t *testing.T) {

	for _, tc := range []struct {
		in  int64
		out []int64
	}{
		{
			in:  0,
			out: []int64{0},
		}, {
			in:  1,
			out: []int64{1},
		}, {
			in:  11,
			out: []int64{1, 1},
		}, {
			in:  123,
			out: []int64{1, 2, 3},
		}, {
			in:  1239875,
			out: []int64{1, 2, 3, 9, 8, 7, 5},
		},
	} {
		t.Run(strconv.FormatInt(tc.in, 10), func(t *testing.T) {
			assert.Equal(t, tc.out, splitDigits(tc.in))
		})
	}
}

func Test_parameterInstruction_opcode(t *testing.T) {
	for _, tc := range []struct {
		pi parameterInstruction
		opcode
	}{
		{
			pi:     parameterInstruction{},
			opcode: 0,
		},
		{
			pi:     parameterInstruction{3},
			opcode: 3,
		},
		{
			pi:     parameterInstruction{1, 2},
			opcode: 12,
		}, {
			pi:     parameterInstruction{9, 4, 6, 5, 7, 1, 6, 5, 4},
			opcode: 54,
		},
	} {
		t.Run(fmt.Sprint(tc.pi), func(t *testing.T) {
			assert.Equal(t, tc.opcode, tc.pi.opcode())
		})
	}
}

func Test_parameterInstruction_parameterMode(t *testing.T) {
	for _, tc := range []struct {
		pi    parameterInstruction
		index int
		want  parameterMode
	}{
		{
			pi:    parameterInstruction{7, 1, 2},
			index: 0,
			want:  7,
		}, {
			pi:    parameterInstruction{7, 1, 2},
			index: 1,
			want:  0,
		}, {
			pi:    parameterInstruction{7, 1, 2},
			index: 2,
			want:  0,
		}, {
			pi:    parameterInstruction{7, 1, 2, 62, 3, 54, 4},
			index: 2,
			want:  2,
		},
	} {
		t.Run(fmt.Sprint(tc.pi), func(t *testing.T) {
			assert.Equal(t, tc.want, tc.pi.parameterMode(tc.index))
		})
	}
}
