package ch2

import (
	"fmt"
	"testing"

	"github.com/asadzeynal/ctci_practice/util"
)

func TestPartition(t *testing.T) {
	var tests = []struct {
		n    *util.Node
		p    int
		want *util.Node
	}{
		// Bad test, tests only specific algorithm, but ok for now
		{util.LinkedListFromSlice([]int{3, 5, 8, 5, 10, 2, 1}), 5, util.LinkedListFromSlice([]int{3, 2, 1, 5, 8, 5, 10})},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v %v", tt.n, tt.p)
		t.Run(testname, func(t *testing.T) {
			ans := partition(tt.n, tt.p)
			ansSlice := util.SliceFromLinkedList(ans)
			wantSlice := util.SliceFromLinkedList(tt.want)
			if !util.EqualSlice(ansSlice, wantSlice) {
				t.Errorf("got %v, want %v", ansSlice, wantSlice)
			}
		})
	}
}
