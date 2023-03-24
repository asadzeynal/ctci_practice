package ch2

import (
	"testing"

	"github.com/asadzeynal/ctci_practice/util"
)

func TestIntersect(t *testing.T) {
	node1 := &util.Node{Data: 1}
	node2 := &util.Node{Data: 2}
	node3 := &util.Node{Data: 3}
	node4 := &util.Node{Data: 4}
	node5 := &util.Node{Data: 5}
	node6 := &util.Node{Data: 6}
	node7 := &util.Node{Data: 7}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	node5.Next = node3

	node6.Next = node7

	res1 := intersect(node1, node5)
	if !res1 {
		t.Errorf("got %v, want %v", res1, true)
	}

	res2 := intersect(node1, node6)
	if res2 {
		t.Errorf("got %v, want %v", res2, false)
	}

}
