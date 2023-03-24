package ch2

import (
	"testing"

	"github.com/asadzeynal/ctci_practice/util"
)

func TestIntersect(t *testing.T) {
	node1 := &util.Node{Data: 0}
	node2 := &util.Node{Data: 0}
	node3 := &util.Node{Data: 0}
	node4 := &util.Node{Data: 0}
	node5 := &util.Node{Data: 0}
	node6 := &util.Node{Data: 0}
	node7 := &util.Node{Data: 0}

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

	res3 := intersectWithReturn(node1, node5)
	if res3 != node3 {
		t.Errorf("got %v, want %v", res3, node3)
	}

	res4 := intersectWithReturn(node1, node6)
	if res4 != nil {
		t.Errorf("got %v, want %v", res4, nil)
	}

}
