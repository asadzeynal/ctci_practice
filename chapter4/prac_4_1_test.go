package ch4

import "testing"

func TestRouteBetweenNodes(t *testing.T) {
	nodes := []*GraphNode{
		{Data: 0},
		{Data: 1},
		{Data: 2},
		{Data: 3},
		{Data: 4},
		{Data: 5},
	}

	nodes[0].Adjacent = []*GraphNode{nodes[1], nodes[2]}
	nodes[2].Adjacent = []*GraphNode{nodes[1], nodes[3]}
	nodes[3].Adjacent = []*GraphNode{nodes[0], nodes[4]}
	nodes[5].Adjacent = []*GraphNode{nodes[0], nodes[4]}

	if !routeBetweenNodes(nodes[0], nodes[4]) {
		t.Errorf("got %v, want %v", false, true)
	}

	if routeBetweenNodes(nodes[0], nodes[4]) {
		t.Errorf("got %v, want %v", true, false)
	}
}
