package ch2

import (
	"fmt"
	"testing"

	"github.com/asadzeynal/ctci_practice/util"
)

var tests = []struct {
	a    *util.Node
	k    int
	want int
}{
	{util.LinkedListFromSlice([]int{1, 2, 3, 4, 5}), 0, 5},
	{util.LinkedListFromSlice([]int{1, 2, 3, 4, 5}), 1, 4},
	{util.LinkedListFromSlice([]int{1, 2, 3, 4, 5}), 2, 3},
	{util.LinkedListFromSlice([]int{1, 2, 3, 4, 5}), 3, 2},
	{util.LinkedListFromSlice([]int{1, 2, 3, 4, 5}), 4, 1},
}

func TestKthToLastIterative(t *testing.T) {
	for _, tt := range tests {
		testname := fmt.Sprintf("%v %v", tt.a, tt.k)
		t.Run(testname, func(t *testing.T) {
			ans := kToLastIterative(tt.a, tt.k)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestKthToLastRecursive(t *testing.T) {
	for _, tt := range tests {
		testname := fmt.Sprintf("%v %v", tt.a, tt.k)
		t.Run(testname, func(t *testing.T) {
			ans := kToLastRecursive(tt.a, tt.k)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
