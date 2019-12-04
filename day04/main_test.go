package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_splitChars(t *testing.T) {
	chars := splitChars(165482)
	assert.Equal(t, [6]int{1, 6, 5, 4, 8, 2}, chars)
}

func Test_num_containsDouble(t *testing.T) {
	tests := []struct {
		name string
		n    num
		want bool
	}{
		{
			n:    num{1, 2, 3, 4, 5, 6},
			want: false,
		},
		{
			n:    num{1, 2, 3, 4, 4, 6},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.containsDouble(); got != tt.want {
				t.Errorf("getDoublesDigitCounts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_num_neverDecreases(t *testing.T) {
	tests := []struct {
		name string
		n    num
		want bool
	}{
		{
			n:    num{1, 1, 1, 1, 1, 1},
			want: true,
		},
		{
			n:    num{1, 1, 1, 3, 5, 6},
			want: true,
		},
		{
			n:    num{1, 1, 1, 3, 1, 6},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.neverDecreases(); got != tt.want {
				t.Errorf("neverDecreases() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_num_getDoublesDigitCounts(t *testing.T) {
	tests := []struct {
		name string
		n    num
		want map[int]int
	}{
		{
			n:    num{0, 1, 2, 3, 4, 5},
			want: map[int]int{},
		},
		{
			n: num{0, 0, 2, 3, 3, 5},
			want: map[int]int{
				0: 1,
				3: 1,
			},
		},
		{
			n: num{0, 0, 0, 3, 3, 5},
			want: map[int]int{
				0: 2,
				3: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.getDoublesDigitCounts(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDoublesDigitCounts() = %v, want %v", got, tt.want)
			}
		})
	}
}
