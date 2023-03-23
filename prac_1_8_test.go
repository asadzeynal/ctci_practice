package main

import (
	"fmt"
	"testing"
)

func TestZeroMatrix(t *testing.T) {
	tests := []struct {
		s    [][]int
		want [][]int
	}{
		{
			[][]int{{0, 0, 0}, {4, 5, 6}, {7, 8, 9}},
			[][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		},
		{
			[][]int{{1, 2, 0}, {4, 5, 6}, {7, 0, 9}},
			[][]int{{0, 0, 0}, {4, 0, 0}, {0, 0, 0}},
		},
		{
			[][]int{{1, 2, 0, 4}, {5, 6, 7, 8}, {9, 10, 0, 12}, {13, 14, 15, 16}},
			[][]int{{0, 0, 0, 0}, {5, 6, 0, 8}, {0, 0, 0, 0}, {13, 14, 0, 16}},
		},
		{
			[][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}},
			[][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.s)
		t.Run(testname, func(t *testing.T) {
			ans := zeroMatrix(tt.s)
			if !EqualMatrix(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
