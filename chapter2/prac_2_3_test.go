package ch2

import (
	"fmt"
	"testing"

	"github.com/asadzeynal/ctci_practice/util"
)

func TestDeleteMiddleNode(t *testing.T) {
	ll := util.LinkedListFromSlice([]int{1, 2, 3, 4, 5})
	node := ll.Next.Next

	var tests = []struct {
		a    *util.Node
		want []int
	}{
		{node, []int{1, 2, 4, 5}},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.a)
		t.Run(testname, func(t *testing.T) {
			deleteMiddleNode(tt.a)
			ans := util.SliceFromLinkedList(ll)
			if !util.EqualSlice(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
