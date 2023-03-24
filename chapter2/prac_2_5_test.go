package ch2

import (
	"fmt"
	"testing"

	"github.com/asadzeynal/ctci_practice/util"
)

func TestSumListsReverse(t *testing.T) {
	var tests = []struct {
		n1   []int
		n2   []int
		want []int
	}{
		{[]int{7, 1, 6}, []int{5, 9, 2}, []int{2, 1, 9}},
		{[]int{9, 9, 9}, []int{9, 9, 9}, []int{8, 9, 9, 1}},
		{[]int{9, 9, 9, 9, 9}, []int{9, 9, 9}, []int{8, 9, 9, 0, 0, 1}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v %v", tt.n1, tt.n2)
		t.Run(testname, func(t *testing.T) {
			node1 := util.LinkedListFromSlice(tt.n1)
			node2 := util.LinkedListFromSlice(tt.n2)

			ans := sumListsReverseRecurse(node1, node2)
			ansSlice := util.SliceFromLinkedList(ans)

			if !util.EqualSlice(ansSlice, tt.want) {
				t.Errorf("got %v, want %v", ansSlice, tt.want)
			}
		})
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v %v", tt.n1, tt.n2)
		t.Run(testname, func(t *testing.T) {
			node1 := util.LinkedListFromSlice(tt.n1)
			node2 := util.LinkedListFromSlice(tt.n2)

			ans := sumListsReverse(node1, node2)
			ansSlice := util.SliceFromLinkedList(ans)

			if !util.EqualSlice(ansSlice, tt.want) {
				t.Errorf("got %v, want %v", ansSlice, tt.want)
			}
		})
	}
}
