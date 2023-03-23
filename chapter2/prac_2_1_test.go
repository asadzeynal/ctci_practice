package ch2

import (
	"fmt"
	"testing"

	"github.com/asadzeynal/ctci_practice/util"
)

func TestRemoveDups(t *testing.T) {
	tests := []struct {
		a    *util.Node
		want *util.Node
	}{
		{util.LinkedListFromSlice([]int{1, 4, 4, 1, 5}), util.LinkedListFromSlice([]int{1, 4, 5})},
		{util.LinkedListFromSlice([]int{1, 1, 1, 1, 1}), util.LinkedListFromSlice([]int{1})},
		{util.LinkedListFromSlice([]int{1, 1, 1, 2, 1}), util.LinkedListFromSlice([]int{1, 2})},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := removeDups(tt.a)
			resAns := util.SliceFromLinkedList(ans)
			resWant := util.SliceFromLinkedList(tt.want)
			if !util.EqualSlice(resAns, resWant) {
				t.Errorf("got %v, want %v", resAns, resWant)
			}
		})
	}
}
